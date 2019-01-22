package printer

import (
	"github.com/heptio/developer-dash/internal/view/component"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func printSelector(selector *metav1.LabelSelector) *component.Selectors {
	s := component.NewSelectors(nil)
	if selector == nil {
		return s
	}

	for k, v := range selector.MatchLabels {
		s.Add(component.NewLabelSelector(k, v))
	}

	for _, e := range selector.MatchExpressions {
		es := component.NewExpressionSelector(e.Key, component.Operator(e.Operator), e.Values)
		s.Add(es)
	}

	return s
}
