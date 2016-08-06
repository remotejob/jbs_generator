package main

import (
	"encoding/json"
	"fmt"
	"github.com/kazarena/json-gold/ld"
	"github.com/remotejob/jbs_generator/dbhandler"
	"github.com/remotejob/jbs_generator/domains"
	"github.com/remotejob/jbs_generator/home_page"
	"github.com/remotejob/jbs_generator/static_site_gen/create_json_ld"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
	"os"
	"path"
	//	"strings"
	"time"
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

	lphead := path.Join("templates", "header.html")
	lp := path.Join("templates", "layout.html")

	funcMap := template.FuncMap{
		"marshal": func(a []byte) template.JS {

			return template.JS(a)
		},
	}

	t, err := template.New("layout.html").Funcs(funcMap).ParseFiles(lp, lphead)
	check(err)

	allarticles := dbhandler.GetAllForStatic(*dbsession)

	for _, site := range sites {

		home_page.Create(allarticles, pwd, site)

	}

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	for _, articlefull := range allarticles {

		articleobj := domains.Article{articlefull.Title, articlefull.Tags, articlefull.Contents, articlefull.Mcontents,articlefull.Author}

		createdstr := articlefull.Created.Format("2006-01-02")
		updatedstr := articlefull.Updated.Format("2006-01-02")

		jsonld := create_json_ld.Create(proc, options, articlefull)

		articletotemplate := domains.Articletotempalte{
			Title:     articlefull.Title,
			Stitle:    articlefull.Stitle,
			Tags:      articlefull.Tags,
			Contents:  articlefull.Contents,
			Mcontents: articlefull.Mcontents,
			Site:      articlefull.Site,
			Created:   createdstr,
			Updated:   updatedstr,
			Jsonld:    jsonld,
		}

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

		err = t.Execute(f, articletotemplate)
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

}
