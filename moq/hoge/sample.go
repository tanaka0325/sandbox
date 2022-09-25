package hoge

//go:generate moq -out=moqs/sample.go -pkg=moqs . Hoge

import "context"

type Ho struct {
	id   string
	name string
}

type Hoge interface {
	Get(ctx context.Context, id string) (Ho, error)
	Update(ctx context.Context, ho Ho) (Ho, error)
}
