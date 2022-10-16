package card

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
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

func countSuits(cards []Card) map[string]int {
	cardSuitsCount := map[string]int{}
	for _, card := range cards {
		if count, ok := cardSuitsCount[card.Suit]; ok {
			cardSuitsCount[card.Suit] = count + 1
		} else {
			cardSuitsCount[card.Suit] = 1
		}
	}
	return cardSuitsCount
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

func isCombinationOfStraight(cards []Card) bool {
	if len(cards) != ValidCombinationSize {
		return false
	}
	numericValues := lo.Map[Card, int](cards, func(card Card, index int) int {
		return card.NumericValue()
	})

	sort.Slice(numericValues, func(i, j int) bool {
		return numericValues[i] > numericValues[j]
	})

	if slices.Contains(numericValues, NumericValueAce) && slices.Contains(numericValues, 2) {
		return numericValues[0] == NumericValueAce &&
			numericValues[1] == 5 &&
			numericValues[2] == 4 &&
			numericValues[3] == 3 &&
			numericValues[4] == 2
	} else {
		for i := 0; i < len(numericValues)-1; i++ {
			if numericValues[i]-numericValues[i+1] != 1 {
				return false
			}
		}
		return true
	}
}

func isCombinationOfFlush(cards []Card) bool {
	if len(cards) != ValidCombinationSize {
		return false
	}
	suits := countSuits(cards)
	suitsCount := lo.Values(suits)
	return len(suitsCount) == 1 && suitsCount[0] == 5
}
