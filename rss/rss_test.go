package rss

import (
	"reflect"
	"testing"

	"github.com/mmcdole/gofeed"
	ext "github.com/mmcdole/gofeed/extensions"
	hfutils "gitlab.com/hyperfocus.systems/hyperfocus-utils"
	"hyperfocus.systems/metacomposite/v2/types"
)

func TestFeedItemToPost(t *testing.T) {
	t.Run("Provides a Post with with an attached Image", func(t *testing.T) {
		item := gofeed.Item{
			Image: &gofeed.Image{
				URL: "https://hyperfocus.systems/someimage.jpg",
			},
			Extensions:  ext.Extensions{},
			Content:     "TestFeed",
			Description: "Test Description",
			Link:        "https://hyperfocus.systems/apost",
			Title:       "A Post Title",
		}

		expected := types.Post{
			Content:     item.Content,
			Description: item.Description,
			FeedID:      "hyperfocusFeed",
			ImageURL:    &item.Image.URL,
			Link:        item.Link,
			Title:       item.Title,
		}

		post := feedItemToPost(&item, expected.FeedID)

		if !reflect.DeepEqual(post, expected) {
			t.Error(hfutils.MismatchError("feedItemToPost", expected, item))
		}
	})

	t.Run("Provides a Post with with a media extension image", func(t *testing.T) {
		imageURL := "https://hyperfocus.systems/someimage.jpg"
		item := gofeed.Item{
			Extensions: ext.Extensions{
				"media": map[string][]ext.Extension{
					"thumbnail": []ext.Extension{
						ext.Extension{
							Attrs: map[string]string{
								"url": imageURL,
							},
						},
					},
				},
			},
			Content:     "TestFeed",
			Description: "Test Description",
			Link:        "https://hyperfocus.systems/apost",
			Title:       "A Post Title",
		}

		expected := types.Post{
			Content:     item.Content,
			Description: item.Description,
			FeedID:      "hyperfocusFeed",
			ImageURL:    &imageURL,
			Link:        item.Link,
			Title:       item.Title,
		}

		post := feedItemToPost(&item, expected.FeedID)

		if !reflect.DeepEqual(post, expected) {
			t.Error(hfutils.MismatchError("feedItemToPost", expected, item))
		}
	})
}
