package firestore_test

import (
	"testing"

	"github.com/jfussion/ignite-attendance-cloud-functions/attendance/repository/firestore"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	tPeople := domain.People{
		ID:       "IGNT-DEMO-0001",
		Name:     "Joe",
		School:   "PUP",
		Course:   "BSCpE",
		IsMember: true,
	}

	tAttendance := domain.Attendance{
		Date:   "September 1, 2019",
		People: tPeople,
	}

	tData := map[string]interface{}{
		"date":     tAttendance.Date,
		"peopleID": tPeople.ID,
		"name":     tPeople.Name,
		"school":   tPeople.School,
		"course":   tPeople.Course,
		"isMember": tPeople.IsMember,
	}

	got := firestore.ToMap(tAttendance)
	assert.Equal(t, tData, got)
}

func TestToCount(t *testing.T) {
	tCount := domain.Count{
		Total:   10,
		Members: 6,
		VIPs:    4,
	}
	tData := map[string]interface{}{
		"total":   int64(tCount.Total),
		"members": int64(tCount.Members),
		"vips":    int64(tCount.VIPs),
	}

	got := firestore.ToCount(tData)
	assert.Equal(t, tCount, got)
}
