package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// feedErrors is a struct used to hold Posts and any errors found
// during the fetch process
type feedErrors struct {
	Posts *[]Post
	Err   Error
}

// API provides an implementation of the API API
type API struct {
	cfg *Config
}

// GetFeeds returns all available feeds
func (api *API) GetFeeds(ctx echo.Context) error {
	resp, err := json.Marshal(api.cfg.Feeds)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(resp))
}

type getFeedsPostsResponse struct {
	posts  []Post
	errors []Error
}

// GetFeedsPosts gets all posts from all feeds
func (api *API) GetFeedsPosts(ctx echo.Context) error {
	feeds := api.cfg.Feeds

	ch := make(chan feedErrors)
	for _, feed := range feeds {
		go func() {
			posts, err := getPostsForFeed(rssLoad{}, &feed)

			var errorReturn Error
			if err != nil {
				ids := []string{feed.ID}
				errMessage := err.Error()
				errorReturn = Error{
					Code:       &errorFeedLoadFail,
					Message:    &errMessage,
					RelatedIDs: &ids,
				}
			}

			ch <- feedErrors{
				Posts: posts,
				Err:   errorReturn,
			}
		}()
	}

	posts := []Post{}
	errors := []Error{}
	for range feeds {
		var out feedErrors
		out = <-ch
		if out.Posts == nil {
			errors = append(errors, out.Err)
		} else {
			posts = append(posts, *out.Posts...)
		}
	}

	resp, err := json.Marshal(getFeedsPostsResponse{
		posts,
		errors,
	})
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(resp))
}

// GetFeedsFeedID returns a feed by it's feedID
func (api *API) GetFeedsFeedID(ctx echo.Context, feedID string) error {
	feed := getFeedByID(feedID, &api.cfg.Feeds)
	if feed == nil {
		return ctx.String(http.StatusNotFound, "Could not find "+feedID)
	}

	resp, err := json.Marshal(feed)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(resp))
}

// GetFeedsFeedIDPosts returns all posts associated with a given feed ID
func (api *API) GetFeedsFeedIDPosts(ctx echo.Context, feedID string) error {
	feed := getFeedByID(feedID, &api.cfg.Feeds)
	posts, err := getPostsForFeed(rssLoad{}, feed)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(posts)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(resp))
}

// GetGroupGroupID returns a group by it's groupID
func (api *API) GetGroupGroupID(ctx echo.Context, groupID string) error {
	group := getGroupByID(groupID, &api.cfg.Groups)
	if group == nil {
		return ctx.String(http.StatusNotFound, "Could not find "+groupID)
	}

	resp, err := json.Marshal(group)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(resp))
}

// GetGroups returns all groups
func (api *API) GetGroups(ctx echo.Context) error {
	resp, err := json.Marshal(api.cfg.Groups)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(resp))
}

// GetGroupsGroupIDFeeds returns a list of feeds associated with a given group ID
func (api *API) GetGroupsGroupIDFeeds(ctx echo.Context, groupID string) error {
	feeds := getFeedsForGroupID(groupID, &api.cfg.Feeds)
	if feeds == nil {
		return ctx.String(http.StatusNotFound, "No feeds are available")
	}

	resp, err := json.Marshal(feeds)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(resp))
}

type getGroupGroupIDPostsResponse struct {
	group  Group
	posts  []Post
	errors []Error
}

// GetGroupGroupIDPosts returns a list of posts associated with a given group
func (api *API) GetGroupGroupIDPosts(ctx echo.Context, groupID string) error {
	group := getGroupByID(groupID, &api.cfg.Groups)
	if group == nil {
		return ctx.String(http.StatusNotFound, fmt.Sprintf("Could not find group with %s", groupID))
	}

	feeds := getFeedsForGroupID(groupID, &api.cfg.Feeds)
	if feeds == nil {
		return ctx.String(http.StatusOK, "[]")

	}

	ch := make(chan feedErrors)
	for _, feed := range *feeds {
		go func() {
			posts, err := getPostsForFeed(rssLoad{}, &feed)

			var errorReturn Error
			if err != nil {
				ids := []string{feed.ID}
				errMessage := err.Error()
				errorReturn = Error{
					Code:       &errorFeedLoadFail,
					Message:    &errMessage,
					RelatedIDs: &ids,
				}
			}

			ch <- feedErrors{
				Posts: posts,
				Err:   errorReturn,
			}
		}()
	}

	posts := []Post{}
	errors := []Error{}
	for range *feeds {
		var out feedErrors
		out = <-ch
		posts = append(posts, *out.Posts...)
		errors = append(errors, out.Err)
	}

	resp, err := json.Marshal(getGroupGroupIDPostsResponse{
		*group,
		posts,
		errors,
	})
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(resp))
}
