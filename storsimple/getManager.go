// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the properties of the specified manager name.
// Azure REST API version: 2017-06-01.
func LookupManager(ctx *pulumi.Context, args *LookupManagerArgs, opts ...pulumi.InvokeOption) (*LookupManagerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagerResult
	err := ctx.Invoke("azure-native:storsimple:getManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The StorSimple Manager.
type LookupManagerResult struct {
	// Represents the type of StorSimple Manager.
	CisIntrinsicSettings *ManagerIntrinsicSettingsResponse `pulumi:"cisIntrinsicSettings"`
	// The etag of the manager.
	Etag *string `pulumi:"etag"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The geo location of the resource.
	Location string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Specifies the Sku.
	Sku *ManagerSkuResponse `pulumi:"sku"`
	// The tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupManagerOutput(ctx *pulumi.Context, args LookupManagerOutputArgs, opts ...pulumi.InvokeOption) LookupManagerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagerResult, error) {
			args := v.(LookupManagerArgs)
			r, err := LookupManager(ctx, &args, opts...)
			var s LookupManagerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagerResultOutput)
}

type LookupManagerOutputArgs struct {
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagerArgs)(nil)).Elem()
}

// The StorSimple Manager.
type LookupManagerResultOutput struct{ *pulumi.OutputState }

func (LookupManagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagerResult)(nil)).Elem()
}

func (o LookupManagerResultOutput) ToLookupManagerResultOutput() LookupManagerResultOutput {
	return o
}

func (o LookupManagerResultOutput) ToLookupManagerResultOutputWithContext(ctx context.Context) LookupManagerResultOutput {
	return o
}

func (o LookupManagerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagerResult] {
	return pulumix.Output[LookupManagerResult]{
		OutputState: o.OutputState,
	}
}

// Represents the type of StorSimple Manager.
func (o LookupManagerResultOutput) CisIntrinsicSettings() ManagerIntrinsicSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupManagerResult) *ManagerIntrinsicSettingsResponse { return v.CisIntrinsicSettings }).(ManagerIntrinsicSettingsResponsePtrOutput)
}

// The etag of the manager.
func (o LookupManagerResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource ID.
func (o LookupManagerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo location of the resource.
func (o LookupManagerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupManagerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
func (o LookupManagerResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Specifies the Sku.
func (o LookupManagerResultOutput) Sku() ManagerSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagerResult) *ManagerSkuResponse { return v.Sku }).(ManagerSkuResponsePtrOutput)
}

// The tags attached to the resource.
func (o LookupManagerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupManagerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagerResultOutput{})
}
