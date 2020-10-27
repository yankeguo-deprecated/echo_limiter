# echo_limiter

an Echo middleware to limit the number of concurrency requests

## Usage

```go
e := echo.New()
e.Use(echo_limiter.New(echo_limiter.FixedBucketLimiter(128)))
```

## Credits

Guo Y.K., MIT License
