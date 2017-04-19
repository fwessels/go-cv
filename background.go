package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

//
func BackgroundGrowRangeSlow(value, lo, hi gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundGrowRangeSlow(value, lo, hi)
}

//
func BackgroundGrowRangeFast(value, lo, hi gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundGrowRangeFast(value, lo, hi)
}

//
func BackgroundIncrementCount(value, lo, hi, loCount, hiCount gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundIncrementCount(value, lo, hi, loCount, hiCount)
}

//
func BackgroundAdjustRange(loCount, loValue, hiCount, hiValue gocvsimd.View, threshold uint8) {

	gocvsimd.SimdSse2BackgroundAdjustRange(loCount, loValue, hiCount, hiValue, uint64(threshold))
}

//
func BackgroundAdjustRangeMasked(loCount, loValue, hiCount, hiValue gocvsimd.View, threshold uint8, mask gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundAdjustRangeMasked(loCount, loValue, hiCount, hiValue, uint64(threshold), mask)
}

//
func BackgroundShiftRange(value, lo, hi gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundShiftRange(value, lo, hi)
}

//
func BackgroundShiftRangeMasked(value, lo, hi, mask gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundShiftRangeMasked(value, lo, hi, mask)
}
