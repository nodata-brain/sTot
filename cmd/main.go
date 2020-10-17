package main

import (
	"github.com/gordonklaus/portaudio"
	"log"
	//"sTot/pkg/googleapi"
)

func main() {
	err := portaudio.Initialize()
	if err != nil {
		log.Printf("Initialize error : %s", err)
		return
	}
	p, err := portaudio.OpenDefaultStream(4, 0, 44100.0, 256, 1)
	if err != nil {
		log.Printf("Open Stream error : %s", err)
		return
	}
	log.Printf("%v", p)
	//stot.New()
}
