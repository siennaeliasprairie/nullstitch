package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	appName = "signal-handler-67a964"
	version = "1.6.0"
)

type Handler struct {
	Name    string
	Version string
}

func NewHandler() *Handler {
	return &Handler{Name: appName, Version: version}
}

func (h *Handler) Process(input string) string {
	fmt.Printf("[%s] Processing: %s\n", h.Name, input)
	return fmt.Sprintf("processed_%s_%d", input, time.Now().Unix())
}

func (h *Handler) Run() {
	fmt.Printf("[%s] v%s starting on %s/%s\n", h.Name, h.Version, runtime.GOOS, runtime.GOARCH)
	result := h.Process("default")
	fmt.Printf("[%s] Result: %s\n", h.Name, result)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("%s v%s\n", appName, version)
		return
	}
	handler := NewHandler()
	handler.Run()
}
