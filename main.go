package main

import (
	"fmt"
	"github.com/BJ-Kim/ricochet_robot/Constants"
	"github.com/BJ-Kim/ricochet_robot/GameResources"
	// "sync"
	"time"
)

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

	startTime := time.Now()
	var startNode = GameResources.Node{0, blueRobot, redRobot, yellowRobot, greenRobot, newMap, nil, "", ""}
	var node *GameResources.Node
	var count = 0
	// var wg sync.WaitGroup
	// messages := make(chan *GameResources.Node)
	for {
		// wg.Add(1)

		// go func() {
		//     defer wg.Done()
		//     messages <- startNode.MoveRobots(&queue, count)
		// }()

		findNode := startNode.MoveRobots(&queue, count)
		if findNode != nil {
			node = findNode
			break
		}
		count++
	}
	// go func() {
	//     for i := range messages {
	//         fmt.Println("??????????????????????")
	//         fmt.Println(i)
	//         fmt.Println("??????????????????????")
	//     }
	// }()
	// wg.Wait()
	nn := node
	for {
		fmt.Println("------------------HISTORY-----------------")
		nn.PrintCurrentPosition()
		if nn.ParentNode == nil {
			fmt.Println("PARENT NIL")
			break
		}
		nn = nn.ParentNode
	}
	endTime := time.Now()
	fmt.Println("################################")
	fmt.Println(endTime.Sub(startTime))
	fmt.Println("COUNT : ", count)
	fmt.Println("DEPTH : ", node.Depth)
	fmt.Println("################################")
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
				newMap[i][j].Color = Constants.COLOR_BLUE
				newMap[i][j].Type = Constants.TYPE_MOON
			}
			if i == 1 && j == 11 {
				newMap[i][j].Color = Constants.COLOR_RED
				newMap[i][j].Type = Constants.TYPE_MOON
			}
			if i == 3 && j == 14 {
				newMap[i][j].Color = Constants.COLOR_GREEN
				newMap[i][j].Type = Constants.TYPE_TRIANGLE
			}
			if i == 4 && j == 3 {
				newMap[i][j].Color = Constants.COLOR_RED
				newMap[i][j].Type = Constants.TYPE_STAR
			}
			if i == 4 && j == 9 {
				newMap[i][j].Color = Constants.COLOR_BLUE
				newMap[i][j].Type = Constants.TYPE_PLANET
			}
			if i == 5 && j == 6 {
				newMap[i][j].Color = Constants.COLOR_GREEN
				newMap[i][j].Type = Constants.TYPE_PLANET
			}
			if i == 6 && j == 1 {
				newMap[i][j].Color = Constants.COLOR_YELLOW
				newMap[i][j].Type = Constants.TYPE_TRIANGLE
			}
			if i == 6 && j == 12 {
				newMap[i][j].Color = Constants.COLOR_YELLOW
				newMap[i][j].Type = Constants.TYPE_STAR
			}
			if i == 9 && j == 12 {
				newMap[i][j].Color = Constants.COLOR_BLUE
				newMap[i][j].Type = Constants.TYPE_STAR
			}
			if i == 10 && j == 3 {
				newMap[i][j].Color = Constants.COLOR_BLUE
				newMap[i][j].Type = Constants.TYPE_TRIANGLE
			}
			if i == 10 && j == 10 {
				newMap[i][j].Color = Constants.COLOR_YELLOW
				newMap[i][j].Type = Constants.TYPE_PLANET
			}
			if i == 11 && j == 5 {
				newMap[i][j].Color = Constants.COLOR_GREEN
				newMap[i][j].Type = Constants.TYPE_STAR
			}
			if i == 12 && j == 2 {
				newMap[i][j].Color = Constants.COLOR_YELLOW
				newMap[i][j].Type = Constants.TYPE_MOON
			}
			if i == 12 && j == 14 {
				newMap[i][j].Color = Constants.COLOR_RED
				newMap[i][j].Type = Constants.TYPE_TRIANGLE
			}
			if i == 13 && j == 4 {
				newMap[i][j].Color = Constants.COLOR_RED
				newMap[i][j].Type = Constants.TYPE_PLANET
			}
			if i == 14 && j == 11 {
				newMap[i][j].Color = Constants.COLOR_GREEN
				newMap[i][j].Type = Constants.TYPE_MOON
			}
		}
	}
}

func SetDefaultRobot(newMap [][]GameResources.GameTile,
	blueRobot *GameResources.Robot,
	redRobot *GameResources.Robot,
	yellowRobot *GameResources.Robot,
	greenRobot *GameResources.Robot) {
	blueRobot.Color = Constants.COLOR_BLUE
	blueRobot.XPosition = 1
	blueRobot.YPosition = 9

	redRobot.Color = Constants.COLOR_RED
	redRobot.XPosition = 4
	redRobot.YPosition = 13

	yellowRobot.Color = Constants.COLOR_YELLOW
	yellowRobot.XPosition = 14
	yellowRobot.YPosition = 13

	greenRobot.Color = Constants.COLOR_GREEN
	greenRobot.XPosition = 12
	greenRobot.YPosition = 3
}
