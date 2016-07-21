package bookgen

import (
	"fmt"
	"github.com/remotejob/jbs_generator/dbhandler"
	"github.com/shogo82148/go-shuffle"
	"gopkg.in/mgo.v2"
	"math/rand"
	"os"
	"time"
)

func Create(session mgo.Session, filename string) {

	if _, err := os.Stat(filename); !os.IsNotExist(err) {

		err := os.Remove(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	joboffers := dbhandler.GetAllUseful(session)

	var numberstoshuffle []int

	for num, _ := range joboffers {

		numberstoshuffle = append(numberstoshuffle, num)

	}
	rand.Seed(time.Now().UTC().UnixNano())

	shuffle.Ints(numberstoshuffle)

	for _, i := range numberstoshuffle {

		moto := ""

		for _, tag := range joboffers[i].Tags {

			moto = moto + " " + tag

		}

		paragraph := joboffers[i].Title + "\n" + moto + "\n" + joboffers[i].Description + "\n"

		if _, err = f.WriteString(paragraph); err != nil {
			panic(err)
		}
	}
	f.Close()
}
