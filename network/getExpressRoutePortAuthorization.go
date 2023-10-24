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

// Gets the specified authorization from the specified express route port.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01.
func LookupExpressRoutePortAuthorization(ctx *pulumi.Context, args *LookupExpressRoutePortAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupExpressRoutePortAuthorizationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExpressRoutePortAuthorizationResult
	err := ctx.Invoke("azure-native:network:getExpressRoutePortAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRoutePortAuthorizationArgs struct {
	// The name of the authorization.
	AuthorizationName string `pulumi:"authorizationName"`
	// The name of the express route port.
	ExpressRoutePortName string `pulumi:"expressRoutePortName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ExpressRoutePort Authorization resource definition.
type LookupExpressRoutePortAuthorizationResult struct {
	// The authorization key.
	AuthorizationKey string `pulumi:"authorizationKey"`
	// The authorization use status.
	AuthorizationUseStatus string `pulumi:"authorizationUseStatus"`
	// The reference to the ExpressRoute circuit resource using the authorization.
	CircuitResourceUri string `pulumi:"circuitResourceUri"`
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

func LookupExpressRoutePortAuthorizationOutput(ctx *pulumi.Context, args LookupExpressRoutePortAuthorizationOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRoutePortAuthorizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRoutePortAuthorizationResult, error) {
			args := v.(LookupExpressRoutePortAuthorizationArgs)
			r, err := LookupExpressRoutePortAuthorization(ctx, &args, opts...)
			var s LookupExpressRoutePortAuthorizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRoutePortAuthorizationResultOutput)
}

type LookupExpressRoutePortAuthorizationOutputArgs struct {
	// The name of the authorization.
	AuthorizationName pulumi.StringInput `pulumi:"authorizationName"`
	// The name of the express route port.
	ExpressRoutePortName pulumi.StringInput `pulumi:"expressRoutePortName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRoutePortAuthorizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRoutePortAuthorizationArgs)(nil)).Elem()
}

// ExpressRoutePort Authorization resource definition.
type LookupExpressRoutePortAuthorizationResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRoutePortAuthorizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRoutePortAuthorizationResult)(nil)).Elem()
}

func (o LookupExpressRoutePortAuthorizationResultOutput) ToLookupExpressRoutePortAuthorizationResultOutput() LookupExpressRoutePortAuthorizationResultOutput {
	return o
}

func (o LookupExpressRoutePortAuthorizationResultOutput) ToLookupExpressRoutePortAuthorizationResultOutputWithContext(ctx context.Context) LookupExpressRoutePortAuthorizationResultOutput {
	return o
}

func (o LookupExpressRoutePortAuthorizationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupExpressRoutePortAuthorizationResult] {
	return pulumix.Output[LookupExpressRoutePortAuthorizationResult]{
		OutputState: o.OutputState,
	}
}

// The authorization key.
func (o LookupExpressRoutePortAuthorizationResultOutput) AuthorizationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.AuthorizationKey }).(pulumi.StringOutput)
}

// The authorization use status.
func (o LookupExpressRoutePortAuthorizationResultOutput) AuthorizationUseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.AuthorizationUseStatus }).(pulumi.StringOutput)
}

// The reference to the ExpressRoute circuit resource using the authorization.
func (o LookupExpressRoutePortAuthorizationResultOutput) CircuitResourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.CircuitResourceUri }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupExpressRoutePortAuthorizationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupExpressRoutePortAuthorizationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupExpressRoutePortAuthorizationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the authorization resource.
func (o LookupExpressRoutePortAuthorizationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupExpressRoutePortAuthorizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRoutePortAuthorizationResultOutput{})
}
