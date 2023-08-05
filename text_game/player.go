package main

type asciiArtGenerator struct{}

func (aag *asciiArtGenerator) GenerateArt(symbol string) string {
	return "ASCII Art of " + symbol
}

type Player struct {
	X, Y              int
	VelocityX         int
	VelocityY         int
	OnGround          bool
	Symbol            string
	AsciiArtGenerator *asciiArtGenerator
}

func NewPlayer(x, y int) *Player {
	aag := &asciiArtGenerator{}
	return &Player{
		X: x, Y: y,
		AsciiArtGenerator: aag,
		Symbol:            aag.GenerateArt("@"),
	}
}

func (p *Player) Update() {
	p.X += p.VelocityX
	p.Y += p.VelocityY
	p.VelocityY--
	if p.Y <= 0 {
		p.Y = 0
		p.VelocityY = 0
		p.OnGround = true
	}
}

func (p *Player) Jump() {
	if p.OnGround {
		p.VelocityY = 10
		p.OnGround = false
	}
}

func (p *Player) MoveLeft() {
	p.VelocityX = -5
}

func (p *Player) MoveRight() {
	p.VelocityX = 5
}

func (p *Player) Stop() {
	p.VelocityX = 0
}
