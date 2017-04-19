package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

//
func Binarization(src gocvsimd.View, value, positive, negative uint8, dst gocvsimd.View, compareType uint8) {

	gocvsimd.SimdSse2Binarization(src, uint64(value), uint64(positive), uint64(negative), dst, uint64(compareType))
}

//
func AveragingBinarization(src gocvsimd.View, value uint8, neighborhood uint64, threshold, positive, negative uint8, dst gocvsimd.View, compareType uint8) {

	gocvsimd.SimdSse2AveragingBinarization(src, uint64(value), neighborhood, uint64(threshold), uint64(positive), uint64(negative), dst, uint64(compareType))
}

//
func ConditionalCount8u(src gocvsimd.View, value uint8, compareType uint8) uint32 {

	return gocvsimd.SimdSse2ConditionalCount8u(src, uint64(value), uint64(compareType))
}

//
func ConditionalCount16i(src gocvsimd.View, value int16, compareType uint8) uint32 {

	return gocvsimd.SimdSse2ConditionalCount16i(src, uint64(value), uint64(compareType))
}

//
func ConditionalSum(src, mask gocvsimd.View, value, compareType uint8) uint64 {

	return gocvsimd.SimdSse2ConditionalSum(src, mask, uint64(value), uint64(compareType))

}

//
func ConditionalSquareSum(src, mask gocvsimd.View, value, compareType uint8) uint64 {

	return gocvsimd.SimdSse2ConditionalSquareSum(src, mask, uint64(value), uint64(compareType))
}

//
func ConditionalSquareGradientSum(src, mask gocvsimd.View, value, compareType uint8) uint64 {

	return gocvsimd.SimdSse2ConditionalSquareGradientSum(src, mask, uint64(value), uint64(compareType))
}

//
func ConditionalFill(src gocvsimd.View, threshold, compareType, value uint8, dst gocvsimd.View) {

	gocvsimd.SimdSse2ConditionalFill(src, uint64(threshold), uint64(compareType), uint64(value), dst)
}

//
func OperationBinary8u(a, b, dst gocvsimd.View, _type uint8)  {

	gocvsimd.SimdSse2OperationBinary8u(a, b, dst, uint64(_type))
}

//
func OperationBinary16i(a, b, dst gocvsimd.View, _type uint8)  {

	gocvsimd.SimdSse2OperationBinary16i(a, b , dst, uint64(_type))
}

//
func VectorProduct(vertical, horizontal, dst gocvsimd.View)  {

	gocvsimd.SimdSse2VectorProduct(vertical, horizontal, dst)
}