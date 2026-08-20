# mod-test

Lightweight black-box testing helpers for Dagger modules.

`mod-test` mounts a workspace-rooted directory view, installs a Dagger CLI
release from `dl.dagger.io`, then runs readable `dagger api call -j` commands
against the target module.

Configure the CLI release with the top-level `daggerCliVersion` setting, either
through `[modules.<alias>.settings]` in `dagger.toml` or by constructing the
module directly — `modTest(daggerCliVersion: "1.0.0-beta.10").target(...)`. The
default is `1.0.0-beta.10`.

Callers provide:

- `workspaceView`: a directory containing every file required to load the target
  module.
- `sourceRootPath`: the target module path inside that directory.

Example:

```dang
pub smoke(ws: Workspace!): Void @check {
  let module = ws.moduleSource(".dagger/modules/fixture")
  let target = modTest.target(module.contextDirectory, module.sourceRootSubpath)

  target.assertJsonString(["echo", "--value", "hello"], "hello")
  target.assertFailure(["fail"], "fail should return a non-zero status")
}
```

The public API follows Go test-style semantics:

- `call(args)` captures stdout, stderr, and exit code, and fails if the command
  exits non-zero.
- `tryCall(args)` captures stdout, stderr, and exit code without requiring
  success.
- `assertSuccess`, `assertFailure`, `assertOutput`, and `assertJson*` helpers
  keep individual checks short and focused. Assertion failures report the
  caller's curated message; when the call itself failed they append its exit
  code and stderr, so a command that never ran is not read as a violated
  expectation. Use `tryCall` when a check needs raw stdout or stderr.
