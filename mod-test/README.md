# mod-test

Lightweight black-box testing helpers for Dagger modules.

`mod-test` mounts a workspace-rooted directory view, installs a pinned Dagger CLI
from `github.com/dagger/dagger/toolchains/cli-dev`, then runs readable
`dagger call -j` commands against the target module.

Callers provide:

- `workspaceView`: a directory containing every file required to load the target
  module.
- `sourceRootPath`: the target module path inside that directory.

Example:

```dang
pub smoke(ws: Workspace!): Void @check {
  let module = polyfill.workspace(ws).moduleSource(".dagger/modules/fixture")
  let target = modTest.target(module.workspaceView, module.sourceRootPath)

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
  keep individual checks short and focused.
