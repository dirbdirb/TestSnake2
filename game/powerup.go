package game

var powerUpList = []rune{
	'🌈',
}

type PowerUp struct {
	char  rune
	coord Coordinate
}

func InitPowerUp() PowerUp {
	return PowerUp{
		char: CharPowerUp(),
		// TODO make coord that isn't on snake.
	}
}

func CharPowerUp() rune {
	if RuneSupport() {
		return powerUpList[0]
	}

	return '?'
}

func DropPowerUpAt(c Coordinate) PowerUp {
	return PowerUp{
		char:  CharPowerUp(),
		coord: c,
	}
}
