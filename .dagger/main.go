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
	"dagger/quarkus-getting-started/internal/dagger"
)

type QuarkusGettingStarted struct{}

func (m *QuarkusGettingStarted) RunJvm(
	// +defaultPath="."
	src *dagger.Directory,
) *dagger.Service {
	return dag.
		Java().
		Quarkus(src).
		JvmImage().
		WithExposedPort(8080).
		AsService(dagger.ContainerAsServiceOpts{UseEntrypoint: true})
}
