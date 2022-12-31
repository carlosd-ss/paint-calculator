package paint

import (
	"digitalrepublic/pkg/entities"
	"errors"
)

const (
	ExtraLargeCan = 18
	LargeCan      = 3.6
	MediumCan     = 2.5
	SmallCan      = 0.5
)

const (
	negativeWindowError = "a quantidade de janelas não pode ser menor do que zero"
	negativeDoorError   = "a quantidade de portas não pode ser menor do que zero"
	wallZeroError       = "é necessario pelo menos 1 parede"
)

type WallInput struct {
	Width          float64 `json:"width"`
	Height         float64 `json:"height"`
	DoorQuantity   int     `json:"door_quantity"`
	WindowQuantity int     `json:"window_quantity"`
}

type CalculateRoomPaintInCansOutput struct {
	ExtraLargeCan int64 `json:"huge_can"`
	LargeCan      int64 `json:"big_can"`
	MediumCan     int64 `json:"medium_can"`
	SmallCan      int64 `json:"small_can"`
}

type CalculateRoomPaintInCansInput struct {
	Walls []WallInput `json:"walls"`
}

type CalculateRoomPaintInCans interface {
	Execute(room CalculateRoomPaintInCansInput) (*CalculateRoomPaintInCansOutput, error)
}

type calculateRoomPaintInCans struct {
}

func NewCalculateRoomPaintInCans() CalculateRoomPaintInCans {
	return &calculateRoomPaintInCans{}
}

func IsDoorNegative(door int) error {
	if door < 0 {
		return errors.New(negativeDoorError)
	}
	return nil
}

func addWindowsAndDoorsToWall(room *entities.Room, input CalculateRoomPaintInCansInput) error {

	for in, wallInput := range input.Walls {

		err := addDoorsToWall(&room.Walls[in], wallInput)
		if err != nil {

			return err
		}

		err = addWindowsToWall(&room.Walls[in], wallInput)
		if err != nil {
			return err
		}

	}

	return nil
}
func addDoorsToWall(wall *entities.Wall, input WallInput) error {

	err := IsDoorNegative(input.DoorQuantity)
	if err != nil {
		return err
	}

	for i := 0; i < input.DoorQuantity; i++ {

		doors := entities.Door{
			Width:  entities.WidthDoor,
			Height: entities.HeightDoor}

		wall.Doors = append(wall.Doors, doors)

	}
	err = wall.ValidateDoors()
	if err != nil {
		return err
	}

	return nil
}
func IsWindowNegative(window int) error {
	if window < 0 {
		return errors.New(negativeWindowError)
	}
	return nil
}

func addWindowsToWall(wall *entities.Wall, input WallInput) error {

	err := IsWindowNegative(input.WindowQuantity)
	if err != nil {
		return err
	}

	for i := 0; i < input.WindowQuantity; i++ {

		doors := entities.Window{
			Width:  entities.WidthWindow,
			Height: entities.HeightWindow}

		wall.Windows = append(wall.Windows, doors)
	}
	err = wall.ValidateWindow()
	if err != nil {
		return err
	}

	return nil
}
func addWallsToRoom(room *entities.Room, input CalculateRoomPaintInCansInput) error {

	if len(input.Walls) == 0 {
		return errors.New(wallZeroError)
	}

	for _, wallInput := range input.Walls {

		wall, err := entities.NewWall(wallInput.Width, wallInput.Height)

		if err != nil {
			return err
		}

		err = room.AddWall(wall)
		if err != nil {
			return err
		}

	}
	err := addWindowsAndDoorsToWall(room, input)
	if err != nil {
		return err
	}

	return nil
}

func formatOutput(cans []entities.Can) CalculateRoomPaintInCansOutput {
	c := CalculateRoomPaintInCansOutput{}
	for _, can := range cans {
		switch can {
		case ExtraLargeCan:
			c.ExtraLargeCan += 1

		case LargeCan:
			c.LargeCan += 1

		case MediumCan:
			c.MediumCan += 1

		case SmallCan:
			c.SmallCan += 1
		}

	}
	return c
}

func (i *calculateRoomPaintInCans) Execute(input CalculateRoomPaintInCansInput) (*CalculateRoomPaintInCansOutput, error) {

	room := entities.Room{}
	err := addWallsToRoom(&room, input)
	if err != nil {
		return nil, err
	}
	paintBudgetCalculator := entities.PaintBudgetCalculator{}
	cans := paintBudgetCalculator.CalculatePaintBudget(room)
	c := formatOutput(cans)
	return &c, nil
}
