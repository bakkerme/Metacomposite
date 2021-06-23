package main

import (
	"github.com/mmcdole/gofeed"
)

type rssLoader interface {
	loadRSS(uri string) (*gofeed.Feed, error)
}

type rssLoad struct{}

func (rsl rssLoad) loadRSS(uri string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	return fp.ParseURL(uri)
}
