package flow

import (
	"context"
	"time"
)

func New() {}

var _ context.Context = (*Context)(nil)

type Context struct {
	Ctx context.Context
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Ctx.Deadline()
}
func (c *Context) Done() <-chan struct{} {
	return c.Ctx.Done()
}
func (c *Context) Err() error {
	return c.Ctx.Err()
}
func (c *Context) Value(key any) any {
	return c.Ctx.Value(key)
}
func (c *Context) GetOutputByIndex(index int, SecondaryIndex... int) any {
	return nil
}
func (c *Context) GetLatestOutput(SecondaryIndex... int) any {
	return nil
}

type FlowCaller func(c *Context) error

type Flow struct {
	F []FlowCaller
	Output []any
}

func (f *Flow) Do(ctx context.Context) error {
	return nil
}

type Config struct {
	AllowFailure bool
}
