package static_site_gen

import (
	"github.com/kazarena/json-gold/ld"
	"github.com/Pallinder/go-randomdata"
	"log"
	"testing"
	"fmt"
	"encoding/json"	
)

//<script type="application/ld+json">
//{
//  "@context": "http://schema.org",
//  "@type": "Article",
//  "author": "John Doe",
//  "name": "How to Tie a Reef Knot",
//  "datePublished": "2010-10-10",
//  "headline": "head line",
//  "image": {"@type": "ImageObject", "url": "http://example.com/image.jpg","height":"10px","width":"20px"},
//  "publisher":{"@type": "Organization", "name": "Alex Mazurov", "logo": {"@type": "ImageObject", "url": "http://example.com/image.jpg","height":"10px","width":"20px"}},
//  "dateModified": "2010-10-10",
//   "mainEntityOfPage": {"@type": "WebPage","@id": "http://cathscafe.example.com"}
//}
//</script>

func TestCreateJsonLD(t *testing.T) {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	//	options.Format = "application/json-ld"

	publisher := map[string]interface{}{"@type": "Organization", "name": "Aleksander Mazurov", "logo": map[string]interface{}{"@type": "ImageObject", "url": "http://example.com/image.jpg", "height": "10px", "width": "20px"}}
	image := map[string]interface{}{"@type": "ImageObject", "url": "http://example.com/image.jpg", "height": "10px", "width": "20px"}
	mainEntityOfPage := map[string]interface{}{"@type": "WebPage", "@id": "http://cathscafe.example.com"}

	doc := map[string]interface{}{
		"@context":         "http://schema.org",
		"@type":            "Article",
		"author":           "Aleksander Mazurov",
		"headline":         "head line",
		"publisher":        publisher,
		"image":            image,
		"datepublished":    "2009-05-08",
		"datemodified":     "2009-05-08",
		"mainEntityOfPage": mainEntityOfPage,
		"keywords":         "seo sales b2b",
		//		"wordcount":           "1120",
		"url": "http://www.example.com",
		//		"description":         "We love to do stuff to help people and stuff",
		"articleBody": "You can paste your entire post in here, and yes it can get really really long.",
	}


	comp, err := proc.Compact(doc, nil, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)
		return
	}
	ld.PrintDocument("", comp)
	b, _ := json.MarshalIndent(comp, "", "  ")
	fmt.Println(string(b))			

}

func TestNameGen(t *testing.T) {
	
// fmt.Println(randomdata.SillyName())

//    // Print a male first name
//    fmt.Println(randomdata.FirstName(randomdata.Male))
//
//    // Print a female first name
//    fmt.Println(randomdata.FirstName(randomdata.Female))
//
//    // Print a last name
//    fmt.Println(randomdata.LastName())
//
//    // Print a male name
//    fmt.Println(randomdata.FullName(randomdata.Male))
//
//    // Print a female name
//    fmt.Println(randomdata.FullName(randomdata.Female))

    // Print a name with random gender
    fmt.Println(randomdata.FullName(randomdata.RandomGender))

    // Print an email
//    fmt.Println(randomdata.Email())
//
//    // Print a country with full text representation
//    fmt.Println(randomdata.Country(randomdata.FullCountry))
//
//    // Print a country using ISO 3166-1 alpha-2
//    fmt.Println(randomdata.Country(randomdata.TwoCharCountry))
//
//    // Print a country using ISO 3166-1 alpha-3
//    fmt.Println(randomdata.Country(randomdata.ThreeCharCountry))
//
//    // Print a currency using ISO 4217
//    fmt.Println(randomdata.Currency())
//
//    // Print the name of a random city
//    fmt.Println(randomdata.City())
//
//    // Print the name of a random american state
//    fmt.Println(randomdata.State(randomdata.Large))
//
//    // Print the name of a random american state using two chars
//    fmt.Println(randomdata.State(randomdata.Small))
//
//    // Print an american sounding street name
//    fmt.Println(randomdata.Street())
//
//    // Print an american sounding address
//    fmt.Println(randomdata.Address())
//
//    // Print a random number >= 10 and <= 20
//    fmt.Println(randomdata.Number(10, 20))
//
//    // Print a number >= 0 and <= 20
//    fmt.Println(randomdata.Number(20))
//
//    // Print a random float >= 0 and <= 20 with decimal point 3
//    fmt.Println(randomdata.Decimal(0, 20, 3))
//
//    // Print a random float >= 10 and <= 20
//    fmt.Println(randomdata.Decimal(10, 20))
//
//    // Print a random float >= 0 and <= 20
//    fmt.Println(randomdata.Decimal(20))
//
//    // Print a bool
//    fmt.Println(randomdata.Boolean())
//
//    // Print a paragraph
//    fmt.Println(randomdata.Paragraph())
//
//    // Print a postal code 
//    fmt.Println(randomdata.PostalCode("SE"))
//
//    // Print a set of 2 random numbers as a string
//    fmt.Println(randomdata.StringNumber(2, "-")) 
//
//    // Print a set of 2 random 3-Digits numbers as a string
//    fmt.Println(randomdata.StringNumberExt(2, "-", 3)) 
//
//    // Print a valid random IPv4 address
//    fmt.Println(randomdata.IpV4Address())
//
//    // Print a day
//    fmt.Println(randomdata.Day())
//
//    // Print a month
//    fmt.Println(randomdata.Month())
//
//    // Print full date like Thursday 22 Aug 2016
//    fmt.Println(randomdata.FullDate())	
	
}

