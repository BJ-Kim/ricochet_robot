package GameResources

import (
// "fmt"
)

type Robot struct {
	Color     string
	XPosition int
	YPosition int
}

// func (rb *Robot) MoveWest(gameMap [][]GameTile) bool {
//     var x = rb.XPosition
//     var y = rb.YPosition
//     gameMap[x][y].ExistRobot = false
//     for i := y; i >= 0; i-- {
//         if i == 0 {
//             rb.YPosition = i
//             gameMap[x][i].ExistRobot = true
//             return true
//         }
//         if gameMap[x][i].West == true || gameMap[x][i-1].ExistRobot == true {
//             if i == y {
//                 return false
//             }
//             rb.YPosition = i
//             gameMap[x][i].ExistRobot = true
//             return true
//         }
//     }
//     return false
// }
//
// func (rb *Robot) MoveEast(gameMap [][]GameTile) bool {
//     var x = rb.XPosition
//     var y = rb.YPosition
//     gameMap[x][y].ExistRobot = false
//     for i := y; i < 16; i++ {
//         if i == 15 {
//             rb.YPosition = i
//             gameMap[x][i].ExistRobot = true
//             return true
//         }
//         if gameMap[x][i].East == true || gameMap[x][i+1].ExistRobot == true {
//             if i == y {
//                 return false
//             }
//             rb.YPosition = i
//             gameMap[x][i].ExistRobot = true
//             return true
//         }
//     }
//     return false
// }
//
// func (rb *Robot) MoveNorth(gameMap [][]GameTile) bool {
//     var x = rb.XPosition
//     var y = rb.YPosition
//     gameMap[x][y].ExistRobot = false
//     for i := x; i >= 0; i-- {
//         if i == 0 {
//             rb.XPosition = i
//             gameMap[i][y].ExistRobot = true
//             return true
//         }
//         if gameMap[i][y].North == true || gameMap[i-1][y].ExistRobot == true {
//             if i == y {
//                 return false
//             }
//             rb.XPosition = i
//             gameMap[i][y].ExistRobot = true
//             return true
//         }
//     }
//     return false
// }

// func (rb *Robot) MoveSouth(gameMap [][]GameTile) bool {
//     var x = rb.XPosition
//     var y = rb.YPosition
//     gameMap[x][y].ExistRobot = false
//     for i := x; i < 16; i++ {
//         if i == 15 {
//             rb.XPosition = i
//             gameMap[i][y].ExistRobot = true
//             return true
//         }
//         if gameMap[i][y].South == true || gameMap[i+1][y].ExistRobot == true {
//             if i == y {
//                 return false
//             }
//             rb.XPosition = i
//             gameMap[i][y].ExistRobot = true
//             return true
//         }
//     }
//     return false
// }

func (rb *Robot) MoveWest(node *Node) bool {
	if node.FromDirection == DIRECTION_WEST && node.MoveRobotColor == rb.Color {
		return false
	}
	var x = rb.XPosition
	var y = rb.YPosition
	node.GameMap[x][y].ExistRobot = false
	for i := y; i >= 0; i-- {
		if i == 0 {
			rb.YPosition = i
			node.GameMap[x][i].ExistRobot = true
			return true
		}
		if node.GameMap[x][i].West == true || node.GameMap[x][i-1].ExistRobot == true {
			if i == y {
				return false
			}
			rb.YPosition = i
			node.GameMap[x][i].ExistRobot = true
			return true
		}
	}
	return false
}

func (rb *Robot) MoveEast(node *Node) bool {
	if node.FromDirection == DIRECTION_EAST && node.MoveRobotColor == rb.Color {
		return false
	}
	var x = rb.XPosition
	var y = rb.YPosition
	node.GameMap[x][y].ExistRobot = false
	for i := y; i < 16; i++ {
		if i == 15 {
			rb.YPosition = i
			node.GameMap[x][i].ExistRobot = true
			return true
		}
		if node.GameMap[x][i].East == true || node.GameMap[x][i+1].ExistRobot == true {
			if i == y {
				return false
			}
			rb.YPosition = i
			node.GameMap[x][i].ExistRobot = true
			return true
		}
	}
	return false
}

func (rb *Robot) MoveNorth(node *Node) bool {
	if node.FromDirection == DIRECTION_NORTH && node.MoveRobotColor == rb.Color {
		return false
	}
	var x = rb.XPosition
	var y = rb.YPosition
	node.GameMap[x][y].ExistRobot = false
	for i := x; i >= 0; i-- {
		if i == 0 {
			rb.XPosition = i
			node.GameMap[i][y].ExistRobot = true
			return true
		}
		if node.GameMap[i][y].North == true || node.GameMap[i-1][y].ExistRobot == true {
			if i == x {
				return false
			}
			rb.XPosition = i
			node.GameMap[i][y].ExistRobot = true
			return true
		}
	}
	return false
}

func (rb *Robot) MoveSouth(node *Node) bool {
	if node.FromDirection == DIRECTION_SOUTH && node.MoveRobotColor == rb.Color {
		return false
	}
	var x = rb.XPosition
	var y = rb.YPosition
	node.GameMap[x][y].ExistRobot = false
	for i := x; i < 16; i++ {
		if i == 15 {
			rb.XPosition = i
			node.GameMap[i][y].ExistRobot = true
			return true
		}
		if node.GameMap[i][y].South == true || node.GameMap[i+1][y].ExistRobot == true {
			if i == x {
				return false
			}
			rb.XPosition = i
			node.GameMap[i][y].ExistRobot = true
			return true
		}
	}
	return false
}
