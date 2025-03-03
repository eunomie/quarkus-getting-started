// A generated module for QuarkusGettingStarted functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"

	"dagger/quarkus-getting-started/internal/dagger"
)

type QuarkusGettingStarted struct{
	Src *dagger.Directory
}

func New(
// +defaultPath="."
src *dagger.Directory) *QuarkusGettingStarted {
	return &QuarkusGettingStarted{
		Src: src,
	}
}

func (m *QuarkusGettingStarted) RunJvm() *dagger.Service {
	return dag.
		Java().
		Quarkus(m.Src).
		JvmImage().
		WithExposedPort(8080).
		AsService(dagger.ContainerAsServiceOpts{UseEntrypoint: true})
}

func (m *QuarkusGettingStarted) CodeMeThatFeature(ctx context.Context, feature string) *dagger.Directory {
	return dag.Toolbox(m.Src).Do(feature)
}

func (m *QuarkusGettingStarted) AddComments(ctx context.Context) *dagger.Directory {
	return dag.Toolbox(m.Src).AddComments()
}

func (m *QuarkusGettingStarted) Refactor(ctx context.Context) *dagger.Directory {
	return dag.Toolbox(m.Src).Refactor()
}

func (m *QuarkusGettingStarted) BumpDeps(ctx context.Context) *dagger.Directory {
	return dag.Toolbox(m.Src).BumpDeps()
}

func (m *QuarkusGettingStarted) Explain(ctx context.Context) (string, error) {
	return dag.Toolbox(m.Src).Explain(ctx)
}

func (m *QuarkusGettingStarted) FindBugs(ctx context.Context) (string, error) {
	return dag.Toolbox(m.Src).FindBugs(ctx)
}
