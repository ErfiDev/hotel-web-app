package utils

import (
	"time"

	"github.com/erfidev/hotel-web-app/models"
)

func ReturnBetweenDates(reservations []models.Reservation) []time.Time {
	timeSlice := []time.Time{}

	for _, value := range reservations {
		sd, ed := value.StartDate, value.EndDate

		if sd.Equal(ed) {
			timeSlice = append(timeSlice, sd)
		} else if sd.Before(ed) {
			timeSlice = append(timeSlice, sd)
			for sd.Before(ed) {
				addToSd := sd.AddDate(0, 0, 1)
				timeSlice = append(timeSlice, addToSd)
				sd = addToSd

				if addToSd.Equal(ed) {
					break
				} else {
					continue
				}
			}
		}
	}

	return timeSlice
}

func DeleteDuplicateValues(slice []string) []string {
	keys := make(map[string]bool)
	removedSlice := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			removedSlice = append(removedSlice, entry)
		}
	}

	return removedSlice
}
