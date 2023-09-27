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

// Get a specified connection created by this subscription.
func LookupSubscriptionNetworkManagerConnection(ctx *pulumi.Context, args *LookupSubscriptionNetworkManagerConnectionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionNetworkManagerConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSubscriptionNetworkManagerConnectionResult
	err := ctx.Invoke("azure-native:network/v20230501:getSubscriptionNetworkManagerConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionNetworkManagerConnectionArgs struct {
	// Name for the network manager connection.
	NetworkManagerConnectionName string `pulumi:"networkManagerConnectionName"`
}

// The Network Manager Connection resource
type LookupSubscriptionNetworkManagerConnectionResult struct {
	// A description of the network manager connection.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Network Manager Id.
	NetworkManagerId *string `pulumi:"networkManagerId"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupSubscriptionNetworkManagerConnectionOutput(ctx *pulumi.Context, args LookupSubscriptionNetworkManagerConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionNetworkManagerConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionNetworkManagerConnectionResult, error) {
			args := v.(LookupSubscriptionNetworkManagerConnectionArgs)
			r, err := LookupSubscriptionNetworkManagerConnection(ctx, &args, opts...)
			var s LookupSubscriptionNetworkManagerConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionNetworkManagerConnectionResultOutput)
}

type LookupSubscriptionNetworkManagerConnectionOutputArgs struct {
	// Name for the network manager connection.
	NetworkManagerConnectionName pulumi.StringInput `pulumi:"networkManagerConnectionName"`
}

func (LookupSubscriptionNetworkManagerConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionNetworkManagerConnectionArgs)(nil)).Elem()
}

// The Network Manager Connection resource
type LookupSubscriptionNetworkManagerConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionNetworkManagerConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionNetworkManagerConnectionResult)(nil)).Elem()
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) ToLookupSubscriptionNetworkManagerConnectionResultOutput() LookupSubscriptionNetworkManagerConnectionResultOutput {
	return o
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) ToLookupSubscriptionNetworkManagerConnectionResultOutputWithContext(ctx context.Context) LookupSubscriptionNetworkManagerConnectionResultOutput {
	return o
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSubscriptionNetworkManagerConnectionResult] {
	return pulumix.Output[LookupSubscriptionNetworkManagerConnectionResult]{
		OutputState: o.OutputState,
	}
}

// A description of the network manager connection.
func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network Manager Id.
func (o LookupSubscriptionNetworkManagerConnectionResultOutput) NetworkManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) *string { return v.NetworkManagerId }).(pulumi.StringPtrOutput)
}

// The system metadata related to this resource.
func (o LookupSubscriptionNetworkManagerConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionNetworkManagerConnectionResultOutput{})
}
