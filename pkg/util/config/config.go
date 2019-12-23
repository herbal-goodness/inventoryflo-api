package config

import (
	"os"
)

var env = os.Getenv("env")

var constants = map[string]string{
	"dbUser": "root",
	"dbPort": "5432",
	"dbName": "inventoryflo",
}

var envVars = map[string]map[string]string{
	"qa": {
		"dbHost": "if-dev-use2-db-instance-1.ctshbytidoqu.us-east-2.rds.amazonaws.com",
	},
	"prod": {
		"dbHost": "todo",
	},
}

var secrets = map[string]map[string]string{
	"qa": {
		"dbPass": "000b8c2df0a6896fa62c4682a0dbef8b00000000000000000000000000000000",
	},
	"prod": {
		"dbPass": "",
	},
}

// Get returns the requested configuration variable given a key
func Get(key string) string {
	if val, ok := constants[key]; ok {
		return val
	}
	if val, ok := envVars[env][key]; ok {
		return val
	}
	if val, ok := secrets[env][key]; ok {
		secret := decrypt(val)
		envVars[env][key] = secret
		return secret
	}
	return ""
}
