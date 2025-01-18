package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/l4m3r/serverless-guard/app/apigw"
)

func LocalHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	_, err = TgHandler(context.Background(), &apigw.APIGatewayRequest{
		HTTPMethod: r.Method,
		Body:       string(body),
	})
	if err != nil {
		fmt.Printf("Decode error: %e", err)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/handler", LocalHandler)
	fmt.Print("Listening")
	http.ListenAndServe(":8080", nil)
}
