package parkinglot

import (
	"container/list"
	"errors"
	"fmt"
	"time"
)

type Size int

const (
	Big Size = iota
	Medium
	Small
)

func (size Size) String() string {
	switch size {
	case Big:
		return "B"
	case Medium:
		return "M"
	case Small:
		return "S"
	}
	return "UNKNOWN"
}

type Vehicle struct {
	NumberPlate string
	Size        Size
}

func (v Vehicle) IsValidVehicle() bool {
	return len(v.NumberPlate) != 0
}

type ParkingSlot struct {
	ID              string
	Size            Size
	AssignedVehicle Vehicle
	AssignedTime    time.Time
}

func (slot *ParkingSlot) IsFree() bool {
	return slot.AssignedVehicle == Vehicle{}
}

func (slot *ParkingSlot) OccupiedSince() int {
	if !slot.IsFree() {
		return int(time.Since(slot.AssignedTime))
	}
	return 0
}

func (slot *ParkingSlot) AssignVehicle(vehicle Vehicle) error {
	if !vehicle.IsValidVehicle() {
		return errors.New("vehicle isn't valid")
	}
	if slot.IsFree() {
		slot.AssignedVehicle = vehicle
		slot.AssignedTime = time.Now()
		return nil
	}

	return errors.New("parking slot isn't empty")
}

func (slot *ParkingSlot) FreeUpSlot() {
	slot.AssignedVehicle = Vehicle{}
}

func CreateParkingSlot(id string, size Size) *ParkingSlot {
	return &ParkingSlot{
		ID:   id,
		Size: size,
	}
}

type ParkingSlots struct {
	Size           Size
	MaxSlots       int
	AvailableSlots int
	FreeSlots      *list.List
	Slots          map[string]*ParkingSlot
}

func CreateParkingSlots(totalSlots int, slotSize Size) *ParkingSlots {
	//Create ParkingSlot for given total size and slotSize

	slots := make(map[string]*ParkingSlot)
	freeSlots := list.New()
	for i := 0; i < totalSlots; i++ {
		id := fmt.Sprintf("%s%d", slotSize.String(), i)
		slots[id] = CreateParkingSlot(id, slotSize)
		freeSlots.PushBack(id)
	}

	return &ParkingSlots{
		Size:           slotSize,
		FreeSlots:      freeSlots,
		MaxSlots:       totalSlots,
		AvailableSlots: totalSlots,
		Slots:          slots,
	}
}

func (pSlots *ParkingSlots) AssignParkingSlot(v Vehicle) (string, error) {
	if pSlots.AvailableSlots != 0 {
		front := pSlots.FreeSlots.Front()
		slotId := fmt.Sprintf("%v", front.Value)
		err := pSlots.Slots[slotId].AssignVehicle(v)
		if err != nil {
			return "", err
		} else {
			pSlots.AvailableSlots -= 1
			pSlots.FreeSlots.Remove(front)
			return slotId, nil
		}
	}
	return "", errors.New("no more free parking slots")
}

func (pSlots *ParkingSlots) FreeupParkingSlot(slotId string) error {
	if _, ok := pSlots.Slots[slotId]; !ok {
		return errors.New("slotid doesn't exist")
	}
	pSlots.Slots[slotId].FreeUpSlot()
	pSlots.AvailableSlots += 1
	pSlots.FreeSlots.PushBack(slotId)
	return nil
}

type ParkingSlotsService struct {
	BigParkingSlots    *ParkingSlots
	MediumParkingSlots *ParkingSlots
	SmallParkingSlots  *ParkingSlots
}

func CreateParkingSlotsService(big, medium, small int) ParkingSlotsService {
	return ParkingSlotsService{
		BigParkingSlots:    CreateParkingSlots(big, Big),
		MediumParkingSlots: CreateParkingSlots(medium, Medium),
		SmallParkingSlots:  CreateParkingSlots(small, Small),
	}
}

func (pss *ParkingSlotsService) AssignParkingSlot(v Vehicle) (string, error) {
	var slotId string = ""
	var err error = nil
	switch v.Size {
	case Big:
		slotId, err = pss.BigParkingSlots.AssignParkingSlot(v)
	case Medium:
		slotId, err = pss.MediumParkingSlots.AssignParkingSlot(v)
	case Small:
		slotId, err = pss.SmallParkingSlots.AssignParkingSlot(v)
	}
	return slotId, err
}

func (pss *ParkingSlotsService) FreeParkingSlot(slotId string) error {
	var err error = nil
	if len(slotId) > 1 {
		slotFirstChar := slotId[0:1]
		switch slotFirstChar {
		case "B":
			err = pss.BigParkingSlots.FreeupParkingSlot(slotId)
		case "M":
			err = pss.MediumParkingSlots.FreeupParkingSlot(slotId)
		case "S":
			err = pss.BigParkingSlots.FreeupParkingSlot(slotId)
		}
	}

	return err
}
