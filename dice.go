// Package dice allow to explore the virtual dice behavior.
// Get a history of the dice throw, the number of time a face was picked, etc.
package dice

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// Dice describe a dice and its history
//
//go:generate mockery --name=Dice --output=mocks --filename=dice.go --outpkg=mocks
type Dice interface {
	// Throw dice simulation.
	Throw() int32
	// SetThrow dice manual setting
	SetThrow(int32)
	// History (getter) return the history of the dice throw
	History() []int32
	// NBPick return the number time the face parameter was picked
	// If the face parameter is not in the dice, return 0
	NBPick(face int32) int64
	// NBThrow return the number of time the dice was thrown
	NBThrow() int64
	// LeastPick return the least picked face
	// If faces have the same number of pick, return all of them
	LeastPick() []Face
	// MostPick return the most picked face
	// If faces have the same number of pick, return all of them
	MorePick() []Face
	// FacesByNBPick return all faces with the same number of pick
	FacesByNBPick(int64) []Face
	// Face return the face detail for the given face(s)
	// If one face is not found, it will be ignored from the return statement
	Faces(faces ...int32) []Face
	// WeakestOrder return the faces in the weakest order
	WeakestOrder() (faces []Face)
	// BestOrder return the faces in the best order
	BestOrder() (faces []Face)
	// String return a string representation of the dice
	String() string
}

type diceEngine struct {
	random  *rand.Rand
	faces   map[int32]int64
	history []int32
	nbThrow int64
	nbFace  int32
}

// New instance dice returnment
func New(nbFace int32, opts ...Option) (Dice, error) {
	if nbFace <= 0 {
		return nil, ErrNewDiceSize
	}

	dice := diceEngine{
		faces:  make(map[int32]int64, nbFace),
		nbFace: nbFace,
		//nolint:gosec // We don't need a cryptographically secure random number generator
		random: rand.New(rand.NewSource(time.Now().Unix())),
	}
	for i := 1; i <= int(nbFace); i++ {
		dice.faces[int32(i)] = 0
	}

	// Apply options
	for _, opt := range opts {
		if err := opt.Apply(&dice); err != nil {
			return nil, err
		}
	}

	return &dice, nil
}

func (d *diceEngine) Throw() int32 {
	d.nbThrow++

	ret := d.random.Int31n(d.nbFace)

	// ret++ to avoid 0 value and allow the included last face value
	ret++

	d.faces[ret]++
	d.history = append(d.history, ret)

	return ret
}

func (d *diceEngine) SetThrow(face int32) {
	d.nbThrow++

	d.faces[face]++
	d.history = append(d.history, face)
}

func (d diceEngine) History() []int32 {
	return d.history
}

func (d diceEngine) NBPick(face int32) int64 {
	return d.faces[face]
}

func (d diceEngine) NBThrow() int64 {
	return d.nbThrow
}

func (d diceEngine) LeastPick() []Face {
	var faces []Face
	var min int64

	min = math.MaxInt64
	for key, val := range d.faces {
		if val < min {
			min = val
			faces = append([]Face{}, Face{
				Value:  key,
				NBPick: val,
			})
		} else if val == min {
			faces = append(faces, Face{
				Value:  key,
				NBPick: val,
			})
		}
	}

	return faces
}

func (d diceEngine) MorePick() []Face {
	var faces []Face
	var max int64

	max = -1
	for key, val := range d.faces {
		if val > max {
			max = val
			faces = append([]Face{}, Face{
				Value:  key,
				NBPick: val,
			})
		} else if val == max {
			faces = append(faces, Face{
				Value:  key,
				NBPick: val,
			})
		}
	}

	return faces
}

func (d diceEngine) Faces(faces ...int32) []Face {
	var list []Face
	for _, face := range faces {
		if val, ok := d.faces[face]; ok {
			list = append(list, Face{
				Value:  face,
				NBPick: val,
			})
		}
	}

	return list
}

func (d diceEngine) FacesByNBPick(nbPick int64) []Face {
	var faces []Face

	for key, val := range d.faces {
		if val == nbPick {
			faces = append(faces, Face{
				Value:  key,
				NBPick: val,
			})
		}
	}

	return faces
}

func (d diceEngine) String() string {
	result := fmt.Sprintf(
		"the dice has %d number face and has be trow %d times.\n",
		d.nbFace,
		d.nbThrow)

	return result
}

// If two faces has the same number pick, the first should be the least pick value
func (d diceEngine) WeakestOrder() []Face {
	faces := d.convertToFace()

	// faces order
	sort.Slice(faces, func(i, j int) bool {
		if faces[i].NBPick < faces[j].NBPick {
			return true
		}
		if faces[i].NBPick == faces[j].NBPick &&
			faces[i].Value < faces[j].Value {

			return true
		}

		return false
	})

	return faces
}

// If two faces has the same number pick, the first should be the least pick value
func (d diceEngine) BestOrder() []Face {
	faces := d.convertToFace()

	// faces order
	sort.Slice(faces, func(i, j int) bool {
		if faces[i].NBPick > faces[j].NBPick {
			return true
		}
		if faces[i].NBPick == faces[j].NBPick &&
			faces[i].Value < faces[j].Value {

			return true
		}

		return false
	})

	return faces
}

func (d diceEngine) convertToFace() []Face {
	faces := []Face{}

	for key, val := range d.faces {
		faces = append(faces, Face{
			Value:  key,
			NBPick: val,
		})
	}

	return faces
}
