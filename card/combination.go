package card

import (
	"github.com/samber/lo"
	"sort"
)

const ValidCombinationSize = 5

const CombinationPairName = "Pair"
const CombinationTwoPairs = "Two Pairs"
const CombinationThreeOfAKind = "Three Of A Kind"
const CombinationStraight = "Straight"
const CombinationFlush = "Flush"
const CombinationFullHouse = "Full House"
const CombinationFourOfAKind = "Four Of A Kind"
const CombinationStraightFlush = "Straight Flush"

type PokerCombination interface {
	Name() string
	Cards() []Card
}

func countFaces(cards []Card) map[string]int {
	cardFaceCount := map[string]int{}
	for _, card := range cards {
		if count, ok := cardFaceCount[card.Face]; ok {
			cardFaceCount[card.Face] = count + 1
		} else {
			cardFaceCount[card.Face] = 1
		}
	}
	return cardFaceCount
}

func isFaceBasedCombination(cards []Card, condition func([]int) bool) bool {
	if len(cards) != ValidCombinationSize {
		return false
	}
	cardFaceCount := countFaces(cards)
	values := lo.Values[string, int](cardFaceCount)
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	return condition(values)
}

func isCombinationOfPair(cards []Card) bool {
	return isFaceBasedCombination(cards, func(values []int) bool {
		return values[0] == 2
	})
}

func isCombinationOfThreeOfAKind(cards []Card) bool {
	return isFaceBasedCombination(cards, func(values []int) bool {
		return values[0] == 3
	})
}

func isCombinationOfFourOfAKind(cards []Card) bool {
	return isFaceBasedCombination(cards, func(values []int) bool {
		return values[0] == 4
	})
}

func isCombinationOfFullHouse(cards []Card) bool {
	return isFaceBasedCombination(cards, func(values []int) bool {
		return values[0] == 3 && values[1] == 2
	})
}

func isCombinationOfTwoPairs(cards []Card) bool {
	return isFaceBasedCombination(cards, func(values []int) bool {
		return values[0] == 2 && values[1] == 2
	})
}
