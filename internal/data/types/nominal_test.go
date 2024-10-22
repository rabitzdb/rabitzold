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
	"time"
)

var _ = Describe("Add", func() {
	var mockData memory.VectorData
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	dataset := uint64(random.Int63n(1000000))
	offset := uint64(random.Int63n(1000000))
	nPosibleFields := int32(30)
	nPosibleValues := int32(20)

	var documents []Document
	var nDocuments int32
	var nFields int32
	var nValues int32

	BeforeEach(func() {
		mockData = memory.NewData()
	})
	JustBeforeEach(func() {
		documents = addNominalDocuments(random, &mockData, nDocuments, nFields, dataset, offset, nPosibleFields, nValues, nPosibleValues)
	})
	Describe("1 document", func() {
		BeforeEach(func() {
			nDocuments = int32(1)
		})
		Context("with 1 nominal field", func() {
			BeforeEach(func() {
				nFields = int32(1)
			})
			Context("with 1 value per field", func() {
				BeforeEach(func() {
					nValues = int32(1)
				})
				It("the document must exist only for the corresponding value", func() {
					for _, document := range documents {
						Expect(
							checkNominalDocument(dataset, offset, &document, &mockData)).To(
							BeTrue())
					}
				})
			})
			Context("with 5 value per field", func() {
				BeforeEach(func() {
					nValues = int32(5)
				})
				It("the document must exist only for each corresponding value", func() {
					for _, document := range documents {
						Expect(
							checkNominalDocument(dataset, offset, &document, &mockData)).To(
							BeTrue())
					}
				})
			})
		})
		Context("with 10 nominal fields", func() {
			BeforeEach(func() {
				nFields = int32(10)
			})
			Context("with 1 value per field", func() {
				BeforeEach(func() {
					nValues = int32(1)
				})
				It("the document must exist only for each corresponding value", func() {
					for _, document := range documents {
						Expect(
							checkNominalDocument(dataset, offset, &document, &mockData)).To(
							BeTrue())
					}
				})
			})
			Context("with 5 value per field", func() {
				BeforeEach(func() {
					nValues = int32(5)
				})
				It("the document must exist only for each corresponding value", func() {
					for _, document := range documents {
						Expect(
							checkNominalDocument(dataset, offset, &document, &mockData)).To(
							BeTrue())
					}
				})
			})
		})
	})
	Describe("100000 documents", func() {
		BeforeEach(func() {
			nDocuments = int32(100000)
		})
		Context("with 1 nominal field", func() {
			BeforeEach(func() {
				nFields = int32(1)
			})
			Context("with 1 value per field", func() {
				BeforeEach(func() {
					nValues = int32(1)
				})
				It("the documents must exist only for each corresponding value", func() {
					for _, document := range documents {
						Expect(
							checkNominalDocument(dataset, offset, &document, &mockData)).To(
							BeTrue())
					}
				})
			})
			Context("with 5 value per field", func() {
				BeforeEach(func() {
					nValues = int32(5)
				})
			})
		})
		Context("with 10 nominal fields", func() {
			BeforeEach(func() {
				nFields = int32(10)
			})
			Context("with 1 value per field", func() {
				BeforeEach(func() {
					nValues = int32(1)
				})
				It("the documents must exist only for each corresponding value", func() {
					for _, document := range documents {
						Expect(
							checkNominalDocument(dataset, offset, &document, &mockData)).To(
							BeTrue())
					}
				})
			})
			Context("with 5 value per field", func() {
				BeforeEach(func() {
					nValues = int32(5)
				})
				It("the documents must exist only for each corresponding value", func() {
					for _, document := range documents {
						Expect(
							checkNominalDocument(dataset, offset, &document, &mockData)).To(
							BeTrue())
					}
				})
			})
		})
	})
})

func addNominalDocuments(random *rand.Rand, data *memory.VectorData, nDocuments int32, nFields int32, datasetId uint64,
	offsetId uint64, nPosibleFields int32, nValues int32, nPosibleValues int32) []Document {
	documents := make([]Document, nDocuments)
	fields := make([]string, nPosibleFields)
	values := make([]string, nPosibleValues)
	for i := int32(0); i < nPosibleFields; i++ {
		fields[i] = "field" + strconv.Itoa(int(i))
	}
	for i := int32(0); i < nPosibleValues; i++ {
		values[i] = "value" + strconv.Itoa(int(i))
	}
	for i := int32(0); i < nDocuments; i++ {
		document := NewDocument(uint32(random.Intn(8000000)))
		documents[i] = document
		fieldIndexes := roaring.Bitmap{}
		for fieldIndexes.GetCardinality() < uint64(nFields) {
			fieldIndexes.Add(uint32(random.Int31n(nPosibleFields)))
		}
		fieldIndexes.Iterate(func(fieldIndex uint32) bool {
			valueIndexes := roaring.Bitmap{}
			fieldValues := make([]string, nValues)
			fieldId := fields[fieldIndex]
			for valueIndexes.GetCardinality() < uint64(nValues) {
				valueIndexes.Add(uint32(random.Int31n(nPosibleValues)))
			}
			pos := 0
			valueIndexes.Iterate(func(valueIndex uint32) bool {
				fieldValues[pos] = fieldId + values[valueIndex]
				return true
			})
			var field DocumentField = types.NewNominalField(fieldId, fieldValues...)
			document.AddField(field)
			return true
		})
		document.Insert(datasetId, offsetId, data)
	}
	return documents
}
func checkNominalDocument(dataset uint64, offset uint64, document *Document, data *memory.VectorData) bool {
	documentId := document.Id
	for _, field := range document.Fields {
		valueMap := nominalValuesToMap(field.Values())
		vectors := data.GetVectors(dataset, offset, field.Id())
		for _, vector := range vectors {
			contains := vector.Bits.Contains(documentId)
			_, valid := valueMap[vector.Value]
			if (valid && !contains) || (!valid && contains) {
				return false
			}
		}
	}
	return true
}
func nominalValuesToMap(values []string) map[string]bool {
	valueMap := make(map[string]bool)
	for _, value := range values {
		valueMap[value] = true
	}
	return valueMap
}
