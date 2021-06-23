package main

import (
	"reflect"
	"testing"

	hfutils "gitlab.com/hyperfocus.systems/hyperfocus-utils"
)

var someFakeID = "fakeID"

func TestGetFeedByID(t *testing.T) {
	feedTestData := *getTestFeedData()

	t.Run("returns nil if no feeds are provided", func(t *testing.T) {
		feed := getFeedByID(someFakeID, &[]Feed{})
		if feed != nil {
			hfutils.MismatchError("getFeedByID", nil, feed)
		}
	})

	t.Run("returns nil if no feeds match is found", func(t *testing.T) {
		feed := getFeedByID(someFakeID, &feedTestData)
		if feed != nil {
			hfutils.MismatchError("getFeedByID", nil, feed)
		}
	})

	t.Run("returns nil if no nil is provided as a feed match is found", func(t *testing.T) {
		feed := getFeedByID(someFakeID, nil)
		if feed != nil {
			hfutils.MismatchError("getFeedByID", nil, feed)
		}
	})

	t.Run("returns matching feed", func(t *testing.T) {
		expected := feedTestData[0]

		feed := getFeedByID("testfeed", &feedTestData)
		if !reflect.DeepEqual(*feed, expected) {
			hfutils.MismatchError("getFeedByID", expected, *feed)
		}
	})
}

func TestGetGroupByID(t *testing.T) {
	groupTestData := *getTestGroupData()

	t.Run("returns nil if no groups are provided", func(t *testing.T) {
		group := getGroupByID(someFakeID, &[]Group{})
		if group != nil {
			hfutils.MismatchError("getGroupByID", nil, group)
		}
	})

	t.Run("returns nil if no groups match is found", func(t *testing.T) {
		group := getGroupByID(someFakeID, &groupTestData)
		if group != nil {
			hfutils.MismatchError("getGroupByID", nil, group)
		}
	})

	t.Run("returns nil if no nil is provided as a group match is found", func(t *testing.T) {
		group := getGroupByID(someFakeID, nil)
		if group != nil {
			hfutils.MismatchError("getGroupByID", nil, group)
		}
	})

	t.Run("returns matching group", func(t *testing.T) {
		expected := groupTestData[0]

		group := getGroupByID("hyperfocus", &groupTestData)
		if !reflect.DeepEqual(*group, expected) {
			hfutils.MismatchError("getGroupByID", expected, *group)
		}
	})
}

func TestGetFeedsForGroupID(t *testing.T) {
	feedTestData := *getTestFeedData()

	t.Run("returns nil if no feeds are provided", func(t *testing.T) {
		feeds := getFeedsForGroupID(someFakeID, &[]Feed{})
		if feeds != nil {
			hfutils.MismatchError("getFeedsForGroupID", nil, feeds)
		}
	})

	t.Run("returns nil if no feeds match is found", func(t *testing.T) {
		feed := getFeedsForGroupID(someFakeID, &feedTestData)
		if feed != nil {
			hfutils.MismatchError("getFeedsForGroupID", nil, feed)
		}
	})

	t.Run("returns nil if no nil is provided as a feed match is found", func(t *testing.T) {
		feed := getFeedsForGroupID(someFakeID, nil)
		if feed != nil {
			hfutils.MismatchError("getFeedsForGroupID", nil, feed)
		}
	})

	t.Run("returns all matching feeds", func(t *testing.T) {
		feeds := getFeedsForGroupID("hyperfocus", &feedTestData)
		if !reflect.DeepEqual(*feeds, feedTestData) {
			hfutils.MismatchError("getFeedsForGroupID", feedTestData, *feeds)
		}
	})

}
