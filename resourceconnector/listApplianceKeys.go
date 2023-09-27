// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourceconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the cluster customer credentials for the dedicated appliance.
// Azure REST API version: 2022-10-27.
func ListApplianceKeys(ctx *pulumi.Context, args *ListApplianceKeysArgs, opts ...pulumi.InvokeOption) (*ListApplianceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListApplianceKeysResult
	err := ctx.Invoke("azure-native:resourceconnector:listApplianceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplianceKeysArgs struct {
	// This sets the type of artifact being returned, when empty no artifact endpoint is returned.
	ArtifactType *string `pulumi:"artifactType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Appliances name.
	ResourceName string `pulumi:"resourceName"`
}

// The List Cluster Keys Results appliance.
type ListApplianceKeysResult struct {
	// Map of artifacts that contains a list of ArtifactProfile used to upload artifacts such as logs.
	ArtifactProfiles map[string]ArtifactProfileResponse `pulumi:"artifactProfiles"`
	// The list of appliance kubeconfigs.
	Kubeconfigs []ApplianceCredentialKubeconfigResponse `pulumi:"kubeconfigs"`
	// Map of Customer User Public, Private SSH Keys and Certificate when available.
	SshKeys map[string]SSHKeyResponse `pulumi:"sshKeys"`
}

func ListApplianceKeysOutput(ctx *pulumi.Context, args ListApplianceKeysOutputArgs, opts ...pulumi.InvokeOption) ListApplianceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApplianceKeysResult, error) {
			args := v.(ListApplianceKeysArgs)
			r, err := ListApplianceKeys(ctx, &args, opts...)
			var s ListApplianceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApplianceKeysResultOutput)
}

type ListApplianceKeysOutputArgs struct {
	// This sets the type of artifact being returned, when empty no artifact endpoint is returned.
	ArtifactType pulumi.StringPtrInput `pulumi:"artifactType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Appliances name.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListApplianceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceKeysArgs)(nil)).Elem()
}

// The List Cluster Keys Results appliance.
type ListApplianceKeysResultOutput struct{ *pulumi.OutputState }

func (ListApplianceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceKeysResult)(nil)).Elem()
}

func (o ListApplianceKeysResultOutput) ToListApplianceKeysResultOutput() ListApplianceKeysResultOutput {
	return o
}

func (o ListApplianceKeysResultOutput) ToListApplianceKeysResultOutputWithContext(ctx context.Context) ListApplianceKeysResultOutput {
	return o
}

func (o ListApplianceKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListApplianceKeysResult] {
	return pulumix.Output[ListApplianceKeysResult]{
		OutputState: o.OutputState,
	}
}

// Map of artifacts that contains a list of ArtifactProfile used to upload artifacts such as logs.
func (o ListApplianceKeysResultOutput) ArtifactProfiles() ArtifactProfileResponseMapOutput {
	return o.ApplyT(func(v ListApplianceKeysResult) map[string]ArtifactProfileResponse { return v.ArtifactProfiles }).(ArtifactProfileResponseMapOutput)
}

// The list of appliance kubeconfigs.
func (o ListApplianceKeysResultOutput) Kubeconfigs() ApplianceCredentialKubeconfigResponseArrayOutput {
	return o.ApplyT(func(v ListApplianceKeysResult) []ApplianceCredentialKubeconfigResponse { return v.Kubeconfigs }).(ApplianceCredentialKubeconfigResponseArrayOutput)
}

// Map of Customer User Public, Private SSH Keys and Certificate when available.
func (o ListApplianceKeysResultOutput) SshKeys() SSHKeyResponseMapOutput {
	return o.ApplyT(func(v ListApplianceKeysResult) map[string]SSHKeyResponse { return v.SshKeys }).(SSHKeyResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApplianceKeysResultOutput{})
}
