// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets all deployments for a remediation at resource scope.
func ListRemediationDeploymentsAtResource(ctx *pulumi.Context, args *ListRemediationDeploymentsAtResourceArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListRemediationDeploymentsAtResourceResult
	err := ctx.Invoke("azure-native:policyinsights/v20211001:listRemediationDeploymentsAtResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtResourceArgs struct {
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
	// Resource ID.
	ResourceId string `pulumi:"resourceId"`
	// Maximum number of records to return.
	Top *int `pulumi:"top"`
}

// List of deployments for a remediation.
type ListRemediationDeploymentsAtResourceResult struct {
	// The URL to get the next set of results.
	NextLink string `pulumi:"nextLink"`
	// Array of deployments for the remediation.
	Value []RemediationDeploymentResponse `pulumi:"value"`
}

func ListRemediationDeploymentsAtResourceOutput(ctx *pulumi.Context, args ListRemediationDeploymentsAtResourceOutputArgs, opts ...pulumi.InvokeOption) ListRemediationDeploymentsAtResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRemediationDeploymentsAtResourceResult, error) {
			args := v.(ListRemediationDeploymentsAtResourceArgs)
			r, err := ListRemediationDeploymentsAtResource(ctx, &args, opts...)
			var s ListRemediationDeploymentsAtResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRemediationDeploymentsAtResourceResultOutput)
}

type ListRemediationDeploymentsAtResourceOutputArgs struct {
	// The name of the remediation.
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
	// Resource ID.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Maximum number of records to return.
	Top pulumi.IntPtrInput `pulumi:"top"`
}

func (ListRemediationDeploymentsAtResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtResourceArgs)(nil)).Elem()
}

// List of deployments for a remediation.
type ListRemediationDeploymentsAtResourceResultOutput struct{ *pulumi.OutputState }

func (ListRemediationDeploymentsAtResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtResourceResult)(nil)).Elem()
}

func (o ListRemediationDeploymentsAtResourceResultOutput) ToListRemediationDeploymentsAtResourceResultOutput() ListRemediationDeploymentsAtResourceResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtResourceResultOutput) ToListRemediationDeploymentsAtResourceResultOutputWithContext(ctx context.Context) ListRemediationDeploymentsAtResourceResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtResourceResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListRemediationDeploymentsAtResourceResult] {
	return pulumix.Output[ListRemediationDeploymentsAtResourceResult]{
		OutputState: o.OutputState,
	}
}

// The URL to get the next set of results.
func (o ListRemediationDeploymentsAtResourceResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtResourceResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Array of deployments for the remediation.
func (o ListRemediationDeploymentsAtResourceResultOutput) Value() RemediationDeploymentResponseArrayOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtResourceResult) []RemediationDeploymentResponse { return v.Value }).(RemediationDeploymentResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRemediationDeploymentsAtResourceResultOutput{})
}
