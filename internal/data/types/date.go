package types

import (
	"fmt"
	. "github.com/rabitzdb/rabitz/internal/data"
	"strconv"
	"time"
)

type DateField struct {
	id string
	value int64
}
func NewDateField(id string,value int64) *DateField {
	return &DateField{id: id, value:value}
}

func (field *DateField) Id() string {
	return field.id
}
func (field *DateField) Values() []string {
	return []string{strconv.Itoa(int(field.value))}
}
func (field *DateField) Insert(datasetId uint64,offsetId uint64,docId uint32,data VectorWriter){
	seconds := field.value / 1000
	date := time.Unix(seconds,0).UTC()
	year := date.Year()
	month := date.Month()
	day := date.Day()
	hour := date.Hour()
	minute := date.Minute()

	yearTime := time.Date(year,1,1,0,0,0,0,time.UTC).Unix()*1000
	monthTime := time.Date(year,month,1,0,0,0,0,time.UTC).Unix()*1000-yearTime
	dayTime := time.Date(year,month,day,0,0,0,0,time.UTC).Unix()*1000-yearTime-monthTime
	hourTime := int64(hour)*60*60*1000
	minuteTime := int64(minute)*60*1000

	data.AddDocument(datasetId,offsetId,field.id,fmt.Sprintf("y_%d",yearTime),docId)
	data.AddDocument(datasetId,offsetId,field.id,fmt.Sprintf("m_%d",monthTime),docId)
	data.AddDocument(datasetId,offsetId,field.id,fmt.Sprintf("d_%d",dayTime),docId)
	data.AddDocument(datasetId,offsetId,field.id,fmt.Sprintf("h_%d",hourTime),docId)
	data.AddDocument(datasetId,offsetId,field.id,fmt.Sprintf("M_%d",minuteTime),docId)
}
