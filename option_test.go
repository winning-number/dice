package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionSeedRandom_Apply(t *testing.T) {
	t.Run("Should return an error because dice is nil", func(t *testing.T) {
		opt := &OptionSeedRandom{}

		err := opt.Apply(nil)
		if assert.Error(t, err) {
			assert.EqualError(t, err, ErrOptionNoDice.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		opt := &OptionSeedRandom{}
		expectedDice := &diceEngine{}

		err := opt.Apply(expectedDice)
		if assert.NoError(t, err) {
			assert.NotNil(t, expectedDice.random)
		}
	})
}
