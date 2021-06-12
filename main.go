package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

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

	homepath := os.Getenv("HOME")

	cfgProvider := FileConfigProvider{}
	cfg, err := cfgProvider.LoadConfig(homepath + "/.config/metacomposite/config.json")
	if err != nil {
		panic(err)
	}

	api := API{cfg: cfg}
	e := echo.New()
	RegisterHandlers(e, &api)
	e.Logger.Fatal(e.Start(":3030"))
}
