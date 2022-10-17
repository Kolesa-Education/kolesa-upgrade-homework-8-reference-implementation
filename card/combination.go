package card

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"sort"
	"strings"
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
	Representation() (string, error)
}

type BasicPokerCombination struct {
	name  string
	cards []Card
}

func (r BasicPokerCombination) Name() string {
	return r.name
}

func (r BasicPokerCombination) Cards() []Card {
	return r.cards
}

func (r BasicPokerCombination) Representation() (string, error) {
	var builder strings.Builder
	for index, card := range r.Cards() {
		representation, err := card.ShortRepresentation()
		if err != nil {
			return "", err
		}
		builder.WriteString(representation)
		if index < ValidCombinationSize-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString(fmt.Sprintf(" | %s", r.Name()))
	return builder.String(), nil
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

func CombinationOf(cards []Card) (PokerCombination, error) {
	if len(cards) != ValidCombinationSize {
		return nil, errors.New("cards is not of valid size")
	}
	switch {
	case isCombinationOfFlush(cards) && isCombinationOfStraight(cards):
		return BasicPokerCombination{name: CombinationStraightFlush, cards: cards}, nil
	case isCombinationOfFourOfAKind(cards):
		return BasicPokerCombination{name: CombinationFourOfAKind, cards: cards}, nil
	case isCombinationOfFullHouse(cards):
		return BasicPokerCombination{name: CombinationFullHouse, cards: cards}, nil
	case isCombinationOfFlush(cards):
		return BasicPokerCombination{name: CombinationFlush, cards: cards}, nil
	case isCombinationOfStraight(cards):
		return BasicPokerCombination{name: CombinationStraight, cards: cards}, nil
	case isCombinationOfThreeOfAKind(cards):
		return BasicPokerCombination{name: CombinationThreeOfAKind, cards: cards}, nil
	case isCombinationOfTwoPairs(cards):
		return BasicPokerCombination{name: CombinationTwoPairs, cards: cards}, nil
	case isCombinationOfPair(cards):
		return BasicPokerCombination{name: CombinationPairName, cards: cards}, nil
	default:
		return nil, nil
	}
}
