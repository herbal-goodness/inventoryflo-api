package data

import (
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/config"
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/db"
)

func GetAll(resource string) (map[string]interface{}, error) {
	tableDetails, exists := config.ResourceToTableMapping[resource]
	if !exists {
		return nil, nil
	}

	var rows []map[string]interface{}
	err := db.Select(&rows, "SELECT * FROM $1", tableDetails.Table)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{resource: rows}, nil
}

func GetResource(resource string, resourceId string) (map[string]interface{}, error) {
	tableDetails, exists := config.ResourceToTableMapping[resource]
	if !exists {
		return nil, nil
	}

	var row map[string]interface{}
	err := db.Get(&row, "SELECT * FROM $1 WHERE $2=$3", tableDetails.Table, tableDetails.Id, resourceId)
	if err != nil {
		return nil, err
	}

	return row, nil
}
