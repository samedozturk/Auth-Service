package internal

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id             primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name           string             `bson:"name" json:"name"`
	Surname        string             `bson:"surname" json:"surname"`
	HashedPassword string             `bson:"hashedPassword" json:"-"`
	Email          string             `bson:"email" json:"email"`
	Tel            string             `bson:"tel" json:"tel"`
	Tc             string             `bson:"tc" json:"tc"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
}

type Tokens struct {
	Id        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"userId"`
	Token     string             `bson:"token" json:"token"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	ExpiredAt time.Time          `bson:"expiredAt" json:"expiredAt"`
	IsActive  bool               `bson:"isActive" json:"isActive"`
}

type Logs struct {
	Id         primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Message    string             `bson:"message" json:"message"`
	Level      string             `bson:"level" json:"level"`
	ActionType string             `bson:"actionType" json:"actionType"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt"`
	Ip         string             `bson:"ip" json:"ip"`
	UserAgent  string             `bson:"userAgent" json:"userAgent"`
}
