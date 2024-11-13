package utils

import "blockhouse_streaming_api/pkg/file/json"

func BindingStruct(src interface{}, desc interface{}) error {
	byteSrc, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteSrc, &desc)
	if err != nil {
		return err
	}
	return nil
}
