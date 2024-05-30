package main

import (
	"bitmap-analyzer/crop"
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"sync"
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

	buff := make([]uint64, 9)
	shouldClose := false
	var wg sync.WaitGroup
	wg.Add(9)
	for idx, tile := range *tiles {
		if idx == len(*tiles)-1 {
			shouldClose = true
		}
		go func(tile crop.Tile, idx int, shouldClose bool) {
			gray := tile.GetGrayLevel()
			buff[idx] = gray
			log.Printf("TILE[%v,%v] %v\n", tile.X, tile.Y, gray)
			wg.Done()
		}(tile, idx, shouldClose)
	}

	var sum uint64 = 0
	wg.Wait()
	for _, value := range buff {
		sum += value
	}
	result := sum / 9

	log.Printf("GRAY_LVL %v\n", result)

	if result < 65535/2 {
		fmt.Printf("DARK\n")
	} else {
		fmt.Printf("LIGHT\n")
	}

	stop := time.Now()
	fmt.Printf("time: %vms\n", stop.UnixMilli()-start.UnixMilli())
}
