package echo_limiter

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewFixedBucketLimiter(t *testing.T) {
	l := NewFixedBucketLimiter(2)
	l.Take()
	require.Equal(t, int64(1), l.Available())
	l.Take()
	require.Equal(t, int64(0), l.Available())
	l.Return()
	require.Equal(t, int64(1), l.Available())
	l.Take()
	require.Equal(t, int64(0), l.Available())
	put := false
	go func() {
		l.Take()
		require.Equal(t, true, put)
	}()
	time.Sleep(time.Millisecond * 100)
	put = true
	l.Return()
}
