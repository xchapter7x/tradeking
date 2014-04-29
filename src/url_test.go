package tradeking

import (
	"testing"
    "io"
    "net/http"
    "bytes"
)

type mockResponseBody struct {
    io.Reader
}

func (mockResponseBody) Close() error { return nil }

type OAuthMock struct {}

func (s *OAuthMock) MakeHttpRequest(verb, url string) (httpResponse *http.Response, err error) {
    httpResponse = &http.Response{}
    httpResponse.Body = mockResponseBody{bytes.NewBufferString("{\"status\":\"connected\"}")}
    return
}

func (s *OAuthMock) GetStreamChannelFromReader(buf io.ReadCloser) (stream *StreamChannel) {
    oauthKey := OauthKeyStorage{
        ConsumerKey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        ConsumerSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        AccessToken: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        AccessSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
    o := NewOAuthConnection(oauthKey,
                       http.NewRequest,
                       http.DefaultClient)
    stream = o.GetStreamChannelFromReader(buf)
    return
}

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

func Test_GetStreamForSymbols(t *testing.T) {
    control := "{\"status\":\"connected\"}"
    oauthC := &OAuthMock{}
    streamChannel := GetStreamForSymbols(oauthC, "FB")
    s := <- streamChannel.Channel

    if control != s {
        t.Errorf("%s != %s; values should be the same", control, s)
    }
}
