// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230707preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a msixpackage.
func LookupMSIXPackage(ctx *pulumi.Context, args *LookupMSIXPackageArgs, opts ...pulumi.InvokeOption) (*LookupMSIXPackageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMSIXPackageResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20230707preview:getMSIXPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMSIXPackageArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName string `pulumi:"hostPoolName"`
	// The version specific package full name of the MSIX package within specified hostpool
	MsixPackageFullName string `pulumi:"msixPackageFullName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Schema for MSIX Package properties.
type LookupMSIXPackageResult struct {
	// User friendly Name to be displayed in the portal.
	DisplayName *string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// VHD/CIM image path on Network Share.
	ImagePath *string `pulumi:"imagePath"`
	// Make this version of the package the active one across the hostpool.
	IsActive *bool `pulumi:"isActive"`
	// Specifies how to register Package in feed.
	IsRegularRegistration *bool `pulumi:"isRegularRegistration"`
	// Date Package was last updated, found in the appxmanifest.xml.
	LastUpdated *string `pulumi:"lastUpdated"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of package applications.
	PackageApplications []MsixPackageApplicationsResponse `pulumi:"packageApplications"`
	// List of package dependencies.
	PackageDependencies []MsixPackageDependenciesResponse `pulumi:"packageDependencies"`
	// Package Family Name from appxmanifest.xml. Contains Package Name and Publisher name.
	PackageFamilyName *string `pulumi:"packageFamilyName"`
	// Package Name from appxmanifest.xml.
	PackageName *string `pulumi:"packageName"`
	// Relative Path to the package inside the image.
	PackageRelativePath *string `pulumi:"packageRelativePath"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Package Version found in the appxmanifest.xml.
	Version *string `pulumi:"version"`
}

func LookupMSIXPackageOutput(ctx *pulumi.Context, args LookupMSIXPackageOutputArgs, opts ...pulumi.InvokeOption) LookupMSIXPackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMSIXPackageResult, error) {
			args := v.(LookupMSIXPackageArgs)
			r, err := LookupMSIXPackage(ctx, &args, opts...)
			var s LookupMSIXPackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMSIXPackageResultOutput)
}

type LookupMSIXPackageOutputArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName pulumi.StringInput `pulumi:"hostPoolName"`
	// The version specific package full name of the MSIX package within specified hostpool
	MsixPackageFullName pulumi.StringInput `pulumi:"msixPackageFullName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMSIXPackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMSIXPackageArgs)(nil)).Elem()
}

// Schema for MSIX Package properties.
type LookupMSIXPackageResultOutput struct{ *pulumi.OutputState }

func (LookupMSIXPackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMSIXPackageResult)(nil)).Elem()
}

func (o LookupMSIXPackageResultOutput) ToLookupMSIXPackageResultOutput() LookupMSIXPackageResultOutput {
	return o
}

func (o LookupMSIXPackageResultOutput) ToLookupMSIXPackageResultOutputWithContext(ctx context.Context) LookupMSIXPackageResultOutput {
	return o
}

func (o LookupMSIXPackageResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMSIXPackageResult] {
	return pulumix.Output[LookupMSIXPackageResult]{
		OutputState: o.OutputState,
	}
}

// User friendly Name to be displayed in the portal.
func (o LookupMSIXPackageResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMSIXPackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) string { return v.Id }).(pulumi.StringOutput)
}

// VHD/CIM image path on Network Share.
func (o LookupMSIXPackageResultOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.ImagePath }).(pulumi.StringPtrOutput)
}

// Make this version of the package the active one across the hostpool.
func (o LookupMSIXPackageResultOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *bool { return v.IsActive }).(pulumi.BoolPtrOutput)
}

// Specifies how to register Package in feed.
func (o LookupMSIXPackageResultOutput) IsRegularRegistration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *bool { return v.IsRegularRegistration }).(pulumi.BoolPtrOutput)
}

// Date Package was last updated, found in the appxmanifest.xml.
func (o LookupMSIXPackageResultOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.LastUpdated }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupMSIXPackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of package applications.
func (o LookupMSIXPackageResultOutput) PackageApplications() MsixPackageApplicationsResponseArrayOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) []MsixPackageApplicationsResponse { return v.PackageApplications }).(MsixPackageApplicationsResponseArrayOutput)
}

// List of package dependencies.
func (o LookupMSIXPackageResultOutput) PackageDependencies() MsixPackageDependenciesResponseArrayOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) []MsixPackageDependenciesResponse { return v.PackageDependencies }).(MsixPackageDependenciesResponseArrayOutput)
}

// Package Family Name from appxmanifest.xml. Contains Package Name and Publisher name.
func (o LookupMSIXPackageResultOutput) PackageFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.PackageFamilyName }).(pulumi.StringPtrOutput)
}

// Package Name from appxmanifest.xml.
func (o LookupMSIXPackageResultOutput) PackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.PackageName }).(pulumi.StringPtrOutput)
}

// Relative Path to the package inside the image.
func (o LookupMSIXPackageResultOutput) PackageRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.PackageRelativePath }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupMSIXPackageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMSIXPackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) string { return v.Type }).(pulumi.StringOutput)
}

// Package Version found in the appxmanifest.xml.
func (o LookupMSIXPackageResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMSIXPackageResultOutput{})
}
