// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsan

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get an VolumeGroups.
// Azure REST API version: 2021-11-20-preview.
//
// Other available API versions: 2022-12-01-preview, 2023-01-01.
func LookupVolumeGroup(ctx *pulumi.Context, args *LookupVolumeGroupArgs, opts ...pulumi.InvokeOption) (*LookupVolumeGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVolumeGroupResult
	err := ctx.Invoke("azure-native:elasticsan:getVolumeGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeGroupArgs struct {
	// The name of the ElasticSan.
	ElasticSanName string `pulumi:"elasticSanName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VolumeGroup.
	VolumeGroupName string `pulumi:"volumeGroupName"`
}

// Response for Volume Group request.
type LookupVolumeGroupResult struct {
	// Type of encryption
	Encryption *string `pulumi:"encryption"`
	// Azure resource identifier.
	Id string `pulumi:"id"`
	// Azure resource name.
	Name string `pulumi:"name"`
	// A collection of rules governing the accessibility from specific network locations.
	NetworkAcls *NetworkRuleSetResponse `pulumi:"networkAcls"`
	// Type of storage target
	ProtocolType *string `pulumi:"protocolType"`
	// State of the operation on the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource metadata required by ARM RPC
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type string `pulumi:"type"`
}

func LookupVolumeGroupOutput(ctx *pulumi.Context, args LookupVolumeGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeGroupResult, error) {
			args := v.(LookupVolumeGroupArgs)
			r, err := LookupVolumeGroup(ctx, &args, opts...)
			var s LookupVolumeGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeGroupResultOutput)
}

type LookupVolumeGroupOutputArgs struct {
	// The name of the ElasticSan.
	ElasticSanName pulumi.StringInput `pulumi:"elasticSanName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VolumeGroup.
	VolumeGroupName pulumi.StringInput `pulumi:"volumeGroupName"`
}

func (LookupVolumeGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeGroupArgs)(nil)).Elem()
}

// Response for Volume Group request.
type LookupVolumeGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeGroupResult)(nil)).Elem()
}

func (o LookupVolumeGroupResultOutput) ToLookupVolumeGroupResultOutput() LookupVolumeGroupResultOutput {
	return o
}

func (o LookupVolumeGroupResultOutput) ToLookupVolumeGroupResultOutputWithContext(ctx context.Context) LookupVolumeGroupResultOutput {
	return o
}

func (o LookupVolumeGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVolumeGroupResult] {
	return pulumix.Output[LookupVolumeGroupResult]{
		OutputState: o.OutputState,
	}
}

// Type of encryption
func (o LookupVolumeGroupResultOutput) Encryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *string { return v.Encryption }).(pulumi.StringPtrOutput)
}

// Azure resource identifier.
func (o LookupVolumeGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure resource name.
func (o LookupVolumeGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of rules governing the accessibility from specific network locations.
func (o LookupVolumeGroupResultOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *NetworkRuleSetResponse { return v.NetworkAcls }).(NetworkRuleSetResponsePtrOutput)
}

// Type of storage target
func (o LookupVolumeGroupResultOutput) ProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *string { return v.ProtocolType }).(pulumi.StringPtrOutput)
}

// State of the operation on the resource.
func (o LookupVolumeGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource metadata required by ARM RPC
func (o LookupVolumeGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o LookupVolumeGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o LookupVolumeGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeGroupResultOutput{})
}
