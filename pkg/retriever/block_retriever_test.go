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

package retriever_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/eth-header-sync/pkg/postgres"
	"github.com/vulcanize/eth-header-sync/pkg/repository"

	"github.com/vulcanize/eth-contract-watcher/pkg/helpers/test_helpers"
	"github.com/vulcanize/eth-contract-watcher/pkg/helpers/test_helpers/mocks"
	"github.com/vulcanize/eth-contract-watcher/pkg/retriever"
)

var _ = Describe("Block Retriever", func() {
	var db *postgres.DB
	var r retriever.BlockRetriever
	var headerRepository repository.HeaderRepository

	BeforeEach(func() {
		db, _ = test_helpers.SetupDBandClient()
		headerRepository = repository.NewHeaderRepository(db)
		r = retriever.NewBlockRetriever(db)
	})

	AfterEach(func() {
		test_helpers.TearDown(db)
	})

	Describe("RetrieveFirstBlock", func() {
		It("Retrieves block number of earliest header in the database", func() {
			_, err := headerRepository.CreateOrUpdateHeader(mocks.MockHeader1)
			Expect(err).ToNot(HaveOccurred())
			_, err = headerRepository.CreateOrUpdateHeader(mocks.MockHeader2)
			Expect(err).ToNot(HaveOccurred())
			_, err = headerRepository.CreateOrUpdateHeader(mocks.MockHeader3)
			Expect(err).ToNot(HaveOccurred())

			i, err := r.RetrieveFirstBlock()
			Expect(err).NotTo(HaveOccurred())
			Expect(i).To(Equal(int64(6194632)))
		})

		It("Fails if no headers can be found in the database", func() {
			_, err := r.RetrieveFirstBlock()
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("RetrieveMostRecentBlock", func() {
		It("Retrieves the latest header's block number", func() {
			_, err := headerRepository.CreateOrUpdateHeader(mocks.MockHeader1)
			Expect(err).ToNot(HaveOccurred())
			_, err = headerRepository.CreateOrUpdateHeader(mocks.MockHeader2)
			Expect(err).ToNot(HaveOccurred())
			_, err = headerRepository.CreateOrUpdateHeader(mocks.MockHeader3)
			Expect(err).ToNot(HaveOccurred())

			i, err := r.RetrieveMostRecentBlock()
			Expect(err).ToNot(HaveOccurred())
			Expect(i).To(Equal(int64(6194634)))
		})

		It("Fails if no headers can be found in the database", func() {
			i, err := r.RetrieveMostRecentBlock()
			Expect(err).To(HaveOccurred())
			Expect(i).To(Equal(int64(0)))
		})
	})
})
