// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the archive version.
func LookupArchiveVersion(ctx *pulumi.Context, args *LookupArchiveVersionArgs, opts ...pulumi.InvokeOption) (*LookupArchiveVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupArchiveVersionResult
	err := ctx.Invoke("azure-native:containerregistry/v20231101preview:getArchiveVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArchiveVersionArgs struct {
	// The name of the archive resource.
	ArchiveName string `pulumi:"archiveName"`
	// The name of the archive version resource.
	ArchiveVersionName string `pulumi:"archiveVersionName"`
	// The type of the package resource.
	PackageType string `pulumi:"packageType"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents an export pipeline for a container registry.
type LookupArchiveVersionResult struct {
	// The detailed error message for the archive version in the case of failure.
	ArchiveVersionErrorMessage *string `pulumi:"archiveVersionErrorMessage"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning state of the archive at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupArchiveVersionOutput(ctx *pulumi.Context, args LookupArchiveVersionOutputArgs, opts ...pulumi.InvokeOption) LookupArchiveVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArchiveVersionResult, error) {
			args := v.(LookupArchiveVersionArgs)
			r, err := LookupArchiveVersion(ctx, &args, opts...)
			var s LookupArchiveVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArchiveVersionResultOutput)
}

type LookupArchiveVersionOutputArgs struct {
	// The name of the archive resource.
	ArchiveName pulumi.StringInput `pulumi:"archiveName"`
	// The name of the archive version resource.
	ArchiveVersionName pulumi.StringInput `pulumi:"archiveVersionName"`
	// The type of the package resource.
	PackageType pulumi.StringInput `pulumi:"packageType"`
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupArchiveVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArchiveVersionArgs)(nil)).Elem()
}

// An object that represents an export pipeline for a container registry.
type LookupArchiveVersionResultOutput struct{ *pulumi.OutputState }

func (LookupArchiveVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArchiveVersionResult)(nil)).Elem()
}

func (o LookupArchiveVersionResultOutput) ToLookupArchiveVersionResultOutput() LookupArchiveVersionResultOutput {
	return o
}

func (o LookupArchiveVersionResultOutput) ToLookupArchiveVersionResultOutputWithContext(ctx context.Context) LookupArchiveVersionResultOutput {
	return o
}

// The detailed error message for the archive version in the case of failure.
func (o LookupArchiveVersionResultOutput) ArchiveVersionErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArchiveVersionResult) *string { return v.ArchiveVersionErrorMessage }).(pulumi.StringPtrOutput)
}

// The resource ID.
func (o LookupArchiveVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArchiveVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupArchiveVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArchiveVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the archive at the time the operation was called.
func (o LookupArchiveVersionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArchiveVersionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupArchiveVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupArchiveVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupArchiveVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArchiveVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArchiveVersionResultOutput{})
}
