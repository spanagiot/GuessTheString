package main

import (
	"fmt"
	"math/rand"
	s "strings"
	"time"
)

var nameToGuess = "To be or not to be"
var population [100]string
var fitnessScore [100]int
var matingPool [100]string
var birthpool [100]string

func main() {
	var maxFitness int = 0
	var fitnessSum int = 0
	var iterations int = 0
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("\tGenetic Algorithm - Guess the name")
	initializePopulation(population[:])
	for true {
		iterations += 1
		maxFitness = 0
		fitnessSum = 0
		for i := 0; i < len(population); i++ {
			fitnessScore[i] = calculateAndReturnFitness(population[i])
			fitnessSum += fitnessScore[i]
			if fitnessScore[i] > maxFitness {
				maxFitness = fitnessScore[i]
				if maxFitness > 2 {
					fmt.Printf("New max fitness: %d with value: %s\r",
						maxFitness, population[i])
				}
			}
			if fitnessScore[i] == len(nameToGuess) {
				fmt.Printf("Found correct offspring: %s with index: %d ",
					population[i], i)
				fmt.Printf("after %d iterations\n", iterations)
				fmt.Printf("Pool currently:")
				fmt.Println(population)
				return
			}
		}
		generateMatingPool(fitnessSum)
		giveBirth()
	}

}

func giveBirth() {
	var firstOffspring string
	var secondOffspring string
	var firstRandomParent int
	var secondRandomParent int
	for i := 0; i < len(matingPool); i += 2 {
		firstRandomParent = rand.Intn(len(matingPool))
		secondRandomParent = rand.Intn(len(matingPool))
		firstOffspring, secondOffspring = createOffspring(
			matingPool[firstRandomParent], matingPool[secondRandomParent])
		birthpool[i] = firstOffspring
		birthpool[i+1] = secondOffspring
	}
	for i := 0; i < len(population); i++ {
		population[i] = birthpool[i]
	}
}

func generateMatingPool(fitnessSum int) {
	var populationIndex int = 0
	var randomInitialPosition int = 0
	for i := 0; i < len(population); i++ {
		randomInitialPosition = rand.Intn(fitnessSum)
		populationIndex = 0
		for j := randomInitialPosition; j < fitnessSum; j += fitnessScore[populationIndex%len(population)] {
			populationIndex += 1
		}
		matingPool[i] = population[populationIndex%len(population)]
	}
}

func createOffspring(firstParent string, secondParent string) (string, string) {
	firstOffspring := ""
	secondOffspring := ""
	for i := 0; i < len(nameToGuess); i++ {
		if rand.Intn(2) == 1 {
			// 1% mutation chance with random character in random position
			if rand.Intn(100) == 5 {
				firstOffspring = s.Join([]string{firstOffspring,
					string(rand.Intn(90) + 32)}, "")
			} else {
				firstOffspring = s.Join([]string{firstOffspring,
					string(secondParent[i])}, "")
			}
			if rand.Intn(100) == 6 {
				secondOffspring = s.Join([]string{secondOffspring,
					string(rand.Intn(90) + 32)}, "")
			} else {
				secondOffspring = s.Join([]string{secondOffspring,
					string(firstParent[i])}, "")
			}
		} else {
			if rand.Intn(100) == 5 {
				firstOffspring = s.Join([]string{firstOffspring,
					string(rand.Intn(90) + 32)}, "")
			} else {
				firstOffspring = s.Join([]string{firstOffspring,
					string(firstParent[i])}, "")
			}
			if rand.Intn(100) == 6 {
				secondOffspring = s.Join([]string{secondOffspring,
					string(rand.Intn(90) + 32)}, "")
			} else {
				secondOffspring = s.Join([]string{secondOffspring,
					string(secondParent[i])}, "")
			}
		}
	}
	return firstOffspring, secondOffspring
}

func calculateAndReturnFitness(name string) int {
	fitness := 0
	for i := 0; i < len(nameToGuess); i++ {
		if i > len(name) {
			break
		}
		if nameToGuess[i] == name[i] {
			fitness++
		}
	}
	return fitness
}

func initializePopulation(initialPopulation []string) {
	var generatedName string
	for i := 0; i < len(initialPopulation); i++ {
		generatedName = ""
		for j := 0; j < len(nameToGuess); j++ {
			generatedName = s.Join([]string{generatedName,
				string(rand.Intn(90) + 32)}, "")
		}
		initialPopulation[i] = generatedName
	}
}
