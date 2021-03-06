// VulcanizeDB
// Copyright © 2019 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package abi_test

import (
	"net/http"

	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	a "github.com/vulcanize/eth-contract-watcher/pkg/abi"
	"github.com/vulcanize/eth-contract-watcher/pkg/testing"
)

var _ = Describe("ABI files", func() {

	Describe("Reading ABI files", func() {

		It("loads a valid ABI file", func() {
			path := testing.TestABIsPath + "valid_abi.json"

			contractAbi, err := a.ParseAbiFile(path)

			Expect(contractAbi).NotTo(BeNil())
			Expect(err).To(BeNil())
		})

		It("reads the contents of a valid ABI file", func() {
			path := testing.TestABIsPath + "valid_abi.json"

			contractAbi, err := a.ReadAbiFile(path)

			Expect(contractAbi).To(Equal("[{\"foo\": \"bar\"}]"))
			Expect(err).To(BeNil())
		})

		It("returns an error when the file does not exist", func() {
			path := testing.TestABIsPath + "missing_abi.json"

			contractAbi, err := a.ParseAbiFile(path)

			Expect(contractAbi).To(Equal(abi.ABI{}))
			Expect(err).To(Equal(a.ErrMissingAbiFile))
		})

		It("returns an error when the file has invalid contents", func() {
			path := testing.TestABIsPath + "invalid_abi.json"

			contractAbi, err := a.ParseAbiFile(path)

			Expect(contractAbi).To(Equal(abi.ABI{}))
			Expect(err).To(Equal(a.ErrInvalidAbiFile))
		})

		Describe("Request ABI from endpoint", func() {

			var (
				server    *ghttp.Server
				client    *a.EtherScanAPI
				abiString string
				err       error
			)

			BeforeEach(func() {
				server = ghttp.NewServer()
				client = a.NewEtherScanClient(server.URL())
				path := testing.TestABIsPath + "sample_abi.json"
				abiString, err = a.ReadAbiFile(path)

				Expect(err).NotTo(HaveOccurred())
				_, err = a.ParseAbi(abiString)
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				server.Close()
			})

			Describe("Fetching ABI from api (etherscan)", func() {
				BeforeEach(func() {

					response := fmt.Sprintf(`{"status":"1","message":"OK","result":%q}`, abiString)
					server.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("GET", "/api", "module=contract&action=getabi&address=0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"),
							ghttp.RespondWith(http.StatusOK, response),
						),
					)
				})

				It("should make a GET request with supplied contract hash", func() {

					abi, err := client.GetAbi("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07")
					Expect(server.ReceivedRequests()).Should(HaveLen(1))
					Expect(err).ShouldNot(HaveOccurred())
					Expect(abi).Should(Equal(abiString))
				})
			})
		})

		Describe("Generating etherscan endpoints based on network", func() {
			It("should return the main endpoint as the default", func() {
				url := a.GenURL("")
				Expect(url).To(Equal("https://api.etherscan.io"))
			})

			It("generates various test network endpoint if test network is supplied", func() {
				ropstenUrl := a.GenURL("ropsten")
				rinkebyUrl := a.GenURL("rinkeby")
				kovanUrl := a.GenURL("kovan")

				Expect(ropstenUrl).To(Equal("https://ropsten.etherscan.io"))
				Expect(kovanUrl).To(Equal("https://kovan.etherscan.io"))
				Expect(rinkebyUrl).To(Equal("https://rinkeby.etherscan.io"))
			})
		})
	})
})
