package queue

type Queue struct {
  queue []any
}

type QueueActions interface {
  Size() int
  IsEmpty() bool
  Push()
  Pop() any
  Seek() any
}

func (q *Queue) Size() int {
  return len(q.queue)
}

func (q *Queue) IsEmpty() bool {
  return q.Size() == 0
}
func (q *Queue) Push(value any) {
  q.queue = append(q.queue, value)
}
func (q *Queue) Pop() any {
  if q.IsEmpty() {
    return nil
  }

  var result = q.queue[0]
  q.queue = q.queue[1:]
  
  return result
}
func (q *Queue) Seek() any {
  if q.IsEmpty() {
    return nil
  }

  return q.queue[0]
}