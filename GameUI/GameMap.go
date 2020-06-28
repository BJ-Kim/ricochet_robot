package GameUI

import (
	// "fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/BJ-Kim/ricochet_robot/Constants"
	"github.com/BJ-Kim/ricochet_robot/GameResources"
	"sync"
)

type gameMapData struct {
	NewApp       fyne.App
	GameWindow   fyne.Window
	AppContainer *widget.Box
	GameMap      Canvas
}

var instance *gameMapData
var once sync.Once

func GetGameMapDataSingleton() *gameMapData {
	once.Do(func() {
		instance = &gameMapData{}
	})
	return instance
}

var TILE_SIZE = 30

func SetGameMap(newMap [][]GameResources.GameTile, cv Canvas) {
	// cv := NewCanvas("HELL", TILE_SIZE*len(newMap), TILE_SIZE*len(newMap))

	for i, arr := range newMap {
		for j, t := range arr {
			x := float64(j * TILE_SIZE)
			y := float64(i * TILE_SIZE)
			if t.North == true && t.West == true {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_west_north_wall.png")
			} else if t.North == true && t.East == true {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_north_east_wall.png")
			} else if t.South == true && t.East == true {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_east_south_wall.png")
			} else if t.South == true && t.West == true {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_west_south_wall.png")
			} else if t.South == true {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_south_wall.png")
			} else if t.North == true {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_north_wall.png")
			} else if t.East == true {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_east_wall.png")
			} else if t.West == true {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_west_wall.png")
			} else {
				cv.Image(x, y, TILE_SIZE, TILE_SIZE, "./Images/tile_empty.png")
			}
			fileName := ""
			if t.Color == Constants.COLOR_BLUE {
				fileName += "blue_"
			} else if t.Color == Constants.COLOR_GREEN {
				fileName += "green_"
			} else if t.Color == Constants.COLOR_RED {
				fileName += "red_"
			} else if t.Color == Constants.COLOR_YELLOW {
				fileName += "yellow_"
			}
			if t.Type == Constants.TYPE_MOON {
				fileName += "moon.png"
			} else if t.Type == Constants.TYPE_PLANET {
				fileName += "planet.png"
			} else if t.Type == Constants.TYPE_STAR {
				fileName += "star.png"
			} else if t.Type == Constants.TYPE_TRIANGLE {
				fileName += "triangle.png"
			}
			if len(fileName) > 5 {
				cv.Image(x+4, y+4, TILE_SIZE-8, TILE_SIZE-8, "./Images/"+fileName)
			}
		}
	}
	// container.Append(cv.Container)
}
