package queue

type Queue[T any] struct {
	Items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.Items) == 0 {
		var z T
		return z
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}
