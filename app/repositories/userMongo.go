package repositories

import (
	"time"

	"github.com/dung13890/api-go/app/contracts"
	"github.com/dung13890/api-go/config"
	"github.com/dung13890/api-go/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserMongo struct {
	collection *mgo.Collection
}

func UserMongoImplement(session *mgo.Session) contracts.UserRepo {
	c := session.DB(config.GetString("mongo.database")).C("users")
	c.EnsureIndex(models.UserModelIndex())

	return &UserMongo{c}
}

func (u *UserMongo) Find(id string) (models.User, error) {
	m := models.User{}
	err := u.collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&m)

	return m, err
}

func (u *UserMongo) FindByEmail(email string) (models.User, error) {
	m := models.User{}
	err := u.collection.Find(bson.M{"email": email}).One(&m)

	return m, err
}

func (u *UserMongo) Store(p models.User) (models.User, error) {
	p.Id = bson.NewObjectId()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := u.collection.Insert(&p)

	rs := models.User{}
	u.collection.FindId(p.Id).One(&rs)

	return rs, err
}

func (u *UserMongo) GetAll() []models.User {
	users := []models.User{}
	iter := u.collection.Find(nil).Iter()
	user := models.User{}
	for iter.Next(&user) {
		users = append(users, user)
	}

	return users
}
