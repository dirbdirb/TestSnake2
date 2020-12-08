package game


var obstacleList = []rune {
    '🦴',
}


type Obstacle struct {
    char    rune
    coord   Coordinate
}

// Function that drops obstacles on the field
func ObstacleAt(c Coordinate) Obstacle {
	return Obstacle{
		char:	CharObstacle(),
		coord:	c,
	}
}

func CharObstacle() rune {
    if RuneSupport() {
        return obstacleList[0]
    }

    return 'X'
}
