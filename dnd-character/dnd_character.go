package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64((score - 10)) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := make([]int, 4)
	for i, _ := range rolls {
		rolls[i] = 1 + rand.Intn(6)
	}

	return maxSubarraySumOfSizeK(rolls, 3)
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	var character Character

	character.Strength = Ability()
	character.Dexterity = Ability()
	character.Constitution = Ability()
	character.Intelligence = Ability()
	character.Wisdom = Ability()
	character.Charisma = Ability()
	character.Hitpoints = 10 + Modifier(character.Constitution)

	return character
}

// Calculate sum of subarray of size k.
func maxSubarraySumOfSizeK(arr []int, k int) int {
	if len(arr) < k {
		return 0
	}

	subarraySum := 0
	for i := 0; i < k; i++ {
		subarraySum += arr[i]
	}

	maxSum := subarraySum
	for i := k; i < len(arr); i++ {
		subarraySum += arr[i]
		subarraySum -= arr[i-k]
		if subarraySum > maxSum {
			maxSum = subarraySum
		}
	}
	return maxSum
}
