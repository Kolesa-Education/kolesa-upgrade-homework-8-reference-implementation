package combinatorics

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8-reference-implementation/card"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
	"log"
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

func Test_combinations(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		ints := []int{1, 2, 3, 4, 5, 6, 7}
		result, err := combinations(ints, 2)
		require.NoError(t, err)
		assert.Equal(t, 21, len(result))
		for i := 0; i < len(result); i++ {
			assert.Equal(t, 2, len(result[i]))
		}
	})
	t.Run("card slice", func(t *testing.T) {
		cards := []card.Card{
			{Face: card.Face2, Suit: card.SuitSpades},
			{Face: card.Face3, Suit: card.SuitDiamonds},
			{Face: card.Face4, Suit: card.SuitSpades},
			{Face: card.Face5, Suit: card.SuitHearts},
			{Face: card.FaceJack, Suit: card.SuitHearts},
			{Face: card.FaceQueen, Suit: card.SuitHearts},
			{Face: card.FaceAce, Suit: card.SuitClubs},
		}
		result, err := combinations(cards, 2)
		log.Println(result)
		require.NoError(t, err)
		assert.Equal(t, 21, len(result))
		for i := 0; i < len(result); i++ {
			assert.Equal(t, 2, len(result[i]))
		}
	})
	t.Run("card slice", func(t *testing.T) {
		cards := []card.Card{
			{Face: card.Face2, Suit: card.SuitSpades},
			{Face: card.Face3, Suit: card.SuitDiamonds},
			{Face: card.Face4, Suit: card.SuitSpades},
			{Face: card.Face5, Suit: card.SuitHearts},
			{Face: card.FaceJack, Suit: card.SuitHearts},
			{Face: card.FaceQueen, Suit: card.SuitHearts},
			{Face: card.FaceAce, Suit: card.SuitClubs},
		}
		result, err := combinations(cards, 5)
		log.Println(result)
		require.NoError(t, err)
		assert.Equal(t, 21, len(result))
		for i := 0; i < len(result); i++ {
			assert.Equal(t, 5, len(result[i]))
		}
	})
}
