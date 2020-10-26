package echo_limiter

import "sync/atomic"

type Limiter interface {
	Take()
	Available() int64
	Return()
}

type fixedBucketLimiter struct {
	tokens    chan struct{}
	available int64
}

func (b *fixedBucketLimiter) Take() {
	atomic.AddInt64(&b.available, -1)
	<-b.tokens
}

func (b *fixedBucketLimiter) Available() int64 {
	return b.available
}

func (b *fixedBucketLimiter) Return() {
	atomic.AddInt64(&b.available, 1)
	b.tokens <- struct{}{}
}

func NewFixedBucketLimiter(capacity int64) Limiter {
	tokens := make(chan struct{}, capacity)
	for i := int64(0); i < capacity; i++ {
		tokens <- struct{}{}
	}
	return &fixedBucketLimiter{
		tokens:    tokens,
		available: capacity,
	}
}
