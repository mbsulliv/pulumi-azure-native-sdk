// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a change data capture.
// Azure REST API version: 2018-06-01.
func LookupChangeDataCapture(ctx *pulumi.Context, args *LookupChangeDataCaptureArgs, opts ...pulumi.InvokeOption) (*LookupChangeDataCaptureResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupChangeDataCaptureResult
	err := ctx.Invoke("azure-native:datafactory:getChangeDataCapture", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChangeDataCaptureArgs struct {
	// The change data capture name.
	ChangeDataCaptureName string `pulumi:"changeDataCaptureName"`
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Change data capture resource type.
type LookupChangeDataCaptureResult struct {
	// A boolean to determine if the vnet configuration needs to be overwritten.
	AllowVNetOverride *bool `pulumi:"allowVNetOverride"`
	// The description of the change data capture.
	Description *string `pulumi:"description"`
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The folder that this CDC is in. If not specified, CDC will appear at the root level.
	Folder *ChangeDataCaptureResponseFolder `pulumi:"folder"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// CDC policy
	Policy MapperPolicyResponse `pulumi:"policy"`
	// List of sources connections that can be used as sources in the CDC.
	SourceConnectionsInfo []MapperSourceConnectionsInfoResponse `pulumi:"sourceConnectionsInfo"`
	// Status of the CDC as to if it is running or stopped.
	Status *string `pulumi:"status"`
	// List of target connections that can be used as sources in the CDC.
	TargetConnectionsInfo []MapperTargetConnectionsInfoResponse `pulumi:"targetConnectionsInfo"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupChangeDataCaptureOutput(ctx *pulumi.Context, args LookupChangeDataCaptureOutputArgs, opts ...pulumi.InvokeOption) LookupChangeDataCaptureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChangeDataCaptureResult, error) {
			args := v.(LookupChangeDataCaptureArgs)
			r, err := LookupChangeDataCapture(ctx, &args, opts...)
			var s LookupChangeDataCaptureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChangeDataCaptureResultOutput)
}

type LookupChangeDataCaptureOutputArgs struct {
	// The change data capture name.
	ChangeDataCaptureName pulumi.StringInput `pulumi:"changeDataCaptureName"`
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupChangeDataCaptureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChangeDataCaptureArgs)(nil)).Elem()
}

// Change data capture resource type.
type LookupChangeDataCaptureResultOutput struct{ *pulumi.OutputState }

func (LookupChangeDataCaptureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChangeDataCaptureResult)(nil)).Elem()
}

func (o LookupChangeDataCaptureResultOutput) ToLookupChangeDataCaptureResultOutput() LookupChangeDataCaptureResultOutput {
	return o
}

func (o LookupChangeDataCaptureResultOutput) ToLookupChangeDataCaptureResultOutputWithContext(ctx context.Context) LookupChangeDataCaptureResultOutput {
	return o
}

// A boolean to determine if the vnet configuration needs to be overwritten.
func (o LookupChangeDataCaptureResultOutput) AllowVNetOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) *bool { return v.AllowVNetOverride }).(pulumi.BoolPtrOutput)
}

// The description of the change data capture.
func (o LookupChangeDataCaptureResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Etag identifies change in the resource.
func (o LookupChangeDataCaptureResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The folder that this CDC is in. If not specified, CDC will appear at the root level.
func (o LookupChangeDataCaptureResultOutput) Folder() ChangeDataCaptureResponseFolderPtrOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) *ChangeDataCaptureResponseFolder { return v.Folder }).(ChangeDataCaptureResponseFolderPtrOutput)
}

// The resource identifier.
func (o LookupChangeDataCaptureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupChangeDataCaptureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) string { return v.Name }).(pulumi.StringOutput)
}

// CDC policy
func (o LookupChangeDataCaptureResultOutput) Policy() MapperPolicyResponseOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) MapperPolicyResponse { return v.Policy }).(MapperPolicyResponseOutput)
}

// List of sources connections that can be used as sources in the CDC.
func (o LookupChangeDataCaptureResultOutput) SourceConnectionsInfo() MapperSourceConnectionsInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) []MapperSourceConnectionsInfoResponse {
		return v.SourceConnectionsInfo
	}).(MapperSourceConnectionsInfoResponseArrayOutput)
}

// Status of the CDC as to if it is running or stopped.
func (o LookupChangeDataCaptureResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// List of target connections that can be used as sources in the CDC.
func (o LookupChangeDataCaptureResultOutput) TargetConnectionsInfo() MapperTargetConnectionsInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) []MapperTargetConnectionsInfoResponse {
		return v.TargetConnectionsInfo
	}).(MapperTargetConnectionsInfoResponseArrayOutput)
}

// The resource type.
func (o LookupChangeDataCaptureResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChangeDataCaptureResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChangeDataCaptureResultOutput{})
}
