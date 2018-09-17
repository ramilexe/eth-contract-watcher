// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package debt_ceiling

import (
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/pit_file/debt_ceiling"
)

type MockPitFileDebtCeilingRepository struct {
	createErr                 error
	missingHeaders            []core.Header
	missingHeadersErr         error
	PassedStartingBlockNumber int64
	PassedEndingBlockNumber   int64
	PassedHeaderID            int64
	PassedModel               debt_ceiling.PitFileDebtCeilingModel
}

func (repository *MockPitFileDebtCeilingRepository) Create(headerID int64, model debt_ceiling.PitFileDebtCeilingModel) error {
	repository.PassedHeaderID = headerID
	repository.PassedModel = model
	return repository.createErr
}

func (repository *MockPitFileDebtCeilingRepository) MissingHeaders(startingBlockNumber, endingBlockNumber int64) ([]core.Header, error) {
	repository.PassedStartingBlockNumber = startingBlockNumber
	repository.PassedEndingBlockNumber = endingBlockNumber
	return repository.missingHeaders, repository.missingHeadersErr
}

func (repository *MockPitFileDebtCeilingRepository) SetMissingHeadersErr(e error) {
	repository.missingHeadersErr = e
}

func (repository *MockPitFileDebtCeilingRepository) SetMissingHeaders(headers []core.Header) {
	repository.missingHeaders = headers
}

func (repository *MockPitFileDebtCeilingRepository) SetCreateError(e error) {
	repository.createErr = e
}