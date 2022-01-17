package builder

const (
	DefaultMaxTotal = 10
	DefaultMaxIdle  = 9
	DefaultMinIdle  = 1
)

type (
	ResourcePoolConfig struct {
		Name     string
		MaxTotal int
		MaxIdle  int
		MinIdle  int
	}
)
