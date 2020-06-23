package GameResources

import (
	// "fmt"
	"github.com/BJ-Kim/ricochet_robot/Constants"
)

type Robot struct {
	Color     string
	XPosition int
	YPosition int
}

func (rb *Robot) MoveWest(node *Node) bool {
	if node.FromDirection == Constants.DIRECTION_WEST && node.MoveRobotColor == rb.Color {
		return false
	}
	var x = rb.XPosition
	var y = rb.YPosition
	for i := y; i >= 0; i-- {
		if i == 0 {
			rb.YPosition = i
			return true
		}
		if node.GameMap[x][i].West == true || node.IsExistRobot(x, i-1) {
			if i == y {
				return false
			}
			rb.YPosition = i
			return true
		}
	}
	return false
}

func (rb *Robot) MoveEast(node *Node) bool {
	if node.FromDirection == Constants.DIRECTION_EAST && node.MoveRobotColor == rb.Color {
		return false
	}
	var x = rb.XPosition
	var y = rb.YPosition
	for i := y; i < 16; i++ {
		if i == 15 {
			rb.YPosition = i
			return true
		}
		if node.GameMap[x][i].East == true || node.IsExistRobot(x, i+1) {
			if i == y {
				return false
			}
			rb.YPosition = i
			return true
		}
	}
	return false
}

func (rb *Robot) MoveNorth(node *Node) bool {
	if node.FromDirection == Constants.DIRECTION_NORTH && node.MoveRobotColor == rb.Color {
		return false
	}
	var x = rb.XPosition
	var y = rb.YPosition
	for i := x; i >= 0; i-- {
		if i == 0 {
			rb.XPosition = i
			return true
		}
		if node.GameMap[i][y].North == true || node.IsExistRobot(i-1, y) {
			if i == x {
				return false
			}
			rb.XPosition = i
			return true
		}
	}
	return false
}

func (rb *Robot) MoveSouth(node *Node) bool {
	if node.FromDirection == Constants.DIRECTION_SOUTH && node.MoveRobotColor == rb.Color {
		return false
	}
	var x = rb.XPosition
	var y = rb.YPosition
	for i := x; i < 16; i++ {
		if i == 15 {
			rb.XPosition = i
			return true
		}
		if node.GameMap[i][y].South == true || node.IsExistRobot(i+1, y) {
			if i == x {
				return false
			}
			rb.XPosition = i
			return true
		}
	}
	return false
}
