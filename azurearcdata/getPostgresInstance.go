// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurearcdata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a postgres Instance resource
// Azure REST API version: 2023-01-15-preview.
//
// Other available API versions: 2024-01-01.
func LookupPostgresInstance(ctx *pulumi.Context, args *LookupPostgresInstanceArgs, opts ...pulumi.InvokeOption) (*LookupPostgresInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPostgresInstanceResult
	err := ctx.Invoke("azure-native:azurearcdata:getPostgresInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPostgresInstanceArgs struct {
	// Name of Postgres Instance
	PostgresInstanceName string `pulumi:"postgresInstanceName"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Postgres Instance.
type LookupPostgresInstanceResult struct {
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// null
	Properties PostgresInstancePropertiesResponse `pulumi:"properties"`
	// Resource sku.
	Sku *PostgresInstanceSkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPostgresInstanceResult
func (val *LookupPostgresInstanceResult) Defaults() *LookupPostgresInstanceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupPostgresInstanceOutput(ctx *pulumi.Context, args LookupPostgresInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupPostgresInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPostgresInstanceResult, error) {
			args := v.(LookupPostgresInstanceArgs)
			r, err := LookupPostgresInstance(ctx, &args, opts...)
			var s LookupPostgresInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPostgresInstanceResultOutput)
}

type LookupPostgresInstanceOutputArgs struct {
	// Name of Postgres Instance
	PostgresInstanceName pulumi.StringInput `pulumi:"postgresInstanceName"`
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPostgresInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPostgresInstanceArgs)(nil)).Elem()
}

// A Postgres Instance.
type LookupPostgresInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupPostgresInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPostgresInstanceResult)(nil)).Elem()
}

func (o LookupPostgresInstanceResultOutput) ToLookupPostgresInstanceResultOutput() LookupPostgresInstanceResultOutput {
	return o
}

func (o LookupPostgresInstanceResultOutput) ToLookupPostgresInstanceResultOutputWithContext(ctx context.Context) LookupPostgresInstanceResultOutput {
	return o
}

// The extendedLocation of the resource.
func (o LookupPostgresInstanceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPostgresInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupPostgresInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPostgresInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// null
func (o LookupPostgresInstanceResultOutput) Properties() PostgresInstancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) PostgresInstancePropertiesResponse { return v.Properties }).(PostgresInstancePropertiesResponseOutput)
}

// Resource sku.
func (o LookupPostgresInstanceResultOutput) Sku() PostgresInstanceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) *PostgresInstanceSkuResponse { return v.Sku }).(PostgresInstanceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPostgresInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPostgresInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPostgresInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPostgresInstanceResultOutput{})
}
