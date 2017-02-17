package utils

import (
	log "github.com/cihub/seelog"
)

func GetStringValue(data map[string]interface{}, field string) string {
	value, ok := data[field]

	if !ok {
		log.Debugf("Unable to read %s value in %+v", field, data)
		return ""
	}

	casted, ok := value.(string)

	if !ok {
		log.Debugf("Unable to cast %s value in %+v", field, value)
	}

	return casted
}

func GetMap(data interface{}) map[string]interface{} {
	content, ok := data.(map[string]interface{})

	if !ok {
		log.Debugf("Unable to cast to map %+v", data)
	}

	return content
}

func GetInterfaceField(data map[string]interface{}, field string) interface{} {
	content, ok := data[field]

	if !ok {
		log.Debugf("Unable to read %s from data %+v", field, data)
	}

	return content
}

func GetArrayInterfaceField(data map[string]interface{}, field string) []interface{} {
	var empty []interface{}
	var check interface{}

	content := GetInterfaceField(data, field)
	if content == check {
		return empty
	}

	arr, ok := content.([]interface{})

	if !ok {
		log.Debugf("Unable to cast []interface for field %s in %+v", field, data)
	} else {
		return arr
	}

	return empty
}
