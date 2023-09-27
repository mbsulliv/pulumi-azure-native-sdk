// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified authorization from the specified express route circuit.
// Azure REST API version: 2023-02-01.
func LookupExpressRouteCircuitAuthorization(ctx *pulumi.Context, args *LookupExpressRouteCircuitAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitAuthorizationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExpressRouteCircuitAuthorizationResult
	err := ctx.Invoke("azure-native:network:getExpressRouteCircuitAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitAuthorizationArgs struct {
	// The name of the authorization.
	AuthorizationName string `pulumi:"authorizationName"`
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Authorization in an ExpressRouteCircuit resource.
type LookupExpressRouteCircuitAuthorizationResult struct {
	// The authorization key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The authorization use status.
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the authorization resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupExpressRouteCircuitAuthorizationOutput(ctx *pulumi.Context, args LookupExpressRouteCircuitAuthorizationOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCircuitAuthorizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCircuitAuthorizationResult, error) {
			args := v.(LookupExpressRouteCircuitAuthorizationArgs)
			r, err := LookupExpressRouteCircuitAuthorization(ctx, &args, opts...)
			var s LookupExpressRouteCircuitAuthorizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCircuitAuthorizationResultOutput)
}

type LookupExpressRouteCircuitAuthorizationOutputArgs struct {
	// The name of the authorization.
	AuthorizationName pulumi.StringInput `pulumi:"authorizationName"`
	// The name of the express route circuit.
	CircuitName pulumi.StringInput `pulumi:"circuitName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCircuitAuthorizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitAuthorizationArgs)(nil)).Elem()
}

// Authorization in an ExpressRouteCircuit resource.
type LookupExpressRouteCircuitAuthorizationResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCircuitAuthorizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitAuthorizationResult)(nil)).Elem()
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) ToLookupExpressRouteCircuitAuthorizationResultOutput() LookupExpressRouteCircuitAuthorizationResultOutput {
	return o
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) ToLookupExpressRouteCircuitAuthorizationResultOutputWithContext(ctx context.Context) LookupExpressRouteCircuitAuthorizationResultOutput {
	return o
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupExpressRouteCircuitAuthorizationResult] {
	return pulumix.Output[LookupExpressRouteCircuitAuthorizationResult]{
		OutputState: o.OutputState,
	}
}

// The authorization key.
func (o LookupExpressRouteCircuitAuthorizationResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

// The authorization use status.
func (o LookupExpressRouteCircuitAuthorizationResultOutput) AuthorizationUseStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.AuthorizationUseStatus }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupExpressRouteCircuitAuthorizationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupExpressRouteCircuitAuthorizationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupExpressRouteCircuitAuthorizationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the authorization resource.
func (o LookupExpressRouteCircuitAuthorizationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupExpressRouteCircuitAuthorizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCircuitAuthorizationResultOutput{})
}
