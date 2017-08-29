package tgbot

import (
  "bytes"
  "errors"
  "reflect"
  "strconv"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "mime/multipart"
)

var httpClient = &http.Client{}

func sendResuest(method string, paramsObject interface{}, t interface{}) error {
  parameters, err := extractParams(paramsObject)
  if err != nil {
    logRequestError(err, method, paramsObject)
    return err
  }

  var requestBytes bytes.Buffer
  var writer = multipart.NewWriter(&requestBytes)
  var withParameters bool = false

  var addStringParameter = func(key string, value string) {
    writer, _ := writer.CreateFormField(key)
    writer.Write([]byte(value))
    withParameters = true
  }

  for key, value := range parameters {
    if v, ok := value.(string); ok {
      addStringParameter(key, v)
    }
  }

  writer.Close()

  request, err := http.NewRequest("POST", apiUrl + method, &requestBytes)
  if err != nil {
    return errors.New("Unable to create request with given parameters")
  }

  if withParameters {
    request.Header.Add("Content-Type", writer.FormDataContentType())
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
    return errors.New("Unable to unserialize response from server")
  }

  if apiResponse.Ok != true {
    return errors.New("API request error")
  }

  err = json.Unmarshal(apiResponse.Result, &t)
  if err != nil {
    return errors.New("Unable to unserialize API result")
  }

  return nil
}

func extractParams(paramsObject interface{}) (map[string]interface{}, error) {
  reflectType := reflect.TypeOf(paramsObject)
  reflectValue := reflect.ValueOf(paramsObject)

  var result = map[string]interface{}{}

  for i := 0; i < reflectType.NumField(); i++ {
    field := reflectType.Field(i)
    option := field.Tag.Get("option")
    var extractedValue string = ""

    switch v := reflectValue.FieldByName(field.Name).Interface().(type) {
    case int:             if v != 0     { extractedValue = strconv.FormatInt(int64(v), 10)}
    case string:          if len(v) > 0 { extractedValue = v }
    case *ChatIdentifier: if v != nil   { extractedValue = v.Get() }
    case *ParseMode:      if v != nil   { extractedValue = v.Get() }
    }

    if len(extractedValue) > 0 {
      result[option] = extractedValue
    } else if len(field.Tag.Get("required")) != 0 {
      return result, errors.New("Missing required option (" + option + ")")
    }
  }

  return result, nil
}

func logRequestError(err error, method string, parameters interface{}) {
  // TODO: Log request error
}
