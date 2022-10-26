# go-stl
go实现常见的数据结构

## 安装

```go
go get -u github.com/melf-xyzh/go-stl
```

## 示例

### 栈

```go
st := stack.NewStack()
// 入栈
st.Push("1")
st.Push("2")
st.Push("3")
st.Push("4")
fmt.Println("打印栈:",st.AllData())
// 出栈
data1 := st.Pop()
fmt.Println("出栈：",data1)
fmt.Println("打印栈:",st.AllData())
```

### 队列

```go
q := queue.NewQueue()
q.Put("第1个")
q.Put("第2个")
q.Put("第3个")
q.Put("第4个")
q.Put("第5个")
q.Put("第6个")
fmt.Println("打印:", q.AllData())
fmt.Println("出队：", q.Get())
fmt.Println("出队：", q.Get())
fmt.Println("出队：", q.Get())
fmt.Println("出队：", q.Get())
fmt.Println("出队：", q.Get())
fmt.Println("出队：", q.Get())
fmt.Println("出队：", q.Get())
fmt.Println("出队：", q.Get())
fmt.Println("出队：", q.Get())
fmt.Println("打印:", q.AllData())
q.Put("第7个")
fmt.Println("打印:", q.AllData())
```

