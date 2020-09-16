package main

import (
	"context"
	"fmt"

	processor "github.com/bernardolm/pipelinedatafilprocessor"
)

var (
	_ processor.IOutput = (*myOutput)(nil)

	MyOutput myOutput
)

type myOutput struct{}

func (s myOutput) Build(ctx context.Context) {
	fmt.Println("not implemented")
}

func (s myOutput) BuildBottomContent(ctx context.Context) {
	fmt.Println("not implemented")
}

func (s myOutput) BuildTopContent(ctx context.Context) {
	fmt.Println("not implemented")
}
