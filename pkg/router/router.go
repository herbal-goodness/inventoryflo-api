package router

import (
	"encoding/json"
	"fmt"
	"github.com/herbal-goodness/inventoryflo-api/pkg/auth"
	"github.com/herbal-goodness/inventoryflo-api/pkg/data"
	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
)

// Route acts like a mux and based on the request routes to the correct function and returns it's response
func Route(method string, path []string, bodyString string) (map[string]interface{}, *model.HttpError) {
	var bodyJSON map[string]interface{}
	if len(path) == 0 || path[0] == "" {
		return nil, errResponse(400, "Empty path not supported. Must specify resource.")
	}
	switch method {
	case "GET":
		return get(path)
	case "POST":
		if err := json.Unmarshal([]byte(bodyString), &bodyJSON); err != nil {
			return nil, errResponse(500, fmt.Sprintf("Unable to unmarshal request body to json: %v", err))
		}
		return post(path, bodyJSON)
	case "PUT":
		if err := json.Unmarshal([]byte(bodyString), &bodyJSON); err != nil {
			return nil, errResponse(500, fmt.Sprintf("Unable to unmarshal request body to json: %v", err))
		}
		return put(path, bodyJSON)
	case "DELETE":
		return del(path)
	default:
		return nil, errResponse(405, "Unsupported method: "+method)
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
	var result map[string]interface{}
	var err error
	if path[0] == "auth" {
		result, err = auth.Perform(body)
	} else {
		result, err = data.AddResource(path[0], body)
	}
	if err != nil {
		return nil, errResponse(500, err.Error())
	}
	return result, nil
}

func put(path []string, body map[string]interface{}) (map[string]interface{}, *model.HttpError) {
	var result map[string]interface{}
	var err error
	if path[0] == "auth" {
		result, err = auth.ChangePassword(body)
		if err != nil {
			return nil, errResponse(500, err.Error())
		}
		return result, nil
	}
	if len(path) > 1 && path[1] != "" {
		result, err = data.UpdateResource(path[0], path[1], body)
		if err != nil {
			return nil, errResponse(500, err.Error())
		}
	} else {
		result, err = data.UpdateResources(path[0], body)
		if err != nil {
			if result != nil {
				resp, _ := json.Marshal(result)
				return nil, errResponse(206, string(resp))
			}
			return nil, errResponse(500, err.Error())
		}
	}
	if result == nil {
		return nil, errResponse(404, "Resource not found.")
	}
	return result, nil
}

func del(path []string) (map[string]interface{}, *model.HttpError) {
	if len(path) == 1 || path[1] == "" {
		return nil, errResponse(400, "Must specify resourceId.")
	}

	result, err := data.DeleteResource(path[0], path[1])
	if err != nil {
		return nil, errResponse(500, err.Error())
	}
	if result == nil {
		return nil, errResponse(204, "Resource not found.")
	}
	return result, nil
}

func errResponse(status int, message string) *model.HttpError {
	return &model.HttpError{
		Status: status,
		Error:  message,
	}
}
