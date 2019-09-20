package functions

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log the interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     AttendanceData `json:"fields"`
	Name       string         `json:"name"`
	UpdateTime time.Time      `json:"updateTime"`
}

type People struct {
	ID, Name, School, Course string
	IsMember                 bool
}

type AttendanceData struct {
	Date struct {
		Value string `json:"StringValue"`
	} `json:"Date"`
	ID struct {
		Value string `json:"StringValue"`
	} `json:"peopleID"`
	Name struct {
		Value string `json:"StringValue"`
	} `json:"name"`
	School struct {
		Value string `json:"StringValue"`
	} `json:"school"`
	Course struct {
		Value string `json:"StringValue"`
	} `json:"course"`
	IsMember struct {
		Value bool `json:"BooleanValue"`
	} `json:"isMember"`
}

type ClientRepo interface {
	Doc(path string) *firestore.DocumentRef
}

type DocSnapshot interface {
	Data() map[string]interface{}
}

type GetterSetter interface {
	Get(ctx context.Context) (doc *firestore.DocumentSnapshot, err error)
	Set(ctx context.Context, data interface{}, opts ...firestore.SetOption) (_ *firestore.WriteResult, err error)
}

func GetPeople(ctx context.Context, id string) (people People, err error) {
	path := fmt.Sprintf("%s/%s", peopleCollection, id)
	r, err := client.Doc(path).Get(ctx)
	if err != nil {
		log.Printf("error: %v", err)
		return people, nil
	}

	data := r.Data()
	people = People{
		ID:       id,
		Name:     data["name"].(string),
		School:   data["school"].(string),
		Course:   data["course"].(string),
		IsMember: data["isMember"].(bool),
	}
	log.Printf("People: %v", people)

	return
}

func PopulateAttendance(ctx context.Context, path string, val AttendanceData) {
	id := val.ID.Value
	people, err := GetPeople(ctx, id)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	addCount(ctx)
	if val.Name.Value != "" {
		log.Print("Skipping.. Attendance is already populated")
		return
	}

	data := map[string]interface{}{
		"date":     val.Date.Value,
		"peopleID": id,
		"name":     people.Name,
		"school":   people.School,
		"course":   people.Course,
		"isMember": people.IsMember,
	}

	_, err = client.Doc(path).Set(ctx, data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func addCount(ctx context.Context) {
	q := client.Doc("attendance/count")
	r, err := q.Get(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	data := r.Data()
	count := int(data["total"].(int64))
	count++

	q.Set(ctx, map[string]int{"total": count})
}

// func isDuplicate(ctx context.Context, collection, date, id string) bool {
// 	q := client.Collection(collection).
// 		Where("date", "==", date).
// 		Where("peopleID", "==", id)
//
// 	return false
// }
