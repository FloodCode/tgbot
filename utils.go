package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
)

var httpClient = &http.Client{}

func sendResuest(method string, apiKey string, paramsObject interface{}, t interface{}) error {
	err := _sendResuest(method, apiKey, paramsObject, t)
	if err != nil {
		logRequestError(err, method, paramsObject)
	}

	return err
}

func _sendResuest(method string, apiKey string, paramsObject interface{}, t interface{}) error {
	parameters, err := extractParams(paramsObject)

	var requestBytes bytes.Buffer
	var writer = multipart.NewWriter(&requestBytes)
	var withParameters = false

	var addFileParameter = func(key string, value []byte, filename string) {
		writer, _ := writer.CreateFormFile(key, filename)
		writer.Write(value)
		withParameters = true
	}

	var addStringParameter = func(key string, value string) {
		writer, _ := writer.CreateFormField(key)
		writer.Write([]byte(value))
		withParameters = true
	}

	for key, value := range parameters {
		if param, ok := value.(string); ok {
			addStringParameter(key, param)
		} else if param, ok := value.(*InputFile); ok {
			fileData := param.Get()
			if stringData, ok := fileData.(string); ok {
				addStringParameter(key, stringData)
			} else if bytesData, ok := fileData.([]byte); ok {
				addFileParameter(key, bytesData, param.GetFilename())
			}
		} else if param, ok := value.([]byte); ok {
			addFileParameter(key, param, "file")
		}
	}

	writer.Close()

	request, err := http.NewRequest("POST", "https://api.telegram.org/bot"+apiKey+"/"+method, &requestBytes)
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Unable to read response body")
	}

	var apiResponse = apiResponse{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return errors.New("Unable to unserialize response from server")
	}

	if apiResponse.Ok != true {
		return errors.New("API request error. Description: " + apiResponse.Description)
	}

	err = json.Unmarshal(apiResponse.Result, &t)
	if err != nil {
		return errors.New("Unable to unserialize API result")
	}

	return nil
}

func extractParams(paramsObject interface{}) (map[string]interface{}, error) {
	var result = map[string]interface{}{}

	if paramsObject == nil {
		return result, nil
	}

	reflectType := reflect.TypeOf(paramsObject)
	reflectValue := reflect.ValueOf(paramsObject)

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		option := field.Tag.Get("option")

		var extractedValue interface{}
		switch v := reflectValue.FieldByName(field.Name).Interface().(type) {
		case bool:
			if v {
				extractedValue = "true"
			}
		case int:
			if v != 0 {
				extractedValue = strconv.FormatInt(int64(v), 10)
			}
		case int64:
			if v != 0 {
				extractedValue = strconv.FormatInt(v, 10)
			}
		case float64:
			if v != 0 {
				extractedValue = strconv.FormatFloat(v, 'f', 6, 64)
			}
		case string:
			if len(v) > 0 {
				extractedValue = v
			}
		case *ChatIdentifier:
			if v != nil {
				extractedValue = v.Get()
			}
		case *ParseMode:
			if v != nil {
				extractedValue = v.Get()
			}
		case *ChatAction:
			if v != nil {
				extractedValue = v.Get()
			}
		case *ReplyMarkup:
			if v != nil {
				extractedValue = v.Get()
			}
		case *InputFile:
			if v != nil {
				extractedValue = v
			}
		}

		if extractedValue != nil {
			result[option] = extractedValue
		} else if len(field.Tag.Get("required")) != 0 {
			return result, errors.New("Missing required option (" + option + ")")
		}
	}

	return result, nil
}

func logRequestError(err error, method string, parameters interface{}) {
	fmt.Println("Request error:", err)
}
