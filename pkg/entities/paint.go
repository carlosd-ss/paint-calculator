package entities

import (
	"errors"
	"math"
)

const (
	maximumRoomWalls      = 4
	minimumRoomWallsArea  = 1.0
	maximumRoomWallsArea  = 50.0
	metersPaintedPerLiter = 5
	maxDoorHeight         = 0.3
	limitWindowAndDoor    = 0.5
	minimumWallAreaPaint  = 0.5
)

const (
	WidthWindow  = 2.0
	HeightWindow = 1.2

	WidthDoor  = 0.8
	HeightDoor = 1.9
)

const (
	doorsAndWindowsAreaInWallError = "a área total de janelas e portas, em metros quadrados, não deve ultrapassar 50% do total da área da parede"
	wallLimitError                 = "não possivel ter mais que 4 paredes"
	wallAreaLimitError             = "tamanho da parede invalido: A parede precisa possuir entre 1 e 50 metros quadrados"
	wallWidhtNegativeError         = "tamanho da parede invalido: A largura da parede não pode ser menor que 0"
	wallHeightNegativeError        = "tamanho da parede invalido: A altura da parede não pode ser menor que 0"
	maxDoorHeightError             = "a altura mínima da parede deve ser 30 centímetros a mais do que a altura da porta"
	minWallAreaPaintError          = "a área minima da parede deve corresponder ao menor tamanho da tinta 0.5L"
)

type Can float64

const (
	HugeCan   Can = 18.0
	BigCan    Can = 3.6
	MediumCan Can = 2.5
	SmallCan  Can = 0.5
)

type Dimensions interface {
	calcArea() float64
}

type Wall struct {
	Width   float64
	Height  float64
	Doors   []Door
	Windows []Window
}

type Door struct {
	Width  float64
	Height float64
}

type Window struct {
	Width  float64
	Height float64
}

type PaintBudgetCalculator struct {
}

func NewWall(width, height float64) (Wall, error) {
	totalAreaInSquareMeters := width * height
	switch {
	case width < 0:
		return Wall{}, errors.New(wallWidhtNegativeError)

	case height < 0:
		return Wall{}, errors.New(wallHeightNegativeError)

	case totalAreaInSquareMeters < minimumRoomWallsArea:
		return Wall{}, errors.New(wallAreaLimitError)

	case totalAreaInSquareMeters > maximumRoomWallsArea:
		return Wall{}, errors.New(wallAreaLimitError)

	case calcLitersPerMeterPainted(totalAreaInSquareMeters) < minimumWallAreaPaint:
		return Wall{}, errors.New(minWallAreaPaintError)

	}

	wall := Wall{Width: width, Height: height}
	return wall, nil

}

func (w *Wall) calcArea() float64 {
	doorsArea := 0.0
	windowsArea := 0.0

	for _, door := range w.Doors {
		doorsArea += door.calcArea()
	}
	for _, window := range w.Windows {
		windowsArea += window.calcArea()
	}

	return (w.Width * w.Height) - (doorsArea + windowsArea)
}

func (d *Door) calcArea() float64 {
	return d.Width * d.Height
}

func (w *Wall) ValidateDoors() error {

	for _, doorActual := range w.Doors {

		err := w.IsDoorHeightWithMax(doorActual)
		if err != nil {
			return err
		}

	}

	if w.isWindowsAndDoorsAreaHigherThanWallArea() {
		return errors.New(doorsAndWindowsAreaInWallError)
	}
	return nil
}

func (w *Wall) IsDoorHeightWithMax(door Door) error {

	if w.Height-door.Height < maxDoorHeight {
		return errors.New(maxDoorHeightError)

	}
	return nil
}

func (w *Wall) ValidateWindow() error {

	if w.isWindowsAndDoorsAreaHigherThanWallArea() {
		return errors.New(doorsAndWindowsAreaInWallError)
	}
	return nil
}

func (w *Window) calcArea() float64 {
	return w.Width * w.Height
}

type Room struct {
	Walls []Wall
}

func (r *Room) AddWall(wall Wall) error {

	r.Walls = append(r.Walls, wall)
	if r.HasMoreThanForWalls() {
		return errors.New(wallLimitError)
	}
	return nil
}

func (r *Room) HasMoreThanForWalls() bool {

	return len(r.Walls) > maximumRoomWalls

}
func (r *Room) calcArea() float64 {

	area := 0.0
	for _, wall := range r.Walls {
		area += wall.calcArea()
	}

	return area

}

func (p *PaintBudgetCalculator) CalculatePaintBudget(room Room) []Can {
	paintCans := []Can{}
	Cans := [4]Can{
		HugeCan,
		BigCan,
		MediumCan,
		SmallCan,
	}
	numberOfCans := 0.0

	roomArea := calcLitersPerMeterPainted(room.calcArea())

	for can := 0; can < len(Cans); can++ {
		if roomArea >= float64(Cans[can]) {

			if can != len(Cans)-1 {
				numberOfCans = math.Floor(roomArea / float64(Cans[can]))
				for i := 0; i < int(numberOfCans); i++ {
					paintCans = append(paintCans, Cans[can])

					roomArea = math.Mod(roomArea, float64(Cans[can]))
				}

			} else {
				numberOfCans = math.Ceil(roomArea / float64(Cans[can]))
				for i := 0; i < int(numberOfCans); i++ {
					paintCans = append(paintCans, Cans[can])

				}
			}
		}
	}

	return paintCans
}

func (w *Wall) isWindowsAndDoorsAreaHigherThanWallArea() bool {

	wallArea := w.calcArea()

	doorsArea := 0.0
	windowsArea := 0.0

	for _, door := range w.Doors {
		doorsArea += door.calcArea()

	}
	for _, window := range w.Windows {
		windowsArea += window.calcArea()

	}
	totalWallArea := wallArea + doorsArea + windowsArea
	windowsAndDoorsArea := windowsArea + doorsArea

	return windowsAndDoorsArea > limitWindowAndDoor*totalWallArea

}

func calcLitersPerMeterPainted(roomArea float64) float64 {
	return roomArea / metersPaintedPerLiter

}
