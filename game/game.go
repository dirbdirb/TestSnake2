package game

import (
	"os"
	"time"
	"fmt"

	"github.com/nsf/termbox-go"
)


type Game struct {
	field	Field
}

type Coordinate struct {
	x int
	y int
}

var (
	cDir = UP
)

// StartGame starts the game of snake.
func StartGame() {
	// Initualizes termbox library
	err := termbox.Init()

	// If an error occur, then panic
	if err != nil {
		// If fail, stop all function and go to defer functions
		panic(err)
	}
	// Ensure that termbox always closes
	defer termbox.Close()

	/*
	   Order of events:
	    1. Show intro screen.
	    2. Get player input so we can decide to start game or whatnot.
        3. Start game when player inputs start.
        4. Show snake screen.
        5. Render snake.
        6. Start moving snake and randomly giving food.
        7. Take player input for direction.
		8. Check if player loses
	*/

	game := Game {
		field:	InitField(),
	}

	// Watch for player input.
	go WatchPlayerInput(&game)

	// Display the field
	game.field.Display()

	for {
		game.field.move()
		game.field.Display()

		// So that we can increase the speed of the snake
		t := time.Duration(game.field.snake.speed)
		time.Sleep(t * time.Millisecond)
	}
}

// WatchPlayerInput watches for player input event
func WatchPlayerInput(game *Game) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		e := termbox.PollEvent()

		// The current direction of the snake
		curDir := game.field.snake.direction

		switch e.Key {
		case termbox.KeyArrowUp:
			// The if statement is so that the snake doesn't go
			// backwards

			if curDir == DOWN {
				game.field.snake.direction = DOWN
			} else {
				game.field.snake.direction = UP
			}

		case termbox.KeyArrowDown:
			if curDir == UP {
				game.field.snake.direction = UP
			} else {
				game.field.snake.direction = DOWN
			}

		case termbox.KeyArrowLeft:
			if curDir == RIGHT {
				game.field.snake.direction = RIGHT
			} else {
				game.field.snake.direction = LEFT
			}

		case termbox.KeyArrowRight:
			if curDir == LEFT {
				game.field.snake.direction = LEFT
			} else {
				game.field.snake.direction = RIGHT
			}

		case termbox.KeyEsc:
			QuitGame()
			return
		}
	}
}

// Function used to close the game when player press esc key
func QuitGame() {
	// Close the termbox
	termbox.Close()

	// Display message to player
	fmt.Println("Thanks for playing!!")

	// Close program without error
	os.Exit(0)
}


// Function for when the player lost
func GameOver(reason string, score int) {
	// Close the termbox
	termbox.Close()

	fmt.Println("Game Over!", reason)
	fmt.Println("Your score is ", score, "points!\n")

	// Close program without error
	os.Exit(0)
}
