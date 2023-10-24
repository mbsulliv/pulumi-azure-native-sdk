// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package policyinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets all deployments for a remediation at subscription scope.
// Azure REST API version: 2021-10-01.
//
// Other available API versions: 2018-07-01-preview.
func ListRemediationDeploymentsAtSubscription(ctx *pulumi.Context, args *ListRemediationDeploymentsAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListRemediationDeploymentsAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights:listRemediationDeploymentsAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtSubscriptionArgs struct {
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
	// Maximum number of records to return.
	Top *int `pulumi:"top"`
}

// List of deployments for a remediation.
type ListRemediationDeploymentsAtSubscriptionResult struct {
	// The URL to get the next set of results.
	NextLink string `pulumi:"nextLink"`
	// Array of deployments for the remediation.
	Value []RemediationDeploymentResponse `pulumi:"value"`
}

func ListRemediationDeploymentsAtSubscriptionOutput(ctx *pulumi.Context, args ListRemediationDeploymentsAtSubscriptionOutputArgs, opts ...pulumi.InvokeOption) ListRemediationDeploymentsAtSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRemediationDeploymentsAtSubscriptionResult, error) {
			args := v.(ListRemediationDeploymentsAtSubscriptionArgs)
			r, err := ListRemediationDeploymentsAtSubscription(ctx, &args, opts...)
			var s ListRemediationDeploymentsAtSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRemediationDeploymentsAtSubscriptionResultOutput)
}

type ListRemediationDeploymentsAtSubscriptionOutputArgs struct {
	// The name of the remediation.
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
	// Maximum number of records to return.
	Top pulumi.IntPtrInput `pulumi:"top"`
}

func (ListRemediationDeploymentsAtSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtSubscriptionArgs)(nil)).Elem()
}

// List of deployments for a remediation.
type ListRemediationDeploymentsAtSubscriptionResultOutput struct{ *pulumi.OutputState }

func (ListRemediationDeploymentsAtSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtSubscriptionResult)(nil)).Elem()
}

func (o ListRemediationDeploymentsAtSubscriptionResultOutput) ToListRemediationDeploymentsAtSubscriptionResultOutput() ListRemediationDeploymentsAtSubscriptionResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtSubscriptionResultOutput) ToListRemediationDeploymentsAtSubscriptionResultOutputWithContext(ctx context.Context) ListRemediationDeploymentsAtSubscriptionResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtSubscriptionResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListRemediationDeploymentsAtSubscriptionResult] {
	return pulumix.Output[ListRemediationDeploymentsAtSubscriptionResult]{
		OutputState: o.OutputState,
	}
}

// The URL to get the next set of results.
func (o ListRemediationDeploymentsAtSubscriptionResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtSubscriptionResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Array of deployments for the remediation.
func (o ListRemediationDeploymentsAtSubscriptionResultOutput) Value() RemediationDeploymentResponseArrayOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtSubscriptionResult) []RemediationDeploymentResponse { return v.Value }).(RemediationDeploymentResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRemediationDeploymentsAtSubscriptionResultOutput{})
}
