package main

// go:generate moq -out=moqs/app.go -pkg=moqs . Application

// Hoge is interface
type Hoge interface{}

// Fuga is struct
type Fuga struct {
	id   int
	name string
}

func Bar(n int) error {
	return nil
}
