package api

import (
	"hyperfocus.systems/metacomposite/v2/types"
)

func getTestFeedData() *[]types.Feed {
	return &[]types.Feed{
		types.Feed{
			ID:          "testfeed",
			URI:         "https://hyperfocus.systems/testfeed",
			Description: "Test Description",
			GroupID:     []string{"hyperfocus", "test"},
			Name:        "Test Feed",
			Type:        RSS,
		},
		types.Feed{
			ID:          "testfeed2",
			URI:         "https://hyperfocus.systems/testfeed2",
			Description: "Test Description 2",
			GroupID:     []string{"hyperfocus", "test", "2"},
			Name:        "Test Feed 2",
			Type:        RSS,
		},
		types.Feed{
			ID:          "testfeed3",
			URI:         "https://hyperfocus.systems/testfeed3",
			Description: "Test Description 3",
			GroupID:     []string{"hyperfocus", "test", "3"},
			Name:        "Test Feed 3",
			Type:        RSS,
		},
	}
}

func getTestGroupData() *[]types.Group {
	return &[]types.Group{
		types.Group{
			ID:   "hyperfocus",
			Name: "Hyperfocus",
		},
		types.Group{
			ID:   "test",
			Name: "Test",
		},
		types.Group{
			ID:   "2",
			Name: "2",
		},
		types.Group{
			ID:   "3",
			Name: "3",
		},
	}
}
