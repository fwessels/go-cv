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

// SegmentationFillSingleHoles fills single holes in mask.
// Mask must has 8-bit gray pixel format.
func SegmentationFillSingleHoles(mask gocvsimd.View, index uint8) {

	gocvsimd.SimdSse2SegmentationFillSingleHoles(mask, uint64(index))
}

// SegmentationChangeIndex changes certain index in mask.
// Mask must has 8-bit gray pixel format.
// 	For every point:
//		if(mask[i] == oldIndex)
//		mask[i] = newIndex;
func SegmentationChangeIndex(mask gocvsimd.View, oldIndex, newIndex uint8) {

	gocvsimd.SimdSse2SegmentationChangeIndex(mask, uint64(oldIndex), uint64(newIndex))
}

// SegmentationPropagate2x2 propagates mask index from parent (upper) to child (lower) level of mask pyramid with using 2x2 scan window.
// For parent and child image the following must be true:  parentWidth = (childWidth + 1)/2, parentHeight = (childHeight + 1)/2.
// All images must have 8-bit gray pixel format. Size of different image is equal to the child image.
func SegmentationPropagate2x2(parent, child, difference gocvsimd.View, currentIndex, invalidIndex, emptyIndex, differenceThreshold uint8) {

	gocvsimd.SimdSse2SegmentationPropagate2x2(parent, child, difference, uint64(currentIndex), uint64(invalidIndex), uint64(emptyIndex), uint64(differenceThreshold))
}