package leetcode

//Below is the interface for Iterator, which is already defined for you.

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	head int
	iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		head: 0,
		iter: iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	if this.head > 0 {
		return true
	} else {
		return this.iter.hasNext()
	}
}

func (this *PeekingIterator) next() int {
	if this.head > 0 {
		next := this.head
		this.head = 0
		return next
	} else {
		return this.iter.next()
	}
}

func (this *PeekingIterator) peek() int {
	if this.head > 0 {
		return this.head
	} else {
		peek := this.iter.next()
		this.head = peek
		return peek
	}
}
