// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a packet capture session by name.
func LookupPacketCapture(ctx *pulumi.Context, args *LookupPacketCaptureArgs, opts ...pulumi.InvokeOption) (*LookupPacketCaptureResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPacketCaptureResult
	err := ctx.Invoke("azure-native:network/v20230501:getPacketCapture", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPacketCaptureArgs struct {
	// The name of the network watcher.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the packet capture session.
	PacketCaptureName string `pulumi:"packetCaptureName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about packet capture session.
type LookupPacketCaptureResult struct {
	// Number of bytes captured per packet, the remaining bytes are truncated.
	BytesToCapturePerPacket *float64 `pulumi:"bytesToCapturePerPacket"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// A list of packet capture filters.
	Filters []PacketCaptureFilterResponse `pulumi:"filters"`
	// ID of the packet capture operation.
	Id string `pulumi:"id"`
	// Name of the packet capture session.
	Name string `pulumi:"name"`
	// The provisioning state of the packet capture session.
	ProvisioningState string `pulumi:"provisioningState"`
	// A list of AzureVMSS instances which can be included or excluded to run packet capture. If both included and excluded are empty, then the packet capture will run on all instances of AzureVMSS.
	Scope *PacketCaptureMachineScopeResponse `pulumi:"scope"`
	// The storage location for a packet capture session.
	StorageLocation PacketCaptureStorageLocationResponse `pulumi:"storageLocation"`
	// The ID of the targeted resource, only AzureVM and AzureVMSS as target type are currently supported.
	Target string `pulumi:"target"`
	// Target type of the resource provided.
	TargetType *string `pulumi:"targetType"`
	// Maximum duration of the capture session in seconds.
	TimeLimitInSeconds *int `pulumi:"timeLimitInSeconds"`
	// Maximum size of the capture output.
	TotalBytesPerSession *float64 `pulumi:"totalBytesPerSession"`
}

// Defaults sets the appropriate defaults for LookupPacketCaptureResult
func (val *LookupPacketCaptureResult) Defaults() *LookupPacketCaptureResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.BytesToCapturePerPacket == nil {
		bytesToCapturePerPacket_ := 0.0
		tmp.BytesToCapturePerPacket = &bytesToCapturePerPacket_
	}
	if tmp.TimeLimitInSeconds == nil {
		timeLimitInSeconds_ := 18000
		tmp.TimeLimitInSeconds = &timeLimitInSeconds_
	}
	if tmp.TotalBytesPerSession == nil {
		totalBytesPerSession_ := 1073741824.0
		tmp.TotalBytesPerSession = &totalBytesPerSession_
	}
	return &tmp
}

func LookupPacketCaptureOutput(ctx *pulumi.Context, args LookupPacketCaptureOutputArgs, opts ...pulumi.InvokeOption) LookupPacketCaptureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPacketCaptureResult, error) {
			args := v.(LookupPacketCaptureArgs)
			r, err := LookupPacketCapture(ctx, &args, opts...)
			var s LookupPacketCaptureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPacketCaptureResultOutput)
}

type LookupPacketCaptureOutputArgs struct {
	// The name of the network watcher.
	NetworkWatcherName pulumi.StringInput `pulumi:"networkWatcherName"`
	// The name of the packet capture session.
	PacketCaptureName pulumi.StringInput `pulumi:"packetCaptureName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPacketCaptureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCaptureArgs)(nil)).Elem()
}

// Information about packet capture session.
type LookupPacketCaptureResultOutput struct{ *pulumi.OutputState }

func (LookupPacketCaptureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCaptureResult)(nil)).Elem()
}

func (o LookupPacketCaptureResultOutput) ToLookupPacketCaptureResultOutput() LookupPacketCaptureResultOutput {
	return o
}

func (o LookupPacketCaptureResultOutput) ToLookupPacketCaptureResultOutputWithContext(ctx context.Context) LookupPacketCaptureResultOutput {
	return o
}

func (o LookupPacketCaptureResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPacketCaptureResult] {
	return pulumix.Output[LookupPacketCaptureResult]{
		OutputState: o.OutputState,
	}
}

// Number of bytes captured per packet, the remaining bytes are truncated.
func (o LookupPacketCaptureResultOutput) BytesToCapturePerPacket() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) *float64 { return v.BytesToCapturePerPacket }).(pulumi.Float64PtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupPacketCaptureResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.Etag }).(pulumi.StringOutput)
}

// A list of packet capture filters.
func (o LookupPacketCaptureResultOutput) Filters() PacketCaptureFilterResponseArrayOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) []PacketCaptureFilterResponse { return v.Filters }).(PacketCaptureFilterResponseArrayOutput)
}

// ID of the packet capture operation.
func (o LookupPacketCaptureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the packet capture session.
func (o LookupPacketCaptureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the packet capture session.
func (o LookupPacketCaptureResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// A list of AzureVMSS instances which can be included or excluded to run packet capture. If both included and excluded are empty, then the packet capture will run on all instances of AzureVMSS.
func (o LookupPacketCaptureResultOutput) Scope() PacketCaptureMachineScopeResponsePtrOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) *PacketCaptureMachineScopeResponse { return v.Scope }).(PacketCaptureMachineScopeResponsePtrOutput)
}

// The storage location for a packet capture session.
func (o LookupPacketCaptureResultOutput) StorageLocation() PacketCaptureStorageLocationResponseOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) PacketCaptureStorageLocationResponse { return v.StorageLocation }).(PacketCaptureStorageLocationResponseOutput)
}

// The ID of the targeted resource, only AzureVM and AzureVMSS as target type are currently supported.
func (o LookupPacketCaptureResultOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) string { return v.Target }).(pulumi.StringOutput)
}

// Target type of the resource provided.
func (o LookupPacketCaptureResultOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) *string { return v.TargetType }).(pulumi.StringPtrOutput)
}

// Maximum duration of the capture session in seconds.
func (o LookupPacketCaptureResultOutput) TimeLimitInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) *int { return v.TimeLimitInSeconds }).(pulumi.IntPtrOutput)
}

// Maximum size of the capture output.
func (o LookupPacketCaptureResultOutput) TotalBytesPerSession() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPacketCaptureResult) *float64 { return v.TotalBytesPerSession }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPacketCaptureResultOutput{})
}
