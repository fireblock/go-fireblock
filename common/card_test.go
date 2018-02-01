package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAllCard1(t *testing.T) {
	_, _, err := CheckAllCard("[ { \"uid\": \"HyPgw0hXf\", \"provider\": \"fireblock\", \"date\": \"20180105\" }, { \"uid\": \"fireblockdev\", \"provider\": \"github\", \"proof\": \"https://gist.github.com/fireblockdev/de9fdfca46e6d707b99ddd7043b21ec6\", \"date\": \"20180116\" }, { \"uid\": \"fireblock_io\", \"provider\": \"twitter\", \"proof\": \"https://twitter.com/fireblock_io/status/949233097060634624\", \"date\": \"20180116\" }, { \"uid\": \"fireblock.io\", \"provider\": \"https\", \"proof\": \"https://fireblock.io/.fireblock/0xe204f3469808079909cafb1e4647f47e6a5c1742\", \"date\": \"20180116\" }, { \"uid\": \"0xe204f3469808079909cafb1e4647f47e6a5c1742\", \"provider\": \"pgp\", \"date\": \"20180116\" } ]", "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.Equal(t, err, nil, "")

}
