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
	"errors"
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -I./include/
// #cgo LDFLAGS: -L${SRCDIR}/../../ -lbls12_377
// #include "msm.h"
import "C"

func Msm(out *G1ProjectivePoint, points []G1PointAffine, scalars []G1ScalarField, device_id int) (*G1ProjectivePoint, error) {
	if len(points) != len(scalars) {
		return nil, errors.New("error on: len(points) != len(scalars)")
	}

	pointsC := (*C.BLS12_377_affine_t)(unsafe.Pointer(&points[0]))
	scalarsC := (*C.BLS12_377_scalar_t)(unsafe.Pointer(&scalars[0]))
	outC := (*C.BLS12_377_projective_t)(unsafe.Pointer(out))
	ret := C.msm_cuda_bls12_377(outC, pointsC, scalarsC, C.size_t(len(points)), C.size_t(device_id))

	if ret != 0 {
		return nil, fmt.Errorf("msm_cuda_bls12_377 returned error code: %d", ret)
	}

	return out, nil
}

func MsmG2(out *G2Point, points []G2PointAffine, scalars []G1ScalarField, device_id int) (*G2Point, error) {
	if len(points) != len(scalars) {
		return nil, errors.New("error on: len(points) != len(scalars)")
	}

	pointsC := (*C.BLS12_377_g2_affine_t)(unsafe.Pointer(&points[0]))
	scalarsC := (*C.BLS12_377_scalar_t)(unsafe.Pointer(&scalars[0]))
	outC := (*C.BLS12_377_g2_projective_t)(unsafe.Pointer(out))

	ret := C.msm_g2_cuda_bls12_377(outC, pointsC, scalarsC, C.size_t(len(points)), C.size_t(device_id))

	if ret != 0 {
		return nil, fmt.Errorf("msm_g2_cuda_bls12_377 returned error code: %d", ret)
	}

	return out, nil
}

func MsmBatch(points *[]G1PointAffine, scalars *[]G1ScalarField, batchSize, deviceId int) ([]*G1ProjectivePoint, error) {
	// Check for nil pointers
	if points == nil || scalars == nil {
		return nil, errors.New("points or scalars is nil")
	}

	if len(*points) != len(*scalars) {
		return nil, errors.New("error on: len(points) != len(scalars)")
	}

	// Check for empty slices
	if len(*points) == 0 || len(*scalars) == 0 {
		return nil, errors.New("points or scalars is empty")
	}

	// Check for zero batchSize
	if batchSize <= 0 {
		return nil, errors.New("error on: batchSize must be greater than zero")
	}

	out := make([]*G1ProjectivePoint, batchSize)

	for i := 0; i < len(out); i++ {
		var p G1ProjectivePoint
		p.SetZero()

		out[i] = &p
	}

	outC := (*C.BLS12_377_projective_t)(unsafe.Pointer(&out[0]))
	pointsC := (*C.BLS12_377_affine_t)(unsafe.Pointer(&(*points)[0]))
	scalarsC := (*C.BLS12_377_scalar_t)(unsafe.Pointer(&(*scalars)[0]))
	msmSizeC := C.size_t(len(*points) / batchSize)
	deviceIdC := C.size_t(deviceId)
	batchSizeC := C.size_t(batchSize)

	ret := C.msm_batch_cuda_bls12_377(outC, pointsC, scalarsC, batchSizeC, msmSizeC, deviceIdC)
	if ret != 0 {
		return nil, fmt.Errorf("msm_batch_cuda_bls12_377 returned error code: %d", ret)
	}

	return out, nil
}

func MsmG2Batch(points *[]G2PointAffine, scalars *[]G1ScalarField, batchSize, deviceId int) ([]*G2Point, error) {
	// Check for nil pointers
	if points == nil || scalars == nil {
		return nil, errors.New("points or scalars is nil")
	}

	if len(*points) != len(*scalars) {
		return nil, errors.New("error on: len(points) != len(scalars)")
	}

	// Check for empty slices
	if len(*points) == 0 || len(*scalars) == 0 {
		return nil, errors.New("points or scalars is empty")
	}

	// Check for zero batchSize
	if batchSize <= 0 {
		return nil, errors.New("error on: batchSize must be greater than zero")
	}

	out := make([]*G2Point, batchSize)

	outC := (*C.BLS12_377_g2_projective_t)(unsafe.Pointer(&out[0]))
	pointsC := (*C.BLS12_377_g2_affine_t)(unsafe.Pointer(&(*points)[0]))
	scalarsC := (*C.BLS12_377_scalar_t)(unsafe.Pointer(&(*scalars)[0]))
	msmSizeC := C.size_t(len(*points) / batchSize)
	deviceIdC := C.size_t(deviceId)
	batchSizeC := C.size_t(batchSize)

	ret := C.msm_batch_g2_cuda_bls12_377(outC, pointsC, scalarsC, batchSizeC, msmSizeC, deviceIdC)
	if ret != 0 {
		return nil, fmt.Errorf("msm_batch_cuda_bls12_377 returned error code: %d", ret)
	}

	return out, nil
}

func Commit(d_out, d_scalars, d_points unsafe.Pointer, count, bucketFactor int) int {
	d_outC := (*C.BLS12_377_projective_t)(d_out)
	scalarsC := (*C.BLS12_377_scalar_t)(d_scalars)
	pointsC := (*C.BLS12_377_affine_t)(d_points)
	countC := (C.size_t)(count)
	largeBucketFactorC := C.uint(bucketFactor)

	ret := C.commit_cuda_bls12_377(d_outC, scalarsC, pointsC, countC, largeBucketFactorC, 0)

	if ret != 0 {
		return -1
	}

	return 0
}

func CommitG2(d_out, d_scalars, d_points unsafe.Pointer, count, bucketFactor int) int {
	d_outC := (*C.BLS12_377_g2_projective_t)(d_out)
	scalarsC := (*C.BLS12_377_scalar_t)(d_scalars)
	pointsC := (*C.BLS12_377_g2_affine_t)(d_points)
	countC := (C.size_t)(count)
	largeBucketFactorC := C.uint(bucketFactor)

	ret := C.commit_g2_cuda_bls12_377(d_outC, scalarsC, pointsC, countC, largeBucketFactorC, 0)

	if ret != 0 {
		return -1
	}

	return 0
}

func CommitBatch(d_out, d_scalars, d_points unsafe.Pointer, count, batch_size int) int {
	d_outC := (*C.BLS12_377_projective_t)(d_out)
	scalarsC := (*C.BLS12_377_scalar_t)(d_scalars)
	pointsC := (*C.BLS12_377_affine_t)(d_points)
	countC := (C.size_t)(count)
	batch_sizeC := (C.size_t)(batch_size)

	ret := C.commit_batch_cuda_bls12_377(d_outC, scalarsC, pointsC, countC, batch_sizeC, 0)

	if ret != 0 {
		return -1
	}

	return 0
}

func CommitG2Batch(d_out, d_scalars, d_points unsafe.Pointer, count, batch_size int) int {
	d_outC := (*C.BLS12_377_g2_projective_t)(d_out)
	scalarsC := (*C.BLS12_377_scalar_t)(d_scalars)
	pointsC := (*C.BLS12_377_g2_affine_t)(d_points)
	countC := (C.size_t)(count)
	batch_sizeC := (C.size_t)(batch_size)

	ret := C.msm_batch_g2_cuda_bls12_377(d_outC, pointsC, scalarsC, countC, batch_sizeC, 0)

	if ret != 0 {
		return -1
	}

	return 0
}
