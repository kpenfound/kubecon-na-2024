// A generated module for KubeconNa2024 functions

package main

import (
	"context"

	"dagger/kubecon-na-2024/internal/dagger"
)

type KubeconNa2024 struct{}

// Returns the goreleaser test pipeline
func (m *KubeconNa2024) Test(ctx context.Context) (string, error) {
	source := dag.Git("github.com/goreleaser/goreleaser").
		Commit("f54fb113ee6415fd2ada737b3ffd71e5e32682ad").
		Tree()
	return dag.Goreleaser(dagger.GoreleaserOpts{Source: source}).
		Test().
		Output(ctx)
}
