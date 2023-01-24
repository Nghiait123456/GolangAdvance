package main

import "fmt"

type Trip int
type Driver struct {
	trips []Trip
}

func (d *Driver) SetTripsBad(trips []Trip) {
	fmt.Println("in bad case")
	d.trips = trips
	fmt.Printf("d.trip = %v, &d.trip = %v, trip =  %v, $trip =  %v \n", d.trips, &d.trips[0], trips, &trips[0])
	d.trips[0] = 3
	fmt.Printf("d.trip = %v, &d.trip = %v, trip =  %v, $trip =  %v \n", d.trips, &d.trips[0], trips, &trips[0])
	fmt.Println("d.trips and trips have same address, one in both change, both changed")
	fmt.Println("end bad case")
}

func (d *Driver) SetTripGood(trips []Trip) {
	fmt.Println("in good case")
	d.trips = make([]Trip, len(trips))
	copy(d.trips, trips)
	fmt.Printf("d.trip = %v, &d.trip = %v, trip =  %v, $trip =  %v \n", d.trips, &d.trips[0], trips, &trips[0])
	d.trips[0] = 3
	fmt.Printf("d.trip = %v, &d.trip = %v, trip =  %v, $trip =  %v \n", d.trips, &d.trips[0], trips, &trips[0])
	fmt.Println("d.trips and trips is independence")
	fmt.Println("end good case")
}

func main() {
	badCase := Driver{}
	tripBad := []Trip{1, 2, 3, 4, 5}
	badCase.SetTripsBad(tripBad)

	goodCase := Driver{}
	tripGood := []Trip{1, 2, 3, 4, 5}
	goodCase.SetTripGood(tripGood)
}
