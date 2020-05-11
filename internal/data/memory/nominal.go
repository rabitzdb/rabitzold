package memory

type NominalField struct {
	id string
	values []string
}
func NewNominalField(id string,values ...string) *NominalField {
	return &NominalField{values:values}
}
func (field *NominalField) Id() string {
	return field.id
}
func (field *NominalField) Values() []string {
	return field.values
}
func (field *NominalField) insert(datasetId uint64,offsetId uint64,docId uint32,data *VectorData){
	for _,value := range field.values {
		data.AddDocument(datasetId, offsetId, field.id, value,docId)
	}
}


