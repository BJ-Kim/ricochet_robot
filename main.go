package main

import (
	"fmt"
	"github.com/BJ-Kim/ricochet_robot/GameResources"
)

var COLOR_BLUE = "BLUE"
var COLOR_RED = "RED"
var COLOR_YELLOW = "YELLOW"
var COLOR_GREEN = "GREEN"

var TYPE_MOON = "MOON"
var TYPE_STAR = "STAR"
var TYPE_PLANET = "PLANET"
var TYPE_TRIANGLE = "TRIANGLE"

var DIRECTION_WEST = "WEST"
var DIRECTION_EAST = "EAST"
var DIRECTION_NORTH = "NORTH"
var DIRECTION_SOUTH = "SOUTH"

var TargetColor = COLOR_GREEN
var TargetType = TYPE_MOON

func main() {
	var queue GameResources.Queue
	var newMap [][]GameResources.GameTile
	var blueRobot GameResources.Robot
	var redRobot GameResources.Robot
	var yellowRobot GameResources.Robot
	var greenRobot GameResources.Robot
	InitMapTile(&newMap)
	SetWall(newMap)
	SetDestination(newMap)
	SetDefaultRobot(newMap, &blueRobot, &redRobot, &yellowRobot, &greenRobot)

	fmt.Println(TargetColor)
	fmt.Println(TargetType)
	/*
		TODO: make node and push to queue
	*/

	var defaultChild []GameResources.Node
	var startNode = GameResources.Node{0, blueRobot, redRobot, yellowRobot, greenRobot, newMap, nil, defaultChild, "", ""}
	var node *GameResources.Node
	var depth int
	for {
		findNode, moveCount := startNode.MoveRobots(&queue)
		// fmt.Println(len(queue.Nodes))
		// if count == 6 {
		//     for i := 0; i < len(queue.Nodes); i++ {
		//         queue.Nodes[i].PrintCurrentPosition()
		//     }
		// }
		// if count == 6 {
		//     queue.Nodes[0].PrintCurrentPosition()
		//     fmt.Println("-------------------------------")
		// }
		// if count == 7 {
		//     queue.Nodes[0].PrintCurrentPosition()
		//     fmt.Println("-------------------------------")
		// }
		// if count == 8 {
		//     queue.Nodes[0].PrintCurrentPosition()
		//     fmt.Println("-------------------------------")
		// }
		// fmt.Println(findNode)
		if findNode != nil {
			node = findNode
			depth = moveCount
			break
		}
	}
	fmt.Println(node.Depth)
	// nn := node
	// for {
	//     fmt.Println("HISTORY")
	//     fmt.Println(nn.BlueRobot.XPosition, nn.BlueRobot.YPosition)
	//     fmt.Println(nn.RedRobot.XPosition, nn.RedRobot.YPosition)
	//     fmt.Println(nn.YellowRobot.XPosition, nn.YellowRobot.YPosition)
	//     fmt.Println(nn.GreenRobot.XPosition, nn.GreenRobot.YPosition)
	//     if nn.ParentNode == nil {
	//         fmt.Println("PARENT NIL")
	//         break
	//     }
	//     nn = nn.ParentNode
	// }
	fmt.Println(depth)
}

func InitMapTile(newMap *[][]GameResources.GameTile) {
	for i := 0; i < 16; i++ {
		var row []GameResources.GameTile
		for j := 0; j < 16; j++ {
			row = append(row, GameResources.GameTile{})
		}
		*newMap = append(*newMap, row)
	}
}

func SetWall(newMap [][]GameResources.GameTile) {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			if i == 0 {
				newMap[i][j].North = true
			}
			if j == 0 {
				newMap[i][j].West = true
			}
			if i == 15 {
				newMap[i][j].South = true
			}
			if j == 15 {
				newMap[i][j].East = true
			}
			if (i == 0 && j == 2) || (i == 0 && j == 8) ||
				(i == 1 && j == 4) || (i == 1 && j == 10) ||
				(i == 2 && j == 7) ||
				(i == 3 && j == 14) ||
				(i == 4 && j == 3) || (i == 4 && j == 9) ||
				(i == 5 && j == 5) ||
				(i == 6 && j == 1) || (i == 6 && j == 11) ||
				(i == 7 && j == 6) || (i == 7 && j == 8) ||
				(i == 8 && j == 6) || (i == 8 && j == 8) ||
				(i == 9 && j == 11) ||
				(i == 10 && j == 3) || (i == 10 && j == 10) ||
				(i == 11 && j == 5) ||
				(i == 12 && j == 1) || (i == 12 && j == 14) ||
				(i == 13 && j == 3) ||
				(i == 14 && j == 10) ||
				(i == 15 && j == 3) || (i == 15 && j == 13) {
				newMap[i][j].East = true
				newMap[i][j+1].West = true
			}
			if (i == 0 && j == 11) ||
				(i == 1 && j == 5) || (i == 1 && j == 15) ||
				(i == 2 && j == 7) || (i == 2 && j == 14) ||
				(i == 3 && j == 0) ||
				(i == 4 && j == 3) || (i == 4 && j == 6) || (i == 4 && j == 9) ||
				(i == 5 && j == 1) ||
				(i == 6 && j == 7) || (i == 6 && j == 8) || (i == 6 && j == 12) ||
				(i == 8 && j == 7) || (i == 8 && j == 8) || (i == 8 && j == 12) ||
				(i == 9 && j == 0) || (i == 9 && j == 15) ||
				(i == 10 && j == 3) || (i == 10 && j == 5) || (i == 10 && j == 10) ||
				(i == 11 && j == 14) ||
				(i == 12 && j == 2) || (i == 12 && j == 4) ||
				(i == 14 && j == 11) {
				newMap[i][j].South = true
				newMap[i+1][j].North = true
			}

		}
	}

}

