package firestore

import (
	"context"

	gFirestore "cloud.google.com/go/firestore"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
	status "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type firestoreRepo struct {
	client *gFirestore.Client
}

func NewFirestoreRepo(c *gFirestore.Client) domain.AttendanceRepository {
	return &firestoreRepo{client: c}
}
func (f *firestoreRepo) Add(ctx context.Context, attendance domain.Attendance) (err error) { return }
func (f *firestoreRepo) UpdateCount(ctx context.Context, id string, count domain.Count) (err error) {
	return
}

func (f *firestoreRepo) Update(ctx context.Context, path string, attendance domain.Attendance) (err error) {
	_, err = f.client.Doc(path).Set(ctx, ToMap(attendance))
	return
}

func (f *firestoreRepo) GetCount(ctx context.Context, path string) (count domain.Count, err error) {
	doc, err := f.client.Doc(path).Get(ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			count = domain.Count{}
			err = nil
		}
		return
	}

	count = ToCount(doc.Data())
	return
}
