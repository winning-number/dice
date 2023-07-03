package dice

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errTest = errors.New("test error")
)

type FakeOption struct{}

func (o *FakeOption) Apply(d *diceEngine) error {
	return errTest
}

func expectedFaces(nbFace int32) map[int32]int64 {
	expected := make(map[int32]int64, nbFace)
	for i := 1; int32(i) <= nbFace; i++ {
		expected[int32(i)] = 0
	}

	return expected
}

func TestNew(t *testing.T) {
	t.Run(
		"Should return an error because we try to instanciate a new dice avec zero face",
		func(t *testing.T) {
			d, err := New(0)

			if assert.Error(t, err) {
				assert.EqualError(t, err, ErrNewDiceSize.Error())
				assert.Nil(t, d)
			}
		})
	t.Run("Should return an error because option fail", func(t *testing.T) {
		d, err := New(10, &FakeOption{})
		if assert.Error(t, err) {
			assert.Nil(t, d)
			assert.EqualError(t, err, errTest.Error())
		}
	})
	t.Run("Should be ok without option", func(t *testing.T) {
		nbFace := 10

		d, err := New(int32(nbFace))
		if assert.NoError(t, err) {
			driverDice, ok := d.(*diceEngine)
			if !ok {
				t.Fatal("New return dice that not implement the *dice type")
			}
			assert.EqualValues(t, driverDice.nbFace, nbFace)
			assert.Len(t, driverDice.faces, nbFace)
			assert.NotNil(t, driverDice.random)
			assert.EqualValues(t, driverDice.nbThrow, 0)
		}
	})
	t.Run("Should be ok with option RandomSeed", func(t *testing.T) {
		nbFace := 3
		seedOption := 42
		expectedDice := diceEngine{
			faces: map[int32]int64{
				1: 0,
				2: 0,
				3: 0,
			},
			nbFace: int32(nbFace),
			//nolint:gosec // We don't need a cryptographically secure random number generator
			random: rand.New(rand.NewSource(42)),
		}

		d, err := New(int32(nbFace), &OptionSeedRandom{
			Seed: int64(seedOption),
		})
		if assert.NoError(t, err) {
			assert.EqualValues(t, &expectedDice, d)
		}
	})
}

func TestDice_Trow(t *testing.T) {
	t.Run("Should be ok with seed equal 2", func(t *testing.T) {
		nbFace := int32(5)
		dice := diceEngine{
			nbFace: nbFace,
			faces:  expectedFaces(nbFace),
			//nolint:gosec // We don't need a cryptographically secure random number generator
			random: rand.New(rand.NewSource(42)),
		}

		// always equal to 1 because the random seed defined to 42
		expectedThrow := int32(1)
		expectedHistory := []int32{1}
		expectedNBThrow := 1
		expectedFacesValues := expectedFaces(nbFace)
		expectedFacesValues[1] = 1

		throw := dice.Throw()
		assert.EqualValues(t, expectedThrow, throw)
		assert.EqualValues(t, expectedHistory, dice.history)
		assert.EqualValues(t, expectedNBThrow, dice.nbThrow)
		assert.EqualValues(t, expectedFacesValues, dice.faces)
	})
}

func TestDice_SetTrow(t *testing.T) {
	t.Run("Should be ok with seed equal 2", func(t *testing.T) {
		nbFace := int32(5)
		dice := diceEngine{
			nbFace: nbFace,
			faces:  expectedFaces(nbFace),
		}
		expectedNBThrow := 1
		expectedHistory := []int32{2}
		expectedFacesValues := expectedFaces(nbFace)
		expectedFacesValues[2] = 1

		dice.SetThrow(2)
		assert.EqualValues(t, expectedHistory, dice.history)
		assert.EqualValues(t, expectedNBThrow, dice.nbThrow)
		assert.EqualValues(t, expectedFacesValues, dice.faces)
	})
}

func TestDice_History(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		dice := diceEngine{
			history: []int32{3, 5, 42},
		}
		expectedHist := []int32{3, 5, 42}

		hist := dice.History()
		assert.EqualValues(t, expectedHist, hist)
	})
}

func TestDice_NBPick(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		nbFace := int32(42)
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		dice.faces[42] = 6
		dice.faces[2] = 1

		assert.EqualValues(t, 6, dice.NBPick(42))
		assert.EqualValues(t, 1, dice.NBPick(2))
		assert.EqualValues(t, 0, dice.NBPick(3))
	})
}

func TestDice_NBThrow(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		dice := diceEngine{
			nbThrow: 10,
		}
		expectedNBThrow := 10

		nbThrow := dice.NBThrow()
		assert.EqualValues(t, expectedNBThrow, nbThrow)
	})
}

