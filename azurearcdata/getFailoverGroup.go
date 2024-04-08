// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurearcdata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a failover group resource
// Azure REST API version: 2023-01-15-preview.
//
// Other available API versions: 2024-01-01.
func LookupFailoverGroup(ctx *pulumi.Context, args *LookupFailoverGroupArgs, opts ...pulumi.InvokeOption) (*LookupFailoverGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFailoverGroupResult
	err := ctx.Invoke("azure-native:azurearcdata:getFailoverGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFailoverGroupArgs struct {
	// The name of the Failover Group
	FailoverGroupName string `pulumi:"failoverGroupName"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of SQL Managed Instance
	SqlManagedInstanceName string `pulumi:"sqlManagedInstanceName"`
}

// A failover group resource.
type LookupFailoverGroupResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// null
	Properties FailoverGroupPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupFailoverGroupResult
func (val *LookupFailoverGroupResult) Defaults() *LookupFailoverGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupFailoverGroupOutput(ctx *pulumi.Context, args LookupFailoverGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFailoverGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFailoverGroupResult, error) {
			args := v.(LookupFailoverGroupArgs)
			r, err := LookupFailoverGroup(ctx, &args, opts...)
			var s LookupFailoverGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFailoverGroupResultOutput)
}

type LookupFailoverGroupOutputArgs struct {
	// The name of the Failover Group
	FailoverGroupName pulumi.StringInput `pulumi:"failoverGroupName"`
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of SQL Managed Instance
	SqlManagedInstanceName pulumi.StringInput `pulumi:"sqlManagedInstanceName"`
}

func (LookupFailoverGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverGroupArgs)(nil)).Elem()
}

// A failover group resource.
type LookupFailoverGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFailoverGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverGroupResult)(nil)).Elem()
}

func (o LookupFailoverGroupResultOutput) ToLookupFailoverGroupResultOutput() LookupFailoverGroupResultOutput {
	return o
}

func (o LookupFailoverGroupResultOutput) ToLookupFailoverGroupResultOutputWithContext(ctx context.Context) LookupFailoverGroupResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFailoverGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFailoverGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// null
func (o LookupFailoverGroupResultOutput) Properties() FailoverGroupPropertiesResponseOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) FailoverGroupPropertiesResponse { return v.Properties }).(FailoverGroupPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFailoverGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFailoverGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFailoverGroupResultOutput{})
}
