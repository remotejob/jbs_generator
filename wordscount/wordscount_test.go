package wordscount

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
	"strconv"	
	//	"unicode/utf8"
)

func TestSingleWordCount(t *testing.T) {

	//	removePunctuation := func(r rune) rune {
	//		if strings.ContainsRune("?!.;,*/:()123456789'-", r) {
	//			return -1
	//		} else {
	//			return r
	//		}
	//	}

	get_words_from := func(text string) []string {

		words := regexp.MustCompile("[\\p{L}\\d_]+")

		return words.FindAllString(text, -1)

	}

	buf := bytes.NewBuffer(nil)

	fcsv, _ := os.Open("/home/juno/neonworkspace/jbs_generator/commonwords.csv")
	r := csv.NewReader(bufio.NewReader(fcsv))
	set_commonwords := make(map[string]struct{})

	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		//	fmt.Println(record[0])

		set_commonwords[record[0]] = struct{}{}

		//	fmt.Println(len(record))
		//	for value := range record {
		//	    fmt.Printf("  %v\n", record[value])
		//	}
	}

	f, _ := os.Open("/tmp/book.txt")

	io.Copy(buf, f) // Error handling elided for brevity.
	f.Close()

	s := buf.String()

	res := get_words_from(s)

	word_counts := make(map[string]int)
	for _, value := range res {
		//		fmt.Println(key, value)
		valuetoinsert := strings.ToLower(value)
		
		if _, err := strconv.Atoi(valuetoinsert); err != nil {		

		if len(valuetoinsert) > 2 {

			if _, ok := set_commonwords[valuetoinsert]; !ok {

				word_counts[valuetoinsert]++
			}
		}
		
		}

	}

	n := map[int][]string{}
	var a []int
	for k, v := range word_counts {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
//	for _, k := range a {
//		for _, s := range n[k] {
//			fmt.Printf("%s, %d\n", s, k)
//		}
//	}

	fmt.Println("count", len(word_counts))

}
