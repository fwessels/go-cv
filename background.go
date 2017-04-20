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

// BackgroundGrowRangeSlow performs background update (initial grow, slow mode).
// All images must have the same width, height and format (8-bit gray).
// 	For every point:
// 		lo[i] -= value[i] < lo[i] ? 1 : 0;
// 		hi[i] += value[i] > hi[i] ? 1 : 0;
// This function is used for background updating in motion detection algorithm.
func BackgroundGrowRangeSlow(value, lo, hi gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundGrowRangeSlow(value, lo, hi)
}

// BackgroundGrowRangeFast performs background update (initial grow, fast mode).
// All images must have the same width, height and format (8-bit gray).
//     For every point:
//         lo[i] = value[i] < lo[i] ? value[i] : lo[i];
//         hi[i] = value[i] > hi[i] ? value[i] : hi[i];
// This function is used for background updating in motion detection algorithm.
func BackgroundGrowRangeFast(value, lo, hi gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundGrowRangeFast(value, lo, hi)
}

// BackgroundIncrementCount performs collection of background statistic.
// All images must have the same width, height and format (8-bit gray).
//     Updates background statistic counters for every point:
//         loCount[i] += (value[i] < loValue[i] && loCount[i] < 255) ? 1 : 0;
//         hiCount[i] += (value[i] > hiValue[i] && hiCount[i] < 255) ? 1 : 0;
// This function is used for background updating in motion detection algorithm.
func BackgroundIncrementCount(value, lo, hi, loCount, hiCount gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundIncrementCount(value, lo, hi, loCount, hiCount)
}

// BackgroundAdjustRange performs adjustment of background range.
// All images must have the same width, height and format (8-bit gray).
//     Adjusts background range for every point:
//         loValue[i] -= (loCount[i] > threshold && loValue[i] > 0) ? 1 : 0;
//         loValue[i] += (loCount[i] < threshold && loValue[i] < 255) ? 1 : 0;
//         loCount[i] = 0;
//         hiValue[i] += (hiCount[i] > threshold && hiValue[i] < 255) ? 1 : 0;
//         hiValue[i] -= (hiCount[i] < threshold && hiValue[i] > 0) ? 1 : 0;
//         hiCount[i] = 0;
// This function is used for background updating in motion detection algorithm.
func BackgroundAdjustRange(loCount, loValue, hiCount, hiValue gocvsimd.View, threshold uint8) {

	gocvsimd.SimdSse2BackgroundAdjustRange(loCount, loValue, hiCount, hiValue, uint64(threshold))
}

// BackgroundAdjustRangeMasked performs adjustment of background range with using adjust range mask.
// All images must have the same width, height and format (8-bit gray).
//     Adjusts background range for every point:
//         if(mask[i])
//         {
//             loValue[i] -= (loCount[i] > threshold && loValue[i] > 0) ? 1 : 0;
//             loValue[i] += (loCount[i] < threshold && loValue[i] < 255) ? 1 : 0;
//             loCount[i] = 0;
//             hiValue[i] += (hiCount[i] > threshold && hiValue[i] < 255) ? 1 : 0;
//             hiValue[i] -= (hiCount[i] < threshold && hiValue[i] > 0) ? 1 : 0;
//             hiCount[i] = 0;
//         }
// This function is used for background updating in motion detection algorithm.
func BackgroundAdjustRangeMasked(loCount, loValue, hiCount, hiValue gocvsimd.View, threshold uint8, mask gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundAdjustRangeMasked(loCount, loValue, hiCount, hiValue, uint64(threshold), mask)
}

// BackgroundShiftRange shifts background range.
//All images must have the same width, height and format (8-bit gray).
// 	For every point:
// 		if (value[i] > hi[i])
// 		{
// 			lo[i] = min(lo[i] + value[i] - hi[i], 255);
// 			hi[i] = value[i];
// 		}
// 		if (lo[i] > value[i])
// 		{
// 			lo[i] = value[i];
// 			hi[i] = max(hi[i] - lo[i] + value[i], 0);
//		}
// This function is used for fast background updating in motion detection algorithm.
func BackgroundShiftRange(value, lo, hi gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundShiftRange(value, lo, hi)
}

// BackgroundShiftRangeMasked shifts background range with using shift range mask.
// All images must have the same width, height and format (8-bit gray).
// 	For every point:
// 		if(mask[i])
// 		{
// 			if (value[i] > hi[i])
// 			{
// 				lo[i] = min(lo[i] + value[i] - hi[i], 255);
// 				hi[i] = value[i];
// 			}
// 			if (lo[i] > value[i])
// 			{
// 				lo[i] = value[i];
// 				hi[i] = max(hi[i] - lo[i] + value[i], 0);
// 			}
// 		}
// This function is used for fast background updating in motion detection algorithm.
func BackgroundShiftRangeMasked(value, lo, hi, mask gocvsimd.View) {

	gocvsimd.SimdSse2BackgroundShiftRangeMasked(value, lo, hi, mask)
}
