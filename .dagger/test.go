package main

import (
	"context"
	"dagger/dagger-env/internal/dagger"
	"strings"
)

func (m *DaggerEnv) Test(ctx context.Context, testData *dagger.Directory) (bool, error) {
	container := dag.Container().From("alpine:latest")

	container, err := m.Load(ctx, container, testData)

	if err != nil {
		return false, err
	}

	container = container.WithExec(
		[]string{
			"/bin/sh",
			"-c",
			"if [[ \"$HELLO\" = \"WORLD\" ]]; then echo \"PASS\" >> results; else echo \"FAIL\" >> results; fi;",
		},
	)

	container = container.WithExec(
		[]string{
			"/bin/sh",
			"-c",
			"if [[ \"$HELLO\" != \"GOODBYE\" ]]; then echo \"PASS\" >> results; else echo \"FAIL\" >> results; fi;",
		},
	)

	container = container.WithExec(
		[]string{
			"/bin/sh",
			"-c",
			"if [[ \"$SECRET\" = \"CLIMATE_CHANGE_IS_REAL\" ]]; then echo \"PASS\" >> results; else echo \"FAIL\" >> results; fi;",
		},
	)

	container = container.WithExec(
		[]string{
			"/bin/sh",
			"-c",
			"if [[ \"$SECRET\" != \"NATRUAL_RHYTHMS\" ]]; then echo \"PASS\" >> results; else echo \"FAIL\" >> results; fi;",
		},
	)

	results, err := container.File("results").Contents(ctx)

	if err != nil {
		return false, err
	}

	if strings.Contains(results, "FAIL") {
		return false, nil
	}

	return true, nil
}
