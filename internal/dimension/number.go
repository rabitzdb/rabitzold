package dimension

import (
	"github.com/rabitzdb/rabitz/internal/data"
)

type HistogramConfig struct {
	Min int64
	Max int64
	Span int64
}

func GetIntegerHistogram(definition DimensionDefinition,data data.Data){

}
