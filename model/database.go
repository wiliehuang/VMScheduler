// database.go
package model

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
)

type DatabaseSession struct {
	*mgo.Session
	databaseName string
}

func NewSession(name string) *DatabaseSession {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	//addIndexToSignatureEmails(session.DB(name))
	return &DatabaseSession{session, name}
}

func addIndexToSignatureEmails(db *mgo.Database) {
	index := mgo.Index{
		Key:      []string{"endsat"},
		Unique:   true,
		DropDups: true,
	}
	indexErr := db.C("reservations").EnsureIndex(index)
	if indexErr != nil {
		panic(indexErr)
	}
}

func (session *DatabaseSession) Database() martini.Handler {
	return func(context martini.Context) {
		s := session.Clone()
		context.Map(s.DB(session.databaseName))
		defer s.Close()
		context.Next()
	}
}
