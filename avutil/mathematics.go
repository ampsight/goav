package avutil

//#cgo pkg-config: libavutil
//#include <stdint.h>
//#include <limits.h>

//#include <libavutil/mathematics.h>
import "C"
import (
	"unsafe"
)

// AvRescaleQ rescales a 64-bit integer by 2 rational numbers.
func AvRescaleQ(a int64, bq Rational, cq Rational) int64 { 
	return (int64)(C.av_rescale_q((C.int64_t)(a), *((*C.struct_AVRational)(unsafe.Pointer(&bq))), *((*C.struct_AVRational)(unsafe.Pointer(&cq)))))
}
