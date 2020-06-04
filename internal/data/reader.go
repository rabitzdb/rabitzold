package data

/*

 */
type VectorReader interface {
	GetVectors(dataset uint64,offset uint64,field string) []Vector
	GetVectorsForValues(dataset uint64,offset uint64,field string,values []string) []Vector
}