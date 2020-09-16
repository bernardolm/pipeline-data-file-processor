package main

import (
	"context"
	"fmt"

	processor "github.com/bernardolm/pipelinedatafilprocessor"
)

var (
	_ processor.ISource = (*mySource)(nil)

	MySource mySource
)

type mySource struct{}

func (s mySource) GetData(ctx context.Context, offset int, limit int) {
	fmt.Println("not implemented")
}

func (s mySource) GetName(_ context.Context) string {
	fmt.Println("not implemented")
	return "not implemented"
}

func (s mySource) PostProcessData(ctx context.Context) {
	fmt.Println("not implemented")
}
