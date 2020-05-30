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

package dag_putters

import (
	"fmt"
	"strings"

	node "github.com/ipfs/go-ipld-format"

	"github.com/vulcanize/eth-contract-watcher/pkg/ipfs"
	"github.com/vulcanize/eth-contract-watcher/pkg/ipfs/ipld"
)

type EthHeaderDagPutter struct {
	adder *ipfs.IPFS
}

func NewEthBlockHeaderDagPutter(adder *ipfs.IPFS) *EthHeaderDagPutter {
	return &EthHeaderDagPutter{adder: adder}
}

func (bhdp *EthHeaderDagPutter) DagPut(n node.Node) (string, error) {
	header, ok := n.(*ipld.EthHeader)
	if !ok {
		return "", fmt.Errorf("EthHeaderDagPutter expected input type %T got %T", &ipld.EthHeader{}, n)
	}
	if err := bhdp.adder.Add(header); err != nil && !strings.Contains(err.Error(), duplicateKeyErrorString) {
		return "", err
	}
	return header.Cid().String(), nil
}
