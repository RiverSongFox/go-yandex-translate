## Convenience class for Yandex Translate API
[**Yandex Translate**](http://api.yandex.com/translate/) is an online API for machine translation service. 
It offers text translation features for over 30 languages.

**Warning:** The asset is currently under development and may be unsuitable for production use.

### Installation
  - Use `go get github.com/yaruson/go-yandex-translate` to download the library
  - Include `github.com/yaruson/go-yandex-translate` into your source

### Usage

> **NOTE:**
> You have to request an [API Key](http://api.yandex.com/key/form.xml?service=trnsl) for Yandex Translate service, if you don't have one yet.

```go
package main
import (
  "os"
  "strings"
  "log"
  "github.com/yaruson/go-yandex-translate"
)

func main() {
  apiKey := os.Args[1]
  lang   := os.Args[2]
  text   := strings.Join(os.Args[3:], " ")

  if result, e := yandex.Translate(apiKey, text, lang); e != nil {
    log.Fatal(e)
  } else {
    log.Printf(`OK
      Yandex.Translate (%s)
        Source: %s
        Result: %s`, lang, text, result)
  }
}
```
