package card

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isCombinationOfPair(t *testing.T) {
	t.Run("[AD, 10S, 10D, 8C, 2C] produce true", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face8,
				Suit: SuitClubs,
			},
			{
				Face: Face2,
				Suit: SuitClubs,
			},
		}
		assert.True(t, isCombinationOfPair(cards))
	})
	t.Run("[AD, 10S, 10D, 8C, 8H] produce true", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face8,
				Suit: SuitClubs,
			},
			{
				Face: Face8,
				Suit: SuitHearts,
			},
		}
		assert.True(t, isCombinationOfPair(cards))
	})
	t.Run("[AD, 10S, 9D, 8C, 2C] produce false", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face9,
				Suit: SuitDiamonds,
			},
			{
				Face: Face8,
				Suit: SuitClubs,
			},
			{
				Face: Face2,
				Suit: SuitClubs,
			},
		}
		assert.False(t, isCombinationOfPair(cards))
	})
	t.Run("[AD, 10S, 10D, 10C, 2C] produce false", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitClubs,
			},
			{
				Face: Face2,
				Suit: SuitClubs,
			},
		}
		assert.False(t, isCombinationOfPair(cards))
	})
	t.Run("empty cards produce false", func(t *testing.T) {
		var cards []Card
		assert.False(t, isCombinationOfPair(cards))
	})
}

func Test_isCombinationOfThreeOfAKind(t *testing.T) {
	t.Run("[AD, 10S, 10D, 10C, 2C] produce true", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitClubs,
			},
			{
				Face: Face2,
				Suit: SuitClubs,
			},
		}
		assert.True(t, isCombinationOfThreeOfAKind(cards))
	})
	t.Run("[AD, 10S, 10D, 10C, AC] produce true", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitClubs,
			},
			{
				Face: FaceAce,
				Suit: SuitClubs,
			},
		}
		assert.True(t, isCombinationOfThreeOfAKind(cards))
	})
	t.Run("[AD, 10S, 10D, 7C, AC] produce false", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face7,
				Suit: SuitClubs,
			},
			{
				Face: FaceAce,
				Suit: SuitClubs,
			},
		}
		assert.False(t, isCombinationOfThreeOfAKind(cards))
	})
	t.Run("empty cards produce false", func(t *testing.T) {
		var cards []Card
		assert.False(t, isCombinationOfThreeOfAKind(cards))
	})
}

func Test_isCombinationOfFourOfAKind(t *testing.T) {
	t.Run("[AD, 10S, 10D, 10C, 10H] produce true", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitClubs,
			},
			{
				Face: Face10,
				Suit: SuitHearts,
			},
		}
		assert.True(t, isCombinationOfFourOfAKind(cards))
	})
	t.Run("[AD, 10S, 10D, 10C, 9H] produce false", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitClubs,
			},
			{
				Face: Face9,
				Suit: SuitHearts,
			},
		}
		assert.False(t, isCombinationOfFourOfAKind(cards))
	})
	t.Run("empty cards produce false", func(t *testing.T) {
		var cards []Card
		assert.False(t, isCombinationOfFourOfAKind(cards))
	})
}

func Test_isCombinationOfTwoPairs(t *testing.T) {
	t.Run("[AD, 10S, 10D, 8C, 8H] produce true", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face8,
				Suit: SuitClubs,
			},
			{
				Face: Face8,
				Suit: SuitHearts,
			},
		}
		assert.True(t, isCombinationOfTwoPairs(cards))
	})
	t.Run("[AD, 10S, 10D, 8C, 2C] produce false", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face8,
				Suit: SuitClubs,
			},
			{
				Face: Face2,
				Suit: SuitClubs,
			},
		}
		assert.False(t, isCombinationOfTwoPairs(cards))
	})
	t.Run("[8D, 10S, 10D, 8C, 8H] produce false", func(t *testing.T) {
		cards := []Card{
			{
				Face: Face8,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face8,
				Suit: SuitClubs,
			},
			{
				Face: Face8,
				Suit: SuitHearts,
			},
		}
		assert.False(t, isCombinationOfTwoPairs(cards))
	})
	t.Run("empty cards produce false", func(t *testing.T) {
		var cards []Card
		assert.False(t, isCombinationOfTwoPairs(cards))
	})
}

func Test_isCombinationOfFullHouse(t *testing.T) {
	t.Run("[8D, 10S, 10D, 8C, 8H] produce true", func(t *testing.T) {
		cards := []Card{
			{
				Face: Face8,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face8,
				Suit: SuitClubs,
			},
			{
				Face: Face8,
				Suit: SuitHearts,
			},
		}
		assert.True(t, isCombinationOfFullHouse(cards))
	})
	t.Run("[AD, 10S, 10D, 10C, 2C] produce false", func(t *testing.T) {
		cards := []Card{
			{
				Face: FaceAce,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitSpades,
			},
			{
				Face: Face10,
				Suit: SuitDiamonds,
			},
			{
				Face: Face10,
				Suit: SuitClubs,
			},
			{
				Face: Face2,
				Suit: SuitClubs,
			},
		}
		assert.False(t, isCombinationOfFullHouse(cards))
	})
	t.Run("empty cards produce false", func(t *testing.T) {
		var cards []Card
		assert.False(t, isCombinationOfFullHouse(cards))
	})
}
