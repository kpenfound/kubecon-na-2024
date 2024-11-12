// A generated module for KubeconNa2024 functions

package main

import (
	"context"
)

type KubeconNa2024 struct{}

// Returns the goreleaser test pipeline
func (m *KubeconNa2024) Test(ctx context.Context) (string, error) {
	return dag.Goreleaser().Test().Output(ctx)
}
