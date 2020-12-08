package game

import (
	"os"
	"strings"
	"math/rand"
)

// Thanks to https://unicode-table.com/en/#supplemental-symbols-and-pictographs
var foodList = []rune{
	'🥓',
	'🥔',
	'🥛',
	'🥑',
}

type Food struct {
	char	rune
	coord	Coordinate
}

func InitFood() Food {
	return Food {
		char: CharFood(),
		// TODO make coord that isn't on snake.
	}
}

func CharFood() rune {
	if RuneSupport() {
		return foodList[rand.Intn(len(foodList))]
	}

	return '*'
}


func RuneSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}

func DropFoodAt(c Coordinate) Food {
	return Food{
		char:	CharFood(),
		coord:	c,
	}
}
