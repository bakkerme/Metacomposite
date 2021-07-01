package reddit

import (
	"fmt"

	"github.com/turnage/graw/reddit"
	"hyperfocus.systems/metacomposite/v2/types"
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
		return nil, fmt.Errorf("Could not get posts for subreddit %s, error: %s", feed.URI, err)
	}

	posts := []types.Post{}
	for _, post := range redditPosts {
		posts = append(posts, redditPostToPost(post, feed.URI))
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
		return nil, fmt.Errorf("Failed to fetch /r/%s, error: %s", subreddit, err)
	}

	return harvest.Posts, nil
}

func redditPostToPost(post *reddit.Post, subreddit string) types.Post {
	return types.Post{
		Content:     post.SelfText,
		Description: "",
		FeedID:      subreddit,
		ImageURL:    &post.Media.OEmbed.ProviderURL,
		Link:        post.URL,
		Title:       post.Title,
	}
}
