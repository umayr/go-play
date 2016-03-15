package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/jesusislam/tldr"
	"github.com/urandom/text-summary/summarize"
	"strings"
	"sync"
)

type Article struct {
	Url     string
	Title   string
	Content string
}

type Result struct {
	LexRank string `json:"lexrank"`
	Natural string `json:"natural"`
}

const (
	MaxSentences = 5
	IdealWordCount = 75
)

func (a *Article) Summary() *Result {
	var (
		cl = make(chan string)
		cs = make(chan string)
	)

	go LexRank(cl, a.Content);
	go Natural(cs, a.Title, a.Content)

	r := &Result{}
	r.LexRank = <-cl
	r.Natural = <-cs

	return r;
}

// LexRank Algorithm
func LexRank(c chan string, content string) {
	t := tldr.New()
	td, _ := t.Summarize(content, MaxSentences)
	c <- td;
}

// Natural Language Processing
func Natural(c chan string, title string, content string) {
	s := summarize.NewFromString(title, content)
	s.IdealWordCount = IdealWordCount
	c <- strings.Join(s.KeyPoints(), "")
}

func try(e error) {
	if e != nil {
		panic(e)
	}
}

func Parse(path string) []Article {
	var (
		raw []byte
		err error
	)
	raw, err = ioutil.ReadFile(path)
	try(err)

	j := []Article{}
	err = json.Unmarshal(raw, &j)
	try(err)

	return j
}

func main() {
	src := Parse("./source.json")
	var wg sync.WaitGroup
	wg.Add(len(src))
	for _, s := range src {
		go func() {
			defer wg.Done()
			r := s.Summary()
			result, err := json.Marshal(r)
			try(err)

			fmt.Printf(strings.Replace(string(result), "\\n", "", -1))
		}()
	}
	wg.Wait()
}
