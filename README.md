# plugin-k8sgen

The Kustomize-generator plugin for [opencharly/charly](https://github.com/opencharly/charly) —
the pure manifest generation that turns a deployment node, a `kind:kubernetes`
cluster template, and the image's ports/uid/gid into a Kustomize
`base/` + `overlays/<inst>/` tree.

`GenerateTree` is **pure**: no disk I/O, no egress — it collects manifest docs
(as JSON) at RELATIVE paths and returns them; the caller owns `RemoveAll`/
`MkdirAll`, the raw manifest copy, the egress validation, and the YAML writes.

## What it provides

| Capability | Surface |
|---|---|
| `verb:k8sgen` | the `OpEmit` manifest generator, invoked peer-to-peer by `candy/plugin-kube`'s `materializeKustomize` |

## What it generates

- `Deployment` / `StatefulSet` / `DaemonSet` / `Job` / `CronJob` by kind
  heuristic.
- `Service` (ClusterIP) from the image Capabilities' ports.
- `PersistentVolumeClaim` from `deployment.Storage` (plus
  `volumeClaimTemplates` for a StatefulSet).
- `Ingress` when `deployment.Expose.Host` is set.
- `kustomization.yaml` wiring.

Out of scope (a deployment needing any of these authors it as an extra overlay
manifest alongside the generated base): ConfigMap / Secret / ExternalSecret,
HorizontalPodAutoscaler / PodDisruptionBudget, NetworkPolicy, ServiceMonitor,
the Gateway API HTTPRoute variant.

## How it is invoked

The plugin is compiled into charly. `candy/plugin-kube`'s `materializeKustomize`
extracts Ports/UID/GID from the image Capabilities, `InvokeProvider`s this
plugin's `OpEmit` with a `spec.KubernetesGenInput`, and does the disk I/O + the
egress gate before the bytes hit disk. There is no authored `plugin:` step.

## Layout

- `candy/plugin-k8sgen/` — the plugin module: `main.go` (the `OpEmit` entry),
  `k8sgen.go` (`GenerateTree`), `schema/k8sgen.cue`, `cmd/serve/main.go`.
- `candy/plugin-k8sgen/charly.yml` — the `plugin-k8sgen:` candy entity.
- `charly.yml` — the root project manifest (`discover: candy`).
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.

## Related

- Owning skill: `/charly-kubernetes:kubernetes` — the Kubernetes deploy surface
  and Kustomize generation. This candy carries no `skill:` entity of its own; the
  gap is tracked in
  [opencharly/opencharly#291](https://github.com/opencharly/opencharly/issues/291).
- `/charly-internals:egress` — the host-side egress gate the caller applies
  before the bytes hit disk.
- `/charly-internals:plugin` — the plugin/provider model.
- [`opencharly/charly`](https://github.com/opencharly/charly) — the charly CLI.
