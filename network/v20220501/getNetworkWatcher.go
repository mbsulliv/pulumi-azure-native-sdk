// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network watcher in a resource group.
func LookupNetworkWatcher(ctx *pulumi.Context, args *LookupNetworkWatcherArgs, opts ...pulumi.InvokeOption) (*LookupNetworkWatcherResult, error) {
	var rv LookupNetworkWatcherResult
	err := ctx.Invoke("azure-native:network/v20220501:getNetworkWatcher", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkWatcherArgs struct {
	// The name of the network watcher.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Network watcher in a resource group.
type LookupNetworkWatcherResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the network watcher resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of running operation IDs.
	RunningOperationIds []int `pulumi:"runningOperationIds"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupNetworkWatcherOutput(ctx *pulumi.Context, args LookupNetworkWatcherOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkWatcherResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkWatcherResult, error) {
			args := v.(LookupNetworkWatcherArgs)
			r, err := LookupNetworkWatcher(ctx, &args, opts...)
			var s LookupNetworkWatcherResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkWatcherResultOutput)
}

type LookupNetworkWatcherOutputArgs struct {
	// The name of the network watcher.
	NetworkWatcherName pulumi.StringInput `pulumi:"networkWatcherName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkWatcherOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkWatcherArgs)(nil)).Elem()
}

// Network watcher in a resource group.
type LookupNetworkWatcherResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkWatcherResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkWatcherResult)(nil)).Elem()
}

func (o LookupNetworkWatcherResultOutput) ToLookupNetworkWatcherResultOutput() LookupNetworkWatcherResultOutput {
	return o
}

func (o LookupNetworkWatcherResultOutput) ToLookupNetworkWatcherResultOutputWithContext(ctx context.Context) LookupNetworkWatcherResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkWatcherResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNetworkWatcherResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupNetworkWatcherResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNetworkWatcherResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the network watcher resource.
func (o LookupNetworkWatcherResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of running operation IDs.
func (o LookupNetworkWatcherResultOutput) RunningOperationIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) []int { return v.RunningOperationIds }).(pulumi.IntArrayOutput)
}

// Resource tags.
func (o LookupNetworkWatcherResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNetworkWatcherResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkWatcherResultOutput{})
}
