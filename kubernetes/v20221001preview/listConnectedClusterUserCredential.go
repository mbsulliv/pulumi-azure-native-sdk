// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets cluster user credentials of the connected cluster with a specified resource group and name.
func ListConnectedClusterUserCredential(ctx *pulumi.Context, args *ListConnectedClusterUserCredentialArgs, opts ...pulumi.InvokeOption) (*ListConnectedClusterUserCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListConnectedClusterUserCredentialResult
	err := ctx.Invoke("azure-native:kubernetes/v20221001preview:listConnectedClusterUserCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedClusterUserCredentialArgs struct {
	// The mode of client authentication.
	AuthenticationMethod string `pulumi:"authenticationMethod"`
	// Boolean value to indicate whether the request is for client side proxy or not
	ClientProxy bool `pulumi:"clientProxy"`
	// The name of the Kubernetes cluster on which get is called.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The list of credential result response.
type ListConnectedClusterUserCredentialResult struct {
	// Contains the REP (rendezvous endpoint) and “Sender” access token.
	HybridConnectionConfig HybridConnectionConfigResponse `pulumi:"hybridConnectionConfig"`
	// Base64-encoded Kubernetes configuration file.
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListConnectedClusterUserCredentialOutput(ctx *pulumi.Context, args ListConnectedClusterUserCredentialOutputArgs, opts ...pulumi.InvokeOption) ListConnectedClusterUserCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConnectedClusterUserCredentialResult, error) {
			args := v.(ListConnectedClusterUserCredentialArgs)
			r, err := ListConnectedClusterUserCredential(ctx, &args, opts...)
			var s ListConnectedClusterUserCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConnectedClusterUserCredentialResultOutput)
}

type ListConnectedClusterUserCredentialOutputArgs struct {
	// The mode of client authentication.
	AuthenticationMethod pulumi.StringInput `pulumi:"authenticationMethod"`
	// Boolean value to indicate whether the request is for client side proxy or not
	ClientProxy pulumi.BoolInput `pulumi:"clientProxy"`
	// The name of the Kubernetes cluster on which get is called.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListConnectedClusterUserCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedClusterUserCredentialArgs)(nil)).Elem()
}

// The list of credential result response.
type ListConnectedClusterUserCredentialResultOutput struct{ *pulumi.OutputState }

func (ListConnectedClusterUserCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedClusterUserCredentialResult)(nil)).Elem()
}

func (o ListConnectedClusterUserCredentialResultOutput) ToListConnectedClusterUserCredentialResultOutput() ListConnectedClusterUserCredentialResultOutput {
	return o
}

func (o ListConnectedClusterUserCredentialResultOutput) ToListConnectedClusterUserCredentialResultOutputWithContext(ctx context.Context) ListConnectedClusterUserCredentialResultOutput {
	return o
}

func (o ListConnectedClusterUserCredentialResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListConnectedClusterUserCredentialResult] {
	return pulumix.Output[ListConnectedClusterUserCredentialResult]{
		OutputState: o.OutputState,
	}
}

// Contains the REP (rendezvous endpoint) and “Sender” access token.
func (o ListConnectedClusterUserCredentialResultOutput) HybridConnectionConfig() HybridConnectionConfigResponseOutput {
	return o.ApplyT(func(v ListConnectedClusterUserCredentialResult) HybridConnectionConfigResponse {
		return v.HybridConnectionConfig
	}).(HybridConnectionConfigResponseOutput)
}

// Base64-encoded Kubernetes configuration file.
func (o ListConnectedClusterUserCredentialResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListConnectedClusterUserCredentialResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectedClusterUserCredentialResultOutput{})
}
