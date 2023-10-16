// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a devcenter.
func LookupDevCenter(ctx *pulumi.Context, args *LookupDevCenterArgs, opts ...pulumi.InvokeOption) (*LookupDevCenterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDevCenterResult
	err := ctx.Invoke("azure-native:devcenter/v20230801preview:getDevCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDevCenterArgs struct {
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a devcenter resource.
type LookupDevCenterResult struct {
	// The URI of the Dev Center.
	DevCenterUri string `pulumi:"devCenterUri"`
	// Encryption settings to be used for server-side encryption for proprietary content (such as catalogs, logs, customizations).
	Encryption *EncryptionResponse `pulumi:"encryption"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Managed identity properties
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDevCenterOutput(ctx *pulumi.Context, args LookupDevCenterOutputArgs, opts ...pulumi.InvokeOption) LookupDevCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDevCenterResult, error) {
			args := v.(LookupDevCenterArgs)
			r, err := LookupDevCenter(ctx, &args, opts...)
			var s LookupDevCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDevCenterResultOutput)
}

type LookupDevCenterOutputArgs struct {
	// The name of the devcenter.
	DevCenterName pulumi.StringInput `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDevCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevCenterArgs)(nil)).Elem()
}

// Represents a devcenter resource.
type LookupDevCenterResultOutput struct{ *pulumi.OutputState }

func (LookupDevCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevCenterResult)(nil)).Elem()
}

func (o LookupDevCenterResultOutput) ToLookupDevCenterResultOutput() LookupDevCenterResultOutput {
	return o
}

func (o LookupDevCenterResultOutput) ToLookupDevCenterResultOutputWithContext(ctx context.Context) LookupDevCenterResultOutput {
	return o
}

func (o LookupDevCenterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDevCenterResult] {
	return pulumix.Output[LookupDevCenterResult]{
		OutputState: o.OutputState,
	}
}

// The URI of the Dev Center.
func (o LookupDevCenterResultOutput) DevCenterUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.DevCenterUri }).(pulumi.StringOutput)
}

// Encryption settings to be used for server-side encryption for proprietary content (such as catalogs, logs, customizations).
func (o LookupDevCenterResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupDevCenterResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDevCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed identity properties
func (o LookupDevCenterResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDevCenterResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupDevCenterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDevCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupDevCenterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDevCenterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDevCenterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDevCenterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDevCenterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDevCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDevCenterResultOutput{})
}