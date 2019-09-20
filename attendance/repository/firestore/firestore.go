package firestore

import (
	"context"
	"fmt"

	gFirestore "cloud.google.com/go/firestore"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
	status "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type firestoreAttendanceRepo struct {
	client     *gFirestore.Client
	collection string
}

func NewFirestoreAttendanceRepo(c *gFirestore.Client, col string) domain.AttendanceRepository {
	return &firestoreAttendanceRepo{
		client:     c,
		collection: col,
	}
}

func CountToMap(count domain.Count) map[string]interface{} {
	return map[string]interface{}{
		"total":   count.Total,
		"members": count.Members,
		"vips":    count.VIPs,
	}
}

func (f *firestoreAttendanceRepo) UpdateCount(ctx context.Context, id string, count domain.Count) (err error) {
	data := CountToMap(count)
	path := fmt.Sprintf("%s/%s", f.collection, id)
	_, err = f.client.Doc(path).Set(ctx, data)
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

func (f *firestoreAttendanceRepo) Add(ctx context.Context, attendance domain.Attendance) (err error) {
	// TODO: remove Add method to interface
	return
}
