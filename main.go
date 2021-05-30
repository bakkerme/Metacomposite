package main

import "github.com/labstack/echo/v4"

var feeds = map[string]string{
	"retrobattlestations": "https://www.reddit.com/r/retrobattlestations/.rss",
}

func main() {
	// fp := gofeed.NewParser()
	// feed, _ := fp.ParseURL(feeds["retrobattlestations"])
	// fmt.Println(feed.Description)

	// for _, post := range feed.Items {
	// fmt.Println(post.Title)
	// fmt.Println(post.Content)
	// fmt.Println(post.Image)
	// }

	// cfg, err := config.GetConfig(&config.FileConfigProvider{})
	// if err != nil {
	// panic(err)
	// }

	api := API{}
	e := echo.New()
	RegisterHandlers(e, &api)
	e.Logger.Fatal(e.Start(":3030"))
}
