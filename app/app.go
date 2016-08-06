package main

import (
	"bytes"
	"fmt"
	"github.com/remotejob/comutils/gen"

	"github.com/remotejob/jbs_generator/bookgen"
	"github.com/remotejob/jbs_generator/dbhandler"
	"github.com/remotejob/jbs_generator/domains"
	"github.com/remotejob/jbs_generator/entryHandler"
	"github.com/remotejob/jbs_generator/sentenses_tokenizer"
	"github.com/remotejob/jbs_generator/wordscount"
//	"github.com/remotejob/jbs_generator/savetags"

	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"time"
	//	"fmt"
	"io"
	"log"
	"os"
)

var addrs []string
var database string
var username string
var password string
var mechanism string

var addrsext []string
var databaseext string
var usernameext string
var passwordext string
var mechanismext string

var sites []string
var commonwords string

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		addrs = cfg.Dbmgo.Addrs
		database = cfg.Dbmgo.Database
		username = cfg.Dbmgo.Username
		password = cfg.Dbmgo.Password
		mechanism = cfg.Dbmgo.Mechanism

		addrsext = cfg.Dbmgoext.Addrs
		databaseext = cfg.Dbmgoext.Database
		usernameext = cfg.Dbmgoext.Username
		passwordext = cfg.Dbmgoext.Password
		mechanismext = cfg.Dbmgoext.Mechanism

		sites = cfg.Sites.Site
		commonwords = cfg.Files.Commonwords

	}

}

func main() {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Timeout:   60 * time.Second,
		Database:  database,
		Username:  username,
		Password:  password,
		Mechanism: mechanism,
	}

	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	mongoDBDialInfoext := &mgo.DialInfo{
		Addrs:     addrsext,
		Timeout:   60 * time.Second,
		Database:  databaseext,
		Username:  usernameext,
		Password:  passwordext,
		Mechanism: mechanism,
	}

	dbsessionext, err := mgo.DialWithInfo(mongoDBDialInfoext)

	if err != nil {
		panic(err)
	}
	defer dbsessionext.Close()

	bookgen.Create(*dbsessionext, "/tmp/book.txt")

	buf := bytes.NewBuffer(nil)

	f, _ := os.Open("/tmp/book.txt")

	io.Copy(buf, f) // Error handling elided for brevity.
	f.Close()

	allsitemaplinks := dbhandler.GetAllSitemaplinks(*dbsession)

	uniq_links := make(map[string]struct{})

	for _, sitemaplink := range allsitemaplinks {
		uniq_links[sitemaplink.Stitle] = struct{}{}

	}

	bestKeywords := wordscount.GetBestKeywords(buf.Bytes(), commonwords, 500)

	sentenses_quant := gen.Random(5, 10)

	sentences := sentenses_tokenizer.Do_tokens(buf.Bytes(), sentenses_quant)

	newArticle := entryHandler.NewEntryarticle()

	stitle := newArticle.AddTitleStitleMcontents(buf.Bytes(), sites, uniq_links)

	if _, ok := uniq_links[stitle]; !ok {

		uniq_links[stitle] = struct{}{}

		newArticle.AddTags(bestKeywords)
		newArticle.AddContents(sentences)
		newArticle.AddAuthor()
		newArticle.InsertIntoDB(*dbsession)

		//	fmt.Println(newArticle.Modarticle.Title)
		//	fmt.Println(newArticle.Modarticle.Stitle)
		fmt.Println(newArticle.Modarticle.Tags)
		//	fmt.Println(newArticle.Modarticle.Contents)

	} else {
		fmt.Println("Creates stitle EXIST!! but it possible", stitle)
	}

}
