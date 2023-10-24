// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets Catalog Devbox Definition error details
// Azure REST API version: 2023-08-01-preview.
//
// Other available API versions: 2023-10-01-preview.
func GetCatalogDevBoxDefinitionErrorDetails(ctx *pulumi.Context, args *GetCatalogDevBoxDefinitionErrorDetailsArgs, opts ...pulumi.InvokeOption) (*GetCatalogDevBoxDefinitionErrorDetailsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetCatalogDevBoxDefinitionErrorDetailsResult
	err := ctx.Invoke("azure-native:devcenter:getCatalogDevBoxDefinitionErrorDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetCatalogDevBoxDefinitionErrorDetailsArgs struct {
	// The name of the Catalog.
	CatalogName string `pulumi:"catalogName"`
	// The name of the Dev Box definition.
	DevBoxDefinitionName string `pulumi:"devBoxDefinitionName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of validator error details. Populated when changes are made to the resource or its dependent resources that impact the validity of the Catalog resource.
type GetCatalogDevBoxDefinitionErrorDetailsResult struct {
	// Errors associated with resources synchronized from the catalog.
	Errors []CatalogErrorDetailsResponse `pulumi:"errors"`
}

func GetCatalogDevBoxDefinitionErrorDetailsOutput(ctx *pulumi.Context, args GetCatalogDevBoxDefinitionErrorDetailsOutputArgs, opts ...pulumi.InvokeOption) GetCatalogDevBoxDefinitionErrorDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCatalogDevBoxDefinitionErrorDetailsResult, error) {
			args := v.(GetCatalogDevBoxDefinitionErrorDetailsArgs)
			r, err := GetCatalogDevBoxDefinitionErrorDetails(ctx, &args, opts...)
			var s GetCatalogDevBoxDefinitionErrorDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCatalogDevBoxDefinitionErrorDetailsResultOutput)
}

type GetCatalogDevBoxDefinitionErrorDetailsOutputArgs struct {
	// The name of the Catalog.
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// The name of the Dev Box definition.
	DevBoxDefinitionName pulumi.StringInput `pulumi:"devBoxDefinitionName"`
	// The name of the devcenter.
	DevCenterName pulumi.StringInput `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetCatalogDevBoxDefinitionErrorDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCatalogDevBoxDefinitionErrorDetailsArgs)(nil)).Elem()
}

// List of validator error details. Populated when changes are made to the resource or its dependent resources that impact the validity of the Catalog resource.
type GetCatalogDevBoxDefinitionErrorDetailsResultOutput struct{ *pulumi.OutputState }

func (GetCatalogDevBoxDefinitionErrorDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCatalogDevBoxDefinitionErrorDetailsResult)(nil)).Elem()
}

func (o GetCatalogDevBoxDefinitionErrorDetailsResultOutput) ToGetCatalogDevBoxDefinitionErrorDetailsResultOutput() GetCatalogDevBoxDefinitionErrorDetailsResultOutput {
	return o
}

func (o GetCatalogDevBoxDefinitionErrorDetailsResultOutput) ToGetCatalogDevBoxDefinitionErrorDetailsResultOutputWithContext(ctx context.Context) GetCatalogDevBoxDefinitionErrorDetailsResultOutput {
	return o
}

func (o GetCatalogDevBoxDefinitionErrorDetailsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetCatalogDevBoxDefinitionErrorDetailsResult] {
	return pulumix.Output[GetCatalogDevBoxDefinitionErrorDetailsResult]{
		OutputState: o.OutputState,
	}
}

// Errors associated with resources synchronized from the catalog.
func (o GetCatalogDevBoxDefinitionErrorDetailsResultOutput) Errors() CatalogErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v GetCatalogDevBoxDefinitionErrorDetailsResult) []CatalogErrorDetailsResponse { return v.Errors }).(CatalogErrorDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCatalogDevBoxDefinitionErrorDetailsResultOutput{})
}
