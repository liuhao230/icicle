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

package bls12377

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNttBLS12_377BBB(t *testing.T) {
	count := 1 << 20
	scalars := GenerateScalars(count, false)

	nttResult := make([]G1ScalarField, len(scalars)) // Make a new slice with the same length
	copy(nttResult, scalars)

	assert.Equal(t, nttResult, scalars)
	NttBatch(&nttResult, false, count, 0)
	assert.NotEqual(t, nttResult, scalars)

	assert.Equal(t, nttResult, nttResult)
}

func TestNttBLS12_377CompareToGnarkDIF(t *testing.T) {
	count := 1 << 2
	scalars := GenerateScalars(count, false)

	nttResult := make([]G1ScalarField, len(scalars)) // Make a new slice with the same length
	copy(nttResult, scalars)

	assert.Equal(t, nttResult, scalars)
	Ntt(&nttResult, false, DIF, 0)
	assert.NotEqual(t, nttResult, scalars)

	assert.Equal(t, nttResult, nttResult)
}

func TestINttBLS12_377CompareToGnarkDIT(t *testing.T) {
	count := 1 << 3
	scalars := GenerateScalars(count, false)

	nttResult := make([]G1ScalarField, len(scalars)) // Make a new slice with the same length
	copy(nttResult, scalars)

	assert.Equal(t, nttResult, scalars)
	Ntt(&nttResult, true, DIT, 0)
	assert.NotEqual(t, nttResult, scalars)

	assert.Equal(t, nttResult, nttResult)
}

func TestNttBLS12_377(t *testing.T) {
	count := 1 << 3

	scalars := GenerateScalars(count, false)

	nttResult := make([]G1ScalarField, len(scalars)) // Make a new slice with the same length
	copy(nttResult, scalars)

	assert.Equal(t, nttResult, scalars)
	Ntt(&nttResult, false, NONE, 0)
	assert.NotEqual(t, nttResult, scalars)

	inttResult := make([]G1ScalarField, len(nttResult))
	copy(inttResult, nttResult)

	assert.Equal(t, inttResult, nttResult)
	Ntt(&inttResult, true, NONE, 0)
	assert.Equal(t, inttResult, scalars)
}

func TestNttBatchBLS12_377(t *testing.T) {
	count := 1 << 5
	batches := 4

	scalars := GenerateScalars(count*batches, false)

	var scalarVecOfVec [][]G1ScalarField = make([][]G1ScalarField, 0)

	for i := 0; i < batches; i++ {
		start := i * count
		end := (i + 1) * count
		batch := make([]G1ScalarField, len(scalars[start:end]))
		copy(batch, scalars[start:end])
		scalarVecOfVec = append(scalarVecOfVec, batch)
	}

	nttBatchResult := make([]G1ScalarField, len(scalars))
	copy(nttBatchResult, scalars)

	NttBatch(&nttBatchResult, false, count, 0)

	var nttResultVecOfVec [][]G1ScalarField

	for i := 0; i < batches; i++ {
		// Clone the slice
		clone := make([]G1ScalarField, len(scalarVecOfVec[i]))
		copy(clone, scalarVecOfVec[i])

		// Add it to the result vector of vectors
		nttResultVecOfVec = append(nttResultVecOfVec, clone)

		// Call the ntt_bls12_377 function
		Ntt(&nttResultVecOfVec[i], false, NONE, 0)
	}

	assert.NotEqual(t, nttBatchResult, scalars)

	// Check that the ntt of each vec of scalars is equal to the intt of the specific batch
	for i := 0; i < batches; i++ {
		if !reflect.DeepEqual(nttResultVecOfVec[i], nttBatchResult[i*count:((i+1)*count)]) {
			t.Errorf("ntt of vec of scalars not equal to intt of specific batch")
		}
	}
}

func BenchmarkNTT(b *testing.B) {
	LOG_NTT_SIZES := []int{12, 15, 20, 21, 22, 23, 24, 25, 26}

	for _, logNTTSize := range LOG_NTT_SIZES {
		nttSize := 1 << logNTTSize
		b.Run(fmt.Sprintf("NTT %d", logNTTSize), func(b *testing.B) {
			scalars := GenerateScalars(nttSize, false)

			nttResult := make([]G1ScalarField, len(scalars)) // Make a new slice with the same length
			copy(nttResult, scalars)
			for n := 0; n < b.N; n++ {
				Ntt(&nttResult, false, NONE, 0)
			}
		})
	}
}
