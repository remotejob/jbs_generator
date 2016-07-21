package mgenerator

import (
	"bytes"
	"github.com/remotejob/comutils/gen"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Write(b []byte) {
	br := bytes.NewReader(b)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]

		words = append(words, next)

		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func Generate( book []byte) string{

	rand.Seed(time.Now().UnixNano())

	prefixLen := 1

	// Seed the random number generator.
	numWords := gen.Random(500, 1500)

	c := NewChain(prefixLen)

	c.Write(book)
	text := c.Generate(numWords)

//	fmt.Println(text)
	return text

}
