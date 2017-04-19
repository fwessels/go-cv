package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

func AlphaBlending(src, alpha, dst gocvsimd.View) {

	gocvsimd.SimdSse2AlphaBlending(src, alpha, dst)
}

func BgraToGray(bgra gocvsimd.View) (gray gocvsimd.View) {

	gray = gocvsimd.View{}
	gray.Recreate(bgra.GetWidth(), bgra.GetHeight(), gocvsimd.GRAY8)

	gocvsimd.SimdSse2BgraToGray(bgra, gray)

	return gray
}

func bgratoyuv(bgra gocvsimd.View, f func(bgra, y, u, v gocvsimd.View)) (y, u, v gocvsimd.View) {

	y, u, v = gocvsimd.View{}, gocvsimd.View{}, gocvsimd.View{}
	y.Recreate(bgra.GetWidth(), bgra.GetHeight(), gocvsimd.GRAY8)
	u.Recreate(bgra.GetWidth(), bgra.GetHeight(), gocvsimd.GRAY8)
	v.Recreate(bgra.GetWidth(), bgra.GetHeight(), gocvsimd.GRAY8)

	gocvsimd.SimdSse2BgraToYuv420p(bgra, y, u, v)

	return y, u, v
}

func BgraToYuv420p(bgra gocvsimd.View) (y, u, v gocvsimd.View) {

	return bgratoyuv(bgra, gocvsimd.SimdSse2BgraToYuv420p)
}

func BgraToYuv422p(bgra gocvsimd.View) (y, u, v gocvsimd.View) {

	return bgratoyuv(bgra, gocvsimd.SimdSse2BgraToYuv422p)
}

func BgraToYuv444p(bgra gocvsimd.View) (y, u, v gocvsimd.View) {

	return bgratoyuv(bgra, gocvsimd.SimdSse2BgraToYuv444p)
}

func GrayToBgra(gray gocvsimd.View) (bgra gocvsimd.View) {

	bgra = gocvsimd.View{}
	bgra.Recreate(gray.GetWidth(), gray.GetHeight(), gocvsimd.BGRA32)

	gocvsimd.SimdSse2GrayToBgra(gray, bgra)

	return bgra
}

func DeinterleaveUv(uv gocvsimd.View) (u, v gocvsimd.View) {

	u, v = gocvsimd.View{}, gocvsimd.View{}
	u.Recreate(uv.GetWidth(), uv.GetHeight(), gocvsimd.GRAY8)
	v.Recreate(uv.GetWidth(), uv.GetHeight(), gocvsimd.GRAY8)

	gocvsimd.SimdSse2DeinterleaveUv(uv, u, v)

	return u, v
}

func FillBgr(dst gocvsimd.View, blue, green, red int) {

	gocvsimd.SimdSse2FillBgr(dst, blue, green, red)
}

func FillBgra(dst gocvsimd.View, blue, green, red, alpha int) {

	gocvsimd.SimdSse2FillBgra(dst, blue, green, red, alpha)
}

func Int16ToGray(src gocvsimd.View) (dst gocvsimd.View) {

	dst = gocvsimd.View{}
	dst.Recreate(dst.GetWidth(), dst.GetHeight(), gocvsimd.GRAY8)

	gocvsimd.SimdSse2Int16ToGray(src, dst)

	return dst
}

func yuvtobgra(y, u, v gocvsimd.View, alpha uint8, f func(y, u, v, bgra gocvsimd.View, alpha uint64)) (bgra gocvsimd.View) {

	bgra = gocvsimd.View{}
	bgra.Recreate(y.GetWidth(), y.GetHeight(), gocvsimd.BGRA32)

	f(y, u, v, bgra, uint64(alpha))

	return bgra
}

func Yuv420pToBgra(y, u, v gocvsimd.View, alpha uint8) (bgra gocvsimd.View) {

	return yuvtobgra(y, u, v, alpha, gocvsimd.SimdSse2Yuv420pToBgra)
}

func Yuv422pToBgra(y, u, v gocvsimd.View, alpha uint8) (bgra gocvsimd.View) {

	return yuvtobgra(y, u, v, alpha, gocvsimd.SimdSse2Yuv422pToBgra)
}

func Yuv444pToBgra(y, u, v gocvsimd.View, alpha uint8) (bgra gocvsimd.View) {

	return yuvtobgra(y, u, v, alpha, gocvsimd.SimdSse2Yuv444pToBgra)
}

// TODO: Add Reorder16bit/32bit/64bit