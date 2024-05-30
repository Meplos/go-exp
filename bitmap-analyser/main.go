package main

import (
	"bitmap-analyzer/crop"
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"
)

func main() {
	analyze("/home/meplos/Documents/sample/fox.jpg")
}

func openImage(filepath string) image.Image {
	fmt.Printf("Start processing image\n")
	file, openingErr := os.Open(filepath)
	if openingErr != nil {
		log.Printf("Cannot open file")
		panic("file not exist")
	}

	filereader := bufio.NewReader(file)

	img, decodeErr := jpeg.Decode(filereader)

	if decodeErr != nil {
		log.Printf("Cannot read file")
		panic("invalid format")
	}
	return img
}

func analyze(filepath string) {
	start := time.Now()

	img := openImage(filepath)

	tiles := crop.Divide(img.(*image.YCbCr), 3)

	c := make(chan int, len(*tiles))
	shouldClose := false
	for idx, tile := range *tiles {
		if idx == len(*tiles)-1 {
			shouldClose = true
		}
		go func(tile crop.Tile, shouldClose bool) {
			gray := tile.GetGrayLevel()
			c <- int(gray)
			if shouldClose {
				close(c)
			}
		}(tile, shouldClose)
	}

	sum := 0
	for i := 0; i < len(*tiles); i++ {
		sum += <-c
	}
	result := sum / len(*tiles)

	if result < 65535/2 {
		fmt.Printf("DARK\n")
	} else {
		fmt.Printf("LIGHT\n")
	}

	stop := time.Now()
	fmt.Printf("time: %vms\n", stop.UnixMilli()-start.UnixMilli())
}
