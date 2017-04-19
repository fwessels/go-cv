# go-cv

go-cv is a computer vision and image processing library for Go using Golang assembly. It is a works-in-progress wrapper around the [Simd](https://github.com/ermig1979/Simd) library. For now most work has been done on the SSE2 version.

## SIMD
The [Simd](https://github.com/ermig1979/Simd) Library is a highly optimized image processing library. It provides many useful high performance algorithms for image processing such as:
- pixel format conversion
- image scaling and filtration
- extraction of statistic information from images
- motion detection
- object detection (HAAR and LBP classifier cascades)
- classification
- neural network

The algorithms are optimized using different SIMD CPU extensions. In particular the library supports following CPU extensions:
- x86/x64: SSE, SSE2, SSE3, SSSE3, SSE4.1, SSE4.2, AVX and AVX2
- ARM: NEON

## Performance compared to OpenCV 2.x

A comparison against [go-opencv](https://github.com/lazywei/go-opencv) shows the following results:

```
                               OpenCV          SSE2
benchmark                   old ns/op     new ns/op      delta
BenchmarkGaussian-8             74338         18481    -75.14%
BenchmarkGaussianRGB-8         186024         57169    -69.27%
BenchmarkBlur-8                110155         16623    -84.91%
BenchmarkBlurRGB-8             293017         53716    -81.67%
BenchmarkMedian3x3-8           129268         23270    -82.00%
BenchmarkMedian3x3RGB-8        169857         65896    -61.21%
BenchmarkMedian5x5-8           883311        131812    -85.08%
BenchmarkMedian5x5RGB-8       1246845        388415    -68.85%
```

## Samples

See the `samples` directory for various sample algorithms.

## go-cv-simd

See the underlying package [go-cv-simd](https://github.com/fwessels/go-cv-simd/) for more information.