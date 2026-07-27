# sdk-test

Black-box contract checks for Dagger SDK modules, such as
`github.com/dagger/go-sdk`, `github.com/dagger/dang-sdk`,
`github.com/dagger/typescript-sdk`, and `github.com/dagger/python-sdk`.

Where `mod-test` calls one module's functions, `sdk-test` exercises an SDK the
way a user does across its whole lifecycle. The checks vendor the SDK module
into a scratch git workspace inside a runner container, install a release
Dagger CLI, then drive the SDK through real CLI commands:

- `dagger sdk install ./<sdk>` registers the SDK and marks it `as-sdk` in
  `dagger.toml`.
- `dagger module init <sdk> test-mod` scaffolds a new module, writes its
  `dagger-module.toml`, installs it in `dagger.toml`, and records the SDK as
  the module's authoring SDK.
- `dagger generate` succeeds on the fresh scaffold.
- After generation the scaffolded module serves functions:
  `dagger api functions test-mod`.
- `dagger sdk module-options <sdk>` introspects the SDK's `initModule`
  capability.
- `dagger module engine required` and `dagger module deps list` work from the
  scaffolded module directory.

Run the checks against an SDK repository:

```sh
dagger -m github.com/dagger/sdk-sdk/sdk-test -W <sdk-repo> check
```

For example:

```sh
dagger -m . -W https://github.com/dagger/go-sdk check
```

When run inside the sdk-sdk repository itself, the checks redirect to the
fixture SDK at `.dagger/modules/sdk-test-e2e/fixtures/sdk-under-test`, so
`dagger -m ./sdk-test -W . check` self-tests the harness.

Configure the CLI release with the top-level `dagger-cli-version` setting; the
default is `1.0.0-beta.7`. Individual targets accept `with-timeout` for slow
SDKs (the default command timeout is `10m`).

Custom checks can reuse the harness through `target`:

```dang
let testTarget = sdkTest.target(module.workspaceView, module.sourceRootPath)
testTarget.install.assertSuccess
testTarget.runInModule(["module", "deps", "list"]).assertSuccess
```
