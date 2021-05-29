// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package api

// Feed defines model for Feed.
type Feed struct {
	Description *string `json:"description,omitempty"`
	Group       *string `json:"group,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// Group defines model for Group.
type Group struct {
	Feeds *[]Feed `json:"feeds,omitempty"`
	Id    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// Post defines model for Post.
type Post struct {
	Content     *string `json:"content,omitempty"`
	Description *string `json:"description,omitempty"`
	Feed        *Feed   `json:"feed,omitempty"`
	ImageURL    *string `json:"imageURL,omitempty"`
	Link        *string `json:"link,omitempty"`
	Title       *string `json:"title,omitempty"`
}
