package card

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_isValidSuit(t *testing.T) {
	assert.True(t, isValidSuit(SuitDiamonds))
	assert.True(t, isValidSuit(SuitSpades))
	assert.True(t, isValidSuit(SuitClubs))
	assert.True(t, isValidSuit(SuitHearts))

	assert.False(t, isValidSuit("Invalid"))
}

func Test_isValidFace(t *testing.T) {
	assert.True(t, isValidFace(Face2))
	assert.True(t, isValidFace(Face3))
	assert.True(t, isValidFace(Face4))
	assert.True(t, isValidFace(Face5))
	assert.True(t, isValidFace(Face6))
	assert.True(t, isValidFace(Face7))
	assert.True(t, isValidFace(Face8))
	assert.True(t, isValidFace(Face9))
	assert.True(t, isValidFace(Face10))
	assert.True(t, isValidFace(FaceJack))
	assert.True(t, isValidFace(FaceQueen))
	assert.True(t, isValidFace(FaceKing))
	assert.True(t, isValidFace(FaceAce))

	assert.False(t, isValidFace("Invalid"))
}

func TestNew(t *testing.T) {
	t.Run("valid creation of cards", func(t *testing.T) {
		suits := []string{SuitDiamonds, SuitSpades, SuitClubs, SuitHearts}
		faces := []string{Face2, Face3, Face4, Face5, Face6, Face7, Face8, Face9, Face10, FaceJack, FaceQueen, FaceKing, FaceAce}

		for _, suit := range suits {
			for _, face := range faces {
				c, err := New(suit, face)
				require.NoError(t, err)
				assert.Equal(t, c.Face, face)
				assert.Equal(t, c.Suit, suit)
			}
		}
	})
	t.Run("invalid face results in error", func(t *testing.T) {
		c, err := New(SuitHearts, "invalid")
		require.Error(t, err)
		assert.Nil(t, c)
	})
	t.Run("invalid suit results in error", func(t *testing.T) {
		c, err := New("invalid", FaceAce)
		require.Error(t, err)
		assert.Nil(t, c)
	})
}

func TestCard_SuitUnicode(t *testing.T) {
	t.Run("spades", func(t *testing.T) {
		c := Card{
			Suit: SuitSpades,
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		require.NoError(t, err)
		assert.Equal(t, unicode, SuitSpadesUnicode)
	})
	t.Run("diamonds", func(t *testing.T) {
		c := Card{
			Suit: SuitDiamonds,
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		require.NoError(t, err)
		assert.Equal(t, unicode, SuitDiamondsUnicode)
	})
	t.Run("hearts", func(t *testing.T) {
		c := Card{
			Suit: SuitHearts,
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		require.NoError(t, err)
		assert.Equal(t, unicode, SuitHeartsUnicode)
	})
	t.Run("clubs", func(t *testing.T) {
		c := Card{
			Suit: SuitClubs,
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		require.NoError(t, err)
		assert.Equal(t, unicode, SuitClubsUnicode)
	})
	t.Run("invalid suit", func(t *testing.T) {
		c := Card{
			Suit: "invalid",
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		assert.Equal(t, "", unicode)
		require.Error(t, err)
	})
}

func TestRandom(t *testing.T) {
	t.Run("seed 1", func(t *testing.T) {
		randomSource := rand.NewSource(1)
		random := rand.New(randomSource)
		card, err := Random(*random)
		require.NoError(t, err)
		// Can be tested because of the fixed seed
		assert.Equal(t, card.Suit, SuitDiamonds)
		assert.Equal(t, card.Face, Face10)
	})
	t.Run("produces no errors", func(t *testing.T) {
		for i := 0; i < 10_000; i++ {
			randomSource := rand.NewSource(time.Now().UnixNano())
			random := rand.New(randomSource)

			_, err := Random(*random)
			require.NoError(t, err)
		}
	})
}

func TestCard_ShortRepresentation(t *testing.T) {
	t.Run("Ace Diamonds", func(t *testing.T) {
		c := Card{
			Suit: SuitDiamonds,
			Face: FaceAce,
		}
		representation, err := c.ShortRepresentation()
		require.NoError(t, err)
		fmt.Println(representation)
		assert.Equal(t, "♦A", representation)
	})
	t.Run("2 Spades", func(t *testing.T) {
		c := Card{
			Suit: SuitSpades,
			Face: Face2,
		}
		representation, err := c.ShortRepresentation()
		require.NoError(t, err)
		fmt.Println(representation)
		assert.Equal(t, "♠2", representation)
	})
	t.Run("invalid suit", func(t *testing.T) {
		c := Card{
			Suit: "invalid",
			Face: FaceAce,
		}
		representation, err := c.ShortRepresentation()
		require.Error(t, err)
		fmt.Println(representation)
		assert.Equal(t, "", representation)
	})
	t.Run("invalid face", func(t *testing.T) {
		t.Run("invalid suit", func(t *testing.T) {
			c := Card{
				Suit: SuitSpades,
				Face: "invalid",
			}
			representation, err := c.ShortRepresentation()
			require.Error(t, err)
			fmt.Println(representation)
			assert.Equal(t, "", representation)
		})
	})
}

func TestCard_IsNumeric(t *testing.T) {
	t.Run("2 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face2}.IsNumeric())
	})
	t.Run("3 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face3}.IsNumeric())
	})
	t.Run("4 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face4}.IsNumeric())
	})
	t.Run("5 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face5}.IsNumeric())
	})
	t.Run("6 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face6}.IsNumeric())
	})
	t.Run("7 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face7}.IsNumeric())
	})
	t.Run("8 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face8}.IsNumeric())
	})
	t.Run("9 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face9}.IsNumeric())
	})
	t.Run("10 is numeric", func(t *testing.T) {
		assert.True(t, Card{Face: Face10}.IsNumeric())
	})
	t.Run("J is not numeric", func(t *testing.T) {
		assert.False(t, Card{Face: FaceJack}.IsNumeric())
	})
	t.Run("Q is not numeric", func(t *testing.T) {
		assert.False(t, Card{Face: FaceQueen}.IsNumeric())
	})
	t.Run("K is not numeric", func(t *testing.T) {
		assert.False(t, Card{Face: FaceKing}.IsNumeric())
	})
	t.Run("A is not numeric", func(t *testing.T) {
		assert.False(t, Card{Face: FaceAce}.IsNumeric())
	})
	t.Run("invalid face is not numeric", func(t *testing.T) {
		assert.False(t, Card{Face: "invalid"}.IsNumeric())
	})
}

