package main

import (
	"context"
	"reflect"
	"testing"
)

func TestModuleSourceIncludeFromConfigsIncludesDeclaredPaths(t *testing.T) {
	configs := map[string]sourceConfig{
		"app/dagger.json": {
			dependencies: []string{"../dep"},
			include:      []string{"../root.txt", "assets/**/*"},
		},
		"dep/dagger.json": {
			include: []string{"../shared.txt", "subdir/**/*"},
		},
	}

	got, err := moduleSourceIncludeFromConfigs(configs, "app")
	if err != nil {
		t.Fatal(err)
	}

	want := []string{
		"app",
		"app/**",
		"app/assets/**/*",
		"app/dagger.json",
		"dep",
		"dep/**",
		"dep/dagger.json",
		"dep/subdir/**/*",
		"root.txt",
		"shared.txt",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("include mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}

func TestModuleSourceIncludeReadsOnlyTargetAndLocalDependencies(t *testing.T) {
	var read []string

	got, err := moduleSourceInclude(context.Background(), "app", func(_ context.Context, p string) (sourceConfig, bool, error) {
		read = append(read, p)
		switch p {
		case "app":
			return sourceConfig{dependencies: []string{"../dep"}}, true, nil
		case "dep":
			return sourceConfig{}, true, nil
		default:
			t.Fatalf("unexpected config read: %s", p)
			return sourceConfig{}, false, nil
		}
	})
	if err != nil {
		t.Fatal(err)
	}

	wantRead := []string{"app", "dep"}
	if !reflect.DeepEqual(read, wantRead) {
		t.Fatalf("read paths mismatch:\n got: %#v\nwant: %#v", read, wantRead)
	}

	want := []string{
		"app",
		"app/**",
		"app/dagger.json",
		"dep",
		"dep/**",
		"dep/dagger.json",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("include mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}
