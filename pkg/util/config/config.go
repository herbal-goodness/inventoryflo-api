package config

import (
	"os"
)

var env = os.Getenv("env")

//// ResourceToTableMapping maps the known resource ids to postgres table details
//var ResourceToTableMapping = map[string]model.TableDetails{
//	"items":    {Table: "itemstemp", Id: "sku", ArrayColumns: map[string]struct{}{"certifications": {}}},
//	"contacts": {Table: "contacts", Id: "id", ArrayColumns: map[string]struct{}{}},
//}

var constants = map[string]string{
	//"dbUser": "root",
	//"dbPort": "5432",
	//"dbName": "inventoryflo",
	"shopifyUrl": "herbal-papaya.myshopify.com/admin/api/2020-01/",
}

var envVars = map[string]map[string]string{
	"qa": {
		//"dbHost": "if-dev-use2-db-instance-1.ctshbytidoqu.us-east-2.rds.amazonaws.com",
	},
	"prod": {
		//"dbHost": "todo",
	},
}

var secrets = map[string]map[string]string{
	"qa": {
		//"dbPass": "60270ce536ba98d849b56b57dbe71025e9d74b77998ed97cb02bad1275bda705e28361cf9306e295",
		"shopifyKey": "f3e5803fd16361b2a3b0adff0c37181c7655ddc8abb762ec16d84e465a65e556c4c077512e1292501e3e53d69da2bc941c06bf6958e353abf208c3f4",
		"shopifyPass": "092f54af6ae24adc90e686bde858608e918daf063bd192c68414562301205e26d054cb3733835688a4db837f486572398d18b12580da6f75e0a58b19",
	},
	"prod": {
		//"dbPass": "dcfefea16f7368ebf1505701e819ab8bd87fa76ac853c48e1644309a1f6bd84502a6fcc963f08f78344e51a4",
		"shopifyKey": "197833dfd3b1fe0437cd6ae0256c1125f5b9066cbbd02b7a16794a94401dbb04423d2861c3e0143ea0e70a55417bb1eca7392246f4bf60a021b18e70",
		"shopifyPass": "859fc854b79623940c38631ea4fbc7022d917c5c574b75ddea4584ba29c45852caa57087c44c8414f0599e18d6e657cf48ec8341d6f6c2eb9e5c975b",
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
