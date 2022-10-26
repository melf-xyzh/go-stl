/**
 * @Time    :2022/6/23 14:19
 * @Author  :Xiaoyu.Zhang
 */

package queue

// node 队列节点
type node struct {
	data interface{}
	next *node
}
