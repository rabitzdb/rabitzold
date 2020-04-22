package data

import (
	"github.com/RoaringBitmap/roaring"
	"math/rand"
	"math"
)

type Vector struct {
	Value string
	Bits *roaring.Bitmap
}

type Data interface {
	GetVectors(dataset uint64,offset uint16,field string) []Vector
	GetVectorsForValues(dataset uint64,offset uint16,field string,values []string) []Vector
}

type MockedData struct {
	offsetSize uint
	density float64
	values []string
}
func (data *MockedData) GetVectors(dataset uint64,offset uint16,field string) []Vector {
	vectors := make([]Vector,len(data.values))
	for i,value := range data.values {
		random := rand.New(rand.NewSource(int64(i)))
		bitset := roaring.New()
		for bitset.GetCardinality() < uint64(math.Round(data.density*float64(data.offsetSize))) {
			bitset.Add(uint32(random.Intn(int(data.offsetSize))))
		}
		vectors[i] = Vector{Value: value,Bits: bitset}
	}
	return vectors
}
func (data *MockedData) GetVectorsForValues(dataset uint64,offset uint16,field string,values []string) []Vector {
	vectors := make([]Vector,len(data.values))
	validValues := make(map[string]bool)
	for _,value := range values {
		validValues[value] = true
	}
	for i,value := range data.values {
		if validValues[value] {
			random := rand.New(rand.NewSource(int64(i)))
			bitset := roaring.New()
			for bitset.GetCardinality() < uint64(math.Round(data.density*float64(data.offsetSize))) {
				bitset.Add(uint32(random.Intn(int(data.offsetSize))))
			}
			vectors[i] = Vector{Value: value, Bits: bitset}
		}
	}
	return vectors
}