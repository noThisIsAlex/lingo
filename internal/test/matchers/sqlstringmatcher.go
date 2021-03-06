package matchers

import (
	"fmt"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// SQLStringMatcher matches a `core.SQL` `String()` value
type SQLStringMatcher struct {
	Expected interface{}
}

func (matcher *SQLStringMatcher) Match(actual interface{}) (success bool, err error) {
	if isNil(actual) {
		return false, fmt.Errorf("expected a core.SQL, got nil")
	}

	if !isSQL(actual) {
		return false, fmt.Errorf("expected a core.SQL.  Got:\n%s", format.Object(actual, 1))
	}

	s := toSQL(actual)

	var subMatcher types.GomegaMatcher
	var hasSubMatcher bool
	if matcher.Expected != nil {
		subMatcher, hasSubMatcher = (matcher.Expected).(types.GomegaMatcher)
		if hasSubMatcher {
			return subMatcher.Match(s.String())
		}
	}

	if exp, ok := toString(matcher.Expected); ok {
		return s.String() == exp, nil
	}

	return false, fmt.Errorf(
		"SQLStringMatcher must be passed a string or a string matcher.  Got:\n%s",
		format.Object(matcher.Expected, 1))
}

func (matcher *SQLStringMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to match core.SQL string", matcher.Expected)
}

func (matcher *SQLStringMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to match core.SQL string", matcher.Expected)
}
