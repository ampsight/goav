// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"
import "github.com/ampsight/goav/avcodec"

func newRational(r C.struct_AVRational) avcodec.Rational {
	return avcodec.NewRational(int(r.num), int(r.den))
}

// AvQ2D converts an AVRational to a float64.
func AvQ2D(r avcodec.Rational) float64 {
	return float64(r.Num()) / float64(r.Den())
}
