package main

import (
	"context"
	"embed"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

//go:embed files/audio.mp3
var file embed.FS

//go:embed files/icon.png
var iconFile embed.FS

// App struct
type App struct {
	ctx       context.Context
	done      chan bool
	tempFiles []string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		tempFiles: []string{},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	a.cleanup()
}

func (a *App) cleanup() {
	for _, file := range a.tempFiles {
		os.Remove(file)
	}
}

func (a *App) extractFile(embedFS embed.FS, path string) (string, error) {
	data, err := embedFS.ReadFile(path)
	if err != nil {
		return "", err
	}

	tmpFile, err := os.CreateTemp("", filepath.Base(path))
	if err != nil {
		return "", err
	}

	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		return "", err
	}

	tmpFile.Close()
	a.tempFiles = append(a.tempFiles, tmpFile.Name())

	return tmpFile.Name(), nil
}

func (a *App) Notify() {
	var iconPath string
	var err error

	if len(a.tempFiles) > 0 {
		iconPath = a.tempFiles[0]
	} else {
		iconPath, err = a.extractFile(iconFile, "files/icon.png")
	}

	if err != nil {
		beeep.Notify("Time is up!", "The timer you set has finished!", "")
	} else {
		beeep.Notify("Time is up!", "The timer you set has finished!", iconPath)
	}

	f, err := file.Open("files/audio.mp3")
	if err != nil {
		log.Fatal(err)
		return
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
		return
	}

	speaker.Init(format.SampleRate, format.SampleRate.N((time.Second / 10)))

	a.done = make(chan bool)

	speaker.Play(streamer)

	<-a.done // Wait until channel closes which is when playback is done

	speaker.Clear()
	streamer.Close()
	f.Close()
}

func (a *App) Stop() {
	if a.done != nil {
		close(a.done)
		a.done = nil
	}
}
