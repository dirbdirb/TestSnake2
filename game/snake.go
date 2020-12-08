package game


type direction int

const InitialSnakeLength = 4
const MAX_SPEED = 60

// Snake the actual snake
type Snake struct {
	body      	[]Coordinate // Snake body
	length    	int          // Snake Length
	direction 	direction    // Direction snake is facing
	speed		int
}

// Got help from https://programming.guide/go/define-enumeration-string.html and https://golangbyexample.com/iota-in-golang/
const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

func InitSnake(w, h int) Snake {
	var temp []Coordinate

	for y := 0; y < InitialSnakeLength; y++ {
		temp = append(temp, Coordinate{x: w / 2, y: (h / 2) + y}) // Essentially, we want body[0] to be the head, pointing upwards.
	}

	return Snake{
		body:      	temp,
		length:    	InitialSnakeLength,
		direction: 	UP,
		speed:		100,
	}
}

func (s *Snake) moveBody(coord Coordinate) {
	// A little jank...
	// Have to make slice without 'make' because it always allocates too much memory over and over due to the len(slice) not actually representing the number of elements in the slice.
	var temp []Coordinate
	temp = append(temp, coord)
	for i := 0; i < len(s.body)-1; i++ { // Copy all except last element over, tail needs to be forgotten.
		temp = append(temp, s.body[i])
	}

	s.body = temp
}

// Function to check if the Coordinate is on the snake body
func (s *Snake) CheckHeadPosition(c Coordinate) bool {
	// Start at s.body[1:] since we want to check the body not including the
	// head
	for _, body := range s.body[1:] {
		// Check if the coord of the head match one coord in body
		if c.x == body.x && c.y == body.y {
			return true
		}
	}
	return false
}

// Check body position for random food drops
func (s *Snake) AvailablePosition(c Coordinate) bool {
	// Traverse through snake body to see if the randomly selected Coordinate
	// is available
	for _, b := range s.body {
		// Check if the coord of the head match one coord in body
		if c.x == b.x && c.y == b.y {
			return false
		}
	}
	return true
}

func (s *Snake) IncreaseSpeed() {
	if s.speed > MAX_SPEED {
		s.speed -= 10
	}
}
