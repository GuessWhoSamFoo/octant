package printer_test

import (
	"testing"
	"time"

	"github.com/heptio/developer-dash/internal/cache"
	"github.com/heptio/developer-dash/internal/overview/printer"
	"github.com/heptio/developer-dash/internal/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_ReplicaSetListHandler(t *testing.T) {
	printOptions := printer.Options{
		Cache: cache.NewMemoryCache(),
	}

	labels := map[string]string{
		"foo": "bar",
	}

	now := time.Unix(1547211430, 0)

	object := &appsv1.ReplicaSetList{
		Items: []appsv1.ReplicaSet{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "replicaset",
					CreationTimestamp: metav1.Time{
						Time: now,
					},
					Labels: labels,
				},
				Status: appsv1.ReplicaSetStatus{
					Replicas:          3,
					AvailableReplicas: 2,
				},
				Spec: appsv1.ReplicaSetSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "myapp",
						},
					},
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								corev1.Container{
									Name:  "nginx",
									Image: "nginx:1.15",
								},
								corev1.Container{
									Name:  "kuard",
									Image: "gcr.io/kuar-demo/kuard-amd64:1",
								},
							},
						},
					},
				},
			},
		},
	}

	got, err := printer.ReplicaSetListHandler(object, printOptions)
	require.NoError(t, err)

	containers := component.NewContainers()
	containers.Add("nginx", "nginx:1.15")
	containers.Add("kuard", "gcr.io/kuar-demo/kuard-amd64:1")

	cols := component.NewTableCols("Name", "Labels", "Status", "Age", "Containers", "Selector")
	expected := component.NewTable("ReplicaSets", cols)
	expected.Add(component.TableRow{
		"Name":       component.NewText("", "replicaset"),
		"Labels":     component.NewLabels(labels),
		"Age":        component.NewTimestamp(now),
		"Selector":   component.NewSelectors([]component.Selector{component.NewLabelSelector("app", "myapp")}),
		"Status":     component.NewText("", "2/3"),
		"Containers": containers,
	})

	assert.Equal(t, expected, got)
}
