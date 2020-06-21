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

var TargetColor = COLOR_GREEN
var TargetType = TYPE_MOON

// type Node struct {
//     Depth       int
//     BlueRobot   GameResources.Robot
//     RedRobot    GameResources.Robot
//     YellowRobot GameResources.Robot
//     GreenRobot  GameResources.Robot
//     GameMap     [][]GameResources.GameTile
//     ParentNode  *Node
//     Child       []Node
// }
//
// func (node *Node) AddChild(child Node) {
//     node.Child = append(node.Child, child)
// }
//
// func (node *Node) MoveRobots(queue *Queue) (*Node, int) {
//     tempNode := *node
//     if tempNode.GreenRobot.MoveWest(tempNode.GameMap) {
//         tempNode.Depth = node.Depth + 1
//         if tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Type == TargetType &&
//             tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Color == TargetColor {
//             return &tempNode, tempNode.Depth
//         }
//         tempNode.ParentNode = node
//         node.AddChild(tempNode)
//         queue.Push(tempNode)
//     }
//     tempNode = *node
//     if tempNode.GreenRobot.MoveEast(tempNode.GameMap) {
//         tempNode.Depth = node.Depth + 1
//         if tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Type == TargetType &&
//             tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Color == TargetColor {
//             return &tempNode, tempNode.Depth
//         }
//         tempNode.ParentNode = node
//         node.AddChild(tempNode)
//         queue.Push(tempNode)
//     }
//     tempNode = *node
//     if tempNode.GreenRobot.MoveNorth(tempNode.GameMap) {
//         tempNode.Depth = node.Depth + 1
//         if tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Type == TargetType &&
//             tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Color == TargetColor {
//             return &tempNode, tempNode.Depth
//         }
//         tempNode.ParentNode = node
//         node.AddChild(tempNode)
//         queue.Push(tempNode)
//     }
//     tempNode = *node
//     if tempNode.GreenRobot.MoveSouth(tempNode) {
//         tempNode.Depth = node.Depth + 1
//         fmt.Println("SOUTH")
//         fmt.Println("DEPTH")
//         fmt.Println(node.Depth)
//         fmt.Println(tempNode.Depth)
//         fmt.Println(tempNode.BlueRobot.XPosition)
//         fmt.Println(tempNode.BlueRobot.YPosition)
//         fmt.Println(tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition])
//         fmt.Println(tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Type == TargetType)
//         fmt.Println(tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Color == TargetColor)
//         if tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Type == TargetType &&
//             tempNode.GameMap[tempNode.BlueRobot.XPosition][tempNode.BlueRobot.YPosition].Color == TargetColor {
//             return &tempNode, tempNode.Depth
//         }
//         tempNode.ParentNode = node
//         node.AddChild(tempNode)
//         queue.Push(tempNode)
//     }
//     return nil, 0
// }
//
// type Queue struct {
//     Nodes []Node
// }
//
// func (queue *Queue) Push(node Node) {
//     queue.Nodes = append(queue.Nodes, node)
// }
//
// func (queue *Queue) Pop() Node {
//     var x Node
//     x = queue.Nodes[0]
//     queue.Nodes = queue.Nodes[1:]
//     return x
// }

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
	var startNode = GameResources.Node{0, blueRobot, redRobot, yellowRobot, greenRobot, newMap, nil, defaultChild}
	node, depth := startNode.MoveRobots(&queue)
	// fmt.Println(node)
	fmt.Println(depth)
	// fmt.Println(node.ParentNode)
	fmt.Println(node.Depth)
	fmt.Println(node.ParentNode.Depth)
	nn := node
	for {
		fmt.Println("HISTORY")
		fmt.Println(nn.BlueRobot.XPosition, nn.BlueRobot.YPosition)
		fmt.Println(nn.RedRobot.XPosition, nn.RedRobot.YPosition)
		fmt.Println(nn.YellowRobot.XPosition, nn.YellowRobot.YPosition)
		fmt.Println(nn.GreenRobot.XPosition, nn.GreenRobot.YPosition)
		if nn.ParentNode == nil {
			fmt.Println("PARENT NIL")
			break
		}
		nn = node.ParentNode
	}
	// var tempNode = startNode
	// if tempNode.BlueRobot.MoveWest(tempNode.GameMap) {
	//     tempNode.ParentNode = &startNode
	//     startNode.AddChild(tempNode)
	//     queue.Push(tempNode)
	// }
	// tempNode = startNode
	// if tempNode.BlueRobot.MoveEast(tempNode.GameMap) {
	//     tempNode.ParentNode = &startNode
	//     startNode.AddChild(tempNode)
	//     queue.Push(tempNode)
	// }
	// tempNode = startNode
	// if tempNode.BlueRobot.MoveNorth(tempNode.GameMap) {
	//     tempNode.ParentNode = &startNode
	//     startNode.AddChild(tempNode)
	//     queue.Push(tempNode)
	// }
	// tempNode = startNode
	// if tempNode.BlueRobot.MoveSouth(tempNode.GameMap) {
	//     tempNode.ParentNode = &startNode
	//     startNode.AddChild(tempNode)
	//     queue.Push(tempNode)
	// }
	// fmt.Println(startNode)
	// fmt.Println(len(startNode.Child))
	// fmt.Println(startNode.Child[0])
	// fmt.Println(startNode.Child[0].GameMap[1][1])
	// fmt.Println(startNode.Child[0].GameMap[1][0])
	// fmt.Println(startNode.Child[1].BlueRobot)
	// fmt.Println(startNode.Child[1].GameMap[1][1])
	// fmt.Println(startNode.Child[1].GameMap[1][4])
	// fmt.Println(startNode.Child[2].BlueRobot)
	// fmt.Println(startNode.Child[2].GameMap[1][1])
	// fmt.Println(startNode.Child[2].GameMap[0][1])
	// fmt.Println(startNode.Child[3].BlueRobot)
	// fmt.Println(startNode.Child[3].GameMap[1][1])
	// fmt.Println(startNode.Child[3].GameMap[5][1])

	// fmt.Println(queue)

	///////////////////////////////////

	// queue.Push(startNode)
	// var nextNode = queue.Pop()
	// fmt.Println(nextNode.Depth)
	// nextNode.BlueRobot.MoveWest(nextNode.GameMap)
	// nextNode.BlueRobot.MoveEast(nextNode.GameMap)
	// nextNode.BlueRobot.MoveNorth(nextNode.GameMap)
	// nextNode.BlueRobot.MoveSouth(nextNode.GameMap)

	// var defaultChild2 []Node
	// var ssn = Node{1, defaultChild2}
	// startNode.AddChild(ssn)
	// fmt.Println(startNode)

	// fmt.Println("-----------------")
	// fmt.Println(newMap[11][11])
	// fmt.Println(blueRobot)
	// fmt.Println(redRobot)
	// fmt.Println(yellowRobot)
	// fmt.Println(greenRobot)
	// fmt.Println("-----------------")

	// greenRobot.MoveEast(newMap)
	// fmt.Println(greenRobot)
	// greenRobot.MoveWest(newMap)
	// fmt.Println(greenRobot)
	// greenRobot.MoveNorth(newMap)
	// fmt.Println(greenRobot)
	// greenRobot.MoveSouth(newMap)
	// fmt.Println(greenRobot)

	// fmt.Println(newMap[14][11])
	// fmt.Println(newMap[12][11])
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
	yellowRobot.XPosition = 14
	yellowRobot.YPosition = 2
	newMap[14][2].ExistRobot = true

	greenRobot.Color = COLOR_GREEN
	greenRobot.XPosition = 10
	greenRobot.YPosition = 11
	newMap[10][11].ExistRobot = true
}
