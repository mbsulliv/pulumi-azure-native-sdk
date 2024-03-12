// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists devices for catalog.
func ListCatalogDevices(ctx *pulumi.Context, args *ListCatalogDevicesArgs, opts ...pulumi.InvokeOption) (*ListCatalogDevicesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListCatalogDevicesResult
	err := ctx.Invoke("azure-native:azuresphere/v20240401:listCatalogDevices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCatalogDevicesArgs struct {
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

// The response of a Device list operation.
type ListCatalogDevicesResult struct {
	// The link to the next page of items
	NextLink string `pulumi:"nextLink"`
	// The Device items on this page
	Value []DeviceResponse `pulumi:"value"`
}

func ListCatalogDevicesOutput(ctx *pulumi.Context, args ListCatalogDevicesOutputArgs, opts ...pulumi.InvokeOption) ListCatalogDevicesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCatalogDevicesResult, error) {
			args := v.(ListCatalogDevicesArgs)
			r, err := ListCatalogDevices(ctx, &args, opts...)
			var s ListCatalogDevicesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCatalogDevicesResultOutput)
}

type ListCatalogDevicesOutputArgs struct {
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

func (ListCatalogDevicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCatalogDevicesArgs)(nil)).Elem()
}

// The response of a Device list operation.
type ListCatalogDevicesResultOutput struct{ *pulumi.OutputState }

func (ListCatalogDevicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCatalogDevicesResult)(nil)).Elem()
}

func (o ListCatalogDevicesResultOutput) ToListCatalogDevicesResultOutput() ListCatalogDevicesResultOutput {
	return o
}

func (o ListCatalogDevicesResultOutput) ToListCatalogDevicesResultOutputWithContext(ctx context.Context) ListCatalogDevicesResultOutput {
	return o
}

// The link to the next page of items
func (o ListCatalogDevicesResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListCatalogDevicesResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// The Device items on this page
func (o ListCatalogDevicesResultOutput) Value() DeviceResponseArrayOutput {
	return o.ApplyT(func(v ListCatalogDevicesResult) []DeviceResponse { return v.Value }).(DeviceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListCatalogDevicesResultOutput{})
}
