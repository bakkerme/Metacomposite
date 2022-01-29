package reddit

import (
	"fmt"
	"html"
	"regexp"
	"strings"

	"github.com/bakkerme/metacomposite/v2/types"
	"github.com/turnage/graw/reddit"
)

//Load is an implementation of types.Loader, allowing for loading reddit posts
type Load struct {
	ID     string
	Secret string
}

//LoadPosts loads all Reddit posts associated with a given feed
func (rdl Load) LoadPosts(feed *types.Feed) (*[]types.Post, error) {
	redditPosts, err := rdl.loadSubredditPost(feed.URI)
	if err != nil {
		return nil, fmt.Errorf("could not get posts for subreddit %s, error: %w", feed.URI, err)
	}

	// fmt.Println(valast.String(redditPosts[1:20]))

	posts := []types.Post{}
	for _, redditPost := range redditPosts {
		post := redditPostToPost(redditPost, feed.URI)
		posts = append(posts, post)
	}

	return &posts, nil
}

func (rdl Load) loadSubredditPost(subreddit string) ([]*reddit.Post, error) {
	cfg := reddit.BotConfig{
		Agent: "linux:metacomposite:v0.0.1 (by /u/dankweedhacker)",
		App: reddit.App{
			ID:     rdl.ID,
			Secret: rdl.Secret,
		},
	}
	bot, _ := reddit.NewBot(cfg)
	harvest, err := bot.Listing("/r/"+subreddit, "")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch /r/%s, error: %w", subreddit, err)
	}

	return harvest.Posts, nil
}

func redditPostToPost(post *reddit.Post, subreddit string) types.Post {
	url := post.URL
	isRedditGallery := false
	if isURLRedditGallery(post.URL) {
		url = getRedditGalleryProxiedURL(post.URL)
		isRedditGallery = true
	}

	imageURLs := getImageURLFromPost(post)

	return types.Post{
		ID:              post.ID,
		Title:           post.Title,
		FeedID:          subreddit,
		Content:         post.SelfTextHTML,
		Description:     post.SelfText,
		ImageURLs:       imageURLs,
		Permalink:       "https://reddit.com" + post.Permalink,
		Link:            url,
		Timestamp:       int(post.CreatedUTC),
		IsRedditGallery: isRedditGallery,
	}
}

func getImageURLFromPost(post *reddit.Post) []string {
	if post.IsSelf {
		if strings.Contains(post.SelfTextHTML, "https://preview.redd.it") {
			previewRegex := regexp.MustCompile(`(?m)\"(https://preview\.redd\.it/.+)\"`)

			previewURLResult := previewRegex.FindStringSubmatch(post.SelfTextHTML)
			if len(previewURLResult) > 1 {
				unescaped := html.UnescapeString(previewURLResult[1])
				return []string{unescaped}
			}

			return []string{}
		} else {
			return []string{}
		}
	}

	if post.IsRedditMediaDomain || strings.Contains(post.URL, "i.imgur.com") {
		return []string{post.URL}
	} else {
		if isThumbnailValid(post.Thumbnail) {
			return []string{post.Thumbnail}
		}
	}

	return []string{}
}

func isURLRedditGallery(url string) bool {
	return strings.Contains(url, "reddit.com/gallery")
}

func getRedditGalleryProxiedURL(url string) string {
	urlComponents := strings.Split(url, "/")
	return "/redditgallery/" + urlComponents[len(urlComponents)-1]
}

func isThumbnailValid(thumbnail string) bool {
	return thumbnail != "self" && thumbnail != "default" && thumbnail != "" && thumbnail != "nsfw"
}
