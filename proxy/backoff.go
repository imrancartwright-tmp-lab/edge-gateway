func ExponentialBackoff(attempt int) time.Duration {
    base := time.Millisecond * 200
    max  := time.Second * 5
    // Full jitter: random value between base*2^attempt and max
    cap := time.Duration(float64(base) * math.Pow(2, float64(attempt)))
    if cap > max {
        cap = max
    }
    jitter := time.Duration(rand.Int63n(int64(cap)))
    if jitter < base {
        jitter = base
    }
    return jitter
}
