// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerbidedicated

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets details about the specified auto scale v-core.
// Azure REST API version: 2021-01-01.
func LookupAutoScaleVCore(ctx *pulumi.Context, args *LookupAutoScaleVCoreArgs, opts ...pulumi.InvokeOption) (*LookupAutoScaleVCoreResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAutoScaleVCoreResult
	err := ctx.Invoke("azure-native:powerbidedicated:getAutoScaleVCore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutoScaleVCoreArgs struct {
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the auto scale v-core. It must be a minimum of 3 characters, and a maximum of 63.
	VcoreName string `pulumi:"vcoreName"`
}

// Represents an instance of an auto scale v-core resource.
type LookupAutoScaleVCoreResult struct {
	// The maximum capacity of an auto scale v-core resource.
	CapacityLimit *int `pulumi:"capacityLimit"`
	// The object ID of the capacity resource associated with the auto scale v-core resource.
	CapacityObjectId *string `pulumi:"capacityObjectId"`
	// An identifier that represents the PowerBI Dedicated resource.
	Id string `pulumi:"id"`
	// Location of the PowerBI Dedicated resource.
	Location string `pulumi:"location"`
	// The name of the PowerBI Dedicated resource.
	Name string `pulumi:"name"`
	// The current deployment state of an auto scale v-core resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SKU of the auto scale v-core resource.
	Sku AutoScaleVCoreSkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
	// The type of the PowerBI Dedicated resource.
	Type string `pulumi:"type"`
}

func LookupAutoScaleVCoreOutput(ctx *pulumi.Context, args LookupAutoScaleVCoreOutputArgs, opts ...pulumi.InvokeOption) LookupAutoScaleVCoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoScaleVCoreResult, error) {
			args := v.(LookupAutoScaleVCoreArgs)
			r, err := LookupAutoScaleVCore(ctx, &args, opts...)
			var s LookupAutoScaleVCoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoScaleVCoreResultOutput)
}

type LookupAutoScaleVCoreOutputArgs struct {
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the auto scale v-core. It must be a minimum of 3 characters, and a maximum of 63.
	VcoreName pulumi.StringInput `pulumi:"vcoreName"`
}

func (LookupAutoScaleVCoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoScaleVCoreArgs)(nil)).Elem()
}

// Represents an instance of an auto scale v-core resource.
type LookupAutoScaleVCoreResultOutput struct{ *pulumi.OutputState }

func (LookupAutoScaleVCoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoScaleVCoreResult)(nil)).Elem()
}

func (o LookupAutoScaleVCoreResultOutput) ToLookupAutoScaleVCoreResultOutput() LookupAutoScaleVCoreResultOutput {
	return o
}

func (o LookupAutoScaleVCoreResultOutput) ToLookupAutoScaleVCoreResultOutputWithContext(ctx context.Context) LookupAutoScaleVCoreResultOutput {
	return o
}

func (o LookupAutoScaleVCoreResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAutoScaleVCoreResult] {
	return pulumix.Output[LookupAutoScaleVCoreResult]{
		OutputState: o.OutputState,
	}
}

// The maximum capacity of an auto scale v-core resource.
func (o LookupAutoScaleVCoreResultOutput) CapacityLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) *int { return v.CapacityLimit }).(pulumi.IntPtrOutput)
}

// The object ID of the capacity resource associated with the auto scale v-core resource.
func (o LookupAutoScaleVCoreResultOutput) CapacityObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) *string { return v.CapacityObjectId }).(pulumi.StringPtrOutput)
}

// An identifier that represents the PowerBI Dedicated resource.
func (o LookupAutoScaleVCoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the PowerBI Dedicated resource.
func (o LookupAutoScaleVCoreResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the PowerBI Dedicated resource.
func (o LookupAutoScaleVCoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current deployment state of an auto scale v-core resource. The provisioningState is to indicate states for resource provisioning.
func (o LookupAutoScaleVCoreResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the auto scale v-core resource.
func (o LookupAutoScaleVCoreResultOutput) Sku() AutoScaleVCoreSkuResponseOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) AutoScaleVCoreSkuResponse { return v.Sku }).(AutoScaleVCoreSkuResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupAutoScaleVCoreResultOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) *SystemDataResponse { return v.SystemData }).(SystemDataResponsePtrOutput)
}

// Key-value pairs of additional resource provisioning properties.
func (o LookupAutoScaleVCoreResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the PowerBI Dedicated resource.
func (o LookupAutoScaleVCoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoScaleVCoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoScaleVCoreResultOutput{})
}
