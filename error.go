package dice

import "errors"

var ErrNewDiceSize = errors.New("the dice can't have a zero or negative face")
var ErrOptionNoDice = errors.New("could not apply an option on a nil dice")
