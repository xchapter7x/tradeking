package tradeking
import (
    "io"
	"bufio"
    "github.com/kurrik/oauth1a"
    "net/http"
)

func NewOAuthConnection(consumerKey, consumerSecret, accessToken, accessSecret string) (oauthConnection *OAuthConnection){
    service := OAuthConnection{}.getService(consumerKey, consumerSecret)
    userConfig := OAuthConnection{}.getOauthConfig(accessToken, accessSecret)
    oauthConnection = &OAuthConnection{service: service,
                                        userConfig: userConfig}
    return
}

type OAuthConnection struct {
    service *oauth1a.Service
    userConfig *oauth1a.UserConfig
}

func (s OAuthConnection) getService(consumerKey, consumerSecret string) (service *oauth1a.Service) {
    service = &oauth1a.Service{
        RequestURL:   URL_OAUTH_REQUEST,
        AuthorizeURL: URL_OAUTH_AUTHORIZE,
        AccessURL:    URL_OAUTH_ACCESS,
        ClientConfig: &oauth1a.ClientConfig{
            ConsumerKey:    consumerKey,
            ConsumerSecret: consumerSecret,
            CallbackURL:    "",
        },
        Signer: new(oauth1a.HmacSha1Signer),
    }
    return
}

func (s OAuthConnection) getOauthConfig(accessToken, accessSecret string) (userConfig *oauth1a.UserConfig) {
    userConfig = oauth1a.NewAuthorizedConfig(accessToken, accessSecret)
    return
}

func (s *OAuthConnection) MakeHttpRequest(verb, url, args string) (httpResponse *http.Response, err error) {
    endPoint := url+args
    httpRequest, _ := http.NewRequest(verb, endPoint, nil)
    s.service.Sign(httpRequest, s.userConfig)
    httpResponse, err = http.DefaultClient.Do(httpRequest)
    return
}

func (s *OAuthConnection) GetStreamChannelFromReader(buf io.ReadCloser) (stream *StreamChannel) {
    stream = NewStreamChannel()
    go streamToChannel(buf, stream)
    return
}

func streamToChannel(buf io.ReadCloser, stream *StreamChannel) {
    var line string
    defer buf.Close()
    reader := bufio.NewReader(buf)
    line += readIntoLine(reader)

    for (*stream).IsAlive() {

        if bufferBreakPoint(reader) {
            (*stream).Write(line)
            line = ""
        }
        line += readIntoLine(reader)
    }
}

func readIntoLine(reader *bufio.Reader) (string) {
    var b byte
    b, _ = (*reader).ReadByte()
    return string(b)
}

func bufferBreakPoint(reader *bufio.Reader) (bool) {
    return (reader.Buffered() == 0)
}
