package main

import (
	"log"
	"os"
	"time"

	// "github.com/gopxl/beep"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/flac"
	"github.com/gopxl/beep/speaker"
)

func main() {
	f, err := os.Open("/home/navit/Music/Lana Del Rey - Cinnamon Girl.flac")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := flac.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	sr := format.SampleRate * 1
	resampled := beep.Resample(4, sr, sr, streamer)

	speaker.Init(sr, sr.N(time.Second/10))
	done := make(chan bool)

	speaker.Play(beep.Seq(resampled, beep.Callback(func() {
		done <- true
	})))
	<-done
}
