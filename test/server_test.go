package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"

	"github.com/SuperNerdIT/telepathic/cmd/server"
)

type serverCtxKey struct{}

type statusCodeCtxKey struct {}

func thisReturnsCode(ctx context.Context, expectedCode int) (error) {
	code, ok := ctx.Value(statusCodeCtxKey{}).(int)
	if !ok {
		return  errors.New("Cannot get status code of context")
	}
	if code != expectedCode {
		return fmt.Errorf("expected %d status code, but there is %d", expectedCode, code)
	}

	return nil
}

func theMainServer(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx,serverCtxKey{}, httptest.NewServer(server.NewServer().Handler) ), nil
}

func iCallHealthEndpoint(ctx context.Context) (context.Context ,error) {
	server, ok := ctx.Value(serverCtxKey{}).(*httptest.Server)
	if !ok { 
		return ctx, errors.New("Unable to get server instance")
	}
	resp, err := http.DefaultClient.Do(func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			e := errors.New("Unable to instantiate a request")
			panic(e)
		}
		return r
	}("GET", server.URL + "/health", nil))

	if err != nil {
		return ctx, err
	}
	defer resp.Body.Close()


	return context.WithValue(ctx, statusCodeCtxKey{}, resp.StatusCode), nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the main server$`, theMainServer)
	ctx.Step(`^I call \/health endpoint$`, iCallHealthEndpoint)
	ctx.Step(`^this returns (\d+) code$`, thisReturnsCode)
}
