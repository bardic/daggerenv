package main

import (
	"context"
	"dagger/dagger-env/internal/dagger"
	"strings"
)

type DaggerEnv struct{}

func (m *DaggerEnv) Load(ctx context.Context, container *dagger.Container, src *dagger.Directory) (*dagger.Container, error) {

	envFiles, err := src.Glob(ctx, "*.env")

	if err != nil {
		return nil, err
	}

	container, err = parse(ctx, container, src, envFiles)

	if err != nil {
		return nil, err
	}

	return container, nil
}

func parse(ctx context.Context, container *dagger.Container, src *dagger.Directory, envFiles []string) (*dagger.Container, error) {
	for _, file := range envFiles {

		fileContent, err := src.File(file).Contents(ctx)

		if err != nil {
			return nil, err
		}

		envPair := strings.Split(fileContent, "\n")

		for _, pair := range envPair {
			envVals := strings.SplitN(pair, "=", 2)
			if strings.Contains(file, "secret") {
				container = container.WithSecretVariable(envVals[0], dag.SetSecret(envVals[0], envVals[1]))
			} else {
				container = container.WithEnvVariable(envVals[0], envVals[1])
			}
		}
	}

	return container, nil
}
