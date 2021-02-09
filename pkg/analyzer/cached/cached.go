package cachedanalyzer

import (
	"github.com/cs3238-tsuzu/kagra/pkg/analyzer"
)

// New returns a new in-memory cached Analyzer
func New(analyzer analyzer.Analyzer) analyzer.Analyzer {
	return &cachedAnalyzer{
		cache:    map[string]interface{}{},
		internal: analyzer,
	}
}

type cachedAnalyzer struct {
	cache    map[string]interface{}
	internal analyzer.Analyzer
}
