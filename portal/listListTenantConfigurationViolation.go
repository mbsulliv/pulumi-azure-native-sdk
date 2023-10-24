// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package portal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets list of items that violate tenant's configuration.
// Azure REST API version: 2020-09-01-preview.
func ListListTenantConfigurationViolation(ctx *pulumi.Context, args *ListListTenantConfigurationViolationArgs, opts ...pulumi.InvokeOption) (*ListListTenantConfigurationViolationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListListTenantConfigurationViolationResult
	err := ctx.Invoke("azure-native:portal:listListTenantConfigurationViolation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListTenantConfigurationViolationArgs struct {
}

// List of list of items that violate tenant's configuration.
type ListListTenantConfigurationViolationResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// The array of violations.
	Value []ViolationResponse `pulumi:"value"`
}

func ListListTenantConfigurationViolationOutput(ctx *pulumi.Context, args ListListTenantConfigurationViolationOutputArgs, opts ...pulumi.InvokeOption) ListListTenantConfigurationViolationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListListTenantConfigurationViolationResult, error) {
			args := v.(ListListTenantConfigurationViolationArgs)
			r, err := ListListTenantConfigurationViolation(ctx, &args, opts...)
			var s ListListTenantConfigurationViolationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListListTenantConfigurationViolationResultOutput)
}

type ListListTenantConfigurationViolationOutputArgs struct {
}

func (ListListTenantConfigurationViolationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListTenantConfigurationViolationArgs)(nil)).Elem()
}

// List of list of items that violate tenant's configuration.
type ListListTenantConfigurationViolationResultOutput struct{ *pulumi.OutputState }

func (ListListTenantConfigurationViolationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListTenantConfigurationViolationResult)(nil)).Elem()
}

func (o ListListTenantConfigurationViolationResultOutput) ToListListTenantConfigurationViolationResultOutput() ListListTenantConfigurationViolationResultOutput {
	return o
}

func (o ListListTenantConfigurationViolationResultOutput) ToListListTenantConfigurationViolationResultOutputWithContext(ctx context.Context) ListListTenantConfigurationViolationResultOutput {
	return o
}

func (o ListListTenantConfigurationViolationResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListListTenantConfigurationViolationResult] {
	return pulumix.Output[ListListTenantConfigurationViolationResult]{
		OutputState: o.OutputState,
	}
}

// The URL to use for getting the next set of results.
func (o ListListTenantConfigurationViolationResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListListTenantConfigurationViolationResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// The array of violations.
func (o ListListTenantConfigurationViolationResultOutput) Value() ViolationResponseArrayOutput {
	return o.ApplyT(func(v ListListTenantConfigurationViolationResult) []ViolationResponse { return v.Value }).(ViolationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListTenantConfigurationViolationResultOutput{})
}
