# sdk-sdk

Shared contract checks for official SDK helper modules.

Run it from an SDK helper module workspace:

```sh
cd ./my/sdk/repo
dagger -m github.com/dagger/sdk-sdk check
```

The checks receive the current `Workspace`, find that workspace's `dagger.json`,
serve the SDK helper module from the workspace, inspect the public GraphQL schema
users see from the CLI, then exercise a small set of user-facing behaviors
without applying returned changesets.

The first coverage is intentionally small:

- `init(ws, name, path, template, ignoreGenerated): Changeset!`
- `mod(ws, path, findUp): Mod!`
- `Mod.path`
- `Mod.dependencies.add(source, name): Changeset!`

The Go SDK currently exposes `Mod.deps`; the contract accepts that alias while
the shared API is being introduced.

Changeset paths are treated as module-relative. For example, `init` is expected
to report `.dagger/modules/<name>/dagger.json`, and dependency updates are
expected to report `dagger.json`.
