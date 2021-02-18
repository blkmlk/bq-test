package listener

import (
	"bq/internal/fixtures"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRecords(t *testing.T) {
	url := "https://min-api.cryptocompare.com/data/pricemultifull?tss=ETH"
	fsyms := []string{fixtures.FSymbolBTC, fixtures.FSymbolETH}
	tsyms := []string{fixtures.TSymbolUSD}

	ctx := context.Background()

	records, err := GetRecords(ctx, url, fsyms, tsyms)
	require.NoError(t, err)
	require.NotNil(t, records)
	require.Len(t, records, len(fsyms))

	for _, fRec := range records {
		require.Len(t, fRec, len(tsyms))
	}
}
