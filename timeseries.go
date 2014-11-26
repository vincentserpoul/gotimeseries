package gotimeseries

import (
	"time"
)

// returns the time series
func GetTimeSeries(begTime time.Time, endTime time.Time, intervalDuration time.Duration) []time.Time {

	if begTime.After(endTime) {
		tempTime := begTime
		begTime = endTime
		endTime = tempTime
	}

	var timeSeries []time.Time

	if begTime == endTime {
		return timeSeries
	}

	if intervalDuration == 0 {
		timeSeries = append(timeSeries, begTime)
		return timeSeries
	}

	for oneTime := begTime.Truncate(intervalDuration); oneTime.Before(endTime); oneTime = oneTime.Add(intervalDuration) {
		timeSeries = append(timeSeries, oneTime)
	}

	return timeSeries
}

// returns the beginning time according to the number of points we need in the serie
func GetBegTimeSeries(endTime time.Time, intervalDuration time.Duration, numberElements int) time.Time {
	begTime := endTime
	for i := 1; i < numberElements; i++ {
		begTime = begTime.Add(-intervalDuration)
	}
	return begTime.Truncate(intervalDuration)
}

// returns the end time according to the number of points we need in the serie
func GetEndTimeSeries(begTime time.Time, intervalDuration time.Duration, numberElements int) time.Time {
	endTime := begTime
	for i := 1; i < numberElements; i++ {
		endTime = endTime.Add(intervalDuration)
	}
	return endTime.Truncate(intervalDuration)
}
