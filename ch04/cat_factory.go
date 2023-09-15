package ch04

import (
	"errors"
	"log"
)

type Car interface {
	BeepBeep()
}
type BMW struct {
	heatedSeatSubscriptionEnabled bool
}

func (B BMW) BeepBeep() {
	panic("implement me")
}

type Tesla struct {
	autoPilotEnabled bool
}

func (t Tesla) BeepBeep() {
	panic("implement me")
}

func BuildCar(carType string) (Car, error) {
	switch carType {
	case "bmw":
		return BMW{heatedSeatSubscriptionEnabled: true}, nil
	case "tesla":
		return Tesla{autoPilotEnabled: true}, nil
	default:
		return nil, errors.New("unknown car type")
	}
}

func main() {
	myCar, err := BuildCar("bmw")
	if err != nil {
		log.Fatal(err)
	}
	_ = myCar
}