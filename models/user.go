package models

type User struct {
	Name string
	Age  int
}

// is this won't be exported due to the lowercase 'u'
type user struct {
	Name string
	Age  int
}
