package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func HttpResponse(data interface{}, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if status >= 400 {
		data = struct {
			Error interface{} `json:"error"`
		}{
			Error: data,
		}
	}

	json.NewEncoder(w).Encode(&data)
}

func HttpRequest(input map[string]interface{}) (map[string]interface{}, error) {
	output := map[string]interface{}{}
	nestedMaps := make(map[string]map[string]interface{})

	pattern := regexp.MustCompile(`(\w+)\[(\d+)\]\[(\w+)\]`)

	for key, value := range input {
		matches := pattern.FindStringSubmatch(key)
		if len(matches) > 0 {
			parentKey := matches[1]
			index, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil, err
			}
			childKey := matches[3]

			if _, exists := nestedMaps[parentKey]; !exists {
				nestedMaps[parentKey] = map[string]interface{}{}
			}

			if _, exists := nestedMaps[parentKey][fmt.Sprint(index)]; !exists {
				nestedMaps[parentKey][fmt.Sprint(index)] = map[string]interface{}{}
			}

			nestedMaps[parentKey][fmt.Sprint(index)].(map[string]interface{})[childKey] = value
		} else {
			output[key] = value
		}
	}

	for parentKey, nestedMap := range nestedMaps {
		output[parentKey] = []map[string]interface{}{}

		for i := 0; i < len(nestedMap); i++ {
			if value, exists := nestedMap[fmt.Sprint(i)]; exists {
				output[parentKey] = append(output[parentKey].([]map[string]interface{}), value.(map[string]interface{}))
			}
		}
	}

	return output, nil
}