func TestCard_StrictNumericValue(t *testing.T) {
	t.Run("CardFace2 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face2, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 2, nv)
	})

	t.Run("CardFace3 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face3, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 3, nv)
	})

	t.Run("CardFace4 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face4, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 4, nv)
	})

	t.Run("CardFace5 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face5, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 5, nv)
	})

	t.Run("CardFace6 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face6, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 6, nv)
	})

	t.Run("CardFace7 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face7, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 7, nv)
	})

	t.Run("CardFace8 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face8, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 8, nv)
	})

	t.Run("CardFace9 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face9, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 9, nv)
	})

	t.Run("CardFace10 produces no error and validly converted", func(t *testing.T) {
		c := Card{Face: Face10, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.NoError(t, err)
		assert.Equal(t, 10, nv)
	})

	t.Run("CardFaceJack produces 0, error", func(t *testing.T) {
		c := Card{Face: FaceJack, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.Error(t, err)
		assert.Equal(t, 0, nv)
	})

	t.Run("CardFaceQueen produces 0, error", func(t *testing.T) {
		c := Card{Face: FaceQueen, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.Error(t, err)
		assert.Equal(t, 0, nv)
	})

	t.Run("CardFaceKing produces 0, error", func(t *testing.T) {
		c := Card{Face: FaceKing, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.Error(t, err)
		assert.Equal(t, 0, nv)
	})

	t.Run("CardFaceAce produces 0, error", func(t *testing.T) {
		c := Card{Face: FaceAce, Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.Error(t, err)
		assert.Equal(t, 0, nv)
	})

	t.Run("invalid card face produces 0, error", func(t *testing.T) {
		c := Card{Face: "invalid", Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.Error(t, err)
		assert.Equal(t, 0, nv)
	})

	t.Run("invalid numeric card face produces 0, error", func(t *testing.T) {
		c := Card{Face: "1", Suit: SuitDiamonds}
		nv, err := c.StrictNumericValue()
		require.Error(t, err)
		assert.Equal(t, 0, nv)
	})

}

func TestCard_NumericValue(t *testing.T) {
	t.Run("Numeric -> StrictNumeric", func(t *testing.T) {
		for i := 2; i <= 10; i++ {
			c := Card{Face: strconv.Itoa(i), Suit: SuitDiamonds}
			nv, err := c.StrictNumericValue()
			require.NoError(t, err)
			assert.Equal(t, i, c.NumericValue())
			assert.Equal(t, nv, c.NumericValue())
		}
	})
	t.Run("Jack -> 11", func(t *testing.T) {
		c := Card{Face: FaceJack, Suit: SuitDiamonds}
		assert.Equal(t, 11, c.NumericValue())
	})
	t.Run("Queen -> 12", func(t *testing.T) {
		c := Card{Face: FaceQueen, Suit: SuitDiamonds}
		assert.Equal(t, 12, c.NumericValue())
	})
	t.Run("King -> 13", func(t *testing.T) {
		c := Card{Face: FaceKing, Suit: SuitDiamonds}
		assert.Equal(t, 13, c.NumericValue())
	})
	t.Run("Ace -> 14", func(t *testing.T) {
		c := Card{Face: FaceAce, Suit: SuitDiamonds}
		assert.Equal(t, 14, c.NumericValue())
	})
	t.Run("invalid -> 0", func(t *testing.T) {
		c := Card{Face: "invalid", Suit: SuitDiamonds}
		assert.Equal(t, 0, c.NumericValue())
	})
}
