// Package main provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package main

// Feed defines model for Feed.
type Feed struct {
	ID          string   `json:"ID"`
	URI         string   `json:"URI"`
	Description string   `json:"description"`
	GroupID     []string `json:"groupID"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
}

// Group defines model for Group.
type Group struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

// Post defines model for Post.
type Post struct {
	Content     string `json:"content"`
	Description string `json:"description"`
	FeedID      string `json:"feedID"`
	ImageURL    string `json:"imageURL"`
	Link        string `json:"link"`
	Title       string `json:"title"`
}
