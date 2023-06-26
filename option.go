package dice

import (
	"errors"
	"math/rand"
)

type Option interface {
	Apply(d *dice) error
}

type OptionSeedRandom struct {
	Seed int64
}

func (o *OptionSeedRandom) Apply(d *dice) error {
	if d == nil {
		return errors.New("dice is nil")
	}

	d.random = rand.New(rand.NewSource(o.Seed))
	return nil
}
