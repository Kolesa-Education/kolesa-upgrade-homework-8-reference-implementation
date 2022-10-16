package card

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_cardsToRepresentation(t *testing.T) {
	t.Run("valid cards", func(t *testing.T) {
		card1, err := New(SuitDiamonds, FaceAce)
		require.NoError(t, err)
		card2, err := New(SuitSpades, FaceJack)
		require.NoError(t, err)
		card3, err := New(SuitSpades, Face10)
		require.NoError(t, err)

		cards := []Card{*card1, *card2, *card3}
		representations := cardsToRepresentations(cards)
		assert.Equal(t, []string{"♦A", "♠J", "♠10"}, representations)
	})

	t.Run("invalid cards produce empty representations", func(t *testing.T) {
		card1 := Card{
			Face: "invalid",
			Suit: SuitSpades,
		}

		card2 := Card{
			Face: FaceAce,
			Suit: "invalid",
		}

		card3 := Card{
			Face: FaceKing,
			Suit: "invalid",
		}

		cards := []Card{card1, card2, card3}
		representations := cardsToRepresentations(cards)
		assert.Equal(t, []string{"", "", ""}, representations)
	})
}
