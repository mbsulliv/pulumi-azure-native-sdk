// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Project resource.
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:storagemover/v20230301:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	// The name of the Project resource.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName string `pulumi:"storageMoverName"`
}

// The Project resource.
type LookupProjectResult struct {
	// A description for the Project.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of this resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource system metadata.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	// The name of the Project resource.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName pulumi.StringInput `pulumi:"storageMoverName"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// The Project resource.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupProjectResult] {
	return pulumix.Output[LookupProjectResult]{
		OutputState: o.OutputState,
	}
}

// A description for the Project.
func (o LookupProjectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of this resource.
func (o LookupProjectResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource system metadata.
func (o LookupProjectResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProjectResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
