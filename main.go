package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bakkerme/metacomposite/v2/api"
	"github.com/bakkerme/metacomposite/v2/env"
	"github.com/bakkerme/metacomposite/v2/reddit"
	"github.com/bakkerme/metacomposite/v2/rss"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

func main() {
	var spec env.Specification
	err := envconfig.Process("mt", &spec)
	if err != nil {
		panic(err.Error())
	}

	configPath := ""

	if spec.Environment == env.Production {
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
	} else if spec.Environment == env.Local {
		configPath = "./api/config.json"
	}

	cfgProvider := api.FileConfigProvider{}
	cfg, err := cfgProvider.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	cfg.Credentials = []api.Credentials{
		{
			Type: "reddit",
			Values: map[string]string{
				"ID":     spec.RedditId,
				"Secret": spec.RedditSecret,
			},
		},
	}

	mAPI := api.API{
		CFG: cfg,
		Loaders: api.Loaders{
			RSS: rss.Load{},
			Reddit: reddit.Load{
				ID:     spec.RedditId,
				Secret: spec.RedditSecret,
			},
		},
	}
	e := echo.New()
	api.RegisterHandlers(e, &mAPI)
	e.Logger.Fatal(e.Start(":3030"))
}
