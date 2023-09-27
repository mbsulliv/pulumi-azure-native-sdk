// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ExpressRoute Circuit Authorization
func LookupAuthorization(ctx *pulumi.Context, args *LookupAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizationResult
	err := ctx.Invoke("azure-native:avs/v20230301:getAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationArgs struct {
	// Name of the ExpressRoute Circuit Authorization in the private cloud
	AuthorizationName string `pulumi:"authorizationName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ExpressRoute Circuit Authorization
type LookupAuthorizationResult struct {
	// The ID of the ExpressRoute Circuit Authorization
	ExpressRouteAuthorizationId string `pulumi:"expressRouteAuthorizationId"`
	// The key of the ExpressRoute Circuit Authorization
	ExpressRouteAuthorizationKey string `pulumi:"expressRouteAuthorizationKey"`
	// The ID of the ExpressRoute Circuit
	ExpressRouteId *string `pulumi:"expressRouteId"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The state of the  ExpressRoute Circuit Authorization provisioning
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupAuthorizationOutput(ctx *pulumi.Context, args LookupAuthorizationOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationResult, error) {
			args := v.(LookupAuthorizationArgs)
			r, err := LookupAuthorization(ctx, &args, opts...)
			var s LookupAuthorizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationResultOutput)
}

type LookupAuthorizationOutputArgs struct {
	// Name of the ExpressRoute Circuit Authorization in the private cloud
	AuthorizationName pulumi.StringInput `pulumi:"authorizationName"`
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAuthorizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationArgs)(nil)).Elem()
}

// ExpressRoute Circuit Authorization
type LookupAuthorizationResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationResult)(nil)).Elem()
}

func (o LookupAuthorizationResultOutput) ToLookupAuthorizationResultOutput() LookupAuthorizationResultOutput {
	return o
}

func (o LookupAuthorizationResultOutput) ToLookupAuthorizationResultOutputWithContext(ctx context.Context) LookupAuthorizationResultOutput {
	return o
}

func (o LookupAuthorizationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAuthorizationResult] {
	return pulumix.Output[LookupAuthorizationResult]{
		OutputState: o.OutputState,
	}
}

// The ID of the ExpressRoute Circuit Authorization
func (o LookupAuthorizationResultOutput) ExpressRouteAuthorizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.ExpressRouteAuthorizationId }).(pulumi.StringOutput)
}

// The key of the ExpressRoute Circuit Authorization
func (o LookupAuthorizationResultOutput) ExpressRouteAuthorizationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.ExpressRouteAuthorizationKey }).(pulumi.StringOutput)
}

// The ID of the ExpressRoute Circuit
func (o LookupAuthorizationResultOutput) ExpressRouteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) *string { return v.ExpressRouteId }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupAuthorizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupAuthorizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the  ExpressRoute Circuit Authorization provisioning
func (o LookupAuthorizationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupAuthorizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationResultOutput{})
}
