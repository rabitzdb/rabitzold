package data

import (
	"github.com/RoaringBitmap/roaring"
)

type Document struct {
	Id uint32
	Fields []DocumentField
}
func NewDocument(id uint32) Document {
	return Document{Id: id, Fields: []DocumentField{}}
}
func (document *Document) Insert(datasetId uint64, offsetId uint64,data VectorWriter){
	for _,field := range document.Fields {
		field.Insert(datasetId,offsetId,document.Id,data)
	}
}
func (document *Document) AddField(field DocumentField){
	document.Fields = append(document.Fields, field)
}
type DocumentField interface {
	Values() []string
	Id() string
	Insert(datasetId uint64, offsetId uint64,docId uint32,data VectorWriter)
}

type Vector struct {
	Value string
	Bits *roaring.Bitmap
}