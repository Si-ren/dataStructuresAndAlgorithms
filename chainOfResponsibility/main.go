package main

import (
	"fmt"
	"math"
)

type Context struct {
	handlers HandlerChain
	index    int8
}

type HandlerFunc func(*Context)
type HandlerChain []HandlerFunc

func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Abort() {
	c.index = math.MaxInt8 >> 1
}
func (c *Context) Add(handler HandlerFunc) *Context {
	c.handlers = append(c.handlers, handler)
	return c
}

func Middleware01(c *Context) {
	fmt.Println("begin Middleware01 ", c.index)
	c.Next()

	fmt.Println("end Middleware01", c.index)
}

func Middleware02(c *Context) {
	fmt.Println("begin Middleware02 ", c.index)

	c.Next()

	fmt.Println("end Middleware02 ", c.index)
}

func Middleware03(c *Context) {
	fmt.Println("begin Middleware03 ", c.index)
	c.Next()
	fmt.Println("end Middleware03 ", c.index)
}

func main() {
	c := Context{index: -1}
	c.Add(Middleware01).Add(Middleware02).Add(Middleware03)
	c.Next()
}
