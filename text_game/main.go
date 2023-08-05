package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Player    *Player
	Items     []*Item
	Score     int
	HighScore int
}

func NewGame() *Game {
	player := NewPlayer(0, 0)
	items := make([]*Item, 5)
	for i := range items {
		items[i] = NewItem(rand.Intn(10), rand.Intn(10))
	}
	return &Game{
		Player: player,
		Items:  items,
	}
}

func (g *Game) StartGame() {
	// Initialize game state
}

func (g *Game) EndGame() {
	// Clean up game state
}

func (g *Game) UpdateScore(points int) {
	g.Score += points
	if g.Score > g.HighScore {
		g.HighScore = g.Score
	}
}

func (g *Game) DisplayScore() {
	fmt.Printf("Current Score: %d, High Score: %d\n", g.Score, g.HighScore)
}

func (g *Game) GetPlayerPosition() (int, int) {
	return g.Player.X, g.Player.Y
}

func (g *Game) GetItemPositions() []int {
	positions := make([]int, len(g.Items))
	for i, item := range g.Items {
		positions[i] = item.X
	}
	return positions
}

func (g *Game) MovePlayer(direction string) {
	if direction == "left" {
		g.Player.MoveLeft()
	} else if direction == "right" {
		g.Player.MoveRight()
	}
}

func (g *Game) JumpPlayer() {
	g.Player.Jump()
}

func (g *Game) CheckCollision() {
	for i := 0; i < len(g.Items); i++ {
		if g.Player.X == g.Items[i].X && g.Player.Y == g.Items[i].Y {
			g.UpdateScore(1)
			g.Items = append(g.Items[:i], g.Items[i+1:]...)
		}
	}
}

func (g *Game) RenderScreen() {
	// Render the game screen
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	game.StartGame()
	for {
		game.MovePlayer("left")
		game.JumpPlayer()
		game.CheckCollision()
		game.DisplayScore()
	}
}
