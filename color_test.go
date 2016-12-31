package imageutil

import (
	"math"
	"testing"
)

func TestByteConversion(t *testing.T) {
	v := FloatToByte(-1)
	if v != 0 {
		t.Errorf("%d != %d", v, 0)
	}
	v = FloatToByte(0)
	if v != 0 {
		t.Errorf("%d != %d", v, 0)
	}
	v = FloatToByte(1.0)
	if v != 255 {
		t.Errorf("%d != %d", v, 255)
	}
	v = FloatToByte(1.1)
	if v != 255 {
		t.Errorf("%d != %d", v, 255)
	}
}

func TestOutOfRangeFloats(t *testing.T) {
	for i := 0; i < 256; i++ {
		roundtrip := int(FloatToByte(ByteToFloat(uint8(i))))
		if i != roundtrip {
			t.Errorf("%d != %d", i, roundtrip)
		}
	}
}

func TestSRGBConversion(t *testing.T) {
	const resolution = 8192
	for i := 0; i <= resolution; i++ {
		v := float32(i) / float32(resolution)

		b := LinearToSRGBByte(v)
		roundtrip := SRGBByteToLinear(b)
		err2 := roundtrip - v
		err2 *= err2
		if err2 >= 1.0/(226*226) {
			t.Errorf("%f != %f (%f / %f) %f", v, roundtrip, SRGBByteToLinear(b-1), SRGBByteToLinear(b+1), 1.0/math.Sqrt(float64(err2)))
		}
	}
}

func TestCompleteConversion(t *testing.T) {
	for i := 0; i < 256; i++ {
		roundtrip := int(LinearToSRGBByte(SRGBByteToLinear(uint8(i))))
		if i != roundtrip {
			t.Errorf("%d != %d", i, roundtrip)
		}
	}
}
