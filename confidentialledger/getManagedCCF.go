// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package confidentialledger

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the properties of a Managed CCF app.
// Azure REST API version: 2023-01-26-preview.
//
// Other available API versions: 2023-06-28-preview.
func LookupManagedCCF(ctx *pulumi.Context, args *LookupManagedCCFArgs, opts ...pulumi.InvokeOption) (*LookupManagedCCFResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedCCFResult
	err := ctx.Invoke("azure-native:confidentialledger:getManagedCCF", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedCCFArgs struct {
	// Name of the Managed CCF
	AppName string `pulumi:"appName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Managed CCF. Contains the properties of Managed CCF Resource.
type LookupManagedCCFResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of Managed CCF Resource.
	Properties ManagedCCFPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupManagedCCFOutput(ctx *pulumi.Context, args LookupManagedCCFOutputArgs, opts ...pulumi.InvokeOption) LookupManagedCCFResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedCCFResult, error) {
			args := v.(LookupManagedCCFArgs)
			r, err := LookupManagedCCF(ctx, &args, opts...)
			var s LookupManagedCCFResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedCCFResultOutput)
}

type LookupManagedCCFOutputArgs struct {
	// Name of the Managed CCF
	AppName pulumi.StringInput `pulumi:"appName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedCCFOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedCCFArgs)(nil)).Elem()
}

// Managed CCF. Contains the properties of Managed CCF Resource.
type LookupManagedCCFResultOutput struct{ *pulumi.OutputState }

func (LookupManagedCCFResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedCCFResult)(nil)).Elem()
}

func (o LookupManagedCCFResultOutput) ToLookupManagedCCFResultOutput() LookupManagedCCFResultOutput {
	return o
}

func (o LookupManagedCCFResultOutput) ToLookupManagedCCFResultOutputWithContext(ctx context.Context) LookupManagedCCFResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedCCFResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupManagedCCFResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagedCCFResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of Managed CCF Resource.
func (o LookupManagedCCFResultOutput) Properties() ManagedCCFPropertiesResponseOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) ManagedCCFPropertiesResponse { return v.Properties }).(ManagedCCFPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupManagedCCFResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupManagedCCFResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagedCCFResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedCCFResultOutput{})
}
