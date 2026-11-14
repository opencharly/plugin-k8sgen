package k8sgen

import "github.com/opencharly/spec/spec"

// clusterspec.go — the plugin's OWN view of the `kubernetes:` cluster template.
//
// opencharly/spec#304 unit 5 moves the authored `#Kubernetes` body (and the Go type generated
// from it) into the plugin that SERVES the kind, so this plugin can no longer borrow
// `spec.Kubernetes` to read the cluster policy it generates against. The kernel already ships
// the body OPAQUELY (Cutover K: `#KubernetesGenInput.cluster_raw`), and the contract's own note
// says the decoder is this plugin — "the full cluster model rides opaquely in Raw and is
// decoded by candy/plugin-k8sgen, never the kernel".
//
// So this file carries exactly the fields the generator READS, with the authored JSON tags.
// It is deliberately not a mirror of the whole body: a plugin that grows a field here is a
// plugin that reads it, and anything else stays in the body, untyped by anybody.

// clusterSpec is the cluster policy the Kustomize tree is generated against.
type clusterSpec struct {
	DefaultNamespace string           `json:"default_namespace"`
	AdmissionPolicy  string           `json:"admission_policy"`
	Defaults         clusterDefaults  `json:"defaults"`
	ImageDefault     imageDefaults    `json:"image_default"`
	Ingress          ingressDefaults  `json:"ingress"`
	Storage          storageDefaults  `json:"storage"`
	PodDefault       podDefaultPolicy `json:"pod_default"`
}

// clusterDefaults carries the labels/annotations every generated object inherits.
type clusterDefaults struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

// imageDefaults carries the container image policy.
type imageDefaults struct {
	PullPolicy  string   `json:"pull_policy"`
	PullSecrets []string `json:"pull_secrets"`
}

// ingressDefaults carries the ingress policy.
type ingressDefaults struct {
	Enabled         bool   `json:"enabled"`
	Class           string `json:"class"`
	CertIssuer      string `json:"cert_issuer"`
	PathTypeDefault string `json:"path_type_default"`
}

// storageDefaults carries the storage-class policy.
type storageDefaults struct {
	ClassDefault      string `json:"class_default"`
	ClassCheap        string `json:"class_cheap"`
	ClassEncrypted    string `json:"class_encrypted"`
	ClassFast         string `json:"class_fast"`
	AccessModeDefault string `json:"access_mode_default"`
}

// podDefaultPolicy carries the pod-level policy.
type podDefaultPolicy struct {
	PriorityClass string            `json:"priority_class"`
	Tolerations   []map[string]any  `json:"tolerations"`
	NodeSelector  map[string]string `json:"node_selector"`
}

// genInput is the plugin's own view of the generation request: the contract wire
// (spec.KubernetesGenInput, embedded so the JSON decodes exactly as it always has) plus the
// cluster body decoded into clusterSpec. The kernel sends the cluster OPAQUELY in
// `cluster_raw` (Cutover K); this plugin owns the decode, so no contract type is needed for it.
type genInput struct {
	spec.KubernetesGenInput
	cluster clusterSpec
}
