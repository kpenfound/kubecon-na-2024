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
		Commit("12155a336baaa705e9363e3001c048238f15ff34").
		Tree()
	return dag.Goreleaser(dagger.GoreleaserOpts{Source: source}).
		Test().
		Output(ctx)
}
