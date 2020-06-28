package GameResources

import (
	"fmt"
	"github.com/BJ-Kim/ricochet_robot/Constants"
	// "github.com/BJ-Kim/ricochet_robot/GameUI"
	"sync"
	"time"
)

// var CURR_NODE Node
// var RESULT_NODE *Node

type nodes struct {
	CurrNode   Node
	ResultNode *Node
}

var instance *nodes
var once sync.Once

func GetNodesSingleton() *nodes {
	once.Do(func() {
		instance = &nodes{}
	})
	return instance
}

func (nds *nodes) SetCurrNode(node Node) {
	nds.CurrNode = node
}

func (nds *nodes) SetResultNode(node *Node) {
	nds.ResultNode = node
}

type Queue struct {
	Nodes []Node
	count int
}

func (queue *Queue) Push(node Node) {
	queue.Nodes = append(queue.Nodes[:queue.count], node)
	queue.count++
}

func (queue *Queue) Pop() *Node {
	var x Node
	x = queue.Nodes[0]
	queue.Nodes = queue.Nodes[1:]
	queue.count--
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

func (node *Node) FindPath(queue *Queue) *Node {
	startTime := time.Now()
	var nd *Node
	var count = 0
	// var wg sync.WaitGroup
	// messages := make(chan *GameResources.Node)
	for {
		// wg.Add(1)

		// go func() {
		//     defer wg.Done()
		//     messages <- startNode.MoveRobots(&queue, count)
		// }()

		findNode := node.MoveRobots(queue, count)
		if findNode != nil {
			nd = findNode
			break
		}
		count++
	}
	endTime := time.Now()
	fmt.Println("################################")
	fmt.Println(endTime.Sub(startTime))
	fmt.Println("DEPTH : ", nd.Depth)
	fmt.Println("################################")
	return nd
}

// func (node *Node) ShowResult() {
//     arr := []*Node{}
//     nn := node
//     for {
//         fmt.Println("------------------HISTORY-----------------")
//         nn.PrintCurrentPosition()
//         if nn.ParentNode == nil {
//             fmt.Println("PARENT NIL")
//             break
//         }
//         nn = nn.ParentNode
//         arr = append(arr, nn)
//     }
//
//     for _, v := range arr {
//         time.Sleep(1 * time.Second)
//         v.ShowNodeState()
//     }
// }
//
// func (node *Node) ShowNodeState() {
//     var GameDataInstance = GameUI.GetGameMapDataSingleton()
//     GameDataInstance.GameMap.Image(
//         float64(node.RedRobot.YPosition*Constants.TILE_SIZE)+4,
//         float64(node.RedRobot.XPosition*Constants.TILE_SIZE)+4,
//         Constants.TILE_SIZE-8, Constants.TILE_SIZE-8,
//         "./Images/robot_r.png",
//     )
//     GameDataInstance.GameMap.Image(
//         float64(node.BlueRobot.YPosition*Constants.TILE_SIZE)+4,
//         float64(node.BlueRobot.XPosition*Constants.TILE_SIZE)+4,
//         Constants.TILE_SIZE-8, Constants.TILE_SIZE-8,
//         "./Images/robot_b.png",
//     )
//     GameDataInstance.GameMap.Image(
//         float64(node.YellowRobot.YPosition*Constants.TILE_SIZE)+4,
//         float64(node.YellowRobot.XPosition*Constants.TILE_SIZE)+4,
//         Constants.TILE_SIZE-8, Constants.TILE_SIZE-8,
//         "./Images/robot_y.png",
//     )
//     GameDataInstance.GameMap.Image(
//         float64(node.GreenRobot.YPosition*Constants.TILE_SIZE)+4,
//         float64(node.GreenRobot.XPosition*Constants.TILE_SIZE)+4,
//         Constants.TILE_SIZE-8, Constants.TILE_SIZE-8,
//         "./Images/robot_g.png",
//     )
// }
//
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

func (node *Node) MoveRobotsGoroutine(queue *Queue, count int, wg *sync.WaitGroup) {
	defer wg.Done()
}

func (node *Node) MoveRobots(queue *Queue, count int) *Node {
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
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
			return &tempNode
		}
		queue.Push(tempNode)
	}
	return nil
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
