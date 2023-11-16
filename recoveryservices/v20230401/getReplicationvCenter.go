// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a registered vCenter server(Add vCenter server).
func LookupReplicationvCenter(ctx *pulumi.Context, args *LookupReplicationvCenterArgs, opts ...pulumi.InvokeOption) (*LookupReplicationvCenterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationvCenterResult
	err := ctx.Invoke("azure-native:recoveryservices/v20230401:getReplicationvCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationvCenterArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
	// vcenter name.
	VcenterName string `pulumi:"vcenterName"`
}

// vCenter definition.
type LookupReplicationvCenterResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// VCenter related data.
	Properties VCenterPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationvCenterOutput(ctx *pulumi.Context, args LookupReplicationvCenterOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationvCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationvCenterResult, error) {
			args := v.(LookupReplicationvCenterArgs)
			r, err := LookupReplicationvCenter(ctx, &args, opts...)
			var s LookupReplicationvCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationvCenterResultOutput)
}

type LookupReplicationvCenterOutputArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// vcenter name.
	VcenterName pulumi.StringInput `pulumi:"vcenterName"`
}

func (LookupReplicationvCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationvCenterArgs)(nil)).Elem()
}

// vCenter definition.
type LookupReplicationvCenterResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationvCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationvCenterResult)(nil)).Elem()
}

func (o LookupReplicationvCenterResultOutput) ToLookupReplicationvCenterResultOutput() LookupReplicationvCenterResultOutput {
	return o
}

func (o LookupReplicationvCenterResultOutput) ToLookupReplicationvCenterResultOutputWithContext(ctx context.Context) LookupReplicationvCenterResultOutput {
	return o
}

// Resource Id
func (o LookupReplicationvCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationvCenterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationvCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

// VCenter related data.
func (o LookupReplicationvCenterResultOutput) Properties() VCenterPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) VCenterPropertiesResponse { return v.Properties }).(VCenterPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationvCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationvCenterResultOutput{})
}
