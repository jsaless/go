package main

import "fmt"

// Handlers Interface
type Handler interface {
	setNext(h Handler) Handler
	handle(s Request) Request
}

// BaseHandler Implementation
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) setNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) passToNext(s Request) Request {
	if h.next != nil {
		return h.next.handle(s)
	}
	return s
}

// Request Struct Implementation
type Request struct {
	name             string
	checkedByOne     bool
	checkedByTwo     bool
	checkedByDefault bool
}

// ConcreteHandlerOne Implementation
type ConcreteHandlerOne struct {
	BaseHandler
}

func (h *ConcreteHandlerOne) handle(s Request) Request {
	if s.checkedByOne {
		fmt.Println("Processed request A")
		h.passToNext(s)
	}
	fmt.Println("Processing request A")
	s.checkedByOne = true
	return h.passToNext(s)
}

// ConcreteHandlerTwo Implementation
type ConcreteHandlerTwo struct {
	BaseHandler
}

func (h *ConcreteHandlerTwo) handle(s Request) Request {
	if s.checkedByTwo {
		fmt.Println("Processed request B")
		h.passToNext(s)
	}
	fmt.Println("Processing request B")
	s.checkedByTwo = true
	return h.passToNext(s)
}

// ConcreteHandlerDefault Implementation
type ConcreteHandlerDefault struct {
	BaseHandler
}

func (h ConcreteHandlerDefault) handle(s Request) Request {
	if s.checkedByDefault {
		fmt.Println("Processed by default handler")
		h.passToNext(s)
	}
	fmt.Println("Processing by default handler")
	s.checkedByDefault = true
	return h.passToNext(s)
}

func main() {

	defaultHandler := &ConcreteHandlerDefault{}

	concreteHandlerTwo := &ConcreteHandlerTwo{}

	concreteHandlerOne := &ConcreteHandlerOne{}

	concreteHandlerOne.setNext(concreteHandlerTwo).setNext(defaultHandler)

	request := Request{name: "abc"}

	concreteHandlerOne.handle(request)
}
