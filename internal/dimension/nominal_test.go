package dimension_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Nominal", func() {
	PDescribe("querying for nominal dimension",func() {
		PContext("vectors are dense", func(){
			PContext("values in deep storage are repeated", func(){
				PIt("should have less values than vectors", func() {

				})
				PIt("should have exactly N values", func() {

				})
				PIt("all documents must be contained", func() {

				})
				PIt("repeated value OR should be the same as the OR of deep storage vectors", func() {

				})
			})
			PContext("values in deep storage are not repeated", func(){
				PIt("should have the same number of values than vectors", func() {

				})
				PIt("should have exactly N values", func() {

				})
				PIt("all documents must be contained", func() {

				})
				PIt("deep storage data for value must be the same than nominal", func() {

				})
			})
		})
		PContext("vectors are sparse (25%)", func(){
			PContext("values in deep storage are repeated", func(){
				PIt("should have less values than vectors", func() {

				})
				PIt("should have exactly N values", func() {

				})
				PIt("25% documents must be contained", func() {

				})
				PIt("repeated value OR should be the same as the OR of deep storage vectors", func() {

				})
			})
			PContext("values in deep storage are not repeated", func(){
				PIt("should have the same number of values than vectors", func() {

				})
				PIt("should have exactly N values", func() {

				})
				PIt("25% documents must be contained", func() {

				})
				PIt("deep storage data for value must be the same than nominal", func() {

				})
			})
		})
	})
})
