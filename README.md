## Convenience class for Yandex Translate API
[**Yandex Translate**](http://api.yandex.com/translate/) is an online API for machine translation service. 
It offers text translation features for over 30 languages.

:warning: **Warning:** The asset is currently under development and may be unsuitable for production use.

### Installation
  - Use `go get github.com/yaruson/go-yandex-translate` to download the library
  - Include `github.com/yaruson/go-yandex-translate` into your source

### Usage

At the moment this library contains a single method `yandex.Translate()` to perform translation.

> **NOTE:**
> You have to get an [API Key](http://api.yandex.com/key/form.xml?service=trnsl) for Yandex Translate service, if you don't have one yet.

#### Basic Example

```go
package main
import (
  "log"
  "github.com/yarson/go-yandex-translate"
)

func main() {
  apiKey := "..."
  text   := "Hello, this is a sample text!"
  lang   := "en-it"
  
  if result, e := yandex.Translate(apiKey, text, lang); e != nil {
    log.Fatal(e)
  } else {
    log.Printf(`Yandex.Translate (%s)
        Source: %s
        Result: %s`, lang, text, result)
  }
}
```
