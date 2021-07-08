package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Dimensions
const (
	N                = 3 // 3 x 3 grid in this case
	SCREEN_WIDTH_PX  = 480
	SCREEN_HEIGHT_PX = 480
	CELL_WIDTH       = SCREEN_WIDTH_PX / N
	CELL_HEIGHT      = SCREEN_HEIGHT_PX / N
)

type state int

// Possible states of state type
const (
	PLAYER_X_WON state = iota
	PLAYER_O_WON
	IS_GAME_RUNNING
	IS_GAME_A_TIE
	IS_GAME_OVER
)

type board [N * N]int

// Possible states of board type
const (
	EMPTY = iota
	PLAYER_X
	PLAYER_O
)

// Colors
var (
	grid_color     = color.RGBA{0xff, 0xff, 0xff, 0xff} // white
	player_x_color = color.RGBA{0x20, 0x4a, 0x87, 0xff} // blue
	player_o_color = color.RGBA{0x73, 0xd2, 0x16, 0xff} // green
)

// Game implements ebiten.Game interface.
type Game struct {
	board  [N * N]int
	player int
	state  int
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	g.MouseClickEvent()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
// Should not mutate the game state, just render the state
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	g.board = [N * N]int{
		EMPTY, EMPTY, PLAYER_X,
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY,
	}
	g.player = PLAYER_X
	g.state = int(IS_GAME_RUNNING)

	g.RenderBoard(screen, player_x_color, player_o_color)
	g.RenderGrid(screen)
}

// ebiten.CursorPosition(x, y int) {
// }

func (g *Game) RenderGrid(screen *ebiten.Image) {
	for i := 1; i < N; i++ {
		ebitenutil.DrawLine(screen, float64(i*CELL_WIDTH), 0, float64(i*CELL_WIDTH), SCREEN_HEIGHT_PX, grid_color)
		ebitenutil.DrawLine(screen, 0, float64(i*CELL_HEIGHT), SCREEN_WIDTH_PX, float64(i*CELL_HEIGHT), grid_color)
	}
}

func (g *Game) RenderBoard(screen *ebiten.Image, player_x_color, player_o_color color.RGBA) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			switch g.board[i*N+j] {
			case PLAYER_X:
				g.RenderX(screen, i, j, player_x_color)

			// case PLAYER_O:
			// g.RenderO(screen, i, j, player_o_color)

			default:
				fmt.Println("default")
			}
		}
	}
}

func (g *Game) RenderX(screen *ebiten.Image, row, column int, player_x_color color.RGBA) {

	var (
		HALF_BOX_SIZE float64 = math.Min(float64(CELL_WIDTH), float64(CELL_HEIGHT)*0.25)
		CENTER_X      float64 = CELL_WIDTH*0.5 + float64(column)*CELL_WIDTH
		CENTER_Y      float64 = CELL_HEIGHT*0.5 + float64(row)*CELL_HEIGHT
	)

	// Draw line from top left to bottom right
	ebitenutil.DrawLine(
		screen,
		CENTER_X-HALF_BOX_SIZE,
		CENTER_Y-HALF_BOX_SIZE,
		CENTER_X+HALF_BOX_SIZE,
		CENTER_Y+HALF_BOX_SIZE,
		player_x_color)

	// Draw line from top right to bottom left to complete the X
	ebitenutil.DrawLine(
		screen,
		CENTER_X+HALF_BOX_SIZE,
		CENTER_Y-HALF_BOX_SIZE,
		CENTER_X-HALF_BOX_SIZE,
		CENTER_Y+HALF_BOX_SIZE,
		player_x_color)
}

// func (g *Game) RenderO(row, column int, player_o_color, screen *ebiten.Image) {
// }

// row, column int
func (g *Game) MouseClickEvent() bool {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		fmt.Println("Mouse pressed")
	}
	return false
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(SCREEN_WIDTH_PX, SCREEN_HEIGHT_PX)
	ebiten.SetWindowTitle("Tic Tac Toe game")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
