# serpprovider-go
This package to ensure uses [SerpProvider](http://serpprovider.com/) service easily.

### Installation

```bash
$ go get github.com/riza/serpprovider-go
```

### Usage

```go
package main

import (
    "github.com/riza/serpprovider-go"
 "fmt")

func main() {

   res := serpprovider.New("API_KEY_HERE").Request(&serpprovider.SearchRequest{
      Query:"webtekno",
	  Country:"tr",
	  Language:"tr",
	  Limit:135,
  });

  fmt.Print(res)
}
```


## Changelog

**Version 1.0**
- Initial
-
## Support
- <img src="https://images.coinviewer.io/currencies/16x16/bitcoin.png" />  **BTC** : 3FSL583ncizWvMrEBZSN2o6K2uudyDZyby *(Min 0.0005 BTC)*
- <img src="https://images.coinviewer.io/currencies/16x16/ethereum.png" /> **ETH** : 0x8e12150a778707729cfd7971e308a9293888549d *(Min 0.001 ETH)*
- <img src="https://images.coinviewer.io/currencies/16x16/ripple.png" /> **XRP** : rNEygqkMv4Vnj8M2eWnYT1TDnV1Sc1X5SN   - **TAG** :  1543091 *(Min 1 XRP)*