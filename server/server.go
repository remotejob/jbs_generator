package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/remotejob/jbs_generator/dbhandler"
		"github.com/remotejob/jbs_generator/handlers/robots"
			"github.com/remotejob/jbs_generator/handlers"
	"github.com/remotejob/jbs_generator/domains"
	//	"github.com/rs/cors"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

var addrs []string
var database string
var username string
var password string
var mechanism string
var sites []string
var webrootdir string
var startpar handlers.Entrystartpar


type Startpar struct {
	
	
	
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
		webrootdir = cfg.Dirs.Webrootdir
//		addrs = cfg.Dbmgo.Addrs
//		database = cfg.Dbmgo.Database
//		username = cfg.Dbmgo.Username
//		password = cfg.Dbmgo.Password
//		mechanism = cfg.Dbmgo.Mechanism
//
//		addrsext = cfg.Dbmgoext.Addrs
//		databaseext = cfg.Dbmgoext.Database
//		usernameext = cfg.Dbmgoext.Username
//		passwordext = cfg.Dbmgoext.Password
//		mechanismext = cfg.Dbmgoext.Mechanism

		sites = cfg.Sites.Site
			
		
		
//		commonwords = cfg.Files.Commonwords

	}
	
	startpar = *handlers.NewEntrystartpar()
	startpar.Modstartpar.Webrootdir = webrootdir		

}





func defaultHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)

	fmt.Println("defaultHandler->", urlParams["article"])

	article := strings.Replace(urlParams["article"], ".json", "", -1)

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

	articleobj := dbhandler.GetOneArticle(*dbsession, article)
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Content-Type", "application/json")

	if articleobj.Title == "" {

		http.NotFound(w, r)

	} else {
		articlejson, _ := json.Marshal(articleobj)

		fmt.Fprintf(w, string(articlejson))

	}

}


func main() {

	

	r := mux.NewRouter()
	//	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/juno/git/remotejob/cv_webpack_desk/static/"))))
	//	r.PathPrefix("/{article}").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/juno/git/remotejob/cv_webpack_desk/static/"))))
	//	r.NotFoundHandler = http.HandlerFunc(handlers.StaticHandler)
	r.HandleFunc("/robots.txt",robots.Generate )
	r.HandleFunc("/sitemap.xml",handlers.CheckServeSitemap)
	r.HandleFunc("/jobs.html",startpar.StaticHandler)
	r.HandleFunc("/contacts.html",startpar.StaticHandler)					
	r.HandleFunc("/api/{article}", defaultHandler)
	r.HandleFunc("/jobs/{article}",startpar.StaticHandler)		
//	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/juno/git/remotejob/cv_webpack_desk/static"))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(webrootdir))))	

	log.Println("Listening at port 8080")


	log.Fatal(http.ListenAndServe(":8080", r))
}
