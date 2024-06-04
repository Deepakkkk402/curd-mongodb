package main

import (
	"net/http"

	"github.com/deepak4020/curd-mongodb/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewuserController(getsession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8000", r)

}

func getsession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
