package main

import (
	"flag"
	"fmt"
	"github.com/remotejob/jbs_generator/domains"
	"github.com/remotejob/jbs_generator/dbutils/updateAuthor"	
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"time"
	"log"
	//	"fmt"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

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
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}
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
	
	updateAuthor.CreateUpdate(*dbsession)
}
