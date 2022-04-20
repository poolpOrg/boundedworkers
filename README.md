# GO-MAXWORKERS

The `maxworkers` package wraps a sync.WaitGroup with a capacity,
allowing to parallelize a function call while ensuring that it only runs on at most 10 workers at all time,
and that all workers are done before continuing:

```go
func main() {
	wg := maxworkers.NewGroup(10)
	for i := 0; i < 1000; i++ {
		wg.Run(func() {
			time.Sleep(1 * time.Second)
		})
	}
	wg.Wait()
	
	fmt.Println("all workers are done")
}
```
