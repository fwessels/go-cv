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
