package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/herbal-goodness/inventoryflo-api/pkg/model"
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
func GetResource(resource string, resourceID string) (map[string]interface{}, error) {
	tableDetails, exists := config.ResourceToTableMapping[resource]
	if !exists {
		return nil, nil
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s=$1", tableDetails.Table, tableDetails.Id), resourceID)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	if len(rows) != 1 {
		return nil, fmt.Errorf("multiple records were found with id: %s", resourceID)
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
		return nil, fmt.Errorf("invalid identifier: %s", i[0])
	}
	idn := strings.Join(i, ", ")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING %s", tableDetails.Table, idn, p, tableDetails.Id)
	rows, err := db.Query(query, v...)
	if err != nil {
		return nil, err
	}
	record[tableDetails.Id] = rows[0][tableDetails.Id]
	return record, nil
}

// UpdateResource updates the record identified by the specified id with the record passed in
func UpdateResource(resource string, resourceID string, record map[string]interface{}) (map[string]interface{}, error) {
	tableDetails, exists := config.ResourceToTableMapping[resource]
	if !exists {
		return nil, nil
	}
	if _, ok := record[tableDetails.Id]; ok {
		delete(record, tableDetails.Id)
	}
	return update(resourceID, record, tableDetails)
}

// UpdateResources updates the passed in array of records
func UpdateResources(resource string, record map[string]interface{}) (map[string]interface{}, error) {
	tableDetails, exists := config.ResourceToTableMapping[resource]
	if !exists {
		return nil, nil
	}
	val, ok := record[resource]
	if !ok {
		return nil, fmt.Errorf("records to be updated not found")
	}
	records := val.([]interface{})

	var successes []map[string]interface{}
	var failures []string
	for _, rec := range records {
		r := rec.(map[string]interface{})
		id, ok := r[tableDetails.Id]
		if !ok {
			failures = append(failures, "ID field not found in record")
			continue
		}

		delete(record, tableDetails.Id)
		updated, err := update(id.(string), r, tableDetails)
		if err != nil {
			failures = append(failures, err.Error())
		} else {
			successes = append(successes, updated)
		}
	}

	response := map[string]interface{}{}
	var err error = nil
	if len(failures) == 0 {
		response[resource] = successes
	} else {
		if len(successes) == 0 {
			err = fmt.Errorf("none of the records passed in were successfully updated")
			response = nil
		} else {
			response[resource] = successes
			response["errors"] = failures
			err = fmt.Errorf("partial success")
		}
	}
	return response, err
}

func update(id string, record map[string]interface{}, td model.TableDetails) (map[string]interface{}, error) {
	i, _, v := deconstruct(record, td.ArrayColumns)
	if v == nil {
		return nil, fmt.Errorf("invalid identifier: %s", i[0])
	}

	var assignments []string
	for idx, idn := range i {
		assignments = append(assignments, fmt.Sprintf("%s = $%d", idn, idx+1))
	}
	assignmentString := strings.Join(assignments, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s=$%d RETURNING *", td.Table, assignmentString, td.Id, len(i)+1)
	v = append(v, id)
	rows, err := db.Query(query, v...)

	if len(rows) == 0 {
		return nil, fmt.Errorf("resource with id %s not found", id)
	}
	return transformRows(rows, td.ArrayColumns)[0], err
}

func deconstruct(record map[string]interface{}, arrayCols map[string]struct{}) ([]string, string, []interface{}) {
	var identifiers []string
	var placeholders []string
	var values []interface{}
	i := 1

	for k, v := range record {
		if !db.IsValidIdentifier(k) {
			return []string{k}, "", nil
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

	return identifiers, strings.Join(placeholders, ", "), values
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
