package paint

import (
	"digitalrepublic/pkg/entities"
	"reflect"
	"testing"
)

func TestIsDoorNegative(t *testing.T) {
	type args struct {
		door int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Should_ReturnPassedDoor_When_DoorValidParameter",
			args:    args{door: 1},
			wantErr: false,
		},
		{
			name:    "Should_NegativeDoorError_When_DoorNegativeParameter",
			args:    args{door: -1},
			wantErr: true,
		},
		{
			name:    "Should_ReturnPassedDoor_When_ZeroDoorParameter",
			args:    args{door: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsDoorNegative(tt.args.door); (err != nil) != tt.wantErr {
				t.Errorf("IsDoorNegative() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsWindowNegative(t *testing.T) {
	type args struct {
		window int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Should_ReturnPassedWindow_When_WindowValidParameter",
			args:    args{window: 1},
			wantErr: false,
		},
		{
			name:    "Should_NegativeWindowError_When_WindowNegativeParameter",
			args:    args{window: -1},
			wantErr: true,
		},
		{
			name:    "Should_ReturnPassedWindow_When_ZeroWindowParameter",
			args:    args{window: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsWindowNegative(tt.args.window); (err != nil) != tt.wantErr {
				t.Errorf("IsWindowNegative() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_addDoorsToWall(t *testing.T) {
	type args struct {
		wall  *entities.Wall
		input WallInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should_ReturnPassedDoors_When_ValidParameters",
			args: args{
				wall: &entities.Wall{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				input: WallInput{
					Width:          5,
					Height:         5,
					DoorQuantity:   2,
					WindowQuantity: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Should_ReturnPassedDoors_When_ZeroParameters",
			args: args{
				wall: &entities.Wall{
					Width:   0,
					Height:  0,
					Doors:   nil,
					Windows: nil,
				},
				input: WallInput{
					Width:          0,
					Height:         0,
					DoorQuantity:   0,
					WindowQuantity: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Should_DoorsAreaInWallError_When_DoorsOverLimit",
			args: args{
				wall: &entities.Wall{
					Width:   3,
					Height:  3,
					Doors:   nil,
					Windows: nil,
				},
				input: WallInput{
					Width:          5,
					Height:         5,
					DoorQuantity:   3,
					WindowQuantity: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "Should_MaxDoorHeightError_When_DoorsOverLimitHeight",
			args: args{
				wall: &entities.Wall{
					Width:   5,
					Height:  2,
					Doors:   nil,
					Windows: nil,
				},
				input: WallInput{
					Width:          5,
					Height:         5,
					DoorQuantity:   2,
					WindowQuantity: 0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addDoorsToWall(tt.args.wall, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("addDoorsToWall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_addWindowsToWall(t *testing.T) {
	type args struct {
		wall  *entities.Wall
		input WallInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should_ReturnPassedWindows_When_ValidParameters",
			args: args{
				wall: &entities.Wall{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				input: WallInput{
					Width:          5,
					Height:         5,
					DoorQuantity:   0,
					WindowQuantity: 2,
				},
			},
			wantErr: false,
		},
		{
			name: "Should_ReturnPassedWindows_When_ZeroParameters",
			args: args{
				wall: &entities.Wall{
					Width:   0,
					Height:  0,
					Doors:   nil,
					Windows: nil,
				},
				input: WallInput{
					Width:          0,
					Height:         0,
					DoorQuantity:   0,
					WindowQuantity: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Should_WindowsAreaInWallError_When_WindowsOverLimit",
			args: args{
				wall: &entities.Wall{
					Width:   3,
					Height:  3,
					Doors:   nil,
					Windows: nil,
				},
				input: WallInput{
					Width:          5,
					Height:         5,
					DoorQuantity:   0,
					WindowQuantity: 3,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addWindowsToWall(tt.args.wall, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("addWindowsToWall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_addWallsToRoom(t *testing.T) {
	type args struct {
		room  *entities.Room
		input CalculateRoomPaintInCansInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should_ReturnPassedWall_When_ValidParameters",
			args: args{
				room: &entities.Room{Walls: []entities.Wall{
					{Width: 5, Height: 5, Doors: nil, Windows: nil},
				}},
				input: CalculateRoomPaintInCansInput{
					[]WallInput{{
						Width:          5,
						Height:         5,
						DoorQuantity:   0,
						WindowQuantity: 0,
					}},
				},
			},
			wantErr: false,
		},
		{
			name: "Should_WallZeroError_When_ZeroParameters",
			args: args{
				room:  nil,
				input: CalculateRoomPaintInCansInput{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addWallsToRoom(tt.args.room, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("addWallsToRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_formatOutput(t *testing.T) {
	type args struct {
		cans []entities.Can
	}
	tests := []struct {
		name string
		args args
		want CalculateRoomPaintInCansOutput
	}{
		{
			name: "Should_ReturnFormatedRoomPaintInCansOutput_When_ValidParameters",
			args: args{cans: []entities.Can{
				ExtraLargeCan, LargeCan, MediumCan, SmallCan,
			}},
			want: CalculateRoomPaintInCansOutput{
				ExtraLargeCan: 1,
				LargeCan:      1,
				MediumCan:     1,
				SmallCan:      1,
			},
		},
		{
			name: "Should_ReturnFormatedRoomPaintInCansOutput_WhenZeroParameter",
			args: args{cans: []entities.Can{}},
			want: CalculateRoomPaintInCansOutput{
				ExtraLargeCan: 0,
				LargeCan:      0,
				MediumCan:     0,
				SmallCan:      0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatOutput(tt.args.cans); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
