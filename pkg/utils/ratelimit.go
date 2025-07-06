package utils

import (
	"sync"
	"time"

	"github.com/chise0904/golang_template_apiserver/pkg/errors"
	lru "github.com/hashicorp/golang-lru"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type tokenBucket struct {
	rate       float64    // tokens added per second
	capacity   float64    // maximum number of tokens
	tokens     float64    // current number of tokens
	lastRefill time.Time  // last refill time
	mu         sync.Mutex // protect the bucket state
}

func newTokenBucket(rate, capacity float64) *tokenBucket {
	return &tokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

// TryConsume attempts to consume a token.
func (tb *tokenBucket) TryConsume() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Refill tokens based on elapsed time.
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens = tb.tokens + elapsed*tb.rate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now

	// Check if a token can be consumed.
	if tb.tokens >= 1 {
		tb.tokens -= 1
		return true
	}
	return false
}

func RateLimitByClientIpMiddleware(poolSize int, rate float64, capacity float64, headerKey string) echo.MiddlewareFunc {
	cache, err := lru.New(poolSize)
	if err != nil {
		panic(err.Error())
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			clientIP := c.Request().Header.Get(headerKey)
			if clientIP == "" {
				clientIP = c.RealIP()
			}

			var bucket *tokenBucket
			if value, ok := cache.Get(clientIP); ok {
				bucket = value.(*tokenBucket)
			} else {
				// Create a new token bucket if not found.
				bucket = newTokenBucket(rate, capacity)
				cache.Add(clientIP, bucket)
			}

			l := log.Ctx(c.Request().Context())
			if !bucket.TryConsume() {
				l.Warn().Msgf("too many request from %s denied", clientIP)
				return errors.ErrorTooManyRequest()
			}

			return next(c)
		}
	}
}
