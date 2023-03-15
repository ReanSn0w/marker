package bulma

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/marker"
)

const (
	// Sizes

	ThreeQuarters = "three-quarters"
	TwoThirds     = "two-thirds"
	Half          = "half"
	OneThird      = "one-third"
	OneQuarter    = "one-quarter"
	Full          = "full"

	FourFifths  = "four-fifths"
	ThreeFifths = "three-fifths"
	TwoFifths   = "two-fifths"
	OneFifth    = "one-fifth"

	One    = "one"
	Two    = "two"
	Three  = "three"
	Four   = "four"
	Five   = "five"
	Six    = "six"
	Seven  = "seven"
	Eight  = "eight"
	Nine   = "nine"
	Ten    = "ten"
	Eleven = "eleven"
	Twelve = "twelve"

	// Keywords
	Narrow    = "narrow"
	Offset    = "offset"
	Gapless   = "gapless"
	Multiline = "multiline"
	Light     = "light"

	VerticalAlignment = "vcentered"
	Centred           = "centred"

	// Screen width
	Mobile     = "mobile"
	Tablet     = "tablet"
	Desktop    = "desktop"
	Widescreen = "widescreen"
	Fullhd     = "fullhd"

	// Size
	SquareX16  = "16x16"
	SquareX24  = "24x24"
	SquareX32  = "32x32"
	SquareX48  = "48x48"
	SquareX64  = "64x64"
	SquareX96  = "96x96"
	SquareX128 = "128x128"

	Rounded = "rounded"

	// Ratio

	RatioSquare = "square"
	Ratio1by1   = "1by1"
	Ratio5by4   = "5by4"
	Ratio4by3   = "4by3"
	Ratio3by2   = "3by2"
	Ratio5by3   = "5by3"
	Ratio16by9  = "16by9"
	Ratio2by1   = "2by1"
	Ratio3by1   = "3by1"
	Ratio4by5   = "4by5"
	Ratio3by4   = "3by4"
	Ratio2by3   = "2by3"
	Ratio3by5   = "3by5"
	Ratio9by16  = "9by16"
	Ratio1by2   = "1by2"
	Ratio1by3   = "1by3"
)

func Is(suffixes ...string) view.Modificator {
	value := "is"

	for _, suffix := range suffixes {
		value += "-" + suffix
	}

	return marker.AppendClass(value)
}

func IsOffset(suffixes ...string) view.Modificator {
	return Is(append([]string{Offset}, suffixes...)...)
}

func IsGapless(suffixes ...string) view.Modificator {
	return Is(Gapless)
}

func IsMultiline(suffixes ...string) view.Modificator {
	return Is(Multiline)
}
