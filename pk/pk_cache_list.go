package pk

import (
	"container/list"
	"sync"
)

type ConcurrentQueue struct {
	l     *list.List
	lock  sync.RWMutex
	count int
}

func NewConcurrentQueue() *ConcurrentQueue {
	return &ConcurrentQueue{
		l:     list.New(),
		count: 0,
	}
}

func (q *ConcurrentQueue) Enqueue(value interface{}) {
	q.lock.Lock() // 使用写锁
	defer q.lock.Unlock()
	q.l.PushBack(value)
	q.count = q.count + 1
}

func (q *ConcurrentQueue) Dequeue() interface{} {
	q.lock.Lock() // 使用写锁
	defer q.lock.Unlock()
	front := q.l.Front()
	if front != nil {
		q.l.Remove(front)
		q.count = q.count - 1
		return front.Value
	}
	return nil
}

func (q *ConcurrentQueue) Front() interface{} {
	q.lock.RLock() // 使用读锁
	defer q.lock.RUnlock()
	front := q.l.Front()
	if front != nil {
		return front.Value
	}
	return nil
}

// 计算Len时可以忍受误差，所以不再加锁
func (q *ConcurrentQueue) Len() int {
	return q.count
}
