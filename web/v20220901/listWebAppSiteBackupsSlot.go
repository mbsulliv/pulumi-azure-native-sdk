// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Gets existing backups of an app.
func ListWebAppSiteBackupsSlot(ctx *pulumi.Context, args *ListWebAppSiteBackupsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSiteBackupsSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppSiteBackupsSlotResult
	err := ctx.Invoke("azure-native:web/v20220901:listWebAppSiteBackupsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSiteBackupsSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get backups of the production slot.
	Slot string `pulumi:"slot"`
}

// Collection of backup items.
type ListWebAppSiteBackupsSlotResult struct {
	// Link to next page of resources.
	NextLink string `pulumi:"nextLink"`
	// Collection of resources.
	Value []BackupItemResponse `pulumi:"value"`
}

func ListWebAppSiteBackupsSlotOutput(ctx *pulumi.Context, args ListWebAppSiteBackupsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSiteBackupsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSiteBackupsSlotResult, error) {
			args := v.(ListWebAppSiteBackupsSlotArgs)
			r, err := ListWebAppSiteBackupsSlot(ctx, &args, opts...)
			var s ListWebAppSiteBackupsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSiteBackupsSlotResultOutput)
}

type ListWebAppSiteBackupsSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get backups of the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppSiteBackupsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSiteBackupsSlotArgs)(nil)).Elem()
}

// Collection of backup items.
type ListWebAppSiteBackupsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSiteBackupsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSiteBackupsSlotResult)(nil)).Elem()
}

func (o ListWebAppSiteBackupsSlotResultOutput) ToListWebAppSiteBackupsSlotResultOutput() ListWebAppSiteBackupsSlotResultOutput {
	return o
}

func (o ListWebAppSiteBackupsSlotResultOutput) ToListWebAppSiteBackupsSlotResultOutputWithContext(ctx context.Context) ListWebAppSiteBackupsSlotResultOutput {
	return o
}

func (o ListWebAppSiteBackupsSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppSiteBackupsSlotResult] {
	return pulumix.Output[ListWebAppSiteBackupsSlotResult]{
		OutputState: o.OutputState,
	}
}

// Link to next page of resources.
func (o ListWebAppSiteBackupsSlotResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSiteBackupsSlotResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Collection of resources.
func (o ListWebAppSiteBackupsSlotResultOutput) Value() BackupItemResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppSiteBackupsSlotResult) []BackupItemResponse { return v.Value }).(BackupItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSiteBackupsSlotResultOutput{})
}
