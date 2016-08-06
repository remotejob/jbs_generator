package wordscount

import (
	//	"strings"
	//	"regexp"
	//	"fmt"
	"bufio"
	//	"bytes"
	"encoding/csv"
//		"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	//	"testing"
)

func GetBestKeywords(bfile []byte, commonwords_file string,quant int) []string {

	var outarray []string
	get_words_from := func(text string) []string {

		words := regexp.MustCompile("[\\p{L}\\d_]+")

		return words.FindAllString(text, -1)

	}

	fcsv, _ := os.Open(commonwords_file)
	r := csv.NewReader(bufio.NewReader(fcsv))
	set_commonwords := make(map[string]struct{})

	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		set_commonwords[record[0]] = struct{}{}

	}

	s := string(bfile)

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
	for i, k := range a {
		if i >= quant {
			break
		}
		
		for _, s := range n[k] {
//			fmt.Printf("%s, %d\n", s, k)
			outarray =append(outarray,s)			
			
		}
	}



	return outarray
}

