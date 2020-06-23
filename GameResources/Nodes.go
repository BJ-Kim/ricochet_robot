package GameResources

import (
	"fmt"
	"github.com/BJ-Kim/ricochet_robot/Constants"
)

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

type Node struct {
	Depth          int
	BlueRobot      Robot
	RedRobot       Robot
	YellowRobot    Robot
	GreenRobot     Robot
	GameMap        [][]GameTile
	ParentNode     *Node
	FromDirection  string
	MoveRobotColor string
}

func (node *Node) IsExistRobot(x int, y int) bool {
	if (node.BlueRobot.XPosition == x && node.BlueRobot.YPosition == y) ||
		(node.RedRobot.XPosition == x && node.RedRobot.YPosition == y) ||
		(node.YellowRobot.XPosition == x && node.YellowRobot.YPosition == y) ||
		(node.GreenRobot.XPosition == x && node.GreenRobot.YPosition == y) {
		return true
	}
	return false
}

func (node *Node) PrintCurrentPosition() {
	fmt.Println(node.Depth)
	fmt.Println("BLUE X:", node.BlueRobot.XPosition, "/ BLUE Y:", node.BlueRobot.YPosition)
	fmt.Println("RED:", node.RedRobot.XPosition, "/ RED Y:", node.RedRobot.YPosition)
	fmt.Println("YELLOW X:", node.YellowRobot.XPosition, "/ YELLOW Y:", node.YellowRobot.YPosition)
	fmt.Println("GREEN X:", node.GreenRobot.XPosition, "/ GREEN Y:", node.GreenRobot.YPosition)
}

func (node *Node) MoveRobots(queue *Queue, count int) (*Node, int) {
	if len(queue.Nodes) != 0 {
		node = queue.Pop()
	}
	tempNode := *node
	if tempNode.GreenRobot.MoveWest(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_EAST
		tempNode.MoveRobotColor = Constants.COLOR_GREEN
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.GreenRobot.MoveEast(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_WEST
		tempNode.MoveRobotColor = Constants.COLOR_GREEN
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.GreenRobot.MoveNorth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_SOUTH
		tempNode.MoveRobotColor = Constants.COLOR_GREEN
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.GreenRobot.MoveSouth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_NORTH
		tempNode.MoveRobotColor = Constants.COLOR_GREEN
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}

	tempNode = *node
	if tempNode.BlueRobot.MoveWest(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_EAST
		tempNode.MoveRobotColor = Constants.COLOR_BLUE
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.BlueRobot.MoveEast(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_WEST
		tempNode.MoveRobotColor = Constants.COLOR_BLUE
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.BlueRobot.MoveNorth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_SOUTH
		tempNode.MoveRobotColor = Constants.COLOR_BLUE
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.BlueRobot.MoveSouth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_NORTH
		tempNode.MoveRobotColor = Constants.COLOR_BLUE
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}

	tempNode = *node
	if tempNode.RedRobot.MoveWest(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_EAST
		tempNode.MoveRobotColor = Constants.COLOR_RED
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.RedRobot.MoveEast(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_WEST
		tempNode.MoveRobotColor = Constants.COLOR_RED
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.RedRobot.MoveNorth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_SOUTH
		tempNode.MoveRobotColor = Constants.COLOR_RED
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.RedRobot.MoveSouth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_NORTH
		tempNode.MoveRobotColor = Constants.COLOR_RED
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}

	tempNode = *node
	if tempNode.YellowRobot.MoveWest(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_EAST
		tempNode.MoveRobotColor = Constants.COLOR_YELLOW
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.YellowRobot.MoveEast(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_WEST
		tempNode.MoveRobotColor = Constants.COLOR_YELLOW
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.YellowRobot.MoveNorth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_SOUTH
		tempNode.MoveRobotColor = Constants.COLOR_YELLOW
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.YellowRobot.MoveSouth(&tempNode) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		tempNode.FromDirection = Constants.DIRECTION_NORTH
		tempNode.MoveRobotColor = Constants.COLOR_YELLOW
		if CheckGoal(tempNode) {
			return &tempNode, tempNode.Depth
		}
		queue.Push(tempNode)
	}
	return nil, 0
}

func CheckGoal(node Node) bool {
	if Constants.TargetColor == Constants.COLOR_GREEN {
		if node.GameMap[node.GreenRobot.XPosition][node.GreenRobot.YPosition].Type == Constants.TargetType &&
			node.GameMap[node.GreenRobot.XPosition][node.GreenRobot.YPosition].Color == Constants.TargetColor {
			return true
		}
	} else if Constants.TargetColor == Constants.COLOR_RED {
		if node.GameMap[node.RedRobot.XPosition][node.RedRobot.YPosition].Type == Constants.TargetType &&
			node.GameMap[node.RedRobot.XPosition][node.RedRobot.YPosition].Color == Constants.TargetColor {
			return true
		}
	} else if Constants.TargetColor == Constants.COLOR_YELLOW {
		if node.GameMap[node.YellowRobot.XPosition][node.YellowRobot.YPosition].Type == Constants.TargetType &&
			node.GameMap[node.YellowRobot.XPosition][node.YellowRobot.YPosition].Color == Constants.TargetColor {
			return true
		}
	} else if Constants.TargetColor == Constants.COLOR_BLUE {
		if node.GameMap[node.BlueRobot.XPosition][node.BlueRobot.YPosition].Type == Constants.TargetType &&
			node.GameMap[node.BlueRobot.XPosition][node.BlueRobot.YPosition].Color == Constants.TargetColor {
			return true
		}
	}

	return false
}
