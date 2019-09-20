package firestore

import (
	"context"

	gFirestore "cloud.google.com/go/firestore"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
	status "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type firestoreAttendanceRepo struct {
	client *gFirestore.Client
}

func NewFirestoreAttendanceRepo(c *gFirestore.Client) domain.AttendanceRepository {
	return &firestoreAttendanceRepo{client: c}
}
func (f *firestoreAttendanceRepo) Add(ctx context.Context, attendance domain.Attendance) (err error) {
	return
}
func (f *firestoreAttendanceRepo) UpdateCount(ctx context.Context, path string, count domain.Count) (err error) {
	return
}

func (f *firestoreAttendanceRepo) Update(ctx context.Context, path string, attendance domain.Attendance) (err error) {
	_, err = f.client.Doc(path).Set(ctx, ToMap(attendance))
	return
}

func (f *firestoreAttendanceRepo) GetCount(ctx context.Context, path string) (count domain.Count, err error) {
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
