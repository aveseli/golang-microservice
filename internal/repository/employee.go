package repository

import (
	"context"
	"time"

	"github.com/aveseli/golang-microservice/internal/cfg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Employee struct
type Employee struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name"`
	Salary float64            `json:"salary"`
	Age    float64            `json:"age"`
}

func db() (*mongo.Collection, context.Context, context.CancelFunc) {
	c := cfg.MongoDb.Db.Collection("employee")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return c, ctx, cancel
}

func GetEmployee(id string) (Employee, error) {
	c, ctx, cancel := db()
	defer cancel()
	objectID, _ := primitive.ObjectIDFromHex(id)
	r := c.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: objectID}})
	e := Employee{}
	err := r.Decode(&e)
	return e, err
}

func GetAllEmployees() ([]Employee, error) {
	var employees []Employee
	c, ctx, cancel := db()

	defer cancel()

	r, err := c.Find(ctx, bson.D{{}})
	if err != nil {
		return employees, err
	}

	for r.Next(ctx) {
		var e Employee
		err := r.Decode(&e)
		if err != nil {
			return employees, err
		}
		employees = append(employees, e)
	}

	if err := r.Err(); err != nil {
		return employees, err
	}

	r.Close(ctx)

	if len(employees) == 0 {
		return employees, mongo.ErrNoDocuments
	}

	return employees, nil
}

func InsertEmployee(e Employee) {
	c, ctx, cancel := db()
	defer cancel()

	c.InsertOne(ctx, e)
}
