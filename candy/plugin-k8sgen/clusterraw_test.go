package k8sgen

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/opencharly/sdk"
	pb "github.com/opencharly/spec/proto"
	"github.com/opencharly/spec/spec"
)

// TestInvoke_DecodesClusterRaw covers the kubernetes substrate-value de-type (Cutover K):
// the kernel ships the cluster body OPAQUELY in ClusterRaw and the plugin must decode it into
// its OWN view (clusterspec.go) before generating — opencharly/spec#304 unit 5 moved the
// authored shape out of the contract, so no contract type is in the path. deployNamespace
// reads that view's DefaultNamespace, so without the decode the overlay carries no namespace
// and this fails.
func TestInvoke_DecodesClusterRaw(t *testing.T) {
	clusterBody := []byte(`{"box":"app","default_namespace":"myns"}`)
	params, err := json.Marshal(spec.KubernetesGenInput{
		DeploymentName: "app",
		ImageRef:       "registry/app:1",
		ClusterRaw:     clusterBody,
	})
	if err != nil {
		t.Fatal(err)
	}
	reply, err := (&provider{}).Invoke(context.Background(), &pb.InvokeRequest{Op: sdk.OpEmit, ParamsJson: params})
	if err != nil {
		t.Fatalf("Invoke: %v", err)
	}
	if !strings.Contains(string(reply.GetResultJson()), "myns") {
		t.Errorf("generated tree did not use ClusterRaw.DefaultNamespace %q: %s", "myns", reply.GetResultJson())
	}
}
