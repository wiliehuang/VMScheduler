// server.go
package controller

import (
	"ScheduleVM/model"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

type Server *martini.ClassicMartini

func NewServer(session *model.DatabaseSession) Server {

	m := Server(martini.Classic())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))
	m.Use(session.Database())

	m.Use(martini.Static("view-modules/angular-bootstrap-calendar"))

	m.Get("/reservation", func(r render.Render, db *mgo.Database) {
		r.JSON(200, model.FetchAllReservations(db))
	})

	m.Post("/reservation", binding.Json(model.Reservation{}),
		func(reservation model.Reservation,
			r render.Render,
			db *mgo.Database) {

			if reservation.Valid() {
				// signature is valid, insert into database
				err := db.C("reservations").Insert(reservation)
				if err == nil {
					// insert successful, 201 Created
					r.JSON(201, reservation)
				} else {
					// insert failed, 400 Bad Request
					r.JSON(400, map[string]string{
						"error": err.Error(),
					})
				}
			} else {
				// signature is invalid, 400 Bad Request
				r.JSON(400, map[string]string{
					"error": "Not a valid signature",
				})
			}
		})

	m.Delete("/reservation", binding.Json(model.Reservation{}),
		func(reservation model.Reservation,
			r render.Render,
			db *mgo.Database) {

			if reservation.Valid() {
				// signature is valid, delete document from database
				err := db.C("reservations").Remove(reservation)
				if err == nil {
					// remove successful, 201 Created
					r.JSON(201, reservation)
				} else {
					// remove failed, 400 Bad Request
					r.JSON(400, map[string]string{
						"error": err.Error(),
					})
				}
			} else {

				// signature is invalid, 400 Bad Request
				r.JSON(400, map[string]string{
					"error": "Not a valid signature",
				})
			}
		})

	m.Put("/reservation", binding.Json(model.Reservation{}),
		func(reservation model.Reservation,
			r render.Render,
			db *mgo.Database) {

			if reservation.Valid() {
				type M map[string]interface{}
				change := M{"$set": model.Reservation{Type: reservation.Type,
					StartsAt: reservation.StartsAt,
					EndsAt:   reservation.EndsAt,
				}}
				// signature is valid, insert into database
				err := db.C("reservations").Update(model.Reservation{Title: reservation.Title}, change)
				if err == nil {
					// insert successful, 201 Created
					r.JSON(201, reservation)
				} else {
					// insert failed, 400 Bad Request
					r.JSON(400, map[string]string{
						"error": err.Error(),
					})
				}
			} else {
				// signature is invalid, 400 Bad Request
				r.JSON(400, map[string]string{
					"error": "Not a valid signature",
				})
			}
		})

	// Return the server. Call Run() on the server to
	// begin listening for HTTP requests.
	return m
}
