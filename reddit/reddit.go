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

	fmt.Println(redditPosts)

	return nil, nil
}

func (rdl Load) loadSubredditPost(subreddit string) ([]*reddit.Post, error) {
	cfg := reddit.BotConfig{
		Agent: "linux:metacomposite:v0.0.1 (by /u/dankweedhacker)",
		App: reddit.App{
			ID:     rdl.ID,
			Secret: rdl.Secret,
			// Username: "yourbotusername",
			// Password: "yourbotspassword",
		},
	}
	bot, _ := reddit.NewBot(cfg)
	harvest, err := bot.Listing("/r/"+subreddit, "")
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch /r/%s, error: %s", subreddit, err)
	}

	// for _, post := range harvest.Posts[:5] {
	// fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	// }

	return harvest.Posts, nil
}

// type mockRedditLoad struct {
// mock.Mock
// }

// func (m *mockRedditLoad) loadPosts(feed *Feed) (*[]Post, error) {
// args := m.Called(feed)
// // return args.
// }
// func (m *mockRedditLoad) loadSubredditPost(subreddit string) (*[]reddit.Post, error) {
// }

// func (m *mockRedditLoad) DoSomething(number int) (bool, error) {

// args := m.Called(number)
// return args.Bool(0), args.Error(1)
// }
