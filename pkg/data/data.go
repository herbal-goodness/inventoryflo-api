package data

import (
	"fmt"
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/config"
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/db"
)

func GetAll(resource string) (map[string]interface{}, error) {
	tableDetails, exists := config.ResourceToTableMapping[resource]
	if !exists {
		return nil, nil
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableDetails.Table))
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

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s=$1", tableDetails.Table, tableDetails.Id), resourceId)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	if len(rows) != 1 {
		return nil, fmt.Errorf("multiple records were found with id: %s", resourceId)
	}

	return rows[0], nil
}