func TestDice_LeastPicks(t *testing.T) {
	nbFace := int32(4)
	t.Run("Should be ok with only one result", func(t *testing.T) {
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		dice.faces[1] = 2
		dice.faces[2] = 5
		dice.faces[3] = 1
		dice.faces[4] = 8
		expectedLeastPick := []Face{{
			PickValue: 3,
			Number:    1,
		}}

		leastPicks := dice.LeastPicks()
		assert.ElementsMatch(t, expectedLeastPick, leastPicks)
	})
	t.Run("Should be ok with any result", func(t *testing.T) {
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		dice.faces[1] = 2
		dice.faces[2] = 5
		dice.faces[3] = 1
		dice.faces[4] = 1
		expectedLeastPick := []Face{{
			PickValue: 3,
			Number:    1,
		}, {
			PickValue: 4,
			Number:    1,
		}}

		leastPicks := dice.LeastPicks()
		assert.ElementsMatch(t, expectedLeastPick, leastPicks)
	})
}

func TestDice_MorePicks(t *testing.T) {
	nbFace := int32(4)
	t.Run("Should be ok with only one result", func(t *testing.T) {
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		dice.faces[1] = 2
		dice.faces[2] = 5
		dice.faces[3] = 1
		dice.faces[4] = 8
		expectedMorePick := []Face{{
			PickValue: 4,
			Number:    8,
		}}

		morePicks := dice.MorePicks()
		assert.ElementsMatch(t, expectedMorePick, morePicks)
	})
	t.Run("Should be ok with any result", func(t *testing.T) {
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		dice.faces[1] = 8
		dice.faces[2] = 5
		dice.faces[3] = 1
		dice.faces[4] = 8
		expectedMorePick := []Face{{
			PickValue: 1,
			Number:    8,
		}, {
			PickValue: 4,
			Number:    8,
		}}

		morePicks := dice.MorePicks()
		assert.ElementsMatch(t, expectedMorePick, morePicks)
	})
}

func TestDice_Faces(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		nbFace := int32(42)
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		dice.faces[1] = 2
		dice.faces[2] = 5
		dice.faces[5] = 13
		dice.faces[4] = 8
		expectedFacesData := []Face{
			{
				PickValue: 4,
				Number:    8,
			},
			{
				PickValue: 1,
				Number:    2,
			}, {
				PickValue: 5,
				Number:    13,
			}, {
				PickValue: 2,
				Number:    5,
			},
		}

		faces := dice.Faces(1, 2, 5, 4, 920)
		assert.ElementsMatch(t, expectedFacesData, faces)
	})
}

func TestDice_FacesByNBPick(t *testing.T) {
	nbFace := int32(42)
	t.Run("Should be ok without result", func(t *testing.T) {
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		expectedFacesData := []Face{}

		faces := dice.FacesByNBPick(14)
		assert.ElementsMatch(t, expectedFacesData, faces)
	})
	t.Run("Should be ok with only one result", func(t *testing.T) {
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		dice.faces[4] = 14
		expectedFacesData := []Face{
			{
				PickValue: 4,
				Number:    14,
			},
		}

		faces := dice.FacesByNBPick(14)
		assert.ElementsMatch(t, expectedFacesData, faces)
	})
	t.Run("Should be ok with any result", func(t *testing.T) {
		dice := diceEngine{
			faces: expectedFaces(nbFace),
		}
		dice.faces[4] = 14
		dice.faces[8] = 14
		expectedFacesData := []Face{
			{
				PickValue: 4,
				Number:    14,
			},
			{
				PickValue: 8,
				Number:    14,
			},
		}

		faces := dice.FacesByNBPick(14)
		assert.ElementsMatch(t, expectedFacesData, faces)
	})
}

func TestDice_String(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		dice := diceEngine{
			nbFace:  42,
			nbThrow: 5,
		}
		expectedSTR := "the dice has 42 number face and has be trow 5 times.\n"
		str := dice.String()

		assert.EqualValues(t, expectedSTR, str)
	})
}

func TestDice_PickAscendingOrder(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		dice := diceEngine{
			faces: map[int32]int64{
				1: 2,
				2: 5,
				3: 1,
				4: 2,
				5: 18,
			},
		}
		expectedFacesData := []Face{
			{
				PickValue: 3,
				Number:    1,
			},
			{
				PickValue: 1,
				Number:    2,
			},
			{
				PickValue: 4,
				Number:    2,
			},
			{
				PickValue: 2,
				Number:    5,
			},
			{
				PickValue: 5,
				Number:    18,
			},
		}

		faces := dice.PickAscendingOrder()
		assert.EqualValues(t, expectedFacesData, faces)
	})
}

func TestDice_PickDescendingOrder(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		dice := diceEngine{
			faces: map[int32]int64{
				1: 2,
				2: 5,
				3: 1,
				4: 2,
				5: 18,
			},
		}
		expectedFacesData := []Face{
			{
				PickValue: 5,
				Number:    18,
			},
			{
				PickValue: 2,
				Number:    5,
			},
			{
				PickValue: 1,
				Number:    2,
			},
			{
				PickValue: 4,
				Number:    2,
			},
			{
				PickValue: 3,
				Number:    1,
			},
		}

		faces := dice.PickDescendingOrder()
		assert.EqualValues(t, expectedFacesData, faces)
	})
}
