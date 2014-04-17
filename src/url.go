package tradeking
import (
    "fmt"
)

const (
    PROTO                           = "https://"
    DOMAIN_DEVELOPERS               = "developers.tradeking.com"
    DOMAIN_STREAM                   = "stream.tradeking.com"
    DOMAIN_API                      = "api.tradeking.com"
    VERSION_1                       = "v1"
    VERSION_CURRENT                 = VERSION_1
    FORMAT_JSON                     = "json"
    FORMAT_XML                      = "xml"
    GET                        = "GET"
    POST                       = "POST"
    MARKET                          = "market"
    UTILITY                         = "utility"
    URL_OAUTH_REQUEST               = PROTO + DOMAIN_DEVELOPERS + "/oauth/request_token"
    URL_OAUTH_AUTHORIZE             = PROTO + DOMAIN_DEVELOPERS + "/oauth/authorize"
    URL_OAUTH_ACCESS                = PROTO + DOMAIN_DEVELOPERS + "/oauth/access_token"
    URL_ACCOUNTS                    = "accounts"
    URL_ACCOUNTS_BALANCES           = URL_ACCOUNTS + "/balances"
    URL_ACCOUNTS_ID_BALANCES        = URL_ACCOUNTS + "/%s/balances"
    URL_ACCOUNTS_ID_HISTORY         = URL_ACCOUNTS + "/%s/history"
    URL_ACCOUNTS_ID_HOLDINGS        = URL_ACCOUNTS + "/%s/holdings"
    URL_ACCOUNTS_ID_ORDERS          = URL_ACCOUNTS + "/%s/orders"
    URL_ACCOUNTS_ID_ORDERS_PREVIEW  = URL_ACCOUNTS + "/%s/orders/preview"
    URL_STREAM_MARKET_QUOTES        = MARKET + "/quotes"
    URL_MARKET_CLOCK                = MARKET + "/clock"
    URL_MARKET_EXT_QUOTES           = MARKET + "/ext/quotes"
    URL_MARKET_NEWS_SEARCH          = MARKET + "/news/search"
    URL_MARKET_NEWS_ID              = MARKET + "/news/%s"
    URL_MARKET_OPTIONS_SEARCH       = MARKET + "/options/search"
    URL_MARKET_OPTIONS_STRIKE       = MARKET + "/options/strikes"
    URL_MARKET_OPTIONS_EXPIRATIONS  = MARKET + "/options/expirations"
    URL_MARKET_TIMESALES            = MARKET + "/timesales"
    URL_MARKET_TOPLISTS             = MARKET + "/toplists"
    URL_MEMBER_PROFILE              = "member/profile"
    URL_UTILITY_STATUS              = UTILITY + "/status"
    URL_UTILITY_VERSION             = UTILITY + "/version"
    URL_WATCHLISTS                  = "watchlists"
    URL_WATCHLISTS_ID               = URL_WATCHLISTS + "/%s"
    URL_WATCHLISTS_ID_SYMBOLS       = URL_WATCHLISTS_ID + "/symbols"
)

func GetStreamForSymbols(oauthC *OAuthConnection, symbols string) (channelBuffer *StreamChannel) {
    url := buildEndPoint(DOMAIN_STREAM, VERSION_CURRENT, FORMAT_JSON, URL_STREAM_MARKET_QUOTES)
    url = url + "?symbols=" + symbols
    res, _ := oauthC.MakeHttpRequest(GET, url, "")
    channelBuffer = oauthC.GetStreamChannelFromReader(res.Body)
    return
}

func buildEndPointWithId(domain, version, format, path, id string) (endPoint string) {
    endPointFormat := buildEndPoint(domain, version, format, path)
    endPoint = fmt.Sprintf(endPointFormat, id)
    return
}

func buildEndPoint(domain, version, format, path string) (endPoint string) {
    endPoint = PROTO + domain + "/" + version + "/" + path + "." + format
    return
}
