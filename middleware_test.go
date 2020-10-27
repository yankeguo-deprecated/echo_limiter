package echo_limiter

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	var coc int64
	var id int64
	e := echo.New()
	h := func(c echo.Context) error {
		atomic.AddInt64(&id, 1)
		t.Logf("started: %d", id)
		defer t.Logf("finished: %d", id)
		atomic.AddInt64(&coc, 1)
		defer atomic.AddInt64(&coc, -1)
		require.True(t, true, coc <= 4)

		a, ok := GetAvailable(c)
		require.Equal(t, true, a <= 4)
		require.Equal(t, true, ok)
		time.Sleep(time.Millisecond * 100)
		return c.String(http.StatusOK, "OK")
	}
	wh := New(FixedBucketLimiter(4))(h)

	for i := 0; i < 20; i++ {
		err := wh(e.NewContext(httptest.NewRequest(http.MethodGet, "/", nil), httptest.NewRecorder()))
		require.NoError(t, err)
	}
}
