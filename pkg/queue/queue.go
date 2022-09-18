package queue

import "errors"

type IQueue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, error)
	Peek() (T, error)
}

type Queue[T any] struct {
	collection map[int]T
	next       int
	last       int
}

func New[T any]() Queue[T] {
	queue := Queue[T]{}
	queue.collection = make(map[int]T)
	queue.next = 0
	queue.last = 0
	return queue
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.collection[queue.last] = value
	queue.last++
}

func (queue *Queue[T]) Dequeue() (T, error) {
	if len(queue.collection) == 0 {
		out := new(T)
		return *out, errors.New("queue is empty")
	}
	ref := queue.collection[queue.next]
	delete(queue.collection, queue.next)
	queue.next++
	return ref, nil
}

func (queue *Queue[T]) Peek() (T, error) {
	if len(queue.collection) == 0 {
		out := new(T)
		return *out, errors.New("queue is empty")
	}
	ref := queue.collection[queue.next]
	return ref, nil
}
