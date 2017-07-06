package databaselayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongodbHandler struct {
	*mgo.Session
}

func NewMongodbHandler(connection string) (*MongodbHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: s,
	}, err
}

func (handler *MongodbHandler) GetAvailableDynos() ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(nil).All(&animals)
	return animals, err
}

func (handler *MongodbHandler) GetDynoByNickname(nickname string) (Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	a := Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"nickname": nickname}).One(&a)
	return a, err
}

func (handler *MongodbHandler) GetDynosByType(dinoType string) ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"animal_type": dinoType}).All(&animals)
	return animals, err
}

func (handler *MongodbHandler) AddAnimal(a Animal) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("Dino").C("animals").Insert(a)
}

func (handler *MongodbHandler) UpdateAnimal(a Animal, nname string) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("Dino").C("animals").Update(bson.M{"nickname": nname}, a)
}

func (handler *MongodbHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}
