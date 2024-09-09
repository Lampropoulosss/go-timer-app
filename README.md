# Go Timer App

Simple project in Go that creates a window and uses webview to load the Svelte code. It allows you to enter after how much time (in HH:MM:SS) you want to get notified and once the timer reaches zero it will send a desktop notification and play a song.

## How to build the application

Make sure you have [Wails](https://wails.io/) installed on your system and in your PATH variable and follow the instructions:

1) `cd frontend && npm i`
2) `cd ../ && go get`
3) `wails build`

## How to change the song
Head to the `files` directory and paste the audio file you want (in MP3) named "audio.mp3" and build again.
