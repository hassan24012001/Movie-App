package db_api

import "github.com/lib/pq"

func scanIntArray(arr *pq.Int64Array) ([]int, error) {
	if arr == nil {
		return nil, nil
	}

	result := make([]int, len(*arr))
	for i, v := range *arr {
		result[i] = int(v)
	}
	return result, nil
}
