package main

import "fmt"

type (
	CommonIteratorImp interface {
		getIterator() CommonIterator
	}
	CommonIterator interface {
		hasMore() bool
		next() interface{}
	}

	Int64Slice struct {
		data []int64
	}

	// Int64SliceIterator int64slice迭代器
	Int64SliceIterator struct {
		index int64
		data  []int64
	}
)

func (s *Int64Slice) getIterator() CommonIterator {
	return &Int64SliceIterator{data: s.data}
}

func (s *Int64SliceIterator) hasMore() bool {
	return s.index < int64(len(s.data))
}

func (s *Int64SliceIterator) next() interface{} {
	next := s.data[s.index]
	s.index++
	return next
}

func main() {
	info := Int64Slice{data: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	iterator := info.getIterator()
	for iterator.hasMore() {
		fmt.Println(iterator.next().(int64))
	}
}
