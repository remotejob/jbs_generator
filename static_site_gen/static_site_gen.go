package main

import (
	"fmt"
	"github.com/remotejob/jbs_generator/dbhandler"
	"github.com/remotejob/jbs_generator/domains"
	"github.com/remotejob/jbs_generator/home_page"	
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
	"os"
	"path"
	"time"
	"encoding/json"	
)

var addrs []string
var database string
var username string
var password string
var mechanism string
var sites []string

var mainroute string

var pwd string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
		mainroute = cfg.Routes.Mainroute
		sites = cfg.Sites.Site

	}

	getpwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//	fmt.Println(pwd)
	pwd = getpwd
	//	fmt.Println(pwd)
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

	lp := path.Join("templates", "layout.html")

	t, err := template.ParseFiles(lp)
	check(err)

	fmt.Println("pwd", pwd)

	allarticles := dbhandler.GetAllForStatic(*dbsession)
	
	
	for _,site :=range sites {
		
		home_page.Create(allarticles, pwd, site)
		
	}
	
	for _, articlefull := range allarticles {

		
		articleobj := domains.Article{articlefull.Title,articlefull.Tags,articlefull.Contents,articlefull.Mcontents}
		
		articlejson, _ := json.Marshal(articleobj)		


		dirstr := path.Join(pwd, "www", articlefull.Site, mainroute)
		filestr := path.Join(pwd, "www", articlefull.Site, mainroute, articlefull.Stitle+".html")
		filestrjson := path.Join(pwd, "www", articlefull.Site, mainroute, articlefull.Stitle+".json")		
		os.MkdirAll(dirstr, 0777)

		f, err := os.Create(filestr)
		if err != nil {
			//    log.Println("create file: ", err)
			check(err)
			return
		}
		err = t.Execute(f, articlefull)
		check(err)

		fmt.Println(filestr)
		
		 jsonFile, err := os.Create(filestrjson)

         if err != nil {
                 fmt.Println(err)
         }
         defer jsonFile.Close()
		 jsonFile.Write(articlejson)
         jsonFile.Close()
		
		

	}

	//	err = t.Execute(f, domains.Article)
	//	check(err)

}
