package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

var feeds = map[string]string{
	"retrobattlestations": "https://www.reddit.com/r/retrobattlestations/.rss",
}

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feeds["retrobattlestations"])
	fmt.Println(feed.Description)

	// for _, post := range feed.Items {
	// fmt.Println(post.Title)
	// fmt.Println(post.Content)
	// fmt.Println(post.Image)
	// }
}
