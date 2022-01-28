package types

// Loader is an interface allowing data associated with a Feed to be loaded and
// output as a Post for the API
type Loader interface {
	LoadPosts(feed *Feed) (*[]Post, error)
}
