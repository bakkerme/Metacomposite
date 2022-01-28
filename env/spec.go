package env

type Environment string

var Local Environment = "local"
var Production Environment = "production"

type Specification struct {
	RedditId     string      `split_words:"true" required:"true"`
	RedditSecret string      `split_words:"true" required:"true"`
	Environment  Environment `split_words:"true" required:"true"`
}
