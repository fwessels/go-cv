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

package main

import (
	_ "log"

	_ "github.com/fwessels/go-cv"
	_ "github.com/fwessels/go-cv-simd/sse2"
)

func main() {

	// TODO: Port back from https://github.com/fwessels/go-cv-cgo/blob/master/samples/face-detect.go

	//objects := []gocv.Object{}
	//
	//img := gocv.View{}
	//
	//// Load image
	//if ok := img.Load("data/images/lena.pgm"); !ok {
	//	log.Fatal("Cannot load image")
	//}
	//
	//// Initialize new detection engine
	//detect := gocv.NewDetection()
	//
	//// Load face detection xml
	//if ok := detect.Load("data/cascades/haar_face_0.xml"); !ok {
	//	log.Fatal("face detection xml not loaded")
	//}
	//
	//// Init
	//if ok := detect.Init(img.Size()); !ok {
	//	log.Fatal("cannot init detect with image size")
	//}
	//
	//// Detect
	//objects, ok := detect.Detect(img, objects)
	//if !ok {
	//	log.Fatal("detection failed")
	//}
	//
	//if len(objects) == 0 {
	//	log.Fatal("no objects found")
	//}
	//
	//// Print found face coordinates
	//log.Printf("%+v", objects[0].Rect)
}
