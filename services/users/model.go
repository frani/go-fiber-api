package models

import (
	"github.com/frani/go-fiber-api/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserCollection | @desc: the user ccollection on the database
var UserCollection *mongo.Collection

// User | @desc: user model struct
type User struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email string             `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
}

/*
CreateUserSchema
@desc: adds schema validation and indexes to collection
*/
func CreateUser(Name string, Email string) result error  {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"name", "email", "password"},
		"properties": bson.M{
			"name": bson.M{
				"bsonType":    "string",
				"description": "must be a string and is required",
			},
			"email": bson.M{
				"bsonType": "string",
				"pattern":  "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			},
			"password": bson.M{
				"bsonType":    "string",
				"description": "must be a string and is required",
			},
		},
	}

	validator := bson.M{
		"$jsonSchema": jsonSchema,
	}

	configs.DB.CreateCollection(configs.Ctx, "users", options.CreateCollection().SetValidator(validator))

	UserCollection = configs.DB.Collection("users")

	newUser := User{
		Id:    primitive.NewObjectID(),
		Name:  Name,
		Email: Email,
	}

	result, err := UserCollection.FindOneAndUpdate(configs.Ctx)
	result, err := UserCollection.InsertOne(configs.Ctx, newUser, options.ups)

	if err != nil {
		return err
	}

	return result
}

func ListUsers(filter bson) error {
	cursor, err := UserCollection.Find(configs.Ctx, filter)

    if err!= nil {
        return err
    }

    defer cursor.Close(configs.Ctx)

    for cursor.Next(configs.Ctx) {
        var user User

        err := cursor.Decode(&user)

        if err!= nil {
            return err
        }

        fmt.Println(user)
    }

    return users
}