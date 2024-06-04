package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/deepak4020/curd-mongodb/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Usercontroller struct {
	session *mgo.Session
}

func NewuserController(s *mgo.Session) *Usercontroller {
	return NewuserController(s)
}

func (uc Usercontroller) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)

	}
	old := bson.ObjectIdHex(id)
	u := models.owner{}
	if err := uc.session.DB("curd-mongodb").C("users").FindId(old).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := bson.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("content-type ", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "%s\n", uj)

	return owner, nil

}

func (uc Usercontroller) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.owner{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	uc.session.DB("curd-mongodb").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(w, "%s\n", uj)

}

func (uc Usercontroller) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	old := bson.ObjectIdHex(id)
	if err := uc.session.DB("curd-mongodb").C("users").RemoveId(old); err != nil {

		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "deleteuser ", old, "\n")
}
