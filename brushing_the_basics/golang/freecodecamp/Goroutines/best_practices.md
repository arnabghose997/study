# Best Practices

- Don't create goroutines in libraries
  - Let consumer control concurrency

- When creating a goroutine, know when it will end
  - Avoid subtle memory leaks

- Check for race condition in complie time
  - Run this to check if your app has Race condition: `go run -race src/main.go`
