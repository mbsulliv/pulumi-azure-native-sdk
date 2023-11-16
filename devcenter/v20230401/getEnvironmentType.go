// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an environment type.
func LookupEnvironmentType(ctx *pulumi.Context, args *LookupEnvironmentTypeArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentTypeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentTypeResult
	err := ctx.Invoke("azure-native:devcenter/v20230401:getEnvironmentType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentTypeArgs struct {
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the environment type.
	EnvironmentTypeName string `pulumi:"environmentTypeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents an environment type.
type LookupEnvironmentTypeResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
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

func LookupEnvironmentTypeOutput(ctx *pulumi.Context, args LookupEnvironmentTypeOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentTypeResult, error) {
			args := v.(LookupEnvironmentTypeArgs)
			r, err := LookupEnvironmentType(ctx, &args, opts...)
			var s LookupEnvironmentTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentTypeResultOutput)
}

type LookupEnvironmentTypeOutputArgs struct {
	// The name of the devcenter.
	DevCenterName pulumi.StringInput `pulumi:"devCenterName"`
	// The name of the environment type.
	EnvironmentTypeName pulumi.StringInput `pulumi:"environmentTypeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEnvironmentTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentTypeArgs)(nil)).Elem()
}

// Represents an environment type.
type LookupEnvironmentTypeResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentTypeResult)(nil)).Elem()
}

func (o LookupEnvironmentTypeResultOutput) ToLookupEnvironmentTypeResultOutput() LookupEnvironmentTypeResultOutput {
	return o
}

func (o LookupEnvironmentTypeResultOutput) ToLookupEnvironmentTypeResultOutputWithContext(ctx context.Context) LookupEnvironmentTypeResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEnvironmentTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEnvironmentTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupEnvironmentTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEnvironmentTypeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupEnvironmentTypeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEnvironmentTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentTypeResultOutput{})
}
