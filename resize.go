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

func StretchGray2x2(src, dst gocvsimd.View) {

	gocvsimd.SimdSse2StretchGray2x2(src, dst)
}

// TODO: Add ReduceGray2x2/3x3/4x4/5x5

// TODO: Add ShiftBilinear