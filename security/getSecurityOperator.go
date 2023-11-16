// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a specific security operator for the requested scope.
// Azure REST API version: 2023-01-01-preview.
func LookupSecurityOperator(ctx *pulumi.Context, args *LookupSecurityOperatorArgs, opts ...pulumi.InvokeOption) (*LookupSecurityOperatorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityOperatorResult
	err := ctx.Invoke("azure-native:security:getSecurityOperator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityOperatorArgs struct {
	// name of the pricing configuration
	PricingName string `pulumi:"pricingName"`
	// name of the securityOperator
	SecurityOperatorName string `pulumi:"securityOperatorName"`
}

// Security operator under a given subscription and pricing
type LookupSecurityOperatorResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupSecurityOperatorOutput(ctx *pulumi.Context, args LookupSecurityOperatorOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityOperatorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityOperatorResult, error) {
			args := v.(LookupSecurityOperatorArgs)
			r, err := LookupSecurityOperator(ctx, &args, opts...)
			var s LookupSecurityOperatorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityOperatorResultOutput)
}

type LookupSecurityOperatorOutputArgs struct {
	// name of the pricing configuration
	PricingName pulumi.StringInput `pulumi:"pricingName"`
	// name of the securityOperator
	SecurityOperatorName pulumi.StringInput `pulumi:"securityOperatorName"`
}

func (LookupSecurityOperatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityOperatorArgs)(nil)).Elem()
}

// Security operator under a given subscription and pricing
type LookupSecurityOperatorResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityOperatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityOperatorResult)(nil)).Elem()
}

func (o LookupSecurityOperatorResultOutput) ToLookupSecurityOperatorResultOutput() LookupSecurityOperatorResultOutput {
	return o
}

func (o LookupSecurityOperatorResultOutput) ToLookupSecurityOperatorResultOutputWithContext(ctx context.Context) LookupSecurityOperatorResultOutput {
	return o
}

// Resource Id
func (o LookupSecurityOperatorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityOperatorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupSecurityOperatorResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSecurityOperatorResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Resource name
func (o LookupSecurityOperatorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityOperatorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o LookupSecurityOperatorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityOperatorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityOperatorResultOutput{})
}
