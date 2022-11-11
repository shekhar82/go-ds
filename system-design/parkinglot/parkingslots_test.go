package parkinglot

import (
	"container/list"
	"reflect"
	"testing"
	"time"
)

func TestSize_String(t *testing.T) {
	tests := []struct {
		name string
		size Size
		want string
	}{
		{"Big", Big, "B"},
		{"Medium", Medium, "M"},
		{"Small", Small, "S"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.size.String(); got != tt.want {
				t.Errorf("Size.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVehicle_IsValidVehicle(t *testing.T) {
	type fields struct {
		NumberPlate string
		Size        Size
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Invalid vehicle test case", fields{NumberPlate: "", Size: Big}, false},
		{"Valid vehicle test case", fields{NumberPlate: "KA01MV5098", Size: Small}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vehicle{
				NumberPlate: tt.fields.NumberPlate,
				Size:        tt.fields.Size,
			}
			if got := v.IsValidVehicle(); got != tt.want {
				t.Errorf("Vehicle.IsValidVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateParkingSlot(t *testing.T) {
	type args struct {
		id   string
		size Size
	}
	tests := []struct {
		name string
		args args
		want *ParkingSlot
	}{
		{"test for big parking slot creation", args{id: "B1", size: Big}, &ParkingSlot{ID: "B1", Size: Big}},
		{"test for medium parking slot creation", args{id: "M1", size: Medium}, &ParkingSlot{ID: "M1", Size: Medium}},
		{"test for small parking slot creation", args{id: "S1", size: Small}, &ParkingSlot{ID: "S1", Size: Small}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateParkingSlot(tt.args.id, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParkingSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSlot_IsFree(t *testing.T) {
	type fields struct {
		ID              string
		Size            Size
		AssignedVehicle Vehicle
		AssignedTime    time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"slot is free", fields{ID: "M1", Size: Big, AssignedVehicle: Vehicle{}, AssignedTime: time.Now()}, true},
		{"slot is not free", fields{ID: "M1", Size: Big, AssignedVehicle: Vehicle{NumberPlate: "MV3456", Size: Big}, AssignedTime: time.Now()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot := &ParkingSlot{
				ID:              tt.fields.ID,
				Size:            tt.fields.Size,
				AssignedVehicle: tt.fields.AssignedVehicle,
				AssignedTime:    tt.fields.AssignedTime,
			}
			if got := slot.IsFree(); got != tt.want {
				t.Errorf("ParkingSlot.IsFree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSlot_AssignVehicle(t *testing.T) {
	type fields struct {
		ID              string
		Size            Size
		AssignedVehicle Vehicle
		AssignedTime    time.Time
	}
	type args struct {
		vehicle Vehicle
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot := &ParkingSlot{
				ID:              tt.fields.ID,
				Size:            tt.fields.Size,
				AssignedVehicle: tt.fields.AssignedVehicle,
				AssignedTime:    tt.fields.AssignedTime,
			}
			if err := slot.AssignVehicle(tt.args.vehicle); (err != nil) != tt.wantErr {
				t.Errorf("ParkingSlot.AssignVehicle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParkingSlot_FreeUpSlot(t *testing.T) {
	type fields struct {
		ID              string
		Size            Size
		AssignedVehicle Vehicle
		AssignedTime    time.Time
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot := &ParkingSlot{
				ID:              tt.fields.ID,
				Size:            tt.fields.Size,
				AssignedVehicle: tt.fields.AssignedVehicle,
				AssignedTime:    tt.fields.AssignedTime,
			}
			slot.FreeUpSlot()
		})
	}
}

func TestCreateParkingSlots(t *testing.T) {
	type args struct {
		totalSlots int
		slotSize   Size
	}
	tests := []struct {
		name string
		args args
		want *ParkingSlots
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateParkingSlots(tt.args.totalSlots, tt.args.slotSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParkingSlots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSlots_AssignParkingSlot(t *testing.T) {
	type fields struct {
		Size           Size
		MaxSlots       int
		AvailableSlots int
		FreeSlots      *list.List
		Slots          map[string]*ParkingSlot
	}
	type args struct {
		v Vehicle
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pSlots := &ParkingSlots{
				Size:           tt.fields.Size,
				MaxSlots:       tt.fields.MaxSlots,
				AvailableSlots: tt.fields.AvailableSlots,
				FreeSlots:      tt.fields.FreeSlots,
				Slots:          tt.fields.Slots,
			}
			got, err := pSlots.AssignParkingSlot(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingSlots.AssignParkingSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParkingSlots.AssignParkingSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSlots_FreeupParkingSlot(t *testing.T) {
	type fields struct {
		Size           Size
		MaxSlots       int
		AvailableSlots int
		FreeSlots      *list.List
		Slots          map[string]*ParkingSlot
	}
	type args struct {
		slotId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pSlots := &ParkingSlots{
				Size:           tt.fields.Size,
				MaxSlots:       tt.fields.MaxSlots,
				AvailableSlots: tt.fields.AvailableSlots,
				FreeSlots:      tt.fields.FreeSlots,
				Slots:          tt.fields.Slots,
			}
			if err := pSlots.FreeupParkingSlot(tt.args.slotId); (err != nil) != tt.wantErr {
				t.Errorf("ParkingSlots.FreeupParkingSlot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateParkingSlotsService(t *testing.T) {
	type args struct {
		big    int
		medium int
		small  int
	}
	tests := []struct {
		name string
		args args
		want ParkingSlotsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateParkingSlotsService(tt.args.big, tt.args.medium, tt.args.small); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParkingSlotsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSlotsService_AssignParkingSlot(t *testing.T) {
	type fields struct {
		BigParkingSlots    *ParkingSlots
		MediumParkingSlots *ParkingSlots
		SmallParkingSlots  *ParkingSlots
	}
	type args struct {
		v Vehicle
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pss := &ParkingSlotsService{
				BigParkingSlots:    tt.fields.BigParkingSlots,
				MediumParkingSlots: tt.fields.MediumParkingSlots,
				SmallParkingSlots:  tt.fields.SmallParkingSlots,
			}
			got, err := pss.AssignParkingSlot(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkingSlotsService.AssignParkingSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParkingSlotsService.AssignParkingSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParkingSlotsService_FreeParkingSlot(t *testing.T) {
	type fields struct {
		BigParkingSlots    *ParkingSlots
		MediumParkingSlots *ParkingSlots
		SmallParkingSlots  *ParkingSlots
	}
	type args struct {
		slotId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pss := &ParkingSlotsService{
				BigParkingSlots:    tt.fields.BigParkingSlots,
				MediumParkingSlots: tt.fields.MediumParkingSlots,
				SmallParkingSlots:  tt.fields.SmallParkingSlots,
			}
			if err := pss.FreeParkingSlot(tt.args.slotId); (err != nil) != tt.wantErr {
				t.Errorf("ParkingSlotsService.FreeParkingSlot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
