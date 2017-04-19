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