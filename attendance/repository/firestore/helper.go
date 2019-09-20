package firestore

import "github.com/jfussion/ignite-attendance-cloud-functions/domain"

func ToMap(attendance domain.Attendance) map[string]interface{} {
	return map[string]interface{}{
		"date":     attendance.Date,
		"peopleID": attendance.People.ID,
		"name":     attendance.People.Name,
		"school":   attendance.People.School,
		"course":   attendance.People.Course,
		"isMember": attendance.People.IsMember,
	}
}

func ToCount(data map[string]interface{}) (count domain.Count) {
	return domain.Count{
		Total:   int(data["total"].(int64)),
		Members: int(data["members"].(int64)),
		VIPs:    int(data["vips"].(int64)),
	}
}
