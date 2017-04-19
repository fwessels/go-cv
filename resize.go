package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

func StretchGray2x2(src, dst gocvsimd.View) {

	gocvsimd.SimdSse2StretchGray2x2(src, dst)
}

// TODO: Add ReduceGray2x2/3x3/4x4/5x5

// TODO: Add ShiftBilinear