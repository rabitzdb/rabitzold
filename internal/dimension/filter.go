package dimension

import (
	"github.com/RoaringBitmap/roaring"
)

type Filter struct {
	docs *roaring.Bitmap
	field string
}

type FilterDefinition struct {
	Dataset uint64
	Offset uint16
	Field string
	Values []string
}
