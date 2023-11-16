// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231114preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a ApplicationResource
func LookupApplicationResource(ctx *pulumi.Context, args *LookupApplicationResourceArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationResourceResult
	err := ctx.Invoke("azure-native:integrationspaces/v20231114preview:getApplicationResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationResourceArgs struct {
	// The name of the Application
	ApplicationName string `pulumi:"applicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the application resource.
	ResourceName string `pulumi:"resourceName"`
	// The name of the space
	SpaceName string `pulumi:"spaceName"`
}

// A resource under application.
type LookupApplicationResourceResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// The Arm id of the application resource.
	ResourceId string `pulumi:"resourceId"`
	// The kind of the application resource.
	ResourceKind *string `pulumi:"resourceKind"`
	// The type of the application resource.
	ResourceType string `pulumi:"resourceType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupApplicationResourceOutput(ctx *pulumi.Context, args LookupApplicationResourceOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResourceResult, error) {
			args := v.(LookupApplicationResourceArgs)
			r, err := LookupApplicationResource(ctx, &args, opts...)
			var s LookupApplicationResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResourceResultOutput)
}

type LookupApplicationResourceOutputArgs struct {
	// The name of the Application
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the application resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The name of the space
	SpaceName pulumi.StringInput `pulumi:"spaceName"`
}

func (LookupApplicationResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResourceArgs)(nil)).Elem()
}

// A resource under application.
type LookupApplicationResourceResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResourceResult)(nil)).Elem()
}

func (o LookupApplicationResourceResultOutput) ToLookupApplicationResourceResultOutput() LookupApplicationResourceResultOutput {
	return o
}

func (o LookupApplicationResourceResultOutput) ToLookupApplicationResourceResultOutputWithContext(ctx context.Context) LookupApplicationResourceResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApplicationResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApplicationResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupApplicationResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Arm id of the application resource.
func (o LookupApplicationResourceResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResourceResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// The kind of the application resource.
func (o LookupApplicationResourceResultOutput) ResourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResourceResult) *string { return v.ResourceKind }).(pulumi.StringPtrOutput)
}

// The type of the application resource.
func (o LookupApplicationResourceResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResourceResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupApplicationResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApplicationResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResourceResultOutput{})
}
