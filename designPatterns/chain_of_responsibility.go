package designpatterns

import "fmt"

type Handler interface {
	SetNext(Handler)
	Handle(msg string)
}

type AbstractHandler struct {
	nextHandler Handler
}

func (h *AbstractHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *AbstractHandler) Handle(msg string) {
	if h.nextHandler != nil {
		h.nextHandler.Handle(msg)
	}
}

type ConsoleLogger struct {
	AbstractHandler
}

func (l *ConsoleLogger) Handle(msg string) {
	fmt.Println("Console Logger: ", msg)
	l.AbstractHandler.Handle(msg)
}

type FileLogger struct {
	AbstractHandler
}

func (l *FileLogger) Handle(msg string) {
	fmt.Println("File Logger: ", msg)
	l.AbstractHandler.Handle(msg)
}

func CORTest() {
	consoleLogger := &ConsoleLogger{}
	fileLogger := &FileLogger{}
	consoleLogger.SetNext(fileLogger)
	consoleLogger.Handle("shimakaze")
}
