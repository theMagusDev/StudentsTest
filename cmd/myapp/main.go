package main

import (
	"fmt"
	"math"
	"Test/internal/utils"
)

const NumberOfExperiments = 20000


func main() {
	runExperiment(false) // random strategy
	runExperiment(true)  // smart strategy
}

func runExperiment(smart bool) {
	successExperimentsCount := 0

	for i := 0; i < NumberOfExperiments; i++ {
		studentsNumbers := utils.GetShuffledArray(50, 50)
		boxToCardMap := utils.CreateBoxToCardMap()

		allStudentsPassed := true

		for _, studentNumber := range studentsNumbers {
			var passed bool
			if smart {
				passed = smartExperimentStudent(studentNumber, boxToCardMap)
			} else {
				passed = randomExperimentStudent(studentNumber, boxToCardMap)
			}

			if !passed {
				allStudentsPassed = false
				break
			}
		}

		if allStudentsPassed {
			successExperimentsCount++
		}
	}

	fmt.Printf(
		"%d of %d experiments were successful, i.e. %.0f%%\n",
		successExperimentsCount,
		NumberOfExperiments,
		math.Round(float64(successExperimentsCount)/float64(NumberOfExperiments)*100),
	)
}


func randomExperimentStudent(studentNumber int, boxToCardMap map[int]int) bool {
	studentChoices := utils.GetShuffledArray(25, 50)
	for _, studentChoice := range studentChoices {
		if boxToCardMap[studentChoice] == studentNumber {
			return true
		}
	}
	return false
}

func smartExperimentStudent(studentNumber int, boxToCardMap map[int]int) bool {
	studentChoice := studentNumber
	attemptsNumber := 1

	for attemptsNumber <= 25 {
		if boxToCardMap[studentChoice] == studentNumber {
			return true
		}
		studentChoice = boxToCardMap[studentChoice]
		attemptsNumber++
	}

	return false
}
