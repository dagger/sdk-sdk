# sdk-sdk

Shared contract checks for official SDK helper modules.

Start a new Dang SDK helper module:

```sh
dagger module init sdk-sdk my-sdk
```

The module name is the Dagger module name. The generated Dang root type is
derived from it, for example `my-sdk` becomes `MySdk`.

Run the checks from an SDK helper module workspace:

```sh
cd ./my/sdk/repo
dagger -m github.com/dagger/sdk-sdk check
```

The checks receive the current `Workspace`, serve the SDK helper module from the
workspace, and exercise its user-facing behavior without applying the returned
changesets.

Under CLI 1.0 the engine owns module bookkeeping — `dagger-module.toml`,
workspace config, and dependency and engine-version edits. An SDK helper module
implements only what is genuinely language-specific:

- `initModule(ws, name, path): Changeset!` — seed the SDK's own files for a new
  module. Because the engine owns module config, `initModule` must **not** write
  `dagger.json` / `dagger-module.toml`.
- a `@generate` hook — regenerate the modules the SDK manages. Managed modules
  are discovered via `currentModule.asSDK.modules`.

`initClient` (typed client generation) is an optional part of the contract and
is not required or exercised here.

Changeset paths are workspace-root-relative. For example, `initModule` for a
module named `my-sdk` is expected to seed files under `.dagger/modules/my-sdk/`.
