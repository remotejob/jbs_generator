package home_page

import (
	"fmt"
	"github.com/remotejob/jbs_generator/domains"
	"github.com/shogo82148/go-shuffle"
	"html/template"
	"math/rand"
	"os"
	"path"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Create(allarticles []domains.Articlefull, pwd string, site string) {

	dirstr := path.Join(pwd, "www", site)
	filestr := path.Join(pwd, "www", site, "index.html")
	fmt.Println(dirstr)
	fmt.Println(filestr)

	os.MkdirAll(dirstr, 0777)

	var numberstoshuffle []int
	for num, _ := range allarticles {

		numberstoshuffle = append(numberstoshuffle, num)

	}
	rand.Seed(time.Now().UTC().UnixNano())

	shuffle.Ints(numberstoshuffle)

	var articles_to_inject []domains.Articlefull

	for c, i := range numberstoshuffle {

		if allarticles[i].Site == site {

			articles_to_inject = append(articles_to_inject, allarticles[i])

		}

		if c > 500 {

			break
		}

	}

	lp := path.Join("templates", "home_page.html")

	t, err := template.ParseFiles(lp)
	check(err)

	f, err := os.Create(filestr)
	if err != nil {
		//    log.Println("create file: ", err)
		check(err)
		return
	}

	err = t.Execute(f, articles_to_inject)
	check(err)

}