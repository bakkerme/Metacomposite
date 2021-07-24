package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/labstack/echo/v4"
	utils "gitlab.com/hyperfocus.systems/hyperfocus-utils"
	"hyperfocus.systems/metacomposite/v2/api"
	"hyperfocus.systems/metacomposite/v2/reddit"
	"hyperfocus.systems/metacomposite/v2/rss"
)

func main() {
	configPath := ""

	currOS := runtime.GOOS
	if currOS == "windows" {
		appdataPath := os.Getenv("appdata")
		configPath = appdataPath + "/Metacomposite/config.json"
	} else if currOS == "linux" {
		homepath := os.Getenv("HOME")
		configPath = homepath + "/.config/metacomposite/config.json"
	} else {
		panic(fmt.Sprintf("OS %s not yet supported", currOS))
	}

	cfgProvider := api.FileConfigProvider{}
	cfg, err := cfgProvider.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	envReader := utils.EnvRead{}
	redditID, found := envReader.LookupEnv("REDDIT_ID")
	if !found {
		panic("Cannot find REDDIT_ID envar")
	}

	redditSecret, found := envReader.LookupEnv("REDDIT_SECRET")
	if !found {
		panic("Cannot find REDDIT_SECRET envar")
	}

	cfg.Credentials = []api.Credentials{
		api.Credentials{
			Type: "reddit",
			Values: map[string]string{
				"ID":     redditID,
				"Secret": redditSecret,
			},
		},
	}

	mAPI := api.API{
		CFG: cfg,
		Loaders: api.Loaders{
			RSS: rss.Load{},
			Reddit: reddit.Load{
				ID:     redditID,
				Secret: redditSecret,
			},
		},
	}
	e := echo.New()
	api.RegisterHandlers(e, &mAPI)
	e.Logger.Fatal(e.Start(":3030"))
}
