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

func SegmentationFillSingleHoles(mask gocvsimd.View, index uint8) {

	gocvsimd.SimdSse2SegmentationFillSingleHoles(mask, uint64(index))
}

func SegmentationChangeIndex(mask gocvsimd.View, oldIndex, newIndex uint8) {

	gocvsimd.SimdSse2SegmentationChangeIndex(mask, uint64(oldIndex), uint64(newIndex))
}

func SegmentationPropagate2x2(parent, child, difference gocvsimd.View, currentIndex, invalidIndex, emptyIndex, differenceThreshold uint8) {

	gocvsimd.SimdSse2SegmentationPropagate2x2(parent, child, difference, uint64(currentIndex), uint64(invalidIndex), uint64(emptyIndex), uint64(differenceThreshold))
}