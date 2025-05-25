package limiter

type Limiter interface {
	Limit() bool
}
