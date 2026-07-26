package main

import (
	"embed"

	"log"
	"time"

	"app/local_store"
	"app/provider_service"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	application.RegisterEvent[string]("time")
}

func main() {

	store, err := local_store.Open()
	if err != nil {
		log.Fatalf("Failed to initialize local store: %v", err)
	}
	defer store.Close()

	app := application.New(application.Options{
		Name:        "hmm",
		Description: "humans messaging machines",
		Services: []application.Service{
			application.NewService(provider_service.New(store)),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "hmm",
		Width:  1000,
		Height: 618,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
	})

	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
