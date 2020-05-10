package memory

import (
	. "github.com/RoaringBitmap/roaring"
	. "github.com/rabitzdb/rabitz/internal/data"
)

type VectorData struct {
	datasets map[uint64]*dataset
}
func NewData() VectorData {
	return VectorData{make(map[uint64]*dataset)}
}
func (data *VectorData) getDataset(datasetId uint64) *dataset {
	dataset, ok := data.datasets[datasetId]
	if !ok {
		dataset = newDataset()
		data.datasets[datasetId] = dataset
	}
	return dataset
}
func (data *VectorData) getValue(datasetId uint64, offsetId uint64, fieldId string, value string) * Bitmap{
	return data.getDataset(datasetId).getValue(offsetId,fieldId,value)
}
func (data *VectorData) GetVectors(dataset uint64,offset uint64,field string) []Vector {
	return data.getDataset(dataset).getValues(offset,field)
}
func (data *VectorData) GetVectorsForValues(dataset uint64,offset uint64,field string,values []string) []Vector {
	vectors := make([]Vector,len(values))
	for index,value := range values {
		bitmap := data.getValue(dataset,offset,field,value)
		vectors[index] = Vector{Value: value,Bits: bitmap}
	}
	return vectors
}
func (data *VectorData) AddDocument(datasetId uint64,offsetId uint64,fieldId string,value string, document uint32) {
	data.getValue(datasetId,offsetId,fieldId,value).Add(document)
}

type dataset struct {
	offsets map[uint64]*offset
}
func newDataset() *dataset {
	return &dataset {make(map[uint64]*offset)}
}
func (dataset *dataset) getOffset(offsetId uint64) *offset {
	offset, ok := dataset.offsets[offsetId]
	if !ok {
		offset = newOffset()
		dataset.offsets[offsetId] = offset
	}
	return offset
}
func (dataset *dataset) getValue(offsetId uint64, fieldId string, value string) *Bitmap{
	return dataset.getOffset(offsetId).getValue(fieldId,value)
}
func (dataset *dataset) getValues(offsetId uint64, fieldId string) []Vector {
	return dataset.getOffset(offsetId).getValues(fieldId)
}

type offset struct {
	fields map[string]*field
}
func newOffset() *offset {
	return &offset {make(map[string]*field)}
}
func (offset *offset) getField(fieldId string) *field {
	field, ok := offset.fields[fieldId]
	if !ok {
		field = newField()
		offset.fields[fieldId] = field
	}
	return field
}
func (offset *offset) getValue(fieldId string,value string) * Bitmap{
	return offset.getField(fieldId).getValue(value)
}
func (offset *offset) getValues(fieldId string) []Vector {
	return offset.getField(fieldId).getValues()
}

type field struct {
	values map[string]*Bitmap
}
func newField() *field {
	return &field {make(map[string]*Bitmap)}
}
func (field *field) getValue(value string) *Bitmap {
	vector, ok := field.values[value]
	if !ok {
		vector = &Bitmap{}
		field.values[value] = vector
	}
	return vector
}
func (field *field) getValues() []Vector {
	vectors := make([]Vector,len(field.values))
	index := 0
	for value,bits := range field.values {
		vectors[index] = Vector{Value:value,Bits:bits}
		index++
	}
	return vectors
}