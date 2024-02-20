package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

type contextKey string

const keyServerAddr contextKey = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	fmt.Printf("%s: got / request first(%t)=%s, second(%t)=%s\n", ctx.Value(keyServerAddr),
		hasFirst, first,
		hasSecond, second)
	_, err := io.WriteString(w, "This is my website!\n")
	if err != nil {
		fmt.Printf("Error writing response: %s\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	_, err := io.WriteString(w, "Hello, HTTP!\n")
	if err != nil {
		fmt.Printf("Error writing response: %s\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx  := context.Background()
	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed \n")
		} else if err != nil {
			fmt.Printf("error listening for server one  %s\n", err)
		}

	fmt.Println("Starting server")

}
