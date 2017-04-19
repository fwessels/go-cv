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
	_ "fmt"
	_ "github.com/fwessels/go-cv"
	_ "io/ioutil"
)

func main() {

	// TODO: Port back from https://github.com/fwessels/go-cv-cgo/blob/master/samples/contour.go
	//buf, err := ioutil.ReadFile("./data/images/lena.jpg")
	//if err != nil {
	//	panic("error loading from file")
	//}
	//mat := gocv.DecodeImageMemM(buf)
	//
	//contourDetector := gocv.ContourDetector{}
	//
	//contour := contourDetector.Init(mat)
	//
	//contours := contourDetector.Detect(contour, mat)
	//
	//fmt.Print(contours)
	//
	////for contour := range contours {
	////	for c := range contour {
	////		fmt.Println(c)
	////		// Simd::DrawLine(image, contours[i][j - 1], contours[i][j], uint8_t(255));
	////	}
	////}
	//
	////image.Save("result.pgm");
}
