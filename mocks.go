package main

func getTestFeedData() *[]Feed {
	return &[]Feed{
		Feed{
			ID:          "testfeed",
			URI:         "https://hyperfocus.systems/testfeed",
			Description: "Test Description",
			GroupID:     []string{"hyperfocus", "test"},
			Name:        "Test Feed",
			Type:        feedTypeRSS,
		},
		Feed{
			ID:          "testfeed2",
			URI:         "https://hyperfocus.systems/testfeed2",
			Description: "Test Description 2",
			GroupID:     []string{"hyperfocus", "test", "2"},
			Name:        "Test Feed 2",
			Type:        feedTypeRSS,
		},
		Feed{
			ID:          "testfeed3",
			URI:         "https://hyperfocus.systems/testfeed3",
			Description: "Test Description 3",
			GroupID:     []string{"hyperfocus", "test", "3"},
			Name:        "Test Feed 3",
			Type:        feedTypeRSS,
		},
	}
}

func getTestGroupData() *[]Group {
	return &[]Group{
		Group{
			ID:   "hyperfocus",
			Name: "Hyperfocus",
		},
		Group{
			ID:   "test",
			Name: "Test",
		},
		Group{
			ID:   "2",
			Name: "2",
		},
		Group{
			ID:   "3",
			Name: "3",
		},
	}
}
