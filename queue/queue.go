/**
 * @Time    :2022/6/23 14:15
 * @Author  :Xiaoyu.Zhang
 */

package queue

import "sync"

type queue struct {
	front *node // 队头
	rear  *node // 对尾
	size  int   // 队列大小
	sync.Mutex
}

// NewQueue
/**
 *  @Description:初始化队列
 *  @return *queue
 */
func NewQueue() *queue {
	q := new(queue)
	q.front = nil
	q.rear = nil
	q.size = 0
	return q
}

// Put
/**
 *  @Description: 入队
 *  @receiver q
 *  @param data
 */
func (q *queue) Put(data interface{}) {
	n := new(node)
	n.data = data
	q.Lock()
	defer q.Unlock()
	if q.rear == nil {
		q.front = n
		q.rear = n
	} else {
		q.rear.next = n
		q.rear = n
	}
	q.size++
}

// Get
/**
 *  @Description: 出队
 *  @receiver q
 *  @return data
 */
func (q *queue) Get() (data interface{}) {
	if q.front == nil {
		return nil
	}
	n := q.front
	q.Lock()
	defer q.Unlock()
	if n.next == nil {
		q.front = nil
		q.rear = nil
	} else {
		q.front = n.next
	}
	q.size--
	return n.data
}

// Size
/**
 *  @Description: 获取队列大小
 *  @receiver q
 *  @return size
 */
func (q *queue) Size() (size int) {
	return q.size
}

// Empty
/**
 *  @Description: 置空
 *  @receiver q
 */
func (q *queue) Empty() {
	q.Lock()
	defer q.Unlock()
	q.front = nil
	q.rear = nil
	q.size = 0
}

// AllData
/**
 *  @Description: 获取队列所有元素
 *  @receiver q
 *  @return allData
 */
func (q *queue) AllData() (allData []interface{}) {
	q.Lock()
	defer q.Unlock()
	if q.Size() == 0 {
		return nil
	}
	allData = make([]interface{}, q.size, q.size)
	index := 0
	n := q.front
	allData[index] = n.data
	for {
		if n.next != nil {
			n = n.next
			index++
			allData[index] = n.data
		} else {
			return
		}
	}
}
