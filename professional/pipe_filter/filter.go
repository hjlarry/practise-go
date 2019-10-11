// Package pipefilter is to define the interfaces and the structures for pipe-filter style implementation
package pipe_filter

type Request interface {
}

type Response interface {
}

type Filter interface {
	Process(data Request) (Response, error)
}
