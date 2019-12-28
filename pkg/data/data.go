package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/herbal-goodness/inventoryflo-api/pkg/util/config"
	"github.com/herbal-goodness/inventoryflo-api/pkg/util/db"
	"github.com/lib/pq"
)

// GetAll gets all records of a resource
func GetAll(resource string) (map[string]interface{}, error) {
	tableDetails, exists := config.ResourceToTableMapping[resource]
	if !exists {
		return nil, nil
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableDetails.Table))
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{resource: transformRows(rows, tableDetails.ArrayColumns)}, nil
}

// GetResource gets a record of the specified resource and id
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

	return transformRows(rows, tableDetails.ArrayColumns)[0], nil
}

// AddResource adds the specified record the specified resource table
func AddResource(resource string, record map[string]interface{}) (map[string]interface{}, error) {
	tableDetails, exists := config.ResourceToTableMapping[resource]
	if !exists {
		return nil, nil
	}

	i, p, v := deconstruct(record, tableDetails.ArrayColumns)
	if v == nil {
		return nil, fmt.Errorf("invalid identifier: %s", i)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING %s", tableDetails.Table, i, p, tableDetails.Id)
	rows, err := db.Query(query, v...)
	if err != nil {
		return nil, err
	}
	record[tableDetails.Id] = rows[0][tableDetails.Id]
	return record, nil
}

func deconstruct(record map[string]interface{}, arrayCols map[string]struct{}) (string, string, []interface{}) {
	var identifiers []string
	var placeholders []string
	var values []interface{}
	i := 1

	for k, v := range record {
		if !db.IsValidIdentifier(k) {
			return k, "", nil
		}
		if v == nil {
			continue
		}
		if _, ok := arrayCols[k]; ok {
			v = pq.Array(v)
		}
		identifiers = append(identifiers, k)
		values = append(values, v)
		placeholders = append(placeholders, "$"+strconv.Itoa(i))
		i++
	}

	return strings.Join(identifiers, ", "), strings.Join(placeholders, ", "), values
}

func transformRows(rows []map[string]interface{}, arrayCols map[string]struct{}) []map[string]interface{} {
	transformed := rows[:0]
	for _, row := range rows {
		for col := range arrayCols {
			if val, ok := row[col]; ok {
				str := string(val.([]byte))
				row[col] = strings.Split(str[1:len(str)-1], ",")
			}
		}
		transformed = append(transformed, row)
	}
	return transformed
}
