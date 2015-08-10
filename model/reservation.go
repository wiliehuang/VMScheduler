//reservation model representation

package model

import (
	"gopkg.in/mgo.v2"
)

type Reservation struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	StartsAt string `json:"startsAt"`
	EndsAt   string `json:"endsAt"`
}

func (reservation *Reservation) Valid() bool {
	return len(reservation.Title) > 0 &&
		len(reservation.Type) > 0 &&
		len(reservation.StartsAt) > 0
}

func FetchAllReservations(db *mgo.Database) []Reservation {
	reservations := []Reservation{}
	err := db.C("reservations").Find(nil).All(&reservations)
	if err != nil {
		panic(err)
	}

	return reservations
}
