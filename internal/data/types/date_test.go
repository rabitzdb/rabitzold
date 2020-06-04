package types_test

import (
	"github.com/RoaringBitmap/roaring"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rabitzdb/rabitz/internal/data"
	"github.com/rabitzdb/rabitz/internal/data/memory"
	"github.com/rabitzdb/rabitz/internal/data/types"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var _ = Describe("Date", func() {
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
		documents = addDateDocuments(random, &mockData, nDocuments, nFields, dataset, offset, nPosibleFields)
	})
	Describe("1 document", func() {
		BeforeEach(func() {
			nDocuments = int32(1)
		})
		Context("with 1 date field", func() {
			BeforeEach(func() {
				nFields = int32(1)
			})
			It("the document must have the correct value in the field", func() {
				for _, document := range documents {
					Expect(
						checkDateDocument(dataset, offset, &document, &mockData)).To(
						BeTrue())
				}
			})
		})
		Context("with 10 date fields", func() {
			BeforeEach(func() {
				nFields = int32(10)
			})
			It("the document must have the correct value in each field", func() {
				for _, document := range documents {
					Expect(
						checkDateDocument(dataset, offset, &document, &mockData)).To(
						BeTrue())
				}
			})
		})
	})
	Describe("100000 documents", func() {
		BeforeEach(func() {
			nDocuments = int32(100000)
		})
		Context("with 1 date field", func() {
			BeforeEach(func() {
				nFields = int32(1)
			})
			It("the document must have the correct value in the field", func() {
				for _, document := range documents {
					Expect(
						checkDateDocument(dataset, offset, &document, &mockData)).To(
						BeTrue())
				}
			})
		})
		Context("with 10 date fields", func() {
			BeforeEach(func() {
				nFields = int32(10)
			})
			It("the document must have the correct value in each field", func() {
				for _, document := range documents {
					Expect(
						checkDateDocument(dataset, offset, &document, &mockData)).To(
						BeTrue())
				}
			})
		})
	})
})
func addDateDocuments(random *rand.Rand, data *memory.VectorData, nDocuments int32, nFields int32, datasetId uint64,
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
			preValue := random.Int63n(3786912000)-2197484613
			date := time.Unix(preValue,0)
			finalValue := time.Date(date.Year(),date.Month(),date.Day(),date.Hour(),date.Minute(),0,0,time.UTC).Unix()*1000
			fieldId := fields[fieldIndex]
			field := types.NewDateField(fieldId,finalValue)
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
func checkDateDocument(dataset uint64, offset uint64, document *Document, data *memory.VectorData) bool {
	documentId := document.Id
	for _, field := range document.Fields {
		vectors := data.GetVectors(dataset, offset, field.Id())
		value := dateVectorsToValue(vectors,documentId)
		if field.Values()[0] != strconv.Itoa(int(value)) {
			return false
		}
	}
	return true
}
func dateVectorsToValue(vectors []Vector,documentId uint32) int64 {
	value := 0
	for _,vector := range vectors {
		if vector.Bits.Contains(documentId){
			value += dateVectorValueToInteger(vector.Value)
		}
	}
	return int64(value)
}
func dateVectorValueToInteger(value string) int {
	split := strings.Split(value,"_")
	val, _ :=  strconv.Atoi(split[1])
	return val
}