package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 123456
	}
	fmt.Printf("ret:%d\n", ret)

	s, ok := ctx.Value("session").(string)
	if !ok {
		s = "qwer"
	}
	fmt.Printf("s:%s\n", s)
}

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", 1358434)
	ctx = context.WithValue(ctx, "session", "asdfasdfasdfasfasdfe")
	process(ctx)
}
