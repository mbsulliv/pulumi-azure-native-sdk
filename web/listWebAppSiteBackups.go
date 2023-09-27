// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Gets existing backups of an app.
// Azure REST API version: 2022-09-01.
func ListWebAppSiteBackups(ctx *pulumi.Context, args *ListWebAppSiteBackupsArgs, opts ...pulumi.InvokeOption) (*ListWebAppSiteBackupsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppSiteBackupsResult
	err := ctx.Invoke("azure-native:web:listWebAppSiteBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSiteBackupsArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Collection of backup items.
type ListWebAppSiteBackupsResult struct {
	// Link to next page of resources.
	NextLink string `pulumi:"nextLink"`
	// Collection of resources.
	Value []BackupItemResponse `pulumi:"value"`
}

func ListWebAppSiteBackupsOutput(ctx *pulumi.Context, args ListWebAppSiteBackupsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSiteBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSiteBackupsResult, error) {
			args := v.(ListWebAppSiteBackupsArgs)
			r, err := ListWebAppSiteBackups(ctx, &args, opts...)
			var s ListWebAppSiteBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSiteBackupsResultOutput)
}

type ListWebAppSiteBackupsOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppSiteBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSiteBackupsArgs)(nil)).Elem()
}

// Collection of backup items.
type ListWebAppSiteBackupsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSiteBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSiteBackupsResult)(nil)).Elem()
}

func (o ListWebAppSiteBackupsResultOutput) ToListWebAppSiteBackupsResultOutput() ListWebAppSiteBackupsResultOutput {
	return o
}

func (o ListWebAppSiteBackupsResultOutput) ToListWebAppSiteBackupsResultOutputWithContext(ctx context.Context) ListWebAppSiteBackupsResultOutput {
	return o
}

func (o ListWebAppSiteBackupsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppSiteBackupsResult] {
	return pulumix.Output[ListWebAppSiteBackupsResult]{
		OutputState: o.OutputState,
	}
}

// Link to next page of resources.
func (o ListWebAppSiteBackupsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSiteBackupsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Collection of resources.
func (o ListWebAppSiteBackupsResultOutput) Value() BackupItemResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppSiteBackupsResult) []BackupItemResponse { return v.Value }).(BackupItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSiteBackupsResultOutput{})
}
