package GameResources

import (
	"fmt"
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

type Node struct {
	Depth          int
	BlueRobot      Robot
	RedRobot       Robot
	YellowRobot    Robot
	GreenRobot     Robot
	GameMap        [][]GameTile
	ParentNode     *Node
	Child          []Node
	FromDirection  string
	MoveRobotColor string
}

func (node *Node) AddChild(child Node) {
	node.Child = append(node.Child, child)
}

func (node *Node) PrintCurrentPosition() {
	fmt.Println(node.Depth)
	fmt.Println("BLUE X:", node.BlueRobot.XPosition, "/ BLUE Y:", node.BlueRobot.YPosition)
	fmt.Println("RED:", node.RedRobot.XPosition, "/ RED Y:", node.RedRobot.YPosition)
	fmt.Println("YELLOW X:", node.YellowRobot.XPosition, "/ YELLOW Y:", node.YellowRobot.YPosition)
	fmt.Println("GREEN X:", node.GreenRobot.XPosition, "/ GREEN Y:", node.GreenRobot.YPosition)
}

func CheckGoal(node Node) bool {
	if TargetColor == COLOR_GREEN {
		if node.GameMap[node.GreenRobot.XPosition][node.GreenRobot.YPosition].Type == TargetType &&
			node.GameMap[node.GreenRobot.XPosition][node.GreenRobot.YPosition].Color == TargetColor {
			fmt.Println("FIND GREEN SOLUTION")
			return true
		}
	} else if TargetColor == COLOR_RED {
		if node.GameMap[node.RedRobot.XPosition][node.RedRobot.YPosition].Type == TargetType &&
			node.GameMap[node.RedRobot.XPosition][node.RedRobot.YPosition].Color == TargetColor {
			fmt.Println("FIND RED SOLUTION")
			return true
		}
	} else if TargetColor == COLOR_YELLOW {
		if node.GameMap[node.YellowRobot.XPosition][node.YellowRobot.YPosition].Type == TargetType &&
			node.GameMap[node.YellowRobot.XPosition][node.YellowRobot.YPosition].Color == TargetColor {
			fmt.Println("FIND YELLOW SOLUTION")
			return true
		}
	} else if TargetColor == COLOR_BLUE {
		if node.GameMap[node.BlueRobot.XPosition][node.BlueRobot.YPosition].Type == TargetType &&
			node.GameMap[node.BlueRobot.XPosition][node.BlueRobot.YPosition].Color == TargetColor {
			fmt.Println("FIND BLUE SOLUTION")
			return true
		}
	}

	return false
}

func (node *Node) MoveRobots(queue *Queue) (*Node, int) {
	if len(queue.Nodes) != 0 {
		node = queue.Pop()
	}
	if node.Depth == 1 {
		node.PrintCurrentPosition()
	}
	tempNode := *node
	// if tempNode.GreenRobot.MoveWest(tempNode.GameMap) {
	if tempNode.GreenRobot.MoveWest(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_EAST
		tempNode.MoveRobotColor = COLOR_GREEN
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		// if tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Type == TargetType &&
		//     tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Color == TargetColor {
		//     fmt.Println("FIND SOLUTION")
		//     return &tempNode, tempNode.Depth
		// }
		// node.AddChild(tempNode)
		queue.Push(tempNode)
	}
	tempNode = *node
	// if tempNode.GreenRobot.MoveEast(tempNode.GameMap) {
	if tempNode.GreenRobot.MoveEast(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_WEST
		tempNode.MoveRobotColor = COLOR_GREEN
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		// if tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Type == TargetType &&
		//     tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Color == TargetColor {
		//     fmt.Println("FIND SOLUTION")
		//     return &tempNode, tempNode.Depth
		// }
		// node.AddChild(tempNode)
		queue.Push(tempNode)
	}
	tempNode = *node
	// if tempNode.GreenRobot.MoveNorth(tempNode.GameMap) {
	if tempNode.GreenRobot.MoveNorth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_SOUTH
		tempNode.MoveRobotColor = COLOR_GREEN
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		// if tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Type == TargetType &&
		//     tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Color == TargetColor {
		//     fmt.Println("FIND SOLUTION")
		//     return &tempNode, tempNode.Depth
		// }
		// node.AddChild(tempNode)
		queue.Push(tempNode)
	}
	tempNode = *node
	// if tempNode.GreenRobot.MoveSouth(tempNode.GameMap) {
	if tempNode.GreenRobot.MoveSouth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_NORTH
		tempNode.MoveRobotColor = COLOR_GREEN
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		// if tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Type == TargetType &&
		//     tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Color == TargetColor {
		//     fmt.Println("FIND SOLUTION")
		//     return &tempNode, tempNode.Depth
		// }
		// node.AddChild(tempNode)
		queue.Push(tempNode)
	}

	tempNode = *node
	if tempNode.BlueRobot.MoveWest(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_EAST
		tempNode.MoveRobotColor = COLOR_BLUE
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.BlueRobot.MoveEast(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_WEST
		tempNode.MoveRobotColor = COLOR_BLUE
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.BlueRobot.MoveNorth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_SOUTH
		tempNode.MoveRobotColor = COLOR_BLUE
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.BlueRobot.MoveSouth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_NORTH
		tempNode.MoveRobotColor = COLOR_BLUE
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}

	tempNode = *node
	if tempNode.RedRobot.MoveWest(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_EAST
		tempNode.MoveRobotColor = COLOR_RED
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.RedRobot.MoveEast(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_WEST
		tempNode.MoveRobotColor = COLOR_RED
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.RedRobot.MoveNorth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_SOUTH
		tempNode.MoveRobotColor = COLOR_RED
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.RedRobot.MoveSouth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_NORTH
		tempNode.MoveRobotColor = COLOR_RED
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}

	tempNode = *node
	if tempNode.YellowRobot.MoveWest(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_EAST
		tempNode.MoveRobotColor = COLOR_YELLOW
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.YellowRobot.MoveEast(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_WEST
		tempNode.MoveRobotColor = COLOR_YELLOW
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.YellowRobot.MoveNorth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_SOUTH
		tempNode.MoveRobotColor = COLOR_YELLOW
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.YellowRobot.MoveSouth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = DIRECTION_NORTH
		tempNode.MoveRobotColor = COLOR_YELLOW
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	return nil, 0
}

type Queue struct {
	Nodes []Node
}

func (queue *Queue) Push(node Node) {
	queue.Nodes = append(queue.Nodes, node)
}

func (queue *Queue) Pop() *Node {
	var x Node
	x = queue.Nodes[0]
	queue.Nodes = queue.Nodes[1:]
	return &x
}
