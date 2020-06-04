package types_test

import (
	"fmt"
	"github.com/RoaringBitmap/roaring"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rabitzdb/rabitz/internal/data"
	"github.com/rabitzdb/rabitz/internal/data/memory"
	"github.com/rabitzdb/rabitz/internal/data/types"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var _ = Describe("Integer", func() {
	var mockData memory.VectorData
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	dataset := uint64(random.Int63n(1000000))
	offset := uint64(random.Int63n(1000000))
	nPosibleFields := int32(30)

	var documents []Document
	var nDocuments int32
	var nFields int32

	BeforeEach(func() {
		mockData = memory.NewData()
	})
	JustBeforeEach(func() {
		documents = addIntegerDocuments(random, &mockData, nDocuments, nFields, dataset, offset, nPosibleFields)
	})
	Describe("1 document", func() {
		BeforeEach(func() {
			nDocuments = int32(1)
		})
		Context("with 1 number field", func() {
			BeforeEach(func() {
				nFields = int32(1)
			})
			It("the document must have the correct value in the field", func() {
				for _, document := range documents {
					Expect(
						checkIntegerDocument(dataset, offset, &document, &mockData)).To(
						BeTrue())
				}
			})
		})
		Context("with 10 number fields", func() {
			BeforeEach(func() {
				nFields = int32(10)
			})
			It("the document must have the correct value in each field", func() {
				for _, document := range documents {
					Expect(
						checkIntegerDocument(dataset, offset, &document, &mockData)).To(
						BeTrue())
				}
			})
		})
	})
	Describe("1000 documents", func() {
		BeforeEach(func() {
			nDocuments = int32(1000)
		})
		Context("with 1 number field", func() {
			BeforeEach(func() {
				nFields = int32(1)
			})
			It("the document must have the correct value in the field", func() {
				for _, document := range documents {
					Expect(
						checkIntegerDocument(dataset, offset, &document, &mockData)).To(
						BeTrue())
				}
			})
		})
		Context("with 10 number fields", func() {
			BeforeEach(func() {
				nFields = int32(10)
			})
			It("the document must have the correct value in each field", func() {
				for _, document := range documents {
					Expect(
						checkIntegerDocument(dataset, offset, &document, &mockData)).To(
						BeTrue())
				}
			})
		})
	})
})
func addIntegerDocuments(random *rand.Rand, data *memory.VectorData, nDocuments int32, nFields int32, datasetId uint64,
	offsetId uint64, nPosibleFields int32) []Document {
	documents := make([]Document, nDocuments)
	fields := make([]string, nPosibleFields)
	for i := int32(0); i < nPosibleFields; i++ {
		fields[i] = "field" + strconv.Itoa(int(i))
	}
	documentsIds := roaring.Bitmap{}
	for documentsIds.GetCardinality() < uint64(nDocuments) {
		documentsIds.Add(uint32(random.Int31()))
	}
	i := 0
	documentsIds.Iterate(func(documentId uint32) bool {
		document := NewDocument(documentId)
		fieldIndexes := roaring.Bitmap{}
		for fieldIndexes.GetCardinality() < uint64(nFields) {
			fieldIndexes.Add(uint32(random.Int31n(nPosibleFields)))
		}
		fieldIndexes.Iterate(func(fieldIndex uint32) bool {
			value := random.Int63()*int64(math.Pow(-1,float64(random.Intn(2))))
			fieldId := fields[fieldIndex]
			field := types.NewIntegerField(fieldId,value)
			document.AddField(field)
			return true
		})
		document.Insert(datasetId, offsetId, data)
		documents[i] = document
		i++
		return true
	})
	return documents
}
func checkIntegerDocument(dataset uint64, offset uint64, document *Document, data *memory.VectorData) bool {
	documentId := document.Id
	for _, field := range document.Fields {
		vectors := data.GetVectors(dataset, offset, field.Id())
		value := integerVectorsToValue(vectors,documentId)
		if field.Values()[0] != strconv.Itoa(int(value)) {
			fmt.Println(field.Id(),value," ",field.Values()[0])
			return false
		}
	}
	return true
}
func integerVectorsToValue(vectors []Vector,documentId uint32) int64 {
	value := 0
	for _,vector := range vectors {
		if vector.Bits.Contains(documentId){
			conv,_ := strconv.Atoi(vector.Value)
			value += conv
		}
	}
	return int64(value)
}
