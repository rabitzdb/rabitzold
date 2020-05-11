package memory

type Document struct {
	Id uint32
	Fields []DocumentField
}
func NewDocument(id uint32) Document {
	return Document{Id: id, Fields: []DocumentField{}}
}
func (document *Document) Insert(datasetId uint64, offsetId uint64,data *VectorData){
	for _,field := range document.Fields {
		field.insert(datasetId,offsetId,document.Id,data)
	}
}
func (document *Document) AddField(field DocumentField){
	document.Fields = append(document.Fields, field)
}
type DocumentField interface {
	Values() []string
	Id() string
	insert(datasetId uint64, offsetId uint64,docId uint32,data *VectorData)
}