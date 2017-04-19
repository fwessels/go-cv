package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

func AbsDifferenceSum(a, b gocvsimd.View) uint64 {

	return gocvsimd.SimdSse2AbsDifferenceSum(a, b)
}

func AbsDifferenceSumMasked(a, b, mask gocvsimd.View, index uint8) uint64 {

	return gocvsimd.SimdSse2AbsDifferenceSumMasked(a, b, mask, uint64(index))
}

func AbsDifferenceSums3x3(current, background gocvsimd.View) [9]uint64 {

	return gocvsimd.SimdSse2AbsDifferenceSums3x3(current, background)
}

func AbsDifferenceSums3x3Masked(current, background, mask gocvsimd.View, index uint8) [9]uint64 {

	return gocvsimd.SimdSse2AbsDifferenceSums3x3Masked(current, background, mask, uint64(index))
}

func AbsGradientSaturatedSum(src, dst gocvsimd.View) {

	gocvsimd.SimdSse2AbsGradientSaturatedSum(src, dst)
}

func AddFeatureDifference(value, lo, hi gocvsimd.View, weight uint16, difference gocvsimd.View) {

	gocvsimd.SimdSse2AddFeatureDifference(value, lo, hi, uint64(weight), difference)
}

func SquaredDifferenceSum(a, b gocvsimd.View) uint64 {

	return gocvsimd.SimdSse2SquaredDifferenceSum(a, b)
}

func SquaredDifferenceSumMasked(a, b, mask gocvsimd.View, index uint8) uint64 {

	return gocvsimd.SimdSse2SquaredDifferenceSumMasked(a, b, mask, uint64(index))
}
