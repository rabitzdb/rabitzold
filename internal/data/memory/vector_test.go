package memory_test

import (
	"github.com/RoaringBitmap/roaring"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rabitzdb/rabitz/internal/data"
	"github.com/rabitzdb/rabitz/internal/data/memory"
	"math/rand"
	"time"
)

var _ = Describe("Number of values is", func() {

	var mockData memory.VectorData
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	dataset := uint64(random.Int63n(1000000))
	offset := uint64(random.Int63n(1000000))
	field := "test"
	var vectors []Vector
	BeforeEach(func(){
		mockData = memory.NewData()
	})
	Describe("one and add",func(){
		value := "value"
		JustBeforeEach(func(){
			vectors = mockData.GetVectorsForValues(dataset,offset,field,[]string{value})
		})
		Describe("one document to vector", func() {
			var documentId uint32
			BeforeEach(func(){
				documentId = uint32(random.Intn(8000000))
				mockData.AddDocument(dataset,offset,field,value,documentId)
			})
			Context("where vector doesn't exist",func(){
				Context("and no other vectors exist",func(){
					It("the vector must exist and unique",func(){
						Expect(len(vectors)).To(Equal(1))
					})
					It("the value of the vector must be the same",func(){
						Expect(vectors[0].Value).To(Equal(value))
					})
					It("cardinality must be one",func(){
						Expect(vectors[0].Bits.GetCardinality()).To(Equal(uint64(1)))
					})
					It("the document must be part",func(){
						Expect(vectors[0].Bits.Contains(documentId)).To(Equal(true))
					})
				})
				Context("and other vectors exist",func(){
					BeforeEach(func() {
						addOtherVectors(dataset,offset,field,mockData,random)
					})
					It("the vector must exist and unique",func(){
						Expect(len(vectors)).To(Equal(1))
					})
					It("the value of the vector must be the same",func(){
						Expect(vectors[0].Value).To(Equal(value))
					})
					It("cardinality must be one",func(){
						Expect(vectors[0].Bits.GetCardinality()).To(Equal(uint64(1)))
					})
					It("the document must be part",func(){
						Expect(vectors[0].Bits.Contains(documentId)).To(Equal(true))
					})
				})
			})
			Context("where vector already exists with other document",func(){
				BeforeEach(func(){
					mockData.AddDocument(dataset,offset,field,value,documentId+1)
				})
				Context("and no other vectors exist",func() {
					It("the vector must exist and unique",func(){
						Expect(len(vectors)).To(Equal(1))
					})
					It("the value of the vector must be the same",func(){
						Expect(vectors[0].Value).To(Equal(value))
					})
					It("cardinality must be two",func(){
						Expect(vectors[0].Bits.GetCardinality()).To(Equal(uint64(2)))
					})
					It("the document must be part",func(){
						Expect(vectors[0].Bits.Contains(documentId)).To(Equal(true))
					})
				})
				Context("and other vectors exist",func(){
					BeforeEach(func() {
						addOtherVectors(dataset,offset,field,mockData,random)
					})
					It("the vector must exist and unique",func(){
						Expect(len(vectors)).To(Equal(1))
					})
					It("the value of the vector must be the same",func(){
						Expect(vectors[0].Value).To(Equal(value))
					})
					It("cardinality must be two",func(){
						Expect(vectors[0].Bits.GetCardinality()).To(Equal(uint64(2)))
					})
					It("the document must be part",func(){
						Expect(vectors[0].Bits.Contains(documentId)).To(Equal(true))
					})
				})
			})
		})
		Describe("100 documents to vector", func() {
			var docs roaring.Bitmap
			BeforeEach(func() {
				docs = addOneHundredDocuments(dataset,offset,field,value,mockData,random)
			})
			Context("where vector doesn't exist",func(){
				Context("and no other vectors exist",func() {
					It("the vector must exist and unique",func(){
						Expect(len(vectors)).To(Equal(1))
					})
					It("the value of the vector must be the same",func(){
						Expect(vectors[0].Value).To(Equal(value))
					})
					It("cardinality must be 100",func(){
						Expect(vectors[0].Bits.GetCardinality()).To(Equal(uint64(100)))
					})
					It("all the documents must be part",func(){
						andVector := roaring.And(vectors[0].Bits,&docs)
						Expect(andVector.GetCardinality()).To(Equal(uint64(100)))
					})
				})
				Context("and other vectors exist",func(){
					BeforeEach(func() {
						addOtherVectors(dataset,offset,field,mockData,random)
					})
					It("the vector must exist and unique",func(){
						Expect(len(vectors)).To(Equal(1))
					})
					It("the value of the vector must be the same",func(){
						Expect(vectors[0].Value).To(Equal(value))
					})
					It("cardinality must be 100",func(){
						Expect(vectors[0].Bits.GetCardinality()).To(Equal(uint64(100)))
					})
					It("the document must be part",func(){
						andVector := roaring.And(vectors[0].Bits,&docs)
						Expect(andVector.GetCardinality()).To(Equal(uint64(100)))
					})
				})
			})
			Context("where vector already exists with other document",func(){
				var documentId uint32
				BeforeEach(func(){
					documentId = uint32(random.Intn(8000000))
					for docs.Contains(documentId) {
						documentId = uint32(random.Intn(8000000))
					}
					mockData.AddDocument(dataset,offset,field,value,documentId)
				})
				Context("and no other vectors exist",func() {
					It("the vector must exist and unique",func(){
						Expect(len(vectors)).To(Equal(1))
					})
					It("the value of the vector must be the same",func(){
						Expect(vectors[0].Value).To(Equal(value))
					})
					It("cardinality must be 101",func(){
						Expect(vectors[0].Bits.GetCardinality()).To(Equal(uint64(101)))
					})
					It("the documents must be part",func(){
						andVector := roaring.And(vectors[0].Bits,&docs)
						Expect(andVector.GetCardinality()).To(Equal(uint64(100)))
					})
				})
				Context("and other vectors exist",func(){
					BeforeEach(func() {
						addOtherVectors(dataset,offset,field,mockData,random)
					})
					It("the vector must exist and unique",func(){
						Expect(len(vectors)).To(Equal(1))
					})
					It("the value of the vector must be the same",func(){
						Expect(vectors[0].Value).To(Equal(value))
					})
					It("cardinality must increase in 100",func(){
						Expect(vectors[0].Bits.GetCardinality()).To(Equal(uint64(101)))
					})
					It("the document must be part",func(){
						andVector := roaring.And(vectors[0].Bits,&docs)
						Expect(andVector.GetCardinality()).To(Equal(uint64(100)))
					})
				})
			})
		})
	})
	Describe("100 and add",func(){
		values := make([]string,100)
		for i := 0;i<100;i++ {
			values[i] = "fvalue"+string(i)
		}
		JustBeforeEach(func(){
			vectors = mockData.GetVectors(dataset,offset,field)
		})
		Describe("one document to each vector", func() {
			var documentId uint32
			BeforeEach(func(){
				documentId = uint32(random.Intn(8000000))
				for _,value := range values {
					mockData.AddDocument(dataset,offset,field,value,documentId)
				}
			})
			Context("where vectors doesn't exist",func(){
				Context("and no other vectors exist",func(){
					It("the vectors must exist and be 100",func(){
						Expect(len(vectors)).To(Equal(100))
					})
					It("must be one vector for each value",func(){
						vectorMap := make(map[string]Vector)
						for _,vector := range vectors {
							vectorMap[vector.Value] = vector
						}
						for _,value := range values {
							Expect(vectorMap[value]).NotTo(BeNil())
						}
					})
					It("the cardinality of each vector must be one",func(){
						for _,vector := range vectors {
							Expect(vector.Bits.GetCardinality()).To(Equal(uint64(1)))
						}
					})
					It("the document must be part of each vector",func(){
						for _,vector := range vectors {
							Expect(vector.Bits.Contains(documentId)).To(Equal(true))
						}
					})
				})
				Context("and other vectors exist",func(){
					var vectorMap map[string]Vector
					BeforeEach(func(){
						addOtherVectors(dataset,offset,field,mockData,random)
					})
					JustBeforeEach(func() {
						vectorMap = make(map[string]Vector)
						for _,vector := range vectors {
							vectorMap[vector.Value] = vector
						}
					})
					It("the vectors must exist and be 200",func(){
						Expect(len(vectors)).To(Equal(200))
					})
					It("must be one vector for each value",func(){
						for _,value := range values {
							Expect(vectorMap[value]).NotTo(BeNil())
						}
					})
					It("the cardinality of each vector must be one",func(){
						for _,value := range values {
							Expect(vectorMap[value].Bits.GetCardinality()).To(Equal(uint64(1)))
						}
					})
					It("the document must be part of each vector",func(){
						for _,value := range values {
							Expect(vectorMap[value].Bits.Contains(documentId)).To(Equal(true))
						}
					})
				})
			})
			Context("where vector already exists with other document",func(){
				BeforeEach(func(){
					for _, value := range values {
						mockData.AddDocument(dataset,offset,field,value,documentId+1)
					}
				})
				Context("and no other vectors exist",func() {
					It("the vectors must exist and be 100",func(){
						Expect(len(vectors)).To(Equal(100))
					})
					It("must be one vector for each value",func(){
						vectorMap := make(map[string]Vector)
						for _,vector := range vectors {
							vectorMap[vector.Value] = vector
						}
						for _,value := range values {
							Expect(vectorMap[value]).NotTo(BeNil())
						}
					})
					It("the cardinality of each vector must be two",func(){
						for _,vector := range vectors {
							Expect(vector.Bits.GetCardinality()).To(Equal(uint64(2)))
						}
					})
					It("the document must be part of each vector",func(){
						for _,vector := range vectors {
							Expect(vector.Bits.Contains(documentId)).To(Equal(true))
							Expect(vector.Bits.Contains(documentId+1)).To(Equal(true))
						}
					})
				})
				Context("and other vectors exist",func(){
					var vectorMap map[string]Vector
					BeforeEach(func(){
						addOtherVectors(dataset,offset,field,mockData,random)
					})
					JustBeforeEach(func() {
						vectorMap = make(map[string]Vector)
						for _,vector := range vectors {
							vectorMap[vector.Value] = vector
						}
					})
					It("the vectors must exist and be 200",func(){
						Expect(len(vectors)).To(Equal(200))
					})
					It("must be one vector for each value",func(){
						for _,value := range values {
							Expect(vectorMap[value]).NotTo(BeNil())
						}
					})
					It("the cardinality of each vector must be two",func(){
						for _,value := range values {
							Expect(vectorMap[value].Bits.GetCardinality()).To(Equal(uint64(2)))
						}
					})
					It("the document must be part of each vector",func(){
						for _,value := range values {
							Expect(vectorMap[value].Bits.Contains(documentId)).To(Equal(true))
							Expect(vectorMap[value].Bits.Contains(documentId+1)).To(Equal(true))
						}
					})
				})
			})
		})
		Describe("100 documents to vector", func() {
			docsMap := make(map[string]roaring.Bitmap)
			BeforeEach(func() {
				for _,value := range values {
					docsMap[value] = addOneHundredDocuments(dataset,offset,field,value,mockData,random)
				}
			})
			Context("where vector doesn't exist",func(){
				Context("and no other vectors exist",func() {
					It("the vectors must exist and be 100",func(){
						Expect(len(vectors)).To(Equal(100))
					})
					It("must be one vector for each value",func(){
						vectorMap := make(map[string]Vector)
						for _,vector := range vectors {
							vectorMap[vector.Value] = vector
						}
						for _,value := range values {
							Expect(vectorMap[value]).NotTo(BeNil())
						}
					})
					It("the cardinality of each vector must be 100",func(){
						for _,vector := range vectors {
							Expect(vector.Bits.GetCardinality()).To(Equal(uint64(100)))
						}
					})
					It("the documents must be part of each vector",func(){
						for _,vector := range vectors {
							docs := docsMap[vector.Value]
							andVector := roaring.And(vector.Bits,&docs)
							Expect(andVector.GetCardinality()).To(Equal(uint64(100)))
						}
					})
				})
				Context("and other vectors exist",func(){
					var vectorMap map[string]Vector
					BeforeEach(func() {
						addOtherVectors(dataset,offset,field,mockData,random)
					})
					JustBeforeEach(func() {
						vectorMap = make(map[string]Vector)
						for _,vector := range vectors {
							vectorMap[vector.Value] = vector
						}
					})
					It("the vectors must exist and be 200",func(){
						Expect(len(vectors)).To(Equal(200))
					})
					It("must be one vector for each value",func(){
						for _,value := range values {
							Expect(vectorMap[value]).NotTo(BeNil())
						}
					})
					It("the cardinality of each vector must be 100",func(){
						for _,value := range values {
							Expect(vectorMap[value].Bits.GetCardinality()).To(Equal(uint64(100)))
						}
					})
					It("the documents must be part of each vector",func(){
						for _,value := range values {
							vector := vectorMap[value]
							docs := docsMap[vector.Value]
							andVector := roaring.And(vector.Bits,&docs)
							Expect(andVector.GetCardinality()).To(Equal(uint64(100)))
						}
					})
				})
			})
			Context("where vector already exists with other document",func(){
				BeforeEach(func(){
					for _,value := range values {
						documentId := uint32(random.Intn(8000000))
						docs := docsMap[value]
						for docs.Contains(documentId) {
							documentId = uint32(random.Intn(8000000))
						}
						mockData.AddDocument(dataset,offset,field,value,documentId)
					}
				})
				Context("and no other vectors exist",func() {
					It("the vectors must exist and be 100",func(){
						Expect(len(vectors)).To(Equal(100))
					})
					It("must be one vector for each value",func(){
						vectorMap := make(map[string]Vector)
						for _,vector := range vectors {
							vectorMap[vector.Value] = vector
						}
						for _,value := range values {
							Expect(vectorMap[value]).NotTo(BeNil())
						}
					})
					It("the cardinality of each vector must be 101",func(){
						for _,vector := range vectors {
							Expect(vector.Bits.GetCardinality()).To(Equal(uint64(101)))
						}
					})
					It("the documents must be part of each vector",func(){
						for _,vector := range vectors {
							docs := docsMap[vector.Value]
							andVector := roaring.And(vector.Bits,&docs)
							Expect(andVector.GetCardinality()).To(Equal(uint64(100)))
						}
					})
				})
				Context("and other vectors exist",func(){
					It("the vectors must exist and be 100",func(){
						Expect(len(vectors)).To(Equal(100))
					})
					It("must be one vector for each value",func(){
						vectorMap := make(map[string]Vector)
						for _,vector := range vectors {
							vectorMap[vector.Value] = vector
						}
						for _,value := range values {
							Expect(vectorMap[value]).NotTo(BeNil())
						}
					})
					It("the cardinality of each vector must be 101",func(){
						for _,vector := range vectors {
							Expect(vector.Bits.GetCardinality()).To(Equal(uint64(101)))
						}
					})
					It("the documents must be part of each vector",func(){
						for _,vector := range vectors {
							docs := docsMap[vector.Value]
							andVector := roaring.And(vector.Bits,&docs)
							Expect(andVector.GetCardinality()).To(Equal(uint64(100)))
						}
					})
				})
			})
		})
	})
})

func addOtherVectors(dataset uint64, offset uint64,field string, data memory.VectorData, random *rand.Rand) {
	for i := 0;i<100;i++ {
		addOneHundredDocuments(dataset,offset,field,"value"+string(i),data,random)
	}
}
func addOneHundredDocuments(dataset uint64, offset uint64,field string, value string, data memory.VectorData, random *rand.Rand) roaring.Bitmap{
	var docs = roaring.Bitmap{}
	for docs.GetCardinality() < 100 {
		document := uint32(random.Int31n(8000000))
		docs.Add(document)
		data.AddDocument(dataset,offset,field,value,document)
	}
	return docs
}
