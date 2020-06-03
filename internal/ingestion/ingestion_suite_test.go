package ingestion_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIngestion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ingestion Suite")
}