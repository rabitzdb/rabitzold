package memory

type Document struct {
	id uint32
	fields []DocumentField
}
func (document *Document) Insert(datasetId uint64, offsetId uint64,data VectorData){
	for _,field := range document.fields {
		field.Insert(datasetId,offsetId,document.id,data)
	}
}
func (document *Document) AddField(field DocumentField){
	document.fields = append(document.fields, field)
}
type DocumentField interface {
	Values() []string
	Insert(datasetId uint64, offsetId uint64,docId uint32,data VectorData)
}