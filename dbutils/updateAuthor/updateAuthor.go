package updateAuthor

import (
	"github.com/remotejob/jbs_generator/domains"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/Pallinder/go-randomdata"
)

func CreateUpdate(session mgo.Session) {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("jbs_generator").C("jbs_generator")
	var results []domains.Articlefull
	err := c.Find(nil).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	for _, article := range results {

		log.Println(article.Stitle)

		colQuerier := bson.M{"_id": bson.ObjectId(article.ID)}

		//		var result domains.Articlefull
		//		err = c.FindId(bson.ObjectIdHex(article.ID)).One(&result)
		author := randomdata.FullName(randomdata.RandomGender)
		change := bson.M{"$set": bson.M{"author": author}}

		err = c.Update(colQuerier, change)
		if err != nil {
			panic(err)
		}

	}

}
