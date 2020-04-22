package dimension

import (
	"github.com/RoaringBitmap/roaring"
)

type Category struct {
	Name string
	Docs *roaring.Bitmap
}

type Dimension struct {
	Data []Category
	Field string
}

type DimensionDefinition struct {
	Dataset uint64
	Offset uint16
	Field string
}