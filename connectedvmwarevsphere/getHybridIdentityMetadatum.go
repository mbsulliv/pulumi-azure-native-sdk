// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Implements HybridIdentityMetadata GET method.
// Azure REST API version: 2022-07-15-preview.
//
// Other available API versions: 2023-03-01-preview.
func LookupHybridIdentityMetadatum(ctx *pulumi.Context, args *LookupHybridIdentityMetadatumArgs, opts ...pulumi.InvokeOption) (*LookupHybridIdentityMetadatumResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHybridIdentityMetadatumResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere:getHybridIdentityMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridIdentityMetadatumArgs struct {
	// Name of the HybridIdentityMetadata.
	MetadataName string `pulumi:"metadataName"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the vm.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// Defines the HybridIdentityMetadata.
type LookupHybridIdentityMetadatumResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity IdentityResponse `pulumi:"identity"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets or sets the Public Key.
	PublicKey *string `pulumi:"publicKey"`
	// The system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Gets or sets the Vm Id.
	VmId *string `pulumi:"vmId"`
}

func LookupHybridIdentityMetadatumOutput(ctx *pulumi.Context, args LookupHybridIdentityMetadatumOutputArgs, opts ...pulumi.InvokeOption) LookupHybridIdentityMetadatumResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridIdentityMetadatumResult, error) {
			args := v.(LookupHybridIdentityMetadatumArgs)
			r, err := LookupHybridIdentityMetadatum(ctx, &args, opts...)
			var s LookupHybridIdentityMetadatumResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridIdentityMetadatumResultOutput)
}

type LookupHybridIdentityMetadatumOutputArgs struct {
	// Name of the HybridIdentityMetadata.
	MetadataName pulumi.StringInput `pulumi:"metadataName"`
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the vm.
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupHybridIdentityMetadatumOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridIdentityMetadatumArgs)(nil)).Elem()
}

// Defines the HybridIdentityMetadata.
type LookupHybridIdentityMetadatumResultOutput struct{ *pulumi.OutputState }

func (LookupHybridIdentityMetadatumResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridIdentityMetadatumResult)(nil)).Elem()
}

func (o LookupHybridIdentityMetadatumResultOutput) ToLookupHybridIdentityMetadatumResultOutput() LookupHybridIdentityMetadatumResultOutput {
	return o
}

func (o LookupHybridIdentityMetadatumResultOutput) ToLookupHybridIdentityMetadatumResultOutputWithContext(ctx context.Context) LookupHybridIdentityMetadatumResultOutput {
	return o
}

func (o LookupHybridIdentityMetadatumResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupHybridIdentityMetadatumResult] {
	return pulumix.Output[LookupHybridIdentityMetadatumResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupHybridIdentityMetadatumResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupHybridIdentityMetadatumResultOutput) Identity() IdentityResponseOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) IdentityResponse { return v.Identity }).(IdentityResponseOutput)
}

// The name of the resource
func (o LookupHybridIdentityMetadatumResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o LookupHybridIdentityMetadatumResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the Public Key.
func (o LookupHybridIdentityMetadatumResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// The system data.
func (o LookupHybridIdentityMetadatumResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupHybridIdentityMetadatumResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the Vm Id.
func (o LookupHybridIdentityMetadatumResultOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridIdentityMetadatumResultOutput{})
}
