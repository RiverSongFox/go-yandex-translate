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