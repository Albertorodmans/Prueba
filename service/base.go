package base

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func Connectiondbase() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conexion realizada")
	}
	return session
}
