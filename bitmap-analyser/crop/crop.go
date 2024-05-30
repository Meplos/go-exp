package crop

import (
	"image"
	"log"
)

type Tile struct {
	X     uint
	Y     uint
	image image.Image
}

func (t *Tile) GetGrayLevel() uint64 {
	var sum uint64 = 0
	var i uint64 = 0

	for x := t.image.Bounds().Min.X; x < t.image.Bounds().Max.X; x++ {
		for y := t.image.Bounds().Min.Y; y < t.image.Bounds().Max.Y; y++ {
			r, g, b, _ := t.image.At(x, y).RGBA()
			grayLvl := ((r + g + b) / 3)
			sum += uint64(grayLvl)
			i++
		}
	}
	gray := sum / i
	return gray
}

func Divide(origin *image.YCbCr, nbTile uint) *[]Tile {
	paddingX := 0
	paddingY := 0
	tileHeight := origin.Bounds().Dy() / int(nbTile)
	tileWidth := origin.Bounds().Dx() / int(nbTile)

	offsetX := origin.Bounds().Dx() % int(nbTile)
	if offsetX != 0 {
		paddingX = getPadding(origin.Bounds().Dx(), offsetX)
	}

	offsetY := origin.Bounds().Dy() % int(nbTile)
	if offsetY != 0 {
		paddingY = getPadding(origin.Bounds().Dy(), offsetY)
	}
	log.Printf("DIV P_X(%v) P_Y(%v)\n", paddingX, paddingY)

	startX := origin.Bounds().Min.X - paddingX
	startY := origin.Bounds().Min.Y - paddingY

	endX := origin.Bounds().Max.X - paddingX
	endY := origin.Bounds().Max.Y - paddingY

	cropImage := origin.SubImage(image.Rect(startX, startY, endX, endY))
	tiles := make([]Tile, 0)
	tileX := 0
	for x := cropImage.Bounds().Min.X; x < cropImage.Bounds().Max.X; x += (tileWidth + 1) {
		tileEndX := x + tileWidth
		tileY := 0
		for y := cropImage.Bounds().Min.Y; y < cropImage.Bounds().Max.Y; y += (tileHeight + 1) {
			tileEndY := y + tileHeight
			tile := Tile{
				X:     uint(tileX),
				Y:     uint(tileY),
				image: origin.SubImage(image.Rect(x, y, tileEndX, tileEndY)),
			}
			log.Printf("TILE[%v, %v] : %v\n", tile.X, tile.Y, tile.image.Bounds())
			tiles = append(tiles, tile)
			tileY++
		}
		tileX++
	}

	log.Printf("len(TILES)=%v\n", len(tiles))
	return &tiles
}

func getPadding(size int, offset int) int {
	return 0
}
