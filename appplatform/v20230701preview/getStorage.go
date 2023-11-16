// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the storage resource.
func LookupStorage(ctx *pulumi.Context, args *LookupStorageArgs, opts ...pulumi.InvokeOption) (*LookupStorageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageResult
	err := ctx.Invoke("azure-native:appplatform/v20230701preview:getStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// The name of the storage resource.
	StorageName string `pulumi:"storageName"`
}

// Storage resource payload.
type LookupStorageResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the storage resource payload.
	Properties StorageAccountResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupStorageOutput(ctx *pulumi.Context, args LookupStorageOutputArgs, opts ...pulumi.InvokeOption) LookupStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageResult, error) {
			args := v.(LookupStorageArgs)
			r, err := LookupStorage(ctx, &args, opts...)
			var s LookupStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageResultOutput)
}

type LookupStorageOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The name of the storage resource.
	StorageName pulumi.StringInput `pulumi:"storageName"`
}

func (LookupStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageArgs)(nil)).Elem()
}

// Storage resource payload.
type LookupStorageResultOutput struct{ *pulumi.OutputState }

func (LookupStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageResult)(nil)).Elem()
}

func (o LookupStorageResultOutput) ToLookupStorageResultOutput() LookupStorageResultOutput {
	return o
}

func (o LookupStorageResultOutput) ToLookupStorageResultOutputWithContext(ctx context.Context) LookupStorageResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupStorageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the storage resource payload.
func (o LookupStorageResultOutput) Properties() StorageAccountResponseOutput {
	return o.ApplyT(func(v LookupStorageResult) StorageAccountResponse { return v.Properties }).(StorageAccountResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupStorageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupStorageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageResultOutput{})
}
