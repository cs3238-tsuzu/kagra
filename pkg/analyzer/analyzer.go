package analyzer

// Analyzer analyzes Kubernetes RBAC resources
type Analyzer struct {
	*Option
}

// NewAnalyzer initializes an Analyzer
func NewAnalyzer(opt *Option) *Analyzer {
	return &Analyzer{}
}
