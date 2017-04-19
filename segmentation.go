package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

func SegmentationFillSingleHoles(mask gocvsimd.View, index uint8) {

	gocvsimd.SimdSse2SegmentationFillSingleHoles(mask, uint64(index))
}

func SegmentationChangeIndex(mask gocvsimd.View, oldIndex, newIndex uint8) {

	gocvsimd.SimdSse2SegmentationChangeIndex(mask, uint64(oldIndex), uint64(newIndex))
}

func SegmentationPropagate2x2(parent, child, difference gocvsimd.View, currentIndex, invalidIndex, emptyIndex, differenceThreshold uint8) {

	gocvsimd.SimdSse2SegmentationPropagate2x2(parent, child, difference, uint64(currentIndex), uint64(invalidIndex), uint64(emptyIndex), uint64(differenceThreshold))
}