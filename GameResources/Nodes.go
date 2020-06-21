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

var TargetColor = COLOR_GREEN
var TargetType = TYPE_MOON

type Node struct {
	Depth       int
	BlueRobot   Robot
	RedRobot    Robot
	YellowRobot Robot
	GreenRobot  Robot
	GameMap     [][]GameTile
	ParentNode  *Node
	Child       []Node
}

func (node *Node) AddChild(child Node) {
	node.Child = append(node.Child, child)
}

func (node *Node) MoveRobots(queue *Queue) (*Node, int) {
	tempNode := *node
	if tempNode.GreenRobot.MoveWest(tempNode.GameMap) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		if tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Type == TargetType &&
			tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Color == TargetColor {
			fmt.Println("FIND SOLUTION")
			return &tempNode, tempNode.Depth
		}
		node.AddChild(tempNode)
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.GreenRobot.MoveEast(tempNode.GameMap) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		if tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Type == TargetType &&
			tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Color == TargetColor {
			fmt.Println("FIND SOLUTION")
			return &tempNode, tempNode.Depth
		}
		node.AddChild(tempNode)
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.GreenRobot.MoveNorth(tempNode.GameMap) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		if tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Type == TargetType &&
			tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Color == TargetColor {
			fmt.Println("FIND SOLUTION")
			return &tempNode, tempNode.Depth
		}
		node.AddChild(tempNode)
		queue.Push(tempNode)
	}
	tempNode = *node
	if tempNode.GreenRobot.MoveSouth(tempNode.GameMap) {
		tempNode.Depth = node.Depth + 1
		tempNode.ParentNode = node
		if tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Type == TargetType &&
			tempNode.GameMap[tempNode.GreenRobot.XPosition][tempNode.GreenRobot.YPosition].Color == TargetColor {
			fmt.Println("FIND SOLUTION")
			return &tempNode, tempNode.Depth
		}
		node.AddChild(tempNode)
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

func (queue *Queue) Pop() Node {
	var x Node
	x = queue.Nodes[0]
	queue.Nodes = queue.Nodes[1:]
	return x
}
