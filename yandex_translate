package yandex

import (
  "fmt"
  "strings"
  "net/url"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func Translate(apiKey, text, direction string) (string, error) {

  type Response struct {
    Code int
    Lang string
    Text[] string
  }

  url := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&text=%s&lang=%s&format=html", apiKey, url.QueryEscape(text), direction)
  if httpResp, e := http.Get(url); e != nil { return "", e } else {
    defer httpResp.Body.Close()
    if jsonResp, e := ioutil.ReadAll(httpResp.Body); e != nil { return "", e } else {
      decResp := json.NewDecoder(strings.NewReader(fmt.Sprintf("%s", jsonResp)))
      var translation Response
      if e := decResp.Decode(&translation); e != nil { return "", e } else {
        return strings.Join(translation.Text, "\n"), nil
      }
    }
    return "", nil
  }

  return "", nil

}
