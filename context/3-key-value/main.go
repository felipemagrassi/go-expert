package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "foo", "bar")

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	foo := ctx.Value("foo").(string)
	fmt.Println(foo)
}
