package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var httpClient = &http.Client{}

func sendResuest(method string, apiKey string, config interface{}, t interface{}) error {
	var reader io.Reader

	buffer, contentType := getRequestBuffer(extractParams(config))
	if buffer != nil {
		reader = buffer
	}

	request, err := http.NewRequest("POST", "https://api.telegram.org/bot"+apiKey+"/"+method, reader)
	if err != nil {
		return errors.New("unable to create request with given parameters")
	}

	if reader != nil {
		request.Header.Add("Content-Type", contentType)
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return errors.New("unable to execute request with given parameters")
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("unable to read response body")
	}

	var apiResponse = apiResponse{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return errors.New("unable to unserialize response from server")
	}

	if apiResponse.Ok != true {
		return errors.New(apiResponse.Description)
	}

	err = json.Unmarshal(apiResponse.Result, &t)
	if err != nil {
		return errors.New("unable to unserialize API result")
	}

	return nil
}

func getRequestBuffer(params map[string]interface{}) (buffer *bytes.Buffer, contentType string) {
	var requestBuffer bytes.Buffer
	var hasParameters = false

	writer := multipart.NewWriter(&requestBuffer)
	defer writer.Close()

	var addFileParameter = func(key string, value []byte, filename string) {
		writer, _ := writer.CreateFormFile(key, filename)
		writer.Write(value)
		hasParameters = true
	}

	var addStringParameter = func(key string, value string) {
		writer, _ := writer.CreateFormField(key)
		writer.Write([]byte(value))
		hasParameters = true
	}

	for key, value := range params {
		if param, ok := value.(string); ok {
			addStringParameter(key, param)
		} else if param, ok := value.(*InputFile); ok {
			fileData := param.getData()
			if stringData, ok := fileData.(string); ok {
				addStringParameter(key, stringData)
			} else if bytesData, ok := fileData.([]byte); ok {
				addFileParameter(key, bytesData, param.getFilename())
			}
		} else if param, ok := value.([]byte); ok {
			addFileParameter(key, param, "file")
		}
	}

	if !hasParameters {
		return nil, ""
	}

	return &requestBuffer, writer.FormDataContentType()
}

func extractParams(config interface{}) map[string]interface{} {
	var result = map[string]interface{}{}

	if config == nil {
		return result
	}

	configType := reflect.TypeOf(config)
	configValue := reflect.ValueOf(config)

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		configField := configValue.FieldByName(field.Name)
		zeroValue := reflect.Zero(configField.Type())

		if reflect.DeepEqual(configField.Interface(), zeroValue.Interface()) {
			continue
		}

		var extractedValue interface{}
		switch v := configField.Interface().(type) {
		case bool:
			extractedValue = "true"
		case int:
			extractedValue = strconv.FormatInt(int64(v), 10)
		case int64:
			extractedValue = strconv.FormatInt(v, 10)
		case float64:
			extractedValue = strconv.FormatFloat(v, 'f', 6, 64)
		case string, *InputFile:
			extractedValue = v
		case stringConfig:
			extractedValue = v.getString()
		case []InlineQueryResult:
			serializedQueryResults := []string{}
			for _, queryResult := range v {
				queryResultJSON := buildInlineQueryResult(queryResult, queryResult.getType())
				serializedQueryResults = append(serializedQueryResults, queryResultJSON)
			}
			extractedValue = "[" + strings.Join(serializedQueryResults, ",") + "]"
		default:
			continue
		}

		result[field.Tag.Get("option")] = extractedValue
	}

	return result
}

func buildInlineQueryResult(queryResult InlineQueryResult, resultType string) string {
	queryResultMap := map[string]interface{}{"type": resultType}
	reflectType := reflect.TypeOf(queryResult)
	reflectValue := reflect.ValueOf(queryResult)

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		if jsonAttribute, ok := field.Tag.Lookup("json"); ok {
			val := reflectValue.FieldByName(field.Name).Interface()
			if val != reflect.Zero(reflect.TypeOf(val)).Interface() {
				queryResultMap[jsonAttribute] = val
			}
		}
	}

	queryResultJSON, _ := json.Marshal(queryResultMap)
	return string(queryResultJSON)
}
