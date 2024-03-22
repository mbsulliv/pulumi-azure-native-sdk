// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of an Azure Site Recovery fabric.
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2023-06-01, 2023-08-01, 2024-01-01.
func LookupReplicationFabric(ctx *pulumi.Context, args *LookupReplicationFabricArgs, opts ...pulumi.InvokeOption) (*LookupReplicationFabricResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationFabricResult
	err := ctx.Invoke("azure-native:recoveryservices:getReplicationFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationFabricArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// OData filter options.
	Filter *string `pulumi:"filter"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Fabric definition.
type LookupReplicationFabricResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Fabric related data.
	Properties FabricPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationFabricOutput(ctx *pulumi.Context, args LookupReplicationFabricOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationFabricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationFabricResult, error) {
			args := v.(LookupReplicationFabricArgs)
			r, err := LookupReplicationFabric(ctx, &args, opts...)
			var s LookupReplicationFabricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationFabricResultOutput)
}

type LookupReplicationFabricOutputArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// OData filter options.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationFabricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationFabricArgs)(nil)).Elem()
}

// Fabric definition.
type LookupReplicationFabricResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationFabricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationFabricResult)(nil)).Elem()
}

func (o LookupReplicationFabricResultOutput) ToLookupReplicationFabricResultOutput() LookupReplicationFabricResultOutput {
	return o
}

func (o LookupReplicationFabricResultOutput) ToLookupReplicationFabricResultOutputWithContext(ctx context.Context) LookupReplicationFabricResultOutput {
	return o
}

// Resource Id
func (o LookupReplicationFabricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationFabricResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationFabricResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) string { return v.Name }).(pulumi.StringOutput)
}

// Fabric related data.
func (o LookupReplicationFabricResultOutput) Properties() FabricPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) FabricPropertiesResponse { return v.Properties }).(FabricPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationFabricResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationFabricResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationFabricResultOutput{})
}
