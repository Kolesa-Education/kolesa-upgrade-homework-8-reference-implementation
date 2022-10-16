package combinatorics

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8-reference-implementation/card"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
	"testing"
)

func Test_deduplicate(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		ints := []int{1, 2, 3, 3, 4, 5, 5, 6}
		deduplicated := deduplicate(ints)
		assert.Equal(t, 6, len(deduplicated))
		assert.True(t, slices.Contains(deduplicated, 1))
		assert.True(t, slices.Contains(deduplicated, 2))
		assert.True(t, slices.Contains(deduplicated, 3))
		assert.True(t, slices.Contains(deduplicated, 4))
		assert.True(t, slices.Contains(deduplicated, 5))
		assert.True(t, slices.Contains(deduplicated, 6))
	})
	t.Run("card slice", func(t *testing.T) {
		cards := []card.Card{
			{Face: card.Face2, Suit: card.SuitSpades},
			{Face: card.Face2, Suit: card.SuitDiamonds},
			{Face: card.Face2, Suit: card.SuitSpades},
			{Face: card.Face10, Suit: card.SuitHearts},
			{Face: card.Face10, Suit: card.SuitHearts},
			{Face: card.Face10, Suit: card.SuitClubs},
		}
		deduplicated := deduplicate(cards)
		assert.Equal(t, 4, len(deduplicated))
		assert.True(t, slices.Contains(cards, card.Card{Face: card.Face2, Suit: card.SuitSpades}))
		assert.True(t, slices.Contains(cards, card.Card{Face: card.Face2, Suit: card.SuitDiamonds}))
		assert.True(t, slices.Contains(cards, card.Card{Face: card.Face10, Suit: card.SuitHearts}))
		assert.True(t, slices.Contains(cards, card.Card{Face: card.Face10, Suit: card.SuitClubs}))
	})
}
