package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/log"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	app := NewApp()

	llm, err := ollama.New(
		ollama.WithModel(app.conf.Model),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	app.banner()

	resp, err := llm.GenerateContent(ctx, app.history)
	if err != nil {
		app.conf.Logger.Fatal(err)
	}

	out, err := glamour.Render(resp.Choices[0].Content, "dark")
	if err != nil {
		app.conf.Logger.Fatal(err)
	}

	fmt.Print(out)
}
