package service

import "context"

type ExampleCommand struct {
}

type ExampleHandler struct {
}

func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{}
}

func (s *ExampleHandler) Handle(ctx context.Context, command ExampleCommand) (string, error) {
	return "hello world", nil
}
