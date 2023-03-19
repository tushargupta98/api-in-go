package middleware

import (
	"github.com/go-chi/chi"
)

// AddRateLimiting adds rate limiting middleware to the router.
func AddRateLimiting(router *chi.Mux) error { /*
		limiter, err := newLimiter()
		if err != nil {
			return err
		}
		httpLimiter, err := throttled.NewGCRARateLimiter(limiter)
		if err != nil {
			return errors.Wrap(err, "failed to create GCRARateLimiter")
		}
		router.Use(throttled.HTTPRateLimiter(limiter)) */
	return nil
}

// newLimiter creates a new rate limiter.
/*func newLimiter() (throttled.HTTPRateLimiter, error) {
	// Connect to Redis.
	redisClient := redis.GetClient()

	// Get the rate limiter settings from the configuration.
	rate, burst, err := config.GetRateLimit()
	if err != nil {
		return nil, err
	}

	// Create a new Redis store.
	store, err := redigostore.New(redisClient, "throttled", rate*2)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Redis store for rate limiter")
	}

	// Create a new limiter.
	limiter, err := throttled.NewGCRARateLimiter(store, throttled.VaryByHeader("User-Agent"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rate limiter")
	}

	// Create a new HTTP rate limiter.
	httpLimiter, err := throttled.NewGCRARateLimiter(limiter, throttled.VaryByHeader("User-Agent"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create HTTP rate limiter")
	}

	// Set the burst limit for the HTTP rate limiter.
	httpLimiter.SetBurst(burst)

	return httpLimiter, nil
}*/
