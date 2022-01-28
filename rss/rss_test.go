package rss

import (
	"reflect"
	"testing"

	hfutils "github.com/bakkerme/hyperfocus-utils"
	"github.com/bakkerme/metacomposite/v2/types"
	"github.com/mmcdole/gofeed"
	ext "github.com/mmcdole/gofeed/extensions"
)

func TestFeedItemToPost(t *testing.T) {
	t.Run("Provides a Post with with an attached Image", func(t *testing.T) {
		item := gofeed.Item{
			Image: &gofeed.Image{
				URL: "https://github.com/bakkerme/someimage.jpg",
			},
			Extensions:  ext.Extensions{},
			Content:     "TestFeed",
			Description: "Test Description",
			Link:        "https://github.com/bakkerme/apost",
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
		imageURL := "https://github.com/bakkerme/someimage.jpg"
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
			Link:        "https://github.com/bakkerme/apost",
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
