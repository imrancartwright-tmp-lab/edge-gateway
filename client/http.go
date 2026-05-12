var sharedTransport = &http.Transport{
    MaxIdleConns:        100,
    MaxIdleConnsPerHost: 10,
    IdleConnTimeout:     90 * time.Second,
    ForceAttemptHTTP2:   true,
}

func NewClient() *http.Client {
    return &http.Client{
        Transport: sharedTransport,
        Timeout:   30 * time.Second,
    }
}
