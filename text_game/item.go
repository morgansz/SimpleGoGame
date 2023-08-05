package main

type Item struct {
	X, Y              int
	Symbol            string
	AsciiArtGenerator *asciiArtGenerator
}

func NewItem(x, y int) *Item {
	aag := &asciiArtGenerator{}
	return &Item{
		X: x, Y: y,
		AsciiArtGenerator: aag,
		Symbol:            aag.GenerateArt("$"),
	}
}
