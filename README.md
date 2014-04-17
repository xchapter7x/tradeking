tradeking api lib
======

this is a lib to use the
tradeking rest api in the go language


NOTE:
THIS LIB IS NOT FULLY FUNCTIONAL


[![wercker status](https://app.wercker.com/status/fad991cb7a12f8e507f62942d95a47bc/m/ "wercker status")](https://app.wercker.com/project/bykey/fad991cb7a12f8e507f62942d95a47bc)


Sample Code to Stream:
```javascript
package main
import (
	"fmt"
	tk "./tradeking/src"
)

func main() {
  consumerKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  consumerSecret := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  accessToken := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  accessSecret := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

  oauthC := tk.NewOAuthConnection(consumerKey,
                                consumerSecret,
                                accessToken,
                                accessSecret)

  streamChannel := tk.GetStreamForSymbols(oauthC, "BLK")
  for s := range streamChannel.Channel {
    fmt.Println("Packet recieved: ", s)
  }
}
```
