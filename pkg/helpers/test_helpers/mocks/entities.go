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

package mocks

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	c2 "github.com/vulcanize/eth-header-sync/pkg/core"

	"github.com/vulcanize/eth-contract-watcher/pkg/config"
	"github.com/vulcanize/eth-contract-watcher/pkg/constants"
	"github.com/vulcanize/eth-contract-watcher/pkg/core"
	"github.com/vulcanize/eth-contract-watcher/pkg/filters"
)

var ExpectedTransferFilter = filters.LogFilter{
	Name:      constants.TusdContractAddress + "_" + "Transfer",
	Address:   constants.TusdContractAddress,
	ToBlock:   -1,
	FromBlock: 6194634,
	Topics:    core.Topics{constants.TransferEvent.Signature()},
}

var ExpectedApprovalFilter = filters.LogFilter{
	Name:      constants.TusdContractAddress + "_" + "Approval",
	Address:   constants.TusdContractAddress,
	ToBlock:   -1,
	FromBlock: 6194634,
	Topics:    core.Topics{constants.ApprovalEvent.Signature()},
}

var rawFakeHeader, _ = json.Marshal(c2.Header{})

var MockHeader1 = c2.Header{
	Hash:        "0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad123ert",
	BlockNumber: 6194632,
	Raw:         rawFakeHeader,
	Timestamp:   "50000000",
}

var MockHeader2 = c2.Header{
	Hash:        "0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad456yui",
	BlockNumber: 6194633,
	Raw:         rawFakeHeader,
	Timestamp:   "50000015",
}

var MockHeader3 = c2.Header{
	Hash:        "0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad234hfs",
	BlockNumber: 6194634,
	Raw:         rawFakeHeader,
	Timestamp:   "50000030",
}

var MockHeader4 = c2.Header{
	Hash:        "0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad234hfs",
	BlockNumber: 6194635,
	Raw:         rawFakeHeader,
	Timestamp:   "50000030",
}

var MockTransferLog1 = types.Log{
	Index:       1,
	Address:     common.HexToAddress(constants.TusdContractAddress),
	BlockNumber: 5488076,
	TxIndex:     110,
	TxHash:      common.HexToHash("0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad6546ae"),
	Topics: []common.Hash{
		common.HexToHash(constants.TransferEvent.Signature()),
		common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000af21"),
		common.HexToHash("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391"),
	},
	Data: hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a23260359000000000000000000000000000000000000000000000000392d2e2bda9c00000000000000000000000000000000000000000000000000927f41fa0a4a418000000000000000000000000000000000000000000000000000000000005adcfebe"),
}

var MockTransferLog2 = types.Log{
	Index:       3,
	Address:     common.HexToAddress(constants.TusdContractAddress),
	BlockNumber: 5488077,
	TxIndex:     2,
	TxHash:      common.HexToHash("0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad6546df"),
	Topics: []common.Hash{
		common.HexToHash(constants.TransferEvent.Signature()),
		common.HexToHash("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391"),
		common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000af21"),
	},
	Data: hexutil.MustDecode("0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000089d24a6b4ccb1b6faa2625fe562bdd9a23260359000000000000000000000000000000000000000000000000392d2e2bda9c00000000000000000000000000000000000000000000000000927f41fa0a4a418000000000000000000000000000000000000000000000000000000000005adcfebe"),
}

var MockNewOwnerLog1 = types.Log{
	Index:       1,
	Address:     common.HexToAddress(constants.EnsContractAddress),
	BlockNumber: 5488076,
	TxIndex:     110,
	TxHash:      common.HexToHash("0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad6546ae"),
	Topics: []common.Hash{
		common.HexToHash(constants.NewOwnerEvent.Signature()),
		common.HexToHash("0x000000000000000000000000c02aaa39b223helloa0e5c4f27ead9083c752553"),
		common.HexToHash("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391"),
	},
	Data: hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000af21"),
}

var MockNewOwnerLog2 = types.Log{
	Index:       3,
	Address:     common.HexToAddress(constants.EnsContractAddress),
	BlockNumber: 5488077,
	TxIndex:     2,
	TxHash:      common.HexToHash("0x135391a0962a63944e5908e6fedfff90fb4be3e3290a21017861099bad6546df"),
	Topics: []common.Hash{
		common.HexToHash(constants.NewOwnerEvent.Signature()),
		common.HexToHash("0x000000000000000000000000c02aaa39b223helloa0e5c4f27ead9083c752553"),
		common.HexToHash("0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba400"),
	},
	Data: hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000af21"),
}

var MockOrderCreatedLog = types.Log{
	Address: common.HexToAddress(constants.MarketPlaceContractAddress),
	Topics: []common.Hash{
		common.HexToHash("0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039"),
		common.HexToHash("0xffffffffffffffffffffffffffffff72ffffffffffffffffffffffffffffffd0"),
		common.HexToHash("0x00000000000000000000000083b7b6f360a9895d163ea797d9b939b9173b292a"),
	},
	Data:        hexutil.MustDecode("0x633f94affdcabe07c000231f85c752c97b9cc43966b432ec4d18641e6d178233000000000000000000000000f87e31492faf9a91b02ee0deaad50d51d56d5d4d0000000000000000000000000000000000000000000003da9fbcf4446d6000000000000000000000000000000000000000000000000000000000016db2524880"),
	BlockNumber: 8587618,
	TxHash:      common.HexToHash("0x7ad9e2f88416738f3c7ad0a6d260f71794532206a0e838299f5014b4fe81e66e"),
	TxIndex:     93,
	BlockHash:   common.HexToHash("0x06a1762b7f2e070793fc24cd785de0fa485e728832c4f3469790153ae51a56a2"),
	Index:       59,
	Removed:     false,
}

var MockSubmitVoteLog = types.Log{
	Address: common.HexToAddress(constants.MolochContractAddress),
	Topics: []common.Hash{
		common.HexToHash("0x29bf0061f2faa9daa482f061b116195432d435536d8af4ae6b3c5dd78223679b"),
		common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000061"),
		common.HexToHash("0x0000000000000000000000006ddf1b8e6d71b5b33912607098be123ffe62ae53"),
		common.HexToHash("0x00000000000000000000000037385081870ef47e055410fefd582e2a95d2960b"),
	},
	Data:        hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001"),
	BlockNumber: 8517621,
	TxHash:      common.HexToHash("0xcc7390a2099812d0dfc9baef201afbc7a44bfae145050c9dc700b77cbd3cd752"),
	TxIndex:     103,
	BlockHash:   common.HexToHash("0x3e82681d8036b1225fcaa8bcd4cdbe757b39f13468286b303cde22146385525e"),
	Index:       132,
	Removed:     false,
}

var MockConfig = config.ContractConfig{
	Network: "",
	Addresses: map[string]bool{
		"0x1234567890abcdef": true,
	},
	Abis: map[string]string{
		"0x1234567890abcdef": "fake_abi",
	},
	Events: map[string][]string{
		"0x1234567890abcdef": {"Transfer"},
	},
	Methods: map[string][]string{
		"0x1234567890abcdef": nil,
	},
	MethodArgs: map[string][]string{
		"0x1234567890abcdef": nil,
	},
	EventArgs: map[string][]string{
		"0x1234567890abcdef": nil,
	},
}
