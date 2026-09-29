# AGENTS.md — plugin-k8sgen

Standalone plugin repo for the Kustomize manifest generator (`verb:k8sgen`). The
plugin is a Go module at `candy/plugin-k8sgen/` (module path
`github.com/opencharly/plugin-k8sgen/candy/plugin-k8sgen`); the root
`charly.yml` only declares `discover: candy` so the repo is a project and its
candy is scanned.

Canonical files:

- `candy/plugin-k8sgen/charly.yml` — the `plugin-k8sgen:` candy entity (`plugin:`
  block, `plan:` check).
- `candy/plugin-k8sgen/` — the Go source: `main.go` (the `OpEmit` entry),
  `k8sgen.go` (`GenerateTree`), `schema/k8sgen.cue`, `cmd/serve/main.go`.
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.
- `README.md` — user overview only; never agent guidance.

## Load these skills first (R0)

- `/charly-kubernetes:kubernetes` — the Kubernetes deploy surface and Kustomize
  generation. Load before changing the generator.
- `/charly-internals:egress` — the host-side egress gate the caller applies
  before the bytes hit disk.
- `/charly-internals:plugin` — the plugin authoring reference: the `plugin:`
  block, the unified Provider model, the per-plugin CUE-schema contract.
- `/charly-internals:git-workflow` — before any git/PR action.

## Build / validate / test

- `go build ./...` in `candy/plugin-k8sgen/` — compile the plugin module.
- `go test ./...` in `candy/plugin-k8sgen/` — the generator + schema tests.
- `charly box validate` at the repo root — the structural check (the candy +
  `plugin:` block, CUE schema).
- The merge gate is the **org-wide** `charly/pr-validator` (required check
  `validate / validate`, defined in `opencharly/.github`); this repo has **no**
  per-repo candy gate.
- The changed path is exercised by the `check-k8s-deploy` bed (via
  `candy/plugin-kube`).

## Modify this repo

- Edit the `plugin-k8sgen:` candy entity, the Go source, and
  `schema/k8sgen.cue` **together**.
- `GenerateTree` MUST stay pure (no disk I/O, no egress): it returns
  RELATIVE-pathed docs; the caller (`candy/plugin-kube`) owns the writes + the
  egress gate.
- There is NO authored `plugin_input`; the verb is invoked with the structured
  `spec.KubernetesGenInput`, so the schema DOCUMENTS the contract (no `#*Input`
  def).

## Landing

- The authoritative rulebook is the umbrella `AGENTS.md` in
  `opencharly/opencharly` and `charly/AGENTS.md` in the charly repo. Load
  `/charly-internals:git-workflow` before any git/PR action; history lives in
  `CHANGELOG/`. Do not restate its rules here.
