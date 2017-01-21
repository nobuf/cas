# TwitCasting API v2 Client in Go

*cas* is a half-baked Go client library for [the TwitCasting API v2](http://apiv2-doc.twitcasting.tv/).

## Installing

```shell
go get github.com/nobuf/cas
```

## Example

### Getting User Information

```go
package main

import (
	"github.com/nobuf/cas"
	"log"
)

func main() {
	api := cas.New("your-client-id", "your-client-secret")
	u, err := api.User("twitcasting_jp")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", u)
}

```
See other examples in the `examples` folder. You can generate your client id and secret via [the Developer page](https://ssl.twitcasting.tv/developer.php).

## Contributing

All are welcome :)