package txtax_test

import (
	"github.com/musinit/txtax"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Analyser_Txn1_FIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txs1, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txs1))
}
