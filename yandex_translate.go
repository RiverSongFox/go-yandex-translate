package yandex

import (
  "fmt"
  "errors"
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
    switch httpResp.StatusCode {
    case 200:
      if jsonResp, e := ioutil.ReadAll(httpResp.Body); e != nil { return "", e } else {
        decResp := json.NewDecoder(strings.NewReader(fmt.Sprintf("%s", jsonResp)))
        var translation Response
        if e := decResp.Decode(&translation); e != nil { return "", e } else {
          return strings.Join(translation.Text, "\n"), nil
        }
      }
    case 401: return "", errors.New("Invalid API key")
    case 402: return "", errors.New("This API key has been blocked")
    case 403: return "", errors.New("You have reached the daily limit for requests (including calls of the detect method)")
    case 404: return "", errors.New("You have reached the daily limit for the volume of translated text (including calls of the detect method)")
    case 413: return "", errors.New("The text size exceeds the maximum")
    case 422: return "", errors.New("The text could not be translated")
    case 501: return "", errors.New("The specified translation direction is not supported")
    }
    
    return "", nil
  }

  return "", nil

}
