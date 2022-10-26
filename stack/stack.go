/**
 * @Time    :2022/6/25 9:55
 * @Author  :Xiaoyu.Zhang
 */

package stack

import (
	"sync"
)

// stack 栈
type stack struct {
	top  *node // 栈顶
	base *node // 栈低
	size int   // 栈的高度
	sync.Mutex
}

// NewStack
/**
 *  @Description: 初始化栈
 *  @return *stack
 */
func NewStack() (s *stack) {
	s = new(stack)
	s.top = nil
	s.base = nil
	s.size = 0
	return s
}

// Push
/**
 *  @Description: 入栈
 *  @receiver s
 *  @param data
 */
func (s *stack) Push(data interface{}) {
	n := new(node)
	n.data = data
	n.next = s.top
	s.Lock()
	defer s.Unlock()
	if s.base == nil {
		s.base = n
		s.top = n
	} else {
		s.top = n
	}
	s.size++
}

// Pop
/**
 *  @Description: 出栈
 *  @receiver s
 *  @return data
 */
func (s *stack) Pop() (data interface{}) {
	if s.top == nil {
		return nil
	}
	n := s.top
	s.Lock()
	defer s.Unlock()
	// 不存在下一个元素，将栈置空
	if n.next == nil {
		s.top = nil
		s.base = nil
	} else {
		s.top = n.next
	}
	s.size--
	return n.data
}

// Size
/**
 *  @Description: 获取栈大小
 *  @receiver q
 *  @return size
 */
func (s *stack) Size() (size int) {
	return s.size
}

// Empty
/**
 *  @Description: 置空
 *  @receiver q
 */
func (s *stack) Empty() {
	s.Lock()
	defer s.Unlock()
	s.top = nil
	s.base = nil
	s.size = 0
}

// AllData
/**
 *  @Description: 获取栈所有元素
 *  @receiver q
 *  @return allData
 */
func (s *stack) AllData() (allData []interface{}) {
	s.Lock()
	defer s.Unlock()
	if s.Size() == 0 {
		return nil
	}
	allData = make([]interface{}, s.size, s.size)
	index := 0
	n := s.top
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
