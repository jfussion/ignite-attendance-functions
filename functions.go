package functions

import (
	"context"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

const peopleCollection = "people"
const attendanceCollection = "attendance-demo"

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

func F(ctx context.Context, e FirestoreEvent) error {
	path := extractPath(e.Value.Name)
	PopulateAttendance(ctx, path, e.Value.Fields)
	return nil
}

func extractPath(path string) string {
	return strings.Split(path, "/documents/")[1]
}
