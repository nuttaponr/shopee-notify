package handler

import (
	"log"
)

type Caller interface {
	Call() (string, error)
}

type Notifier interface {
	Notify(msg string) error
}

type Handler struct {
	caller   Caller
	notifier Notifier
}

func New(c Caller, n Notifier) *Handler {
	return &Handler{c, n}
}

func (h *Handler) DoIt() {
	log.Printf("do it")
	msg, err := h.caller.Call()
	if err != nil {
		log.Fatal(err.Error())
	}
	if msg == "" {
		return
	}
	if err := h.notifier.Notify(msg); err != nil {
		log.Fatal(err.Error())
	}
}
