// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A global reach connection resource
func LookupGlobalReachConnection(ctx *pulumi.Context, args *LookupGlobalReachConnectionArgs, opts ...pulumi.InvokeOption) (*LookupGlobalReachConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGlobalReachConnectionResult
	err := ctx.Invoke("azure-native:avs/v20220501:getGlobalReachConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalReachConnectionArgs struct {
	// Name of the global reach connection in the private cloud
	GlobalReachConnectionName string `pulumi:"globalReachConnectionName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A global reach connection resource
type LookupGlobalReachConnectionResult struct {
	// The network used for global reach carved out from the original network block provided for the private cloud
	AddressPrefix string `pulumi:"addressPrefix"`
	// Authorization key from the peer express route used for the global reach connection
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The connection status of the global reach connection
	CircuitConnectionStatus string `pulumi:"circuitConnectionStatus"`
	// The ID of the Private Cloud's ExpressRoute Circuit that is participating in the global reach connection
	ExpressRouteId *string `pulumi:"expressRouteId"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Identifier of the ExpressRoute Circuit to peer with in the global reach connection
	PeerExpressRouteCircuit *string `pulumi:"peerExpressRouteCircuit"`
	// The state of the  ExpressRoute Circuit Authorization provisioning
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupGlobalReachConnectionOutput(ctx *pulumi.Context, args LookupGlobalReachConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalReachConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalReachConnectionResult, error) {
			args := v.(LookupGlobalReachConnectionArgs)
			r, err := LookupGlobalReachConnection(ctx, &args, opts...)
			var s LookupGlobalReachConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalReachConnectionResultOutput)
}

type LookupGlobalReachConnectionOutputArgs struct {
	// Name of the global reach connection in the private cloud
	GlobalReachConnectionName pulumi.StringInput `pulumi:"globalReachConnectionName"`
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGlobalReachConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalReachConnectionArgs)(nil)).Elem()
}

// A global reach connection resource
type LookupGlobalReachConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalReachConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalReachConnectionResult)(nil)).Elem()
}

func (o LookupGlobalReachConnectionResultOutput) ToLookupGlobalReachConnectionResultOutput() LookupGlobalReachConnectionResultOutput {
	return o
}

func (o LookupGlobalReachConnectionResultOutput) ToLookupGlobalReachConnectionResultOutputWithContext(ctx context.Context) LookupGlobalReachConnectionResultOutput {
	return o
}

// The network used for global reach carved out from the original network block provided for the private cloud
func (o LookupGlobalReachConnectionResultOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

// Authorization key from the peer express route used for the global reach connection
func (o LookupGlobalReachConnectionResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

// The connection status of the global reach connection
func (o LookupGlobalReachConnectionResultOutput) CircuitConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.CircuitConnectionStatus }).(pulumi.StringOutput)
}

// The ID of the Private Cloud's ExpressRoute Circuit that is participating in the global reach connection
func (o LookupGlobalReachConnectionResultOutput) ExpressRouteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) *string { return v.ExpressRouteId }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupGlobalReachConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupGlobalReachConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Identifier of the ExpressRoute Circuit to peer with in the global reach connection
func (o LookupGlobalReachConnectionResultOutput) PeerExpressRouteCircuit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) *string { return v.PeerExpressRouteCircuit }).(pulumi.StringPtrOutput)
}

// The state of the  ExpressRoute Circuit Authorization provisioning
func (o LookupGlobalReachConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupGlobalReachConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalReachConnectionResultOutput{})
}
