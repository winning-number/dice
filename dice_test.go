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

func (o *FakeOption) Apply(d *dice) error {
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
	t.Run("Should return an error because we try to instanciate a new dice avec zero face", func(t *testing.T) {
		expectedErr := "The dice can't have a zero face or negative face"
		d, err := New(0)

		if assert.Error(t, err) {
			assert.EqualError(t, err, expectedErr)
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
			driverDice, ok := d.(*dice)
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
		expectedDice := dice{
			faces: map[int32]int64{
				1: 0,
				2: 0,
				3: 0,
			},
			nbFace: int32(nbFace),
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
		d := dice{
			nbFace: int32(nbFace),
			faces:  expectedFaces(nbFace),
			random: rand.New(rand.NewSource(42)),
		}

		// always equal to 1 because the random seed defined to 42
		expectedThrow := int32(1)
		expectedHistory := []int32{1}
		expectedNBThrow := 1
		expectedFacesValues := expectedFaces(nbFace)
		expectedFacesValues[1] = 1

		throw := d.Throw()
		assert.EqualValues(t, expectedThrow, throw)
		assert.EqualValues(t, expectedHistory, d.history)
		assert.EqualValues(t, expectedNBThrow, d.nbThrow)
		assert.EqualValues(t, expectedFacesValues, d.faces)
	})
}

func TestDice_SetTrow(t *testing.T) {
	t.Run("Should be ok with seed equal 2", func(t *testing.T) {
		nbFace := int32(5)
		d := dice{
			nbFace: nbFace,
			faces:  expectedFaces(nbFace),
		}
		expectedNBThrow := 1
		expectedHistory := []int32{2}
		expectedFacesValues := expectedFaces(nbFace)
		expectedFacesValues[2] = 1

		d.SetThrow(2)
		assert.EqualValues(t, expectedHistory, d.history)
		assert.EqualValues(t, expectedNBThrow, d.nbThrow)
		assert.EqualValues(t, expectedFacesValues, d.faces)
	})
}

func TestDice_History(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		d := dice{
			history: []int32{3, 5, 42},
		}
		expectedHist := []int32{3, 5, 42}

		hist := d.History()
		assert.EqualValues(t, expectedHist, hist)
	})
}

func TestDice_NBPick(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		nbFace := int32(42)
		d := dice{
			faces: expectedFaces(nbFace),
		}
		d.faces[42] = 6
		d.faces[2] = 1

		assert.EqualValues(t, 6, d.NBPick(42))
		assert.EqualValues(t, 1, d.NBPick(2))
		assert.EqualValues(t, 0, d.NBPick(3))
	})
}

func TestDice_NBThrow(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		d := dice{
			nbThrow: 10,
		}
		expectedNBThrow := 10

		nbThrow := d.NBThrow()
		assert.EqualValues(t, expectedNBThrow, nbThrow)
	})
}

func TestDice_LeastPicks(t *testing.T) {
	nbFace := int32(4)
	t.Run("Should be ok with only one result", func(t *testing.T) {
		d := dice{
			faces: expectedFaces(nbFace),
		}
		d.faces[1] = 2
		d.faces[2] = 5
		d.faces[3] = 1
		d.faces[4] = 8
		expectedLeastPick := []Face{{
			PickValue: 3,
			Number:    1,
		}}

		leastPicks := d.LeastPicks()
		assert.ElementsMatch(t, expectedLeastPick, leastPicks)
	})
	t.Run("Should be ok with any result", func(t *testing.T) {
		d := dice{
			faces: expectedFaces(nbFace),
		}
		d.faces[1] = 2
		d.faces[2] = 5
		d.faces[3] = 1
		d.faces[4] = 1
		expectedLeastPick := []Face{{
			PickValue: 3,
			Number:    1,
		}, {
			PickValue: 4,
			Number:    1,
		}}

		leastPicks := d.LeastPicks()
		assert.ElementsMatch(t, expectedLeastPick, leastPicks)
	})
}

func TestDice_MorePicks(t *testing.T) {
	nbFace := int32(4)
	t.Run("Should be ok with only one result", func(t *testing.T) {
		d := dice{
			faces: expectedFaces(nbFace),
		}
		d.faces[1] = 2
		d.faces[2] = 5
		d.faces[3] = 1
		d.faces[4] = 8
		expectedMorePick := []Face{{
			PickValue: 4,
			Number:    8,
		}}

		morePicks := d.MorePicks()
		assert.ElementsMatch(t, expectedMorePick, morePicks)
	})
	t.Run("Should be ok with any result", func(t *testing.T) {
		d := dice{
			faces: expectedFaces(nbFace),
		}
		d.faces[1] = 8
		d.faces[2] = 5
		d.faces[3] = 1
		d.faces[4] = 8
		expectedMorePick := []Face{{
			PickValue: 1,
			Number:    8,
		}, {
			PickValue: 4,
			Number:    8,
		}}

		morePicks := d.MorePicks()
		assert.ElementsMatch(t, expectedMorePick, morePicks)
	})
}

func TestDice_Faces(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		nbFace := int32(42)
		d := dice{
			faces: expectedFaces(nbFace),
		}
		d.faces[1] = 2
		d.faces[2] = 5
		d.faces[5] = 13
		d.faces[4] = 8
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

		faces := d.Faces(1, 2, 5, 4, 920)
		assert.ElementsMatch(t, expectedFacesData, faces)
	})
}

func TestDice_FacesByNBPick(t *testing.T) {
	nbFace := int32(42)
	t.Run("Should be ok without result", func(t *testing.T) {
		d := dice{
			faces: expectedFaces(nbFace),
		}
		expectedFacesData := []Face{}

		faces := d.FacesByNBPick(14)
		assert.ElementsMatch(t, expectedFacesData, faces)
	})
	t.Run("Should be ok with only one result", func(t *testing.T) {
		d := dice{
			faces: expectedFaces(nbFace),
		}
		d.faces[4] = 14
		expectedFacesData := []Face{
			{
				PickValue: 4,
				Number:    14,
			},
		}

		faces := d.FacesByNBPick(14)
		assert.ElementsMatch(t, expectedFacesData, faces)
	})
	t.Run("Should be ok with any result", func(t *testing.T) {
		d := dice{
			faces: expectedFaces(nbFace),
		}
		d.faces[4] = 14
		d.faces[8] = 14
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

		faces := d.FacesByNBPick(14)
		assert.ElementsMatch(t, expectedFacesData, faces)
	})
}

func TestDice_String(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		d := dice{
			nbFace:  42,
			nbThrow: 5,
		}
		expectedSTR := "the dice has 42 number face and has be trow 5 times.\n"
		str := d.String()

		assert.EqualValues(t, expectedSTR, str)
	})
}

func TestDice_PickAscendingOrder(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		d := dice{
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

		faces := d.PickAscendingOrder()
		assert.EqualValues(t, expectedFacesData, faces)
	})
}

func TestDice_PickDescendingOrder(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		d := dice{
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

		faces := d.PickDescendingOrder()
		assert.EqualValues(t, expectedFacesData, faces)
	})
}
