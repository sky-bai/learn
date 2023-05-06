在Golang中进行开发时，当互斥锁始终试图获得一个它永远无法获得的锁时，它可能会遇到饥饿问题。
在本文中，我们将研究影响Go 1.8的饥饿问题，这个问题在Go 1.9中得到了解决。

## starvation

为了说明互斥体的饥饿情况，我将以Russ Cox提出的关于互斥改进的例子为例:

```
func main() {
	done := make(chan bool, 1) // 构造一个是否成功的chan 信号
	var mu sync.Mutex

	// goroutine 1
	go func() { // 开一个协程监听事件 go func select
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock() // 唯一执行一段代码逻辑
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()

	// goroutine 2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		mu.Unlock()
	}
	done <- true
}
```

goroutine 1 长时间获取锁并短暂释放它
goroutine 2 短暂获取锁并长时间释放它
两者的周期都是100微秒，但是由于goroutine 1不断地请求锁，我们可以预期它会更频繁地获得锁。

![https://miro.medium.com/v2/resize:fit:640/format:webp/1*B1atM-b6GPDS0_Q_TPEUBw.png](img.png)