package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

//
func TextureBoostedSaturatedGradient(src gocvsimd.View, saturation, boost uint8, dx, dy gocvsimd.View) {

	gocvsimd.SimdSse2TextureBoostedSaturatedGradient(src, uint64(saturation), uint64(boost), dx, dy)
}

//
func TextureBoostedUv(src gocvsimd.View, boost uint8, dst gocvsimd.View) {

	gocvsimd.SimdSse2TextureBoostedUv(src, uint64(boost), dst)
}

//
func TextureGetDifferenceSum(src, lo, hi gocvsimd.View) int64 {

	return gocvsimd.SimdSse2TextureGetDifferenceSum(src, lo, hi)
}

//
func TexturePerformCompensation(src gocvsimd.View, shift int, dst gocvsimd.View) {

	gocvsimd.SimdSse2TexturePerformCompensation(src, shift, dst)
}