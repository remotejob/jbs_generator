package entryHandler

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/gosimple/slug"
	"github.com/remotejob/comutils/gen"
	"github.com/remotejob/comutils/str"
	"github.com/remotejob/jbs_generator/mgenerator"
	//	"github.com/remotejob/jbs_generator/dbhandler/Aprinter"
	"gopkg.in/mgo.v2"
	"math/rand"
	//	"github.com/remotejob/jbs_generator/wordscount"
	"github.com/remotejob/jbs_generator/savetags"
	"log"
	"strings"
	"time"
)

type Article struct {
	Title     string
	Stitle    string
	Tags      string
	Contents  string
	Mcontents string
	Site      string
	Author    string
	Created   time.Time
	Updated   time.Time
}

type Entryarticle struct {
	Modarticle Article
}

func NewEntryarticle() *Entryarticle {
	return &Entryarticle{Article{}}
}

func (article *Entryarticle) AddAuthor() {

	authorName := randomdata.FullName(randomdata.RandomGender)

	article.Modarticle.Author = authorName

}

func (article *Entryarticle) AddTitleStitleMcontents(bfile []byte, sites []string, uniqlinks map[string]struct{}) string {

	rand.Seed(time.Now().UTC().UnixNano())
	siteid := gen.Random(0, len(sites))

	mtext := mgenerator.Generate(bfile)

	mtexttokens := strings.Fields(mtext)

	var title string = ""

	for i, token := range mtexttokens {

		title = title + " " + token
		if i >= 10 {
			break
		}

	}

	stitle := slug.Make(title)
	article.Modarticle.Title = str.UpcaseInitial(title)
	article.Modarticle.Stitle = slug.Make(title)
	article.Modarticle.Mcontents = mtext
	article.Modarticle.Site = sites[siteid]

	return stitle

}

func (article *Entryarticle) AddTags(tags []string) {

	var tagsquant = len(tags)
	var tags_str string = ""
	var tags_to_save []string
	for i := 0; i < 10; i++ {

		tagint := gen.Random(0, tagsquant)
		tags_to_save = append(tags_to_save, tags[tagint])
		tags_str = tags_str + " " + tags[tagint]

	}

	savetags.Saveinfile("createdtags.csv", tags_to_save)

	article.Modarticle.Tags = strings.TrimSpace(tags_str)

}

func (article *Entryarticle) AddContents(sentenses []string) {

	var contents string = ""

	for _, sentens := range sentenses {

		contents = contents + " " + strings.Replace(sentens, "\n", "", -1)

	}

	article.Modarticle.Contents = str.UpcaseInitial(contents)

}

func (article *Entryarticle) InsertIntoDB(session mgo.Session) {

	now := time.Now()
	articletodb := Article{article.Modarticle.Title, article.Modarticle.Stitle, article.Modarticle.Tags, article.Modarticle.Contents, article.Modarticle.Mcontents, article.Modarticle.Site, article.Modarticle.Author, now, now}
	//	dbhandler.InsetArticle(session, articletodb)
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("jbs_generator").C("jbs_generator")

	err := c.Insert(articletodb)
	if err != nil {
		log.Fatal(err)
	}

}
