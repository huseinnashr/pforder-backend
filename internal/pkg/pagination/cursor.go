package pagination

import (
	"encoding/base64"
	"encoding/json"
)

func DecodeCursor(cursor string, data interface{}) error {
	marshalledData, err := base64.RawStdEncoding.DecodeString(cursor)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(marshalledData, data); err != nil {
		return err
	}

	return nil
}

func EncodeCursor(data interface{}) string {
	marshalledData, _ := json.Marshal(data)
	return base64.RawStdEncoding.EncodeToString(marshalledData)
}
