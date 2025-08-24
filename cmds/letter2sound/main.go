package main

import (
	"log"

	"galuma.net/synthetic/sound"
)

func program() error {
	//phrase := "Salut Martin, le petit lapin"
	phrase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz"
	//phrase := "Martin Guillaume Anne-Laure Gaelle Lucie"
	streamer := string2streamer(phrase)
	//save(streamer, format(defaultSampleRate), "output.wav")
	return sound.Play(streamer)
}

func main() {
	if err := program(); err != nil {
		log.Fatal(err)
	}
}
