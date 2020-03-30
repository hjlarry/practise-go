package pipe_filter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tf *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	res := []int{}
	for _, item := range parts {
		s, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}
	return res, nil
}
