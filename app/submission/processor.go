package submission

import (
	"math"
	"math/cmplx"
	"encoding/json"

	"api/app/submission/store"
)

type OriginalPoint struct {
	X int
	Y int
	Time float64
}

type Vector struct {
	N int
	Real float64
	Imaginary float64
}

func Process(submissionId int) error {
	store := store.New()
	originalPoints := createOriginalPoints(submissionId)
	n := 0
	maxVectorCount := 101
	vectors := []Vector{}

	for (len(vectors) < maxVectorCount) {// && vectorsOutsideThreshold(vectors, originalPoints) {
		vectors = append(vectors, buildVector(n, originalPoints))

		if n != 0 {
			vectors = append(vectors, buildVector(n * -1, originalPoints))
		}

		n++
	}

	store.AddVectors(submissionId, vectors)

	return nil
}

func createOriginalPoints(submissionId int) *[]OriginalPoint {
	store := store.New()
	submission, err := store.Get(submissionId)

	if err != nil {
		panic("The submission could not be found in storage.")
	}

	originalPoints := []OriginalPoint{}
	err = json.Unmarshal(submission.OriginalPoints, &originalPoints)

	if err != nil {
		panic("The input points seem to be improperly formatted.")
	}

	normalizeTime(originalPoints)

	return &originalPoints
}

func normalizeTime(originalPoints []OriginalPoint) {
	finalPoint := originalPoints[len(originalPoints) - 1]

	for i := 0; i < len(originalPoints); i++ {
		originalPoints.Time = originalPoints.Time / finalPoint.Time
	}
}

func vectorsOutsideThreshold(originalPoints *[]OriginalPoint, vectors *[]Vector) bool {
	errorThreshold := 0.02
	averageError := 1

	return true
}

func buildVector(n int, originalPoints *[]OriginalPoint) Vector {
	time := 0.00
	timeDelta := 0.01
	originalPointsIndex := 0
	cumulativeValue := 0 + 0i

	for time <= 1 {
		originalPoint, originalPointsIndex := findOriginalPoint(time, originalPoints[originalPointsIndex:])
		originalComplexValue := complex(OriginalPoint.X, OriginalPoint.Y)
		cumulativeValue := originalComplexValue * cmplx.Exp(float64(-n) * math.Pi * 2i * time)

		time += timeDelta
	}

	return Vector{N: n, Real: real(cumulativeValue), Imaginary: imag(cumulativeValue)}
}

func findOriginalPoint(time float64, originalPoints []OriginalPoint) (OriginalPoint, int) {
	for i := 0; i < len(originalPoints); i++ {
		if originalPoints[i].Time >= time {
			return originalPoints[i], i
		}
	}
	'
	return originalPoints[len(originalPoints) - 1], len(originalPoints) - 1
}