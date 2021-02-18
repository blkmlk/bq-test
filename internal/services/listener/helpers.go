package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func GetRecords(ctx context.Context, dataUrl string, fsyms, tsyms []string) (Records, error) {
	u, err := url.Parse(dataUrl)
	if err != nil {
		return nil, err
	}

	u.RawQuery = ""
	query := u.Query()
	query.Set("fsyms", strings.Join(fsyms, ","))
	query.Set("tsyms", strings.Join(tsyms, ","))
	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	type RawData struct {
		Raw map[string]map[string]*Record `json:"RAW"`
	}

	var data RawData
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var records = make(Records)

	for fSymbol, r1 := range data.Raw {
		fRec, ok := records[fSymbol]
		if !ok {
			fRec = make(map[string]*Record)
			records[fSymbol] = fRec
		}

		for tSymbol, record := range r1 {
			fRec[tSymbol] = record
		}
	}

	return records, nil
}
