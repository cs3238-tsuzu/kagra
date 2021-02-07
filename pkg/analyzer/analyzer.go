package analyzer

// Analyzer analyzes Kubernetes RBAC resources
type Analyzer struct {
	opt *Option
}

// NewAnalyzer initializes an Analyzer
func NewAnalyzer(opt *Option) *Analyzer {
	return &Analyzer{}
}
