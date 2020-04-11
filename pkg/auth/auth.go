package auth

import (
	"fmt"
	"strings"
)

// In Progress

var actionToMethodMapping = map[string]func(map[string]interface{}) (map[string]interface{}, error){
	"register":   register,
	"deregister": deregister,
	"login":      logIn,
}

var validActions = ""

func ChangePassword(loginInfo map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func Perform(authAction map[string]interface{}) (map[string]interface{}, error) {
	action, ok := authAction["action"]
	if !ok {
		return nil, fmt.Errorf("cannot find 'action' in request body")
	}
	f, ok := actionToMethodMapping[action.(string)]
	if !ok {
		if validActions == "" {
			actions := make([]string, len(actionToMethodMapping))
			for k := range actionToMethodMapping {
				actions = append(actions, k)
			}
			validActions = strings.Join(actions, ", ")
		}
		return nil, fmt.Errorf("action %s is not recognized, must be one of [%s]", action, validActions)
	}
	user, ok := authAction["user"]
	if !ok {
		return nil, fmt.Errorf("cannot find 'user' in request body")
	}
	return f(user.(map[string]interface{}))
}

func register(user map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func deregister(user map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func logIn(user map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
