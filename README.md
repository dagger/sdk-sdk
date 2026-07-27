# sdk-sdk

Shared contract checks for official SDK helper modules.

For black-box lifecycle checks that drive an SDK module through the real CLI
(`dagger sdk install`, `dagger module init`, `dagger generate`), see
[`sdk-test`](./sdk-test):

```sh
dagger -m github.com/dagger/sdk-sdk/sdk-test -W <sdk-repo> check
```

Start a new Dang SDK helper module:

```sh
dagger call sdk-sdk init --name my-sdk
```

`--name` is the Dagger module name. The generated Dang root type is derived
from it, for example `my-sdk` becomes `MySdk`.

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
- `Mod.engine.required`
- `Mod.engine.require(version): Changeset!`

The Go SDK currently exposes `Mod.deps`; the contract accepts that alias while
the shared API is being introduced.

Changeset paths are treated as module-relative. For example, `init` is expected
to report `.dagger/modules/<name>/dagger.json`, and dependency updates are
expected to report `dagger.json`.
