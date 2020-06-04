package data

type VectorWriter interface {
	AddDocument(datasetId uint64,offsetId uint64,fieldId string,value string, document uint32)
}
