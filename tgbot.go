package tgbot

import (
  "fmt"
  "time"
  "bytes"
  "errors"
  "strconv"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "mime/multipart"
)

var apiUrl = ""
var poolDelay = 300
var lastUpdateId int = -1
var updatesCallback = func([]Update) { }
var httpClient = &http.Client{}

func SetPoolDelay(delay int) {
  poolDelay = delay
}

func SetAPIKey(key string) {
  apiUrl = "https://api.telegram.org/bot" + key + "/"
}

func SetUpdatesCallback(callback func([]Update)) {
  updatesCallback = callback
}

func Poll() {
  for true {
    var updates = getUpdates(lastUpdateId + 1)

    if len(updates) != 0 {
      var lastUpdate = updates[len(updates) - 1]
      updatesCallback(updates)
      lastUpdateId = lastUpdate.UpdateId
    }

    time.Sleep(time.Duration(poolDelay) * time.Millisecond)
  }
}

func getUpdates(offset int) []Update {
  var updates []Update
  sendResuest("getUpdates", map[string]interface{}{
    "offset": offset,
  }, &updates)

  return updates
}

func logRequestError(err error, method string, parameters map[string]interface{}) {
  fmt.Println("Error!")
  // TODO: Log request error
}

func sendResuest(method string, parameters map[string]interface{}, t interface{}) error {
  err := _sendResuest(method, parameters, &t)
  if err != nil {
    logRequestError(err, method, parameters)
  }

  return err
}

func _sendResuest(method string, parameters map[string]interface{}, t interface{}) error {
  var requestBytes bytes.Buffer
  var writer = multipart.NewWriter(&requestBytes)

  for key, value := range parameters {
    if v, ok := value.(int); ok {
      writer, _ := writer.CreateFormField(key)
      writer.Write([]byte(strconv.FormatInt(int64(v), 10)))
      continue
    }

    if v, ok := value.(int64); ok {
      writer, _ := writer.CreateFormField(key)
      writer.Write([]byte(strconv.FormatInt(v, 10)))
      continue
    }

    if v, ok := value.(string); ok {
      writer, _ := writer.CreateFormField(key)
      writer.Write([]byte(v))
      continue
    }
  }

  writer.Close()

  request, err := http.NewRequest("POST", apiUrl + method, &requestBytes)
  request.Header.Add("Content-Type", writer.FormDataContentType())
  if err != nil {
    return errors.New("Unable to create request with given parameters")
  }

  response, err := httpClient.Do(request)
  if err != nil {
    return errors.New("Unable to execute request with given parameters")
  }

  defer response.Body.Close()

  if response.StatusCode != http.StatusOK {
    return errors.New("Response status code is not OK")
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return errors.New("Unable to read response body")
  }

  var apiResponse = APIResponse{}
  err = json.Unmarshal(body, &apiResponse)
  if err != nil {
    return errors.New("Unable unserialize response from server")
  }

  if apiResponse.Ok != true {
    return errors.New("API request error")
  }

  err = json.Unmarshal(apiResponse.Result, &t)
  if err != nil {
    return errors.New("Unable unserialize API result")
  }

  return nil
}
