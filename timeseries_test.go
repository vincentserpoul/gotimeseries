package gotimeseries

import (
	// "log"
	"testing"
	"time"
)

func TestGetTimeSeries15m(t *testing.T) {
	begTime, _ := time.Parse(time.RFC3339, "2014-01-01T22:14:03Z")
	endTime, _ := time.Parse(time.RFC3339, "2014-01-01T22:15:03Z")
	intervalDuration, _ := time.ParseDuration("15m")

	timeSeries := GetTimeSeries(begTime, endTime, intervalDuration)
	if len(timeSeries) != 2 {
		t.Error("The size of resulting serie should be 2")
	}
}

func TestGetTimeSeries24h(t *testing.T) {
	begTime, _ := time.Parse(time.RFC3339, "2014-01-01T22:14:03Z")
	endTime, _ := time.Parse(time.RFC3339, "2014-01-07T22:15:03Z")
	intervalDuration, _ := time.ParseDuration("24h")

	timeSeries := GetTimeSeries(begTime, endTime, intervalDuration)
	if len(timeSeries) != 7 {
		t.Error("The size of resulting serie should be 7")
	}
}

func TestGetBegTimeSeries(t *testing.T) {
	endTime, _ := time.Parse(time.RFC3339, "2014-01-01T22:15:03Z")
	intervalDuration, _ := time.ParseDuration("1h")

	begTime := GetBegTimeSeries(endTime, intervalDuration, 2)
	expBegTime, _ := time.Parse(time.RFC3339, "2014-01-01T21:00:00Z")
	if begTime != expBegTime {
		t.Error("The serie starting point should be 2014-01-01T21:00:00Z, the output was ", begTime.Format(time.RFC3339))
	}
}

func TestGetEndTimeSeries(t *testing.T) {
	begTime, _ := time.Parse(time.RFC3339, "2014-01-01T22:15:03Z")
	intervalDuration, _ := time.ParseDuration("1h")

	endTime := GetEndTimeSeries(begTime, intervalDuration, 2)
	expEndTime, _ := time.Parse(time.RFC3339, "2014-01-01T23:00:00Z")
	if endTime != expEndTime {
		t.Error("The serie ending point should be 2014-01-01T23:00:00Z, the output was ", endTime.Format(time.RFC3339))
	}
}

func BenchmarkGetTimeSeries(b *testing.B) {
	// worse case scenario, 6 years of data, every 168h
	begTime, _ := time.Parse(time.RFC3339, "2009-01-01T22:14:03Z")
	endTime, _ := time.Parse(time.RFC3339, "2014-01-01T22:15:03Z")
	intervalDuration, _ := time.ParseDuration("168h")
	benchmarkGetTimeSeries(begTime, endTime, intervalDuration, b)
}

func benchmarkGetTimeSeries(begTime time.Time, endTime time.Time, intervalDuration time.Duration, b *testing.B) {
	// run the GetTimeSerie function b.N times
	for n := 0; n < b.N; n++ {
		GetTimeSeries(begTime, endTime, intervalDuration)
	}
}

func BenchmarkGetBegTimeSeries(b *testing.B) {
	// 1000 points
	endTime, _ := time.Parse(time.RFC3339, "2014-01-01T22:15:03Z")
	intervalDuration, _ := time.ParseDuration("168h")
	benchmarkGetBegTimeSeries(endTime, intervalDuration, 1000, b)
}

func benchmarkGetBegTimeSeries(endTime time.Time, intervalDuration time.Duration, numberElements int, b *testing.B) {
	// run the GetTimeSerie function b.N times
	for n := 0; n < b.N; n++ {
		GetBegTimeSeries(endTime, intervalDuration, numberElements)
	}
}

func BenchmarkGetEndTimeSeries(b *testing.B) {
	begTime, _ := time.Parse(time.RFC3339, "2009-01-01T22:14:03Z")
	intervalDuration, _ := time.ParseDuration("168h")
	benchmarkGetEndTimeSeries(begTime, intervalDuration, 1000, b)
}

func benchmarkGetEndTimeSeries(begTime time.Time, intervalDuration time.Duration, numberElements int, b *testing.B) {
	// run the GetTimeSerie function b.N times
	for n := 0; n < b.N; n++ {
		GetEndTimeSeries(begTime, intervalDuration, numberElements)
	}
}

// to launch benchmark, do:
// go test -bench=.
