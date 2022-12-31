package entities

import (
	"reflect"
	"testing"
)

func TestNewWall(t *testing.T) {
	type args struct {
		width  float64
		height float64
	}
	tests := []struct {
		name    string
		args    args
		want    Wall
		wantErr bool
	}{

		{name: "Should_ReturnPassedNewWall_When_ValidParameters", args: args{
			width:  5,
			height: 5,
		}, want: Wall{
			Width:   5,
			Height:  5,
			Doors:   nil,
			Windows: nil,
		},
			wantErr: false},
		{name: "Should_WallWidhtNegativeError_When_NegativeWidhtParameter", args: args{
			width:  -1,
			height: 1,
		}, want: Wall{
			Width:   0.0,
			Height:  0.0,
			Doors:   nil,
			Windows: nil,
		},
			wantErr: true},
		{name: "Should_WallHeightNegativeError_When_NegativeHeightParameter", args: args{
			width:  1,
			height: -1,
		}, want: Wall{
			Width:   0.0,
			Height:  0.0,
			Doors:   nil,
			Windows: nil,
		},
			wantErr: true},

		{name: "Should_MinWallAreaLimitError_When_SmallAreaParameters", args: args{
			width:  0.5,
			height: 0.5,
		}, want: Wall{
			Width:   0.0,
			Height:  0.0,
			Doors:   nil,
			Windows: nil,
		},
			wantErr: true},

		{name: "Should_MaxWallAreaLimitError_When_LargeAreaParameters", args: args{
			width:  10,
			height: 10,
		}, want: Wall{
			Width:   0,
			Height:  0,
			Doors:   nil,
			Windows: nil,
		},
			wantErr: true},

		{name: "Should_WallAreaLimitError_When_ZeroAsParameters", args: args{
			width:  0,
			height: 0,
		}, want: Wall{
			Width:   0,
			Height:  0,
			Doors:   nil,
			Windows: nil,
		},
			wantErr: true},

		{name: "Should_MinWallAreaPaintError_When_AreaParametersBelow", args: args{
			width:  1,
			height: 1,
		}, want: Wall{
			Width:   0,
			Height:  0,
			Doors:   nil,
			Windows: nil,
		},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWall(tt.args.width, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWall() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWall_IsDoorHeightWithMax(t *testing.T) {
	type fields struct {
		Width   float64
		Height  float64
		doors   []Door
		windows []Window
	}
	type args struct {
		door Door
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Shoul_ReturnPassedHeightWall_When_ValidParameters", fields: fields{
			Width:   5,
			Height:  2.20,
			doors:   nil,
			windows: nil,
		}, args: args{door: Door{
			Width:  WidthDoor,
			Height: HeightDoor,
		}}, wantErr: false},
		{name: "Shoul_MaxDoorHeightError_When_HeightWallBelowTheLimit", fields: fields{
			Width:   5,
			Height:  2.19,
			doors:   nil,
			windows: nil,
		}, args: args{door: Door{
			Width:  WidthDoor,
			Height: HeightDoor,
		}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wall{
				Width:   tt.fields.Width,
				Height:  tt.fields.Height,
				Doors:   tt.fields.doors,
				Windows: tt.fields.windows,
			}
			if err := w.IsDoorHeightWithMax(tt.args.door); (err != nil) != tt.wantErr {
				t.Errorf("IsDoorHeightWithMax() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWall_calcArea(t *testing.T) {

	type fields struct {
		Width   float64
		Height  float64
		doors   []Door
		windows []Window
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should_ReturnPassedWallArea_When_ValidParameters",
			fields: fields{
				Width:   5,
				Height:  5,
				doors:   nil,
				windows: nil,
			},
			want: 25,
		},
		{
			name: "Should_ReturnPassedWallArea_When_ZeroParameters",
			fields: fields{
				Width:   0,
				Height:  0,
				doors:   nil,
				windows: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wall{
				Width:   tt.fields.Width,
				Height:  tt.fields.Height,
				Doors:   tt.fields.doors,
				Windows: tt.fields.windows,
			}
			if got := w.calcArea(); got != tt.want {
				t.Errorf("calcArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWall_isWindowsAndDoorsAreaHigherThanWallArea(t *testing.T) {
	type fields struct {
		Width   float64
		Height  float64
		doors   []Door
		windows []Window
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Should_ReturnPassedWindowsAndDoorsArea_When_ValidParameters",
			fields: fields{
				Width:  5,
				Height: 5,
				doors: []Door{{
					Width:  WidthDoor,
					Height: HeightDoor,
				}},
				windows: []Window{{
					Width:  WidthWindow,
					Height: HeightWindow,
				}},
			},
			want: false,
		},
		{
			name: "Should_WindowsAndDoorsAreaHigherThanWallAreaTrue_When_WindowsOverLimitParameters",
			fields: fields{
				Width:  2,
				Height: 2,
				doors: []Door{{
					Width:  WidthDoor,
					Height: HeightDoor,
				}},
				windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			want: true,
		},
		{
			name: "Should_WindowsAndDoorsAreaHigherThanWallAreaTrue_When_DoorsOverLimitParameters",
			fields: fields{
				Width:  2,
				Height: 2,
				doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wall{
				Width:   tt.fields.Width,
				Height:  tt.fields.Height,
				Doors:   tt.fields.doors,
				Windows: tt.fields.windows,
			}
			if got := w.isWindowsAndDoorsAreaHigherThanWallArea(); got != tt.want {
				t.Errorf("isWindowsAndDoorsAreaHigherThanWallArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_AddWall(t *testing.T) {
	type fields struct {
		walls []Wall
	}
	type args struct {
		wall Wall
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{

		{
			name: "Should_ReturnPassedWall_When_WallValidParameters",
			fields: fields{walls: []Wall{
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
			}},
			args: args{Wall{
				Width:   5,
				Height:  5,
				Doors:   nil,
				Windows: nil,
			}},
			wantErr: false,
		},
		{
			name: "Should_WallLimitError_When_OverLimitWalls",
			fields: fields{walls: []Wall{
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
			}},
			args: args{Wall{
				Width:   5,
				Height:  5,
				Doors:   nil,
				Windows: nil,
			}},
			wantErr: true,
		},
		{
			name:    "Should_ReturnPassedWall_When_ZeroParameters",
			fields:  fields{walls: []Wall{}},
			args:    args{Wall{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				Walls: tt.fields.walls,
			}
			if err := r.AddWall(tt.args.wall); (err != nil) != tt.wantErr {
				t.Errorf("AddWall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRoom_HasMoreThanForWalls(t *testing.T) {
	type fields struct {
		walls []Wall
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Should_ReturnPassedLimitWalls_When_ValidQuantityWallsAsParameter",
			fields: fields{walls: []Wall{
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
			}}, want: false},

		{name: "Should_MaximumRoomWallsError_When_InvalidQuantityWallsAsParameter",
			fields: fields{walls: []Wall{
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
				{
					Width:   5,
					Height:  5,
					Doors:   nil,
					Windows: nil,
				},
			}}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				Walls: tt.fields.walls,
			}
			if got := r.HasMoreThanForWalls(); got != tt.want {
				t.Errorf("HasMoreThanForWalls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_calcArea(t *testing.T) {
	type fields struct {
		walls []Wall
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should_ReturnPassedRoomArea_When_ValidParameters",
			fields: fields{
				walls: []Wall{
					{
						Width:   5,
						Height:  5,
						Doors:   nil,
						Windows: nil,
					},
					{
						Width:   5,
						Height:  5,
						Doors:   nil,
						Windows: nil,
					},
					{
						Width:   5,
						Height:  5,
						Doors:   nil,
						Windows: nil,
					},
					{
						Width:   5,
						Height:  5,
						Doors:   nil,
						Windows: nil,
					},
				},
			},
			want: 100,
		},
		{
			name: "Should_ReturnPassedRoomArea_When_ZeroParameters",
			fields: fields{
				walls: []Wall{
					{
						Width:   0,
						Height:  0,
						Doors:   nil,
						Windows: nil,
					},
					{
						Width:   0,
						Height:  0,
						Doors:   nil,
						Windows: nil,
					},
					{
						Width:   0,
						Height:  0,
						Doors:   nil,
						Windows: nil,
					},
					{
						Width:   0,
						Height:  0,
						Doors:   nil,
						Windows: nil,
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Room{
				Walls: tt.fields.walls,
			}
			if got := r.calcArea(); got != tt.want {
				t.Errorf("calcArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWall_ValidateDoors(t *testing.T) {
	type fields struct {
		Width   float64
		Height  float64
		Doors   []Door
		Windows []Window
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "Should_ReturnPassedDoor_When_ValidParameters",
			fields: fields{
				Width:  2,
				Height: 4,
				Doors: []Door{{
					Width:  WidthDoor,
					Height: HeightDoor,
				}},
				Windows: nil,
			},
			wantErr: false,
		},
		{
			name: "Should_MaxDoorHeightError_When_DoorsOverLimitHeight",
			fields: fields{
				Width:  2,
				Height: 1,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: nil,
			},
			wantErr: true,
		},
		{
			name: "Should_DoorsAreaInWallError_When_DoorsOverLimit",
			fields: fields{
				Width:  1,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: nil,
			},
			wantErr: true,
		},
		{
			name: "Should_ReturnPassedDoor_When_ZeroDoorsParameters",
			fields: fields{
				Width:   2,
				Height:  2,
				Doors:   nil,
				Windows: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wall{
				Width:   tt.fields.Width,
				Height:  tt.fields.Height,
				Doors:   tt.fields.Doors,
				Windows: tt.fields.Windows,
			}
			if err := w.ValidateDoors(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateDoors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoor_calcArea(t *testing.T) {

	doorAreaResult := WidthDoor * HeightDoor

	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{

		{name: "Should_ReturnPassedDoorArea_When_ValidParameters", fields: fields{
			Width:  WidthDoor,
			Height: HeightDoor,
		}, want: doorAreaResult},

		{name: "Should_ReturnPassedDoorArea_When_ZeroParameters", fields: fields{
			Width:  0,
			Height: 0,
		}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Door{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := d.calcArea(); got != tt.want {
				t.Errorf("calcArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWall_ValidateWindow(t *testing.T) {
	type fields struct {
		Width   float64
		Height  float64
		Doors   []Door
		Windows []Window
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Should_ReturnPassedWindow_When_ValidParameters",
			fields: fields{
				Width:  2,
				Height: 5,
				Doors:  nil,
				Windows: []Window{{
					Width:  WidthWindow,
					Height: HeightWindow,
				}},
			},
			wantErr: false,
		},
		{
			name: "Should_WindowsAreaInWallError_When_WindowsOverLimitParameters",
			fields: fields{
				Width:  2,
				Height: 3,
				Doors:  nil,
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should_ReturnPassedWindow_When_ZeroWindowsParameters",
			fields: fields{
				Width:   2,
				Height:  2,
				Doors:   nil,
				Windows: []Window{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wall{
				Width:   tt.fields.Width,
				Height:  tt.fields.Height,
				Doors:   tt.fields.Doors,
				Windows: tt.fields.Windows,
			}
			if err := w.ValidateWindow(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateWindow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWindow_calcArea(t *testing.T) {

	windowAreaResult := WidthWindow * HeightWindow

	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{

		{name: "Should_ReturnPassedWindowArea_When_ValidParameters", fields: fields{
			Width:  WidthWindow,
			Height: HeightWindow,
		}, want: windowAreaResult},

		{name: "Should_ReturnPassedWindowArea_When_ZeroParameters", fields: fields{
			Width:  0,
			Height: 0,
		}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Window{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := w.calcArea(); got != tt.want {
				t.Errorf("calcArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcLitersPerMeterPainted(t *testing.T) {

	roomArea := 10.0
	roomAreaLitersResult := roomArea / metersPaintedPerLiter
	type args struct {
		roomArea float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Should_ReturnPassedLitersPerMeterPainted_When_ValidParameters", args: args{roomArea}, want: roomAreaLitersResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcLitersPerMeterPainted(tt.args.roomArea); got != tt.want {
				t.Errorf("calcLitersPerMeterPainted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaintBudgetCalculator_CalculatePaintBudget(t *testing.T) {
	type args struct {
		room Room
	}
	tests := []struct {
		name string
		args args
		want []Can
	}{
		{name: "Should_ReturnPassedPainBudget_When_1WallParameters", args: args{room: Room{Walls: []Wall{
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
		}}}, want: []Can{BigCan, SmallCan, SmallCan}},

		{name: "Should_ReturnPassedPainBudget_When_2WallParameters", args: args{room: Room{Walls: []Wall{
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
		}}}, want: []Can{BigCan, BigCan, SmallCan, SmallCan, SmallCan}},
		{name: "Should_ReturnPassedPainBudget_When_3WallParameters", args: args{room: Room{Walls: []Wall{
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
		}}}, want: []Can{BigCan, BigCan, BigCan, SmallCan, SmallCan, SmallCan, SmallCan}},
		{name: "Should_ReturnPassedPainBudget_When_4WallParameters", args: args{room: Room{Walls: []Wall{
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
			{
				Width:  5,
				Height: 5,
				Doors: []Door{
					{
						Width:  WidthDoor,
						Height: HeightDoor,
					},
				},
				Windows: []Window{
					{
						Width:  WidthWindow,
						Height: HeightWindow,
					},
				},
			},
		}}}, want: []Can{BigCan, BigCan, BigCan, BigCan, SmallCan, SmallCan, SmallCan, SmallCan, SmallCan}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaintBudgetCalculator{}
			if got := p.CalculatePaintBudget(tt.args.room); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePaintBudget() = %v, want %v", got, tt.want)
			}
		})
	}
}
