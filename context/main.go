package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "my-key", "Hola mundo")

	viewContext(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go myProcess(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Time expired")
		fmt.Println(ctx.Err())
	}
}

func viewContext(ctx context.Context) {
	fmt.Printf("My values is: %s\n", ctx.Value("my-key"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out")
			return
		default:
			fmt.Println("Doing some process")
		}

		time.Sleep(500 * time.Millisecond)
	}
}
