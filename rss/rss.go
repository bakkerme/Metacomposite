package rss

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bakkerme/metacomposite/v2/types"
	"github.com/mmcdole/gofeed"
)

// Load is an implementation of Loader, for loading RSS feed posts
type Load struct{}

// LoadPosts loads all posts associated with a given feed
func (rsl Load) LoadPosts(feed *types.Feed) (*[]types.Post, error) {
	rssFeed, err := rsl.LoadRSS(feed.URI)
	if err != nil {
		return nil, fmt.Errorf("Could not load posts from %s, error: %s", feed.URI, err)
	}

	posts := []types.Post{}
	for _, item := range rssFeed.Items {
		posts = append(posts, feedItemToPost(item, feed.ID))
	}

	return &posts, nil
}

// LoadRSS specifically loads RSS feeds given a URI and outputs the internal gofeed.Feeds
func (rsl Load) LoadRSS(uri string) (*gofeed.Feed, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("user-agent", "linux:metacomposite:v0.0.1")
	req.Header.Set("accept", "application/xml")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Did not get 200. %d %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fp := gofeed.NewParser()
	rss, err := fp.ParseString(string(body))
	if err != nil {
		return nil, err
	}

	return rss, nil
}

func feedItemToPost(item *gofeed.Item, feedID string) types.Post {
	imageURL := []string{}
	if item.Image != nil {
		imageURL = []string{item.Image.URL}
	}
	if item.Extensions["media"] != nil {
		u := item.Extensions["media"]["thumbnail"][0].Attrs["url"]
		imageURL = []string{u}
	}

	return types.Post{
		Content:     item.Content,
		Description: item.Description,
		FeedID:      feedID,
		ImageURLs:   imageURL,
		Link:        item.Link,
		Title:       item.Title,
	}
}
