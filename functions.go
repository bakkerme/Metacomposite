package main

func getFeedByID(feedID string, feeds *[]Feed) *Feed {
	if feeds == nil {
		return nil
	}

	for _, feed := range *feeds {
		if feed.ID == feedID {
			return &feed
		}
	}

	return nil
}

func getGroupByID(groupID string, groups *[]Group) *Group {
	if groups == nil {
		return nil
	}

	for _, group := range *groups {
		if group.ID == groupID {
			return &group
		}
	}

	return nil
}

func getFeedsForGroupID(groupID string, feeds *[]Feed) *[]Feed {
	if feeds == nil {
		return nil
	}

	matchingFeeds := []Feed{}
	for _, feed := range *feeds {
		for _, feedGroupID := range feed.GroupID {
			if feedGroupID == groupID {
				matchingFeeds = append(matchingFeeds, feed)
			}
		}
	}

	return &matchingFeeds
}

func getPostsForFeed(rsl rssLoader, feed *Feed) (*[]Post, error) {
	rssFeed, err := rsl.loadRSS(feed.URI)

	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, item := range rssFeed.Items {
		var imageURL *string
		if item.Image != nil {
			imageURL = &item.Image.URL
		}
		if item.Extensions["media"] != nil {
			u := item.Extensions["media"]["thumbnail"][0].Attrs["url"]
			imageURL = &u
		}

		posts = append(posts, Post{
			Content:     item.Content,
			Description: item.Description,
			FeedID:      feed.ID,
			ImageURL:    imageURL,
			Link:        item.Link,
			Title:       item.Title,
		})
	}

	return &posts, nil
}
