package tradeking

import (
	"testing"
)

func Test_buildEndPoint(t *testing.T) {
    controlEndPoint := "https://stream.tradeking.com/v1/market/quotes.json"
    testEndPoint := buildEndPoint(DOMAIN_STREAM, VERSION_CURRENT, FORMAT_JSON, URL_STREAM_MARKET_QUOTES)
    if controlEndPoint != testEndPoint {
        t.Errorf("%s != %s; values should be the same", controlEndPoint, testEndPoint)
    }
}

func Test_buildEndPointWithId(t *testing.T) {
    controlEndPoint := "https://api.tradeking.com/v1/accounts/12345/balances.json"
    testEndPoint := buildEndPointWithId(DOMAIN_API, VERSION_CURRENT, FORMAT_JSON, URL_ACCOUNTS_ID_BALANCES, "12345")
    if controlEndPoint != testEndPoint {
        t.Errorf("%s != %s; values should be the same", controlEndPoint, testEndPoint)
    }
}
