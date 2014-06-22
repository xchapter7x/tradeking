package tradeking

import (
    "testing"
    "net/http"
    "errors"
)
var controlResponseCode = 500
type mockDoObject struct {}
func (s *mockDoObject) Do (req *http.Request) (resp *http.Response, err error) {
    err = errors.New("passedError")
    resp = new(http.Response)
    resp.StatusCode = controlResponseCode
    return
}

func getOauthC() (*OAuthConnection) {
    oauthKey := OauthKeyStorage{
        ConsumerKey: "",
        ConsumerSecret: "",
        AccessToken: "",
        AccessSecret: ""}
    mock := new(mockDoObject)
    oauthC := NewOAuthConnection(oauthKey,
                                  http.NewRequest,
                                  mock)
    return oauthC
}

func getEndPoint() (string) {
    url := buildEndPoint(DOMAIN_STREAM, VERSION_CURRENT, FORMAT_JSON, URL_STREAM_MARKET_QUOTES)
    url = url + "?symbols=FB"
    args := ""
    endPoint := url+args
    return endPoint
}

func Test_NewOAuthConnection(t *testing.T) {
    //t.Errorf("; values should be the same")
}

func Test_GetStreamChannelFromReader(t *testing.T) {
    //t.Errorf("; values should be the same")
}

func Test_GetChannelFromReader(t *testing.T) {
    //t.Errorf("; values should be the same")
}

func Test_MakeHttpRequestError(t *testing.T) {
    oauthC := getOauthC()
    endPoint := getEndPoint()
    _, err := oauthC.MakeHttpRequest(GET, endPoint)

    if err == nil {
        t.Errorf("an error should have been returned")
    }
}

func Test_MakeHttpRequestResponse(t *testing.T) {
    oauthC := getOauthC()
    endPoint := getEndPoint()
    resp, _ := oauthC.MakeHttpRequest(GET, endPoint)

    if resp.StatusCode != controlResponseCode {
        t.Errorf("response code should be 500")
    }
}
