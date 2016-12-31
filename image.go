package imageutil

import (
	"image"
)

type Image4 struct {
	Width         int
	Height        int
	Pix           []float32
	Linear        bool
	Premultiplied bool
}

func (img *Image4) ConvertPremultiplied(state bool) {
	w := img.Width
	h := img.Height
	if img.Premultiplied && !state {
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				o := (j*w + i) * 4
				a := img.Pix[o+3]
				if a > 0.0 {
					img.Pix[o+0] = img.Pix[o+0] / a
					img.Pix[o+1] = img.Pix[o+1] / a
					img.Pix[o+2] = img.Pix[o+2] / a
				}
			}
		}
		img.Premultiplied = false
	} else if !img.Premultiplied && state {
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				o := (j*w + i) * 4
				a := img.Pix[o+3]
				img.Pix[o+0] = img.Pix[o+0] * a
				img.Pix[o+1] = img.Pix[o+1] * a
				img.Pix[o+2] = img.Pix[o+2] * a
			}
		}
		img.Premultiplied = true
	}
}

func (img *Image4) ConvertLinear(state bool) {
	w := img.Width
	h := img.Height
	if img.Linear && !state {
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				o := (j*w + i) * 4
				img.Pix[o+0] = LinearToSRGB(img.Pix[o+0])
				img.Pix[o+1] = LinearToSRGB(img.Pix[o+1])
				img.Pix[o+2] = LinearToSRGB(img.Pix[o+2])
			}
		}
		img.Linear = false
	} else if !img.Linear && state {
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				o := (j*w + i) * 4
				img.Pix[o+0] = SRGBToLinear(img.Pix[o+0])
				img.Pix[o+1] = SRGBToLinear(img.Pix[o+1])
				img.Pix[o+2] = SRGBToLinear(img.Pix[o+2])
			}
		}
		img.Linear = true
	}
}

func FromNativeImage(src image.Image) *Image4 {
	b := src.Bounds()
	w := b.Size().X
	h := b.Size().Y
	dst := &Image4{
		Width:  w,
		Height: h,
		Pix:    make([]float32, w*h*4),
	}

	switch src := src.(type) {
	case *image.NRGBA:
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				so := src.PixOffset(i, j)
				do := (j*w + i) * 4
				dst.Pix[do+0] = ByteToFloat(src.Pix[so+0])
				dst.Pix[do+1] = ByteToFloat(src.Pix[so+1])
				dst.Pix[do+2] = ByteToFloat(src.Pix[so+2])
				dst.Pix[do+3] = ByteToFloat(src.Pix[so+3])
			}
		}
	case *image.RGBA:
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				so := src.PixOffset(i, j)
				do := (j*w + i) * 4
				dst.Pix[do+0] = ByteToFloat(src.Pix[so+0])
				dst.Pix[do+1] = ByteToFloat(src.Pix[so+1])
				dst.Pix[do+2] = ByteToFloat(src.Pix[so+2])
				dst.Pix[do+3] = ByteToFloat(src.Pix[so+3])
			}
			dst.Premultiplied = !src.Opaque()
		}
	default:
		panic(src)
	}
	return dst
}

func (img *Image4) ToNativeImage() *image.NRGBA {
	dst := image.NewNRGBA(image.Rect(0, 0, img.Width, img.Height))
	for j := 0; j < img.Height; j++ {
		for i := 0; i < img.Width; i++ {
			so := (j*img.Width + i) * 4
			do := dst.PixOffset(i, j)
			dst.Pix[do+0] = FloatToByte(img.Pix[so+0])
			dst.Pix[do+1] = FloatToByte(img.Pix[so+1])
			dst.Pix[do+2] = FloatToByte(img.Pix[so+2])
			dst.Pix[do+3] = FloatToByte(img.Pix[so+3])
		}
	}
	return dst
}
