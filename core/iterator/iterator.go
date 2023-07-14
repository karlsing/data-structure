package iterator

type Iterable[T any] interface {
	HasNext() bool
	Next() T
}

// ForeachB iterates an iterable object with given function "f" of each element and its index and stop when "f" returns ture
func ForeachB[T any](iter Iterable[T], f func(index int, elem T) bool) {
	for i := 0; iter.HasNext(); i++ {
		if f(i, iter.Next()) {
			break
		}
	}
}

// ForeachI iterates an iterable object with given function of each element and its index
func ForeachI[T any](iter Iterable[T], f func(index int, elem T)) {
	for i := 0; iter.HasNext(); i++ {
		f(i, iter.Next())
	}
}

// Foreach iterates an iterable object with given function of each element
func Foreach[T any](iter Iterable[T], f func(elem T)) {
	for iter.HasNext() {
		f(iter.Next())
	}
}
