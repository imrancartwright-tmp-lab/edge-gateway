// Configure retry policy and circuit breaker
RetryConfig: RetryConfig{
    MaxAttempts: 3,
    RetryOn: []int{429, 500, 502, 503, 504},
    BaseDelay: 200 * time.Millisecond,
    MaxDelay: 2 * time.Second,
    WithJitter: true,
},
CircuitBreakerConfig: CircuitBreakerConfig{
    ErrorThreshold:  50,          // percent
    Window:          10 * time.Second,
    HalfOpenRequests: 3,
    RecoveryTimeout: 30 * time.Second,
},
