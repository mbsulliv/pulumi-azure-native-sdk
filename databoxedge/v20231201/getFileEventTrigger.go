// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a specific trigger by name.
func LookupFileEventTrigger(ctx *pulumi.Context, args *LookupFileEventTriggerArgs, opts ...pulumi.InvokeOption) (*LookupFileEventTriggerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFileEventTriggerResult
	err := ctx.Invoke("azure-native:databoxedge/v20231201:getFileEventTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileEventTriggerArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The trigger name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Trigger details.
type LookupFileEventTriggerResult struct {
	// A custom context tag typically used to correlate the trigger against its usage. For example, if a periodic timer trigger is intended for certain specific IoT modules in the device, the tag can be the name or the image URL of the module.
	CustomContextTag *string `pulumi:"customContextTag"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// Trigger Kind.
	// Expected value is 'FileEvent'.
	Kind string `pulumi:"kind"`
	// The object name.
	Name string `pulumi:"name"`
	// Role sink info.
	SinkInfo RoleSinkInfoResponse `pulumi:"sinkInfo"`
	// File event source details.
	SourceInfo FileSourceInfoResponse `pulumi:"sourceInfo"`
	// Metadata pertaining to creation and last modification of Trigger
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}

func LookupFileEventTriggerOutput(ctx *pulumi.Context, args LookupFileEventTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupFileEventTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileEventTriggerResult, error) {
			args := v.(LookupFileEventTriggerArgs)
			r, err := LookupFileEventTrigger(ctx, &args, opts...)
			var s LookupFileEventTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileEventTriggerResultOutput)
}

type LookupFileEventTriggerOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The trigger name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFileEventTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileEventTriggerArgs)(nil)).Elem()
}

// Trigger details.
type LookupFileEventTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupFileEventTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileEventTriggerResult)(nil)).Elem()
}

func (o LookupFileEventTriggerResultOutput) ToLookupFileEventTriggerResultOutput() LookupFileEventTriggerResultOutput {
	return o
}

func (o LookupFileEventTriggerResultOutput) ToLookupFileEventTriggerResultOutputWithContext(ctx context.Context) LookupFileEventTriggerResultOutput {
	return o
}

// A custom context tag typically used to correlate the trigger against its usage. For example, if a periodic timer trigger is intended for certain specific IoT modules in the device, the tag can be the name or the image URL of the module.
func (o LookupFileEventTriggerResultOutput) CustomContextTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) *string { return v.CustomContextTag }).(pulumi.StringPtrOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupFileEventTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Trigger Kind.
// Expected value is 'FileEvent'.
func (o LookupFileEventTriggerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The object name.
func (o LookupFileEventTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Role sink info.
func (o LookupFileEventTriggerResultOutput) SinkInfo() RoleSinkInfoResponseOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) RoleSinkInfoResponse { return v.SinkInfo }).(RoleSinkInfoResponseOutput)
}

// File event source details.
func (o LookupFileEventTriggerResultOutput) SourceInfo() FileSourceInfoResponseOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) FileSourceInfoResponse { return v.SourceInfo }).(FileSourceInfoResponseOutput)
}

// Metadata pertaining to creation and last modification of Trigger
func (o LookupFileEventTriggerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupFileEventTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileEventTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileEventTriggerResultOutput{})
}
