/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

// Binarization performs binarization of 8-bit gray image.
// All images must have 8-bit gray format and must have the same width and height.
// 	For every point:
//		dst[i] = compare(src[i], value) ? positive : negative;
// where compare(a, b) depends from compareType (see ::SimdCompareType).
func Binarization(src gocvsimd.View, value, positive, negative uint8, dst gocvsimd.View, compareType uint8) {

	gocvsimd.SimdSse2Binarization(src, uint64(value), uint64(positive), uint64(negative), dst, uint64(compareType))
}

// AveragingBinarization performs averaging binarization of 8-bit gray image.
// All images must have 8-bit gray format and must have the same width and height.
// 	For every point:
//		sum = 0; area = 0;
//		for(dy = -neighborhood; dy <= neighborhood; ++dy)
//		{
//			for(dx = -neighborhood; dx <= neighborhood; ++dx)
//			{
//				if(x + dx >= 0 && x + dx < width && y + dy >= 0 && y + dy < height)
//				{
//					area++;
//					if(compare(src[x + dx, x + dy], value))
//						sum++;
//				}
//			}
//		}
//		dst[x, y] = sum*255 > area*threshold ? positive : negative;
//where compare(a, b) depends from compareType (see ::SimdCompareType).
func AveragingBinarization(src gocvsimd.View, value uint8, neighborhood uint64, threshold, positive, negative uint8, dst gocvsimd.View, compareType uint8) {

	gocvsimd.SimdSse2AveragingBinarization(src, uint64(value), neighborhood, uint64(threshold), uint64(positive), uint64(negative), dst, uint64(compareType))
}

// ConditionalCount8u calculates number of points satisfying certain condition for 8-bit gray image.
//	For every point:
//		if(compare(src[i], value))
//		count++;
// where compare(a, b) depends from compareType (see ::SimdCompareType).
func ConditionalCount8u(src gocvsimd.View, value uint8, compareType uint8) uint32 {

	return gocvsimd.SimdSse2ConditionalCount8u(src, uint64(value), uint64(compareType))
}

// ConditionalCount16i calculates the number of points satisfying certain condition for 16-bit signed integer image.
// 	For every point:
//		if(compare(src[i], value))
//		count++;
// where compare(a, b) depends from compareType (see ::SimdCompareType).
func ConditionalCount16i(src gocvsimd.View, value int16, compareType uint8) uint32 {

	return gocvsimd.SimdSse2ConditionalCount16i(src, uint64(value), uint64(compareType))
}

// ConditionalSum calculates sum of image points when mask points satisfying certain condition.
// All images must have 8-bit gray format and must have the same width and height.
// 	For every point:
//		if(compare(mask[i], value))
//		sum += src[i];
// where compare(a, b) depends from compareType (see ::SimdCompareType).
func ConditionalSum(src, mask gocvsimd.View, value, compareType uint8) uint64 {

	return gocvsimd.SimdSse2ConditionalSum(src, mask, uint64(value), uint64(compareType))

}

// ConditionalSquareSum calculates sum of squared image points when mask points satisfying certain condition.
// All images must have 8-bit gray format and must have the same width and height.
//	For every point:
//		if(compare(mask[i], value))
//		sum += src[i]*src[i];
// where compare(a, b) depends from compareType (see ::SimdCompareType).
func ConditionalSquareSum(src, mask gocvsimd.View, value, compareType uint8) uint64 {

	return gocvsimd.SimdSse2ConditionalSquareSum(src, mask, uint64(value), uint64(compareType))
}

// ConditionalSquareGradientSum calculates sum of squared gradient of image points when mask points satisfying certain condition.
// All images must have 8-bit gray format and must have the same width and height. The image height and width must be equal or greater 3.
// 	For every point except border:
//		if(compare(mask[x, y], value))
//		{
//			dx = src[x + 1, y] - src[x - 1, y];
//			dy = src[x, y + 1] - src[x, y - 1];
//			sum += dx*dx + dy*dy;
//		}
// where compare(a, b) depends from compareType (see ::SimdCompareType).
func ConditionalSquareGradientSum(src, mask gocvsimd.View, value, compareType uint8) uint64 {

	return gocvsimd.SimdSse2ConditionalSquareGradientSum(src, mask, uint64(value), uint64(compareType))
}

// ConditionalFill fills pixels of 8-bit gray image by given value if corresponding pixels of input 8-bit gray image satisfy certain condition.
// All images must have the same width and height.
// 	For every point:
//		if(compare(src[i], threshold))
//		dst[i] = value;
// where compare(a, b) depends from compareType (see ::SimdCompareType).
func ConditionalFill(src gocvsimd.View, threshold, compareType, value uint8, dst gocvsimd.View) {

	gocvsimd.SimdSse2ConditionalFill(src, uint64(threshold), uint64(compareType), uint64(value), dst)
}

// OperationBinary8u performs given operation between two images.
// All images must have the same width, height and format (8-bit gray, 16-bit UV (UV plane of NV12 pixel format), 24-bit BGR or 32-bit BGRA).
func OperationBinary8u(a, b, dst gocvsimd.View, _type uint8)  {

	gocvsimd.SimdSse2OperationBinary8u(a, b, dst, uint64(_type))
}

// OperationBinary16i performs given operation between two images.
// All images must have the same width, height and ::SimdPixelFormatInt16 pixel format.
func OperationBinary16i(a, b, dst gocvsimd.View, _type uint8)  {

	gocvsimd.SimdSse2OperationBinary16i(a, b , dst, uint64(_type))
}

// VectorProduct calculates result 8-bit gray image as product of two vectors.
//	For all points:
//		dst[x, y] = horizontal[x]*vertical[y]/255;
func VectorProduct(vertical, horizontal, dst gocvsimd.View)  {

	gocvsimd.SimdSse2VectorProduct(vertical, horizontal, dst)
}