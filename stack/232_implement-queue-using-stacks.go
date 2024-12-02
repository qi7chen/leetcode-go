package leetcode_stack

// 仅使用两个栈实现先入先出队列
// https://leetcode.cn/problems/implement-queue-using-stacks/

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Empty() bool {
	return len(q.stack1) == 0
}

// 将元素 x 推到队列的末尾
func (q *MyQueue) Push(x int) {
	q.stack1 = append(q.stack1, x)
}

func (q *MyQueue) peekPop(pop bool) int {
	// 将stack1的元素全部转移到stack2
	for i := len(q.stack1) - 1; i >= 0; i-- {
		q.stack2 = append(q.stack2, q.stack1[i])
	}
	q.stack1 = q.stack1[:0]

	// 取出stack2的栈顶元素
	var top = q.stack2[len(q.stack2)-1]
	if pop {
		q.stack2 = q.stack2[:len(q.stack2)-1]
	}

	// 将stack2的元素全部转移到stack1
	for i := len(q.stack2) - 1; i >= 0; i-- {
		q.stack1 = append(q.stack1, q.stack2[i])
	}
	q.stack2 = q.stack2[:0]
	return top
}

// 从队列的开头移除并返回元素
func (q *MyQueue) Pop() int {
	return q.peekPop(true)
}

// 返回队列开头的元素
func (q *MyQueue) Peek() int {
	return q.peekPop(false)
}
