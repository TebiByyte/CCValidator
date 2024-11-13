package main

import (
	"ccvalidator/validate"
	"context"
	"fmt"
	"net"
	"net/http"
)

func main() {
    handler := http.NewServeMux()
    handler.HandleFunc("/creditcard/validate", validate.HandleValidateRequest)

    server := http.Server {
        Addr: ":8080",
        BaseContext: func(l net.Listener) context.Context {
            return context.Background()
        }, 
        Handler: handler,
    }

    err := server.ListenAndServe()
    if err != nil {
        fmt.Printf("err: %v\n", err)
    }

}
