package api

import (
	"fmt"

	"github.com/bakkerme/metacomposite/v2/types"
	"github.com/labstack/echo/v4"
)

func getFeedByID(feedID string, feeds *[]types.Feed) *types.Feed {
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

func getGroupByID(groupID string, groups *[]types.Group) *types.Group {
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

func getFeedsForGroupID(groupID string, feeds *[]types.Feed) []types.Feed {
	matchingFeeds := []types.Feed{}
	if feeds == nil {
		return matchingFeeds
	}

	for _, feed := range *feeds {
		for _, feedGroupID := range feed.GroupID {
			if feedGroupID == groupID {
				matchingFeeds = append(matchingFeeds, feed)
			}
		}
	}

	return matchingFeeds
}

func getPostsForFeed(lds Loaders, feed *types.Feed) (*[]types.Post, error) {
	switch feed.Type {
	case RSS:
		return lds.RSS.LoadPosts(feed)
	case Reddit:
		return lds.Reddit.LoadPosts(feed)
	}

	return nil, fmt.Errorf("%s is not implemented as a feed type", feed.Type)
}

func getHostFromEchoContext(ctx echo.Context) string {
	return fmt.Sprintf("%s://%s", ctx.Request().URL.Scheme, ctx.Request().Host)
}
