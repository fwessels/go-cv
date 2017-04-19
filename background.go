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
