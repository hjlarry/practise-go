package pipe_filter

type StraigtPipeline struct {
	Name    string
	Filters *[]Filter
}

func NewStraigtPipeline(name string, filters ...Filter) *StraigtPipeline {
	return &StraigtPipeline{
		Name:    name,
		Filters: &filters,
	}
}

func (s *StraigtPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *s.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
