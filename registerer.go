package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {
	fmt.Println("batata here")
	// check the passed configuration and initialize the plugin
	name, ok := extra["name"].([]interface{})
	if !ok {
		return nil, errors.New("wrong config")
	}
	if v, ok := name[2].(string); !ok || v != string(r) {
		return nil, fmt.Errorf("unknown register %s", name)
	}

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(readJSON("./internal/handler/docs_index.html")))

		return
	}), nil
}

func readJSON(filePath string) string {
	c, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read content data from %s, err: %v", filePath, err)
	}

	return string(c)
}
