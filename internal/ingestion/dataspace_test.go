package ingestion

import (
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"math/rand"
	"time"
)

//Characters to generate a random name
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//Get a random number
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//Create a random string with a defined length and a input a charset
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//Create a random string with a defined lenght
func String(length int) string {
	return stringWithCharset(length, charset)
}

//Init Database
var _ = BeforeSuite(func() {
	InitLog()
	user:= GetConfigValue("db.user")
	pass:= GetConfigValue("db.password")
	host:= GetConfigValue("db.host")
	port:= GetConfigValue("db.port")
	database:= GetConfigValue("db.database")
	InitDb(user,pass,host,port,database)
	db.Ping()
})

var _ = Describe("dataspaces", func() {

	Describe("list dataspaces", func() {
		Context("when a request ask to list dataspaces", func() {
			Listds(1)
			//time.Now().UnixNano()
			PIt("list available dataspaces and should be zero", func() {

			})

			PIt("list available dataspaces and should be one", func() {

			})

		})
	})

	Describe("create a new dataspace", func() {
		//Dataspace name
		//n := String(10)

		Context("dataspace doesn't exist", func() {
			PIt("create a new dataspace ", func() {

			})
		})

		Context("a dataspace with the same name exist", func() {
			PIt("get an error creating a dataspace that exist", func() {

			})
		})

		Context("dataspace should contain a valid structure", func() {
			PIt("create a new dataspace with a valid structure", func() {

			})

			PIt("get an error creating a dataspace with same name attributes", func() {

			})

			PIt("get an error creating a dataspace with no structure", func() {

			})
		})
	})
})
