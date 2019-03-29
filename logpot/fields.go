package logpot

import (
	"encoding/json"
	"fmt"
)

type Fields map[string]interface{}

func segregateFields(fields Fields) (Fields, Fields) {
	if logOptions.PrintFieldsInsideMessage {
		return sepFields(fields, logOptions.PrintAsFields)
	}
	return fields, Fields{}
}

func (fields Fields) toString() string {
	if len(fields) == 0 {
		return ""
	}

	//b, err := json.MarshalIndent(fields, "", "  ") // MarshalIndent should be better, if not, use regular Marshal
	b, err := json.Marshal(fields)
	if err != nil {
		return fmt.Sprintf("%v (failed marshal log! error: %v)", fields, err)
	}
	return string(b)
}

func sepFields(allFields Fields, separatedFields []string) (Fields, Fields) {
	all := allFields.clone()
	asFields := Fields{}

	for _, field := range separatedFields {
		if val, exist := all[field]; exist {
			asFields[field] = val
			delete(all, field)
		}
	}

	return asFields, all
}

func (fields Fields) clone() Fields {
	cloned := Fields{}
	for k, v := range fields {
		cloned[k] = v
	}
	return cloned
}
