package main

import (
	"flag"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
)

type App struct {
	conf      *Config
	imageFile string
	imageData []byte
	userQuery string
	history   []llms.MessageContent
}

func NewApp() *App {
	conf := LoadConfig()

	app := &App{
		conf: conf,
	}

	flag.StringVar(&app.imageFile, "image", "", "Filename of the image.")
	flag.StringVar(&app.userQuery, "query", "What is in this image? Use markdown format for your response.", "Query about the image.")
	flag.Parse()

	if app.imageFile == "" {
		flag.Usage()

		conf.Logger.Fatalf("image parameter is required.")
	}

	var err error

	app.imageData, err = os.ReadFile(app.imageFile)
	if err != nil {
		conf.Logger.Fatalf("failed to load image %q: %v", app.imageFile, err)
	}

	app.push(llms.MessageContent{
		Role: schema.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{
			llms.BinaryPart("image/png", app.imageData),
			llms.TextPart(app.userQuery),
		},
	})

	return app
}

func (a *App) banner() {
	a.conf.Logger.Infof("Model: %v", a.conf.Model)
	a.conf.Logger.Infof("Query: %v", a.userQuery)
}

func (a *App) push(msg llms.MessageContent) {
	a.history = append(a.history, msg)
}
