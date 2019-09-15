package user

import (
	"errors"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

//User holds data for single User
//Language conventions :
type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const (
	dbPath = "users.db"
)

var (
	ErrRecordInvalid = errors.New("record is invalid")
)

func All() ([]User, error) {

	db, err := storm.Open(dbPath)

	if err != nil {
		return nil, err
	}

	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

//One return a single user record from database
func One(id bson.ObjectId) (*User, error) {

	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := new(User)

	err = db.One("ID", id, user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

//Delete record removes a given record from the database
func Delete(id bson.ObjectId) error {

	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)

	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}

//Save updates or creates a given record in the database
func Save(u *User) error {

	if err := validate(u); err != nil {
		return err
	}

	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Save(u)
}

//validate method is used to validate data before saving or update
func validate(u *User) error {

	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
