package types

import (
	. "github.com/rabitzdb/rabitz/internal/data"
	"strconv"
)

type IntegerField struct {
	id string
	value int64
}
func NewIntegerField(id string,value int64) *IntegerField {
	return &IntegerField{id: id, value:value}
}

func (field *IntegerField) Id() string {
	return field.id
}
func (field *IntegerField) Values() []string {
	return []string{strconv.Itoa(int(field.value))}
}
func (field *IntegerField) Insert(datasetId uint64,offsetId uint64,docId uint32,data VectorWriter){
	sign := int64(signum(field.value))
	abs := field.value*sign
	if abs == 0 {
		data.AddDocument(datasetId,offsetId,field.id,"0",docId)
	} else {
		indexValue := int64(1)
		for i := int64(0);i<64;i++ {
			if indexValue > abs {
				break
			}
			valid := abs & indexValue
			if valid == indexValue {
				fieldValue := strconv.Itoa(int(indexValue*sign))
				data.AddDocument(datasetId,offsetId,field.id,fieldValue,docId)
			}
			indexValue = indexValue << 1
		}
	}
}
func signum(x int64) int {
	return int((x >> 63) | int64(uint64(-x)>>63))
}
