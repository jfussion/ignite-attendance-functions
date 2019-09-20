package firestore

import (
	"context"
	"log"

	gFirestore "cloud.google.com/go/firestore"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
	status "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type firestorePeopleRepo struct {
	client     *gFirestore.Client
	collection string
}

func NewFirestorePeopleRepo(c *gFirestore.Client, col string) *firestorePeopleRepo {
	return &firestorePeopleRepo{
		client:     c,
		collection: col,
	}
}

func (f *firestorePeopleRepo) Get(ctx context.Context, id string) (people domain.People, err error) {
	doc, err := f.client.Collection(f.collection).Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Printf("error: document %s not found", id)
		}

		return
	}

	data := doc.Data()
	data["id"] = id
	people = ToPeople(data)
	return
}

func ToPeople(data map[string]interface{}) domain.People {
	return domain.People{
		ID:       data["id"].(string),
		Name:     data["name"].(string),
		School:   data["school"].(string),
		Course:   data["course"].(string),
		IsMember: data["isMember"].(bool),
	}
}
