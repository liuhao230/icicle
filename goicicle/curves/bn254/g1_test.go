// Copyright 2023 Ingonyama
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by Ingonyama DO NOT EDIT

package bn254

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFieldBN254One(t *testing.T) {
	var oneField G1BaseField
	oneField.SetOne()

	rawOneField := [8]uint32([8]uint32{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})

	assert.Equal(t, oneField.S, rawOneField)
}

func TestNewFieldBN254Zero(t *testing.T) {
	var zeroField G1BaseField
	zeroField.SetZero()

	rawZeroField := [8]uint32([8]uint32{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})

	assert.Equal(t, zeroField.S, rawZeroField)
}

func TestFieldBN254ToBytesLe(t *testing.T) {
	var p G1ProjectivePoint
	p.Random()

	expected := make([]byte, len(p.X.S)*4) // each uint32 takes 4 bytes
	for i, v := range p.X.S {
		binary.LittleEndian.PutUint32(expected[i*4:], v)
	}

	assert.Equal(t, p.X.ToBytesLe(), expected)
	assert.Equal(t, len(p.X.ToBytesLe()), 32)
}

func TestNewPointBN254Zero(t *testing.T) {
	var pointZero G1ProjectivePoint
	pointZero.SetZero()

	var baseOne G1BaseField
	baseOne.SetOne()

	var zeroSanity G1BaseField
	zeroSanity.SetZero()

	assert.Equal(t, pointZero.X, zeroSanity)
	assert.Equal(t, pointZero.Y, baseOne)
	assert.Equal(t, pointZero.Z, zeroSanity)
}

func TestBN254Eq(t *testing.T) {
	var p1 G1ProjectivePoint
	p1.Random()
	var p2 G1ProjectivePoint
	p2.Random()

	assert.Equal(t, p1.Eq(&p1), true)
	assert.Equal(t, p1.Eq(&p2), false)
}

func TestBN254StripZ(t *testing.T) {
	var p1 G1ProjectivePoint
	p1.Random()

	p2ZLess := p1.StripZ()

	assert.IsType(t, G1PointAffine{}, *p2ZLess)
	assert.Equal(t, p1.X, p2ZLess.X)
	assert.Equal(t, p1.Y, p2ZLess.Y)
}

func TestPointBN254fromLimbs(t *testing.T) {
	var p G1ProjectivePoint
	p.Random()

	x := p.X.Limbs()
	y := p.Y.Limbs()
	z := p.Z.Limbs()

	xSlice := x[:]
	ySlice := y[:]
	zSlice := z[:]

	var pFromLimbs G1ProjectivePoint
	pFromLimbs.FromLimbs(&xSlice, &ySlice, &zSlice)

	assert.Equal(t, pFromLimbs, p)
}

func TestNewPointAffineNoInfinityBN254Zero(t *testing.T) {
	var zeroP G1PointAffine
	zeroP.SetZero()

	var zeroSanity G1BaseField
	zeroSanity.SetZero()

	assert.Equal(t, zeroP.X, zeroSanity)
	assert.Equal(t, zeroP.Y, zeroSanity)
}

func TestPointAffineNoInfinityBN254FromLimbs(t *testing.T) {
	// Initialize your test values
	x := [8]uint32{1, 2, 3, 4, 5, 6, 7, 8}
	y := [8]uint32{9, 10, 11, 12, 13, 14, 15, 16}
	xSlice := x[:]
	ySlice := y[:]

	// Execute your function
	var result G1PointAffine
	result.FromLimbs(&xSlice, &ySlice)

	var xBase G1BaseField
	var yBase G1BaseField
	xBase.FromLimbs(x)
	yBase.FromLimbs(y)

	// Define your expected result
	expected := &G1PointAffine{
		X: xBase,
		Y: yBase,
	}

	// Test if result is as expected
	assert.Equal(t, result, expected)
}

func TestGetFixedLimbs(t *testing.T) {
	t.Run("case of valid input of length less than 8", func(t *testing.T) {
		slice := []uint32{1, 2, 3, 4, 5, 6, 7}
		expected := [8]uint32{1, 2, 3, 4, 5, 6, 7, 0}

		result := GetFixedLimbs(&slice)
		assert.Equal(t, result, expected)
	})

	t.Run("case of valid input of length 8", func(t *testing.T) {
		slice := []uint32{1, 2, 3, 4, 5, 6, 7, 8}
		expected := [8]uint32{1, 2, 3, 4, 5, 6, 7, 8}

		result := GetFixedLimbs(&slice)
		assert.Equal(t, result, expected)
	})

	t.Run("case of empty input", func(t *testing.T) {
		slice := []uint32{}
		expected := [8]uint32{0, 0, 0, 0, 0, 0, 0, 0}

		result := GetFixedLimbs(&slice)
		assert.Equal(t, result, expected)
	})

	t.Run("case of input length greater than 8", func(t *testing.T) {
		slice := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("the code did not panic")
			}
		}()

		GetFixedLimbs(&slice)
	})
}
