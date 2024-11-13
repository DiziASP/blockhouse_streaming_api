package utils

import (
	"blockhouse_streaming_api/pkg/merge"
)

func MergeStruct(dst interface{}, src interface{}) error {
	if err := merge.Merge(dst, src, merge.WithOverride); err != nil {
		return err
	}

	return nil
}
