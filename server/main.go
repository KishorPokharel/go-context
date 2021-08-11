package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/KishorPokharel/go-context/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(42), int64(100))
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <- ctx.Done():
		log.Println(ctx, ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}