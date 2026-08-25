# plugin-k8sgen

The `plugin-k8sgen` plugin candy of the [opencharly/charly](https://github.com/opencharly/charly)
candy library, as a standalone repo (the candy de-submodule cutover, plugin
kind). The Go module lives at `candy/plugin-k8sgen/` with module path
`github.com/opencharly/plugin-k8sgen/candy/plugin-k8sgen`; the charly resolver fetches this repo at the pinned tag and
the compiled-in wiring imports the module at that path.
