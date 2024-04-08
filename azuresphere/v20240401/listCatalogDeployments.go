// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists deployments for catalog.
func ListCatalogDeployments(ctx *pulumi.Context, args *ListCatalogDeploymentsArgs, opts ...pulumi.InvokeOption) (*ListCatalogDeploymentsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListCatalogDeploymentsResult
	err := ctx.Invoke("azure-native:azuresphere/v20240401:listCatalogDeployments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCatalogDeploymentsArgs struct {
	// Name of catalog
	CatalogName string `pulumi:"catalogName"`
	// Filter the result list using the given expression
	Filter *string `pulumi:"filter"`
	// The maximum number of result items per page.
	Maxpagesize *int `pulumi:"maxpagesize"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of result items to skip.
	Skip *int `pulumi:"skip"`
	// The number of result items to return.
	Top *int `pulumi:"top"`
}

// The response of a Deployment list operation.
type ListCatalogDeploymentsResult struct {
	// The link to the next page of items
	NextLink *string `pulumi:"nextLink"`
	// The Deployment items on this page
	Value []DeploymentResponse `pulumi:"value"`
}

func ListCatalogDeploymentsOutput(ctx *pulumi.Context, args ListCatalogDeploymentsOutputArgs, opts ...pulumi.InvokeOption) ListCatalogDeploymentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCatalogDeploymentsResult, error) {
			args := v.(ListCatalogDeploymentsArgs)
			r, err := ListCatalogDeployments(ctx, &args, opts...)
			var s ListCatalogDeploymentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCatalogDeploymentsResultOutput)
}

type ListCatalogDeploymentsOutputArgs struct {
	// Name of catalog
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// Filter the result list using the given expression
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The maximum number of result items per page.
	Maxpagesize pulumi.IntPtrInput `pulumi:"maxpagesize"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The number of result items to skip.
	Skip pulumi.IntPtrInput `pulumi:"skip"`
	// The number of result items to return.
	Top pulumi.IntPtrInput `pulumi:"top"`
}

func (ListCatalogDeploymentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCatalogDeploymentsArgs)(nil)).Elem()
}

// The response of a Deployment list operation.
type ListCatalogDeploymentsResultOutput struct{ *pulumi.OutputState }

func (ListCatalogDeploymentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCatalogDeploymentsResult)(nil)).Elem()
}

func (o ListCatalogDeploymentsResultOutput) ToListCatalogDeploymentsResultOutput() ListCatalogDeploymentsResultOutput {
	return o
}

func (o ListCatalogDeploymentsResultOutput) ToListCatalogDeploymentsResultOutputWithContext(ctx context.Context) ListCatalogDeploymentsResultOutput {
	return o
}

// The link to the next page of items
func (o ListCatalogDeploymentsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCatalogDeploymentsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// The Deployment items on this page
func (o ListCatalogDeploymentsResultOutput) Value() DeploymentResponseArrayOutput {
	return o.ApplyT(func(v ListCatalogDeploymentsResult) []DeploymentResponse { return v.Value }).(DeploymentResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListCatalogDeploymentsResultOutput{})
}
