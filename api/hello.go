package helloworld

import (
	"context"
	"log"

	hello "github.com/piotrostr/full-auto-gke/api/gen/hello"
)

// hello service example implementation.
// The example methods log the requests and return zero values.
type hellosrvc struct {
	logger *log.Logger
}

// NewHello returns the hello service implementation.
func NewHello(logger *log.Logger) hello.Service {
	return &hellosrvc{logger}
}

// SayHello implements say-hello.
func (s *hellosrvc) SayHello(ctx context.Context, p *hello.SayHelloPayload) (res string, err error) {
	s.logger.Print("hello.say-hello")
	return "Hello " + p.Name, nil
}
