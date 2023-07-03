package dice

import (
	"math/rand"
)

type Option interface {
	Apply(d *diceEngine) error
}

type OptionSeedRandom struct {
	Seed int64
}

func (o *OptionSeedRandom) Apply(d *diceEngine) error {
	if d == nil {
		return ErrOptionNoDice
	}

	//nolint:gosec // We don't need a cryptographically secure random number generator
	d.random = rand.New(rand.NewSource(o.Seed))

	return nil
}
