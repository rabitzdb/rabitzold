package data

import (
	"github.com/RoaringBitmap/roaring"
)

type Vector struct {
	Value string
	Bits *roaring.Bitmap
}

type VectorReader interface {
	GetVectors(dataset uint64,offset uint64,field string) []Vector
	GetVectorsForValues(dataset uint64,offset uint64,field string,values []string) []Vector
}