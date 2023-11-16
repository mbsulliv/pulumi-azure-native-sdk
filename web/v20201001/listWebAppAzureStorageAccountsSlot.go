// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Azure storage account configurations of an app.
func ListWebAppAzureStorageAccountsSlot(ctx *pulumi.Context, args *ListWebAppAzureStorageAccountsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppAzureStorageAccountsSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppAzureStorageAccountsSlotResult
	err := ctx.Invoke("azure-native:web/v20201001:listWebAppAzureStorageAccountsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAzureStorageAccountsSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the Azure storage account configurations for the production slot.
	Slot string `pulumi:"slot"`
}

// AzureStorageInfo dictionary resource.
type ListWebAppAzureStorageAccountsSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Azure storage accounts.
	Properties map[string]AzureStorageInfoValueResponse `pulumi:"properties"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppAzureStorageAccountsSlotOutput(ctx *pulumi.Context, args ListWebAppAzureStorageAccountsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppAzureStorageAccountsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppAzureStorageAccountsSlotResult, error) {
			args := v.(ListWebAppAzureStorageAccountsSlotArgs)
			r, err := ListWebAppAzureStorageAccountsSlot(ctx, &args, opts...)
			var s ListWebAppAzureStorageAccountsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppAzureStorageAccountsSlotResultOutput)
}

type ListWebAppAzureStorageAccountsSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the Azure storage account configurations for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppAzureStorageAccountsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAzureStorageAccountsSlotArgs)(nil)).Elem()
}

// AzureStorageInfo dictionary resource.
type ListWebAppAzureStorageAccountsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppAzureStorageAccountsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAzureStorageAccountsSlotResult)(nil)).Elem()
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) ToListWebAppAzureStorageAccountsSlotResultOutput() ListWebAppAzureStorageAccountsSlotResultOutput {
	return o
}

func (o ListWebAppAzureStorageAccountsSlotResultOutput) ToListWebAppAzureStorageAccountsSlotResultOutputWithContext(ctx context.Context) ListWebAppAzureStorageAccountsSlotResultOutput {
	return o
}

// Resource Id.
func (o ListWebAppAzureStorageAccountsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppAzureStorageAccountsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppAzureStorageAccountsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure storage accounts.
func (o ListWebAppAzureStorageAccountsSlotResultOutput) Properties() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) map[string]AzureStorageInfoValueResponse {
		return v.Properties
	}).(AzureStorageInfoValueResponseMapOutput)
}

// The system metadata relating to this resource.
func (o ListWebAppAzureStorageAccountsSlotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o ListWebAppAzureStorageAccountsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAzureStorageAccountsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppAzureStorageAccountsSlotResultOutput{})
}
