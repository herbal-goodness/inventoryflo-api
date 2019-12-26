package router

import (
	"encoding/json"
	"fmt"
	"github.com/herbal-goodness/inventoryflo-api/pkg/data"
	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
)

// Route acts like a mux and based on the request routes to the correct function and returns it's response
func Route(method string, path []string, bodyString string) (map[string]interface{}, *model.HttpError) {
	var bodyJson map[string]interface{}
	switch method {
	case "GET":
		return get(path)
	case "POST":
		if err := json.Unmarshal([]byte(bodyString), &bodyJson); err != nil {
			return nil, errResponse(500, fmt.Sprintf("Unable to unmarshal request body to json: %v", err))
		}
		return post(path, bodyJson)
	case "PUT":
		if err := json.Unmarshal([]byte(bodyString), &bodyJson); err != nil {
			return nil, errResponse(500, fmt.Sprintf("Unable to unmarshal request body to json: %v", err))
		}
		return put(path, bodyJson)
	case "DELETE":
		return del(path)
	default:
		return nil, errResponse(404, "Unsupported method: "+method)
	}
}

func get(path []string) (map[string]interface{}, *model.HttpError) {
	var result map[string]interface{}
	var err error
	if len(path) == 1 || path[1] == "" {
		result, err = data.GetAll(path[0])
	} else {
		result, err = data.GetResource(path[0], path[1])
	}
	if err != nil {
		return nil, errResponse(500, err.Error())
	}
	if result == nil {
		return nil, errResponse(404, "Resource not found.")
	}
	return result, nil
}

func post(path []string, body map[string]interface{}) (map[string]interface{}, *model.HttpError) {
	return nil, nil
}

func put(path []string, body map[string]interface{}) (map[string]interface{}, *model.HttpError) {
	return nil, nil
}

func del(path []string) (map[string]interface{}, *model.HttpError) {
	return nil, nil
}

func errResponse(status int, message string) *model.HttpError {
	return &model.HttpError{
		Status: status,
		Error:  message,
	}
}
