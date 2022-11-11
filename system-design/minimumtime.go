package systemdesign

func numberOfTripsForAGivenTime(time []int, givenTime int64) int64 {
	var totalTrips int64 = 0
	for i := range time {
		totalTrips += givenTime / int64(time[i])
	}
	return totalTrips
}

func MinimumTime(time []int, totalTrips int) int64 {
	var lowestTime int64 = 0
	var highestTime int64 = 9223372036854775807

	for lowestTime < highestTime {
		mid := lowestTime + (highestTime-lowestTime)/2

		if int(numberOfTripsForAGivenTime(time, mid)) >= totalTrips {
			highestTime = mid
		} else {
			lowestTime = mid + 1
		}
	}
	return lowestTime
}
