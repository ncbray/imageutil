package imageutil

import (
	"math"
)

func LinearToSRGB(l float32) float32 {
	if l <= 0.0031308 {
		return 12.92 * l
	} else {
		return 1.055*float32(math.Pow(float64(l), 1/2.4)) - 0.055
	}
}

func SRGBToLinear(g float32) float32 {
	if g <= 0.04045 {
		return g / 12.92
	} else {
		return float32(math.Pow((float64(g)+0.055)/1.055, 2.4))
	}
}

func FloatToByte(v float32) uint8 {
	v *= 255
	if v < 0 {
		v = 0
	} else if v > 255 {
		v = 255
	}
	// Rounding?
	r := math.Floor(float64(v + 0.5))
	return uint8(r)
}

func ByteToFloat(v uint8) float32 {
	return float32(v) / float32(255)
}

func LinearToSRGBByte(v float32) uint8 {
	return FloatToByte(LinearToSRGB(v))
}

func SRGBByteToLinear(v uint8) float32 {
	return SRGBToLinear(ByteToFloat(v))
}
