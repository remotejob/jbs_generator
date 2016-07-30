package domains

import (
	"encoding/xml"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Articlefull struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Stitle    string
	Tags      string
	Contents  string
	Mcontents string
	Site      string
	Created   time.Time
	Updated   time.Time
}

type Article struct {
	Title     string
	Tags      string
	Contents  string
	Mcontents string
}

type Sitemap_from_db struct {
	Stitle  string
	Site    string
	Updated time.Time
}

type ServerConfig struct {
	Dbmgo struct {
		Addrs     []string
		Database  string
		Username  string
		Password  string
		Mechanism string
	}

	Dbmgoext struct {
		Addrs     []string
		Database  string
		Username  string
		Password  string
		Mechanism string
	}
	Sites struct {
		Site []string
	}
	Dirs struct {
		Sitemapsdir string
		Webrootdir  string
	}

	Routes struct {
		Mainroute string
	}

	Files struct {
		Commonwords string
	}
}

type JobOffer struct {
	Title       string
	Tags        []string
	Description string
}

//type BlogIndex struct {
//	Stopic string
//	Topic  string
//}
//
//type Md struct {
//	Locale string
//	Themes string
//	Site   string
//	Menu   string
//	Stopic string
//	Topic  string
//	Items  []BlogItem
//}
//
//type Rssresors struct {
//	Topic string
//	Link  string
//}

type SitemapObj struct {
	Changefreq    string
	Hoursduration float64
	Loc           string
	Lastmod       string
}

//type BlogItem struct {
//	Stopic     string
//	Topic      string
//	Stitle     string
//	Title      string
//	Contents   string
//	Created_at time.Time
//	Updated_at time.Time
//}
//
//type Blog struct {
//	//	Topic string
//	Items map[string][]BlogItem
//}
//
//type KeywordObj struct {
//	Keyword string
//	Cl      int
//	Lang    string
//}
//
//type Contents struct {
//	Title      string
//	Moto       string
//	Contents   string
//	Created_at time.Time
//	Updated_at time.Time
//}

type Pages struct {
	//	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	//	XmlImageNS string   `xml:"xmlns:image,attr"`
	//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
	Pages []*Page `xml:"url"`
}

type Page struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	Changefreq string   `xml:"changefreq"`
	//	Name       string   `xml:"news:news>news:publication>news:name"`
	//	Language   string   `xml:"news:news>news:publication>news:language"`
	//	Title      string   `xml:"news:news>news:title"`
	//	Keywords   string   `xml:"news:news>news:keywords"`
	//	Image      string   `xml:"image:image>image:loc"`
}

type Config struct {
	Maintitle string
	Subtitle  string
	Cv        []struct {
		Name string
		Path string
		Img  string
		Item []struct {
			Title    string
			Rank     int
			Duration int
			Link     string
			Extra    string
			Img      string
		}
	}
}

type Job struct {
	Maintitle string
	Subtitle  string
	Jobs      []struct {
		Name string
		Path string
		Img  string
		Item []struct {
			Title    string
			Rank     int
			Duration string
			Position string
			Details  string
			Location string
			Country  string
		}
	}
}
