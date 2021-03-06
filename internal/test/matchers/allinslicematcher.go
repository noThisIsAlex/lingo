package matchers

import (
	"fmt"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// AllInSliceMatcher matches a slice of entries to matchers.
type AllInSliceMatcher struct {
	Expected  []interface{}
	idxFailed int
}

func (matcher *AllInSliceMatcher) Match(actual interface{}) (success bool, err error) {
	if isNil(actual) {
		return false, fmt.Errorf("expected a array/slice, got nil")
	}

	if !isArrayOrSlice(actual) {
		return false, fmt.Errorf("expected an array/slice.  Got:\n%s", format.Object(actual, 1))
	}

	s, err := ToSliceE(actual)
	if err != nil {
		return false, fmt.Errorf("unable to convert slice: %w", err)
	}

	if len(s) != len(matcher.Expected) {
		return false, fmt.Errorf("had length %d but expected %d elements",
			len(s), len(matcher.Expected))
	}

	var subMatcher types.GomegaMatcher
	var hasSubMatcher bool
	if matcher.Expected != nil {
		for idx, exp := range matcher.Expected {
			subMatcher, hasSubMatcher = exp.(types.GomegaMatcher)
			if !hasSubMatcher {
				subMatcher = gomega.Equal(exp)
			}
			valueToMatch := s[idx]
			subSuccess, subErr := subMatcher.Match(valueToMatch)
			if !subSuccess {
				matcher.idxFailed = idx
				return subSuccess, subErr
			}
		}
		return true, nil
	}

	return false, fmt.Errorf(
		"AllInSliceMatcher must be passed zero or more multiple values (for Equal or) matchers.  Got:\n%s",
		format.Object(matcher.Expected, 1))
}

func (matcher *AllInSliceMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, fmt.Sprintf(
		"to match a slice at index %d", matcher.idxFailed), matcher.Expected)
}

func (matcher *AllInSliceMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, fmt.Sprintf(
		"not to match a slice but did at index %d", matcher.idxFailed), matcher.Expected)
}

// ToSliceE casts an interface to a []interface{} type.
func ToSliceE(i interface{}) ([]interface{}, error) {
	var s []interface{}

	switch v := i.(type) {
	case [][]interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	case []interface{}:
		return append(s, v...), nil
	case []map[string]interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	default:
		return s, fmt.Errorf("unable to cast %#v of type %T to []interface{}", i, i)
	}
}
