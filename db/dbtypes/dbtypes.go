package dbtypes

// Collection is an abstraction for a group of documents retrived by a key
type Collection interface {
	Get(string) (interface{}, error)
	Save(string, interface{})
}
