package dice

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// Dice describe a dice and the history
type Dice interface {
	Throw() int32
	SetThrow(int32)
	History() []int32
	NBPick(face int32) int64
	NBThrow() int64
	LeastPicks() []Face
	MorePicks() []Face
	FacesByNBPick(int64) []Face
	Faces(faces ...int32) []Face
	PickAscendingOrder() (faces []Face)
	PickDescendingOrder() (faces []Face)
	String() string
}

type dice struct {
	nbThrow int64
	faces   map[int32]int64
	history []int32
	nbFace  int32
	random  *rand.Rand
}

// New instance dice returnment
func New(nbFace int32, opts ...Option) (Dice, error) {
	if nbFace <= 0 {
		return nil, errors.New("The dice can't have a zero face or negative face")
	}

	d := dice{
		faces:  make(map[int32]int64, nbFace),
		nbFace: nbFace,
		random: rand.New(rand.NewSource(time.Now().Unix())),
	}
	for i := 1; i <= int(nbFace); i++ {
		d.faces[int32(i)] = 0
	}

	// Apply options
	for _, opt := range opts {
		if err := opt.Apply(&d); err != nil {
			return nil, err
		}
	}

	return &d, nil
}

// Throw dice simulation.
func (d *dice) Throw() int32 {
	d.nbThrow++
	ret := rand.Int31n(d.nbFace)
	if d.random != nil {
		ret = d.random.Int31n(d.nbFace)
	}

	// ret++ to avoid 0 value and allow the included last face value
	ret++

	d.faces[ret]++
	d.history = append(d.history, ret)
	return ret
}

// SetThrow dice manual setting
func (d *dice) SetThrow(face int32) {
	d.nbThrow++

	d.faces[face]++
	d.history = append(d.history, face)
}

// History (getter)
func (d dice) History() []int32 {
	return d.history
}

// NBPick return the number time the face parameter was picked
func (d dice) NBPick(face int32) int64 {
	return d.faces[face]
}

// NBThrow getter value
func (d dice) NBThrow() int64 {
	return d.nbThrow
}

// LeastPicks return the least picked faces
func (d dice) LeastPicks() []Face {
	var faces []Face
	var min int64

	min = math.MaxInt64
	for key, val := range d.faces {
		if val < min {
			min = val
			faces = append([]Face{}, Face{
				PickValue: key,
				Number:    val,
			})
		} else if val == min {
			faces = append(faces, Face{
				PickValue: key,
				Number:    val,
			})
		}
	}

	return faces
}

// MorePicks return the more picked faces
func (d dice) MorePicks() []Face {
	var faces []Face
	var max int64

	max = -1
	for key, val := range d.faces {
		if val > max {
			max = val
			faces = append([]Face{}, Face{
				PickValue: key,
				Number:    val,
			})
		} else if val == max {
			faces = append(faces, Face{
				PickValue: key,
				Number:    val,
			})
		}
	}

	return faces
}

// Faces return face list to the given faces.
// If one face is not found, it will be ignored from the return statement
func (d dice) Faces(faces ...int32) []Face {
	var list []Face
	for _, face := range faces {
		if val, ok := d.faces[face]; ok {
			list = append(list, Face{
				PickValue: face,
				Number:    val,
			})
		}
	}

	return list
}

// FacesByNBPick return a face list which match with the nb pick parameter
func (d dice) FacesByNBPick(nbPick int64) []Face {
	var faces []Face

	for key, val := range d.faces {
		if val == nbPick {
			faces = append(faces, Face{
				PickValue: key,
				Number:    val,
			})
		}
	}

	return faces
}

// String return a printable data info about the dice
func (d dice) String() string {
	result := fmt.Sprintf("the dice has %d number face and has be trow %d times.\n", d.nbFace, d.nbThrow)
	return result
}

// PickAscendingOrder return the faces in the acending order
// If two faces has the same number pick, the first should be the least pick value
func (d dice) PickAscendingOrder() []Face {
	faces := d.convertToFace()

	// faces order
	sort.Slice(faces, func(i, j int) bool {
		if faces[i].Number < faces[j].Number {
			return true
		}
		if faces[i].Number == faces[j].Number &&
			faces[i].PickValue < faces[j].PickValue {
			return true
		}
		return false
	})

	return faces
}

// PickDescendingOrder return the faces in the decsending order
// If two faces has the same number pick, the first should be the least pick value
func (d dice) PickDescendingOrder() []Face {
	faces := d.convertToFace()

	// faces order
	sort.Slice(faces, func(i, j int) bool {
		if faces[i].Number > faces[j].Number {
			return true
		}
		if faces[i].Number == faces[j].Number &&
			faces[i].PickValue < faces[j].PickValue {
			return true
		}
		return false
	})

	return faces
}

func (d dice) convertToFace() []Face {
	var faces []Face

	for key, val := range d.faces {
		faces = append(faces, Face{
			PickValue: key,
			Number:    val,
		})
	}
	return faces
}
