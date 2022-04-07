package lib

import "time"

func DateAsTime(dob string) time.Time {
	t, _ := time.Parse("2006-01-02", dob)
	return t
}
