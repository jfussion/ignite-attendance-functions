package functions

import (
	"context"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	firestoreAttendanceRepo "github.com/jfussion/ignite-attendance-cloud-functions/attendance/repository/firestore"
	attendanceUsecase "github.com/jfussion/ignite-attendance-cloud-functions/attendance/usecase"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
	firestorePeopleRepo "github.com/jfussion/ignite-attendance-cloud-functions/people/repository/firestore"
	peopleUsecase "github.com/jfussion/ignite-attendance-cloud-functions/people/usecase"
)

const peopleCollection = "people"

var client *firestore.Client

var projectID = os.Getenv("GCLOUD_PROJECT")

func init() {
	conf := &firebase.Config{ProjectID: projectID}

	// Use context.Background() because the app/client should persist across
	// invocations.
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

func F(ctx context.Context, e FirestoreEvent) (err error) {
	data := e.Value.Fields
	path := extractPath(e.Value.Name)
	attendanceCollection := strings.Split(path, "/")[0]

	attendanceRepo := firestoreAttendanceRepo.NewFirestoreAttendanceRepo(client, attendanceCollection)
	peopleRepo := firestorePeopleRepo.NewFirestorePeopleRepo(client, peopleCollection)
	peopleUcase := peopleUsecase.NewPeopleUsecase(peopleRepo)
	attendanceUcase := attendanceUsecase.NewAttendanceUsecase(attendanceRepo)

	people, err := peopleUcase.Get(ctx, data.ID.Value)
	if err != nil {
		log.Printf("something went wrong when getting people data: %v", err)
		return
	}

	attendance := domain.Attendance{
		Date:   data.Date.Value,
		People: people,
	}

	if data.Name.Value != "" {
		log.Print("info: skipping, attendance has already populated")
	} else {
		err = attendanceUcase.Update(ctx, path, attendance)
		if err != nil {
			log.Printf("something went updating attendance: %v", err)
			return
		}
	}

	err = attendanceUcase.IncrementCount(ctx, people.IsMember)
	if err != nil {
		log.Printf("something went updating attendance: %v", err)
	}

	return
}

func extractPath(path string) string {
	return strings.Split(path, "/documents/")[1]
}
