package dbhandler

import (
	//	"fmt"
	"github.com/remotejob/jbs_generator/domains"
	"github.com/remotejob/jbs_generator/entryHandler"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	//	"time"
)

func GetAllForStatic(session mgo.Session) []domains.Articlefull {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("jbs_generator").C("jbs_generator")
	var results []domains.Articlefull
	err := c.Find(nil).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results

}

func GetOneArticle(session mgo.Session, stitle string) domains.Article {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("jbs_generator").C("jbs_generator")

	var result domains.Article

	err := c.Find(bson.M{"stitle": stitle}).Select(bson.M{"created": 0, "updated": 0, "stitle": 0, "site": 0}).One(&result)
	if err != nil {

		log.Fatal(err)
		//		return
	}

	return result

}

func GetAllSitemaplinks(session mgo.Session) []domains.Sitemap_from_db {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("jbs_generator").C("jbs_generator")
	var results []domains.Sitemap_from_db
	err := c.Find(nil).Select(bson.M{"stitle": 1, "site": 1, "updated": 1}).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results
}

func GetAllUseful(session mgo.Session) []domains.JobOffer {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cv_employers").C("employers")

	var results []domains.JobOffer

	err := c.Find(nil).Select(bson.M{"description": 1, "title": 1, "tags": 1}).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	return results
}

func InsetArticle(session mgo.Session, article entryHandler.Article) {
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("jbs_generator").C("jbs_generator")

	err := c.Insert(article)
	if err != nil {
		log.Fatal(err)
	}

}
