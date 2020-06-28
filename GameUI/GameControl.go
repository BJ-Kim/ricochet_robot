package GameUI

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/BJ-Kim/ricochet_robot/Constants"
	"github.com/BJ-Kim/ricochet_robot/GameResources"
	"time"
)

func SetControlPannel(container *widget.Box, queue *GameResources.Queue) {
	nodesInstance := GameResources.GetNodesSingleton()
	resetBtn := widget.NewButton("RESET Robot Position", func() {
		fmt.Println("reset robot position")
	})
	aiBtn := widget.NewButton("Find Path", func() {
		fmt.Println("find button")
		nodesInstance.SetResultNode(nodesInstance.CurrNode.FindPath(queue))
	})
	showResultBtn := widget.NewButton("Show Result", func() {
		fmt.Println("show result button")
		// nodesInstance.ResultNode.ShowResult()
		ShowResult(nodesInstance.ResultNode)
	})
	controllPannel := widget.NewVBox(resetBtn, aiBtn, showResultBtn)
	controllPannel.Resize(fyne.Size{Width: 500, Height: 500})

	container.Append(controllPannel)
}

func ShowResult(node *GameResources.Node) {
	arr := []*GameResources.Node{}
	nn := node
	for {
		fmt.Println("------------------HISTORY-----------------")
		nn.PrintCurrentPosition()
		if nn.ParentNode == nil {
			fmt.Println("PARENT NIL")
			break
		}
		nn = nn.ParentNode
		arr = append(arr, nn)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		ShowNodeState(arr[i])
	}
	// for _, v := range arr {
	//     time.Sleep(1 * time.Second)
	// }
}

func ShowNodeState(node *GameResources.Node) {
	var GameDataInstance = GetGameMapDataSingleton()
	GameDataInstance.GameMap.Container.Refresh()
	GameDataInstance.GameMap.Image(
		float64(node.RedRobot.YPosition*Constants.TILE_SIZE)+4,
		float64(node.RedRobot.XPosition*Constants.TILE_SIZE)+4,
		Constants.TILE_SIZE-8, Constants.TILE_SIZE-8,
		"./Images/robot_r.png",
	)
	GameDataInstance.GameMap.Image(
		float64(node.BlueRobot.YPosition*Constants.TILE_SIZE)+4,
		float64(node.BlueRobot.XPosition*Constants.TILE_SIZE)+4,
		Constants.TILE_SIZE-8, Constants.TILE_SIZE-8,
		"./Images/robot_b.png",
	)
	GameDataInstance.GameMap.Image(
		float64(node.YellowRobot.YPosition*Constants.TILE_SIZE)+4,
		float64(node.YellowRobot.XPosition*Constants.TILE_SIZE)+4,
		Constants.TILE_SIZE-8, Constants.TILE_SIZE-8,
		"./Images/robot_y.png",
	)
	GameDataInstance.GameMap.Image(
		float64(node.GreenRobot.YPosition*Constants.TILE_SIZE)+4,
		float64(node.GreenRobot.XPosition*Constants.TILE_SIZE)+4,
		Constants.TILE_SIZE-8, Constants.TILE_SIZE-8,
		"./Images/robot_g.png",
	)
}