func SetDestination(newMap [][]GameResources.GameTile) {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			if i == 1 && j == 5 {
				newMap[i][j].Color = COLOR_BLUE
				newMap[i][j].Type = TYPE_MOON
			}
			if i == 1 && j == 11 {
				newMap[i][j].Color = COLOR_RED
				newMap[i][j].Type = TYPE_MOON
			}
			if i == 3 && j == 14 {
				newMap[i][j].Color = COLOR_GREEN
				newMap[i][j].Type = TYPE_TRIANGLE
			}
			if i == 4 && j == 3 {
				newMap[i][j].Color = COLOR_RED
				newMap[i][j].Type = TYPE_STAR
			}
			if i == 4 && j == 9 {
				newMap[i][j].Color = COLOR_BLUE
				newMap[i][j].Type = TYPE_PLANET
			}
			if i == 5 && j == 6 {
				newMap[i][j].Color = COLOR_GREEN
				newMap[i][j].Type = TYPE_PLANET
			}
			if i == 6 && j == 1 {
				newMap[i][j].Color = COLOR_YELLOW
				newMap[i][j].Type = TYPE_TRIANGLE
			}
			if i == 6 && j == 12 {
				newMap[i][j].Color = COLOR_YELLOW
				newMap[i][j].Type = TYPE_STAR
			}
			if i == 9 && j == 12 {
				newMap[i][j].Color = COLOR_BLUE
				newMap[i][j].Type = TYPE_STAR
			}
			if i == 10 && j == 3 {
				newMap[i][j].Color = COLOR_BLUE
				newMap[i][j].Type = TYPE_TRIANGLE
			}
			if i == 10 && j == 10 {
				newMap[i][j].Color = COLOR_YELLOW
				newMap[i][j].Type = TYPE_PLANET
			}
			if i == 11 && j == 5 {
				newMap[i][j].Color = COLOR_GREEN
				newMap[i][j].Type = TYPE_STAR
			}
			if i == 12 && j == 2 {
				newMap[i][j].Color = COLOR_YELLOW
				newMap[i][j].Type = TYPE_MOON
			}
			if i == 12 && j == 14 {
				newMap[i][j].Color = COLOR_RED
				newMap[i][j].Type = TYPE_TRIANGLE
			}
			if i == 13 && j == 4 {
				newMap[i][j].Color = COLOR_RED
				newMap[i][j].Type = TYPE_PLANET
			}
			if i == 14 && j == 11 {
				newMap[i][j].Color = COLOR_GREEN
				newMap[i][j].Type = TYPE_MOON
			}
		}
	}
}

func SetDefaultRobot(newMap [][]GameResources.GameTile,
	blueRobot *GameResources.Robot,
	redRobot *GameResources.Robot,
	yellowRobot *GameResources.Robot,
	greenRobot *GameResources.Robot) {
	blueRobot.Color = COLOR_BLUE
	blueRobot.XPosition = 1
	blueRobot.YPosition = 1
	newMap[1][1].ExistRobot = true

	redRobot.Color = COLOR_RED
	redRobot.XPosition = 8
	redRobot.YPosition = 13
	newMap[8][13].ExistRobot = true

	yellowRobot.Color = COLOR_YELLOW
	// yellowRobot.XPosition = 14
	// yellowRobot.YPosition = 2
	// newMap[14][2].ExistRobot = true
	yellowRobot.XPosition = 14
	yellowRobot.YPosition = 12
	newMap[15][12].ExistRobot = true

	greenRobot.Color = COLOR_GREEN
	// greenRobot.XPosition = 10
	// greenRobot.YPosition = 11
	// newMap[10][11].ExistRobot = true
	greenRobot.XPosition = 11
	greenRobot.YPosition = 12
	newMap[11][12].ExistRobot = true
}
