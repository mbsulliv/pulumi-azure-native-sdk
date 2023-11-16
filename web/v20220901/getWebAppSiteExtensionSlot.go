// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Get site extension information by its ID for a web site, or a deployment slot.
func LookupWebAppSiteExtensionSlot(ctx *pulumi.Context, args *LookupWebAppSiteExtensionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSiteExtensionSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppSiteExtensionSlotResult
	err := ctx.Invoke("azure-native:web/v20220901:getWebAppSiteExtensionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSiteExtensionSlotArgs struct {
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site extension name.
	SiteExtensionId string `pulumi:"siteExtensionId"`
	// Name of the deployment slot. If a slot is not specified, the API uses the production slot.
	Slot string `pulumi:"slot"`
}

// Site Extension Information.
type LookupWebAppSiteExtensionSlotResult struct {
	// List of authors.
	Authors []string `pulumi:"authors"`
	// Site Extension comment.
	Comment *string `pulumi:"comment"`
	// Detailed description.
	Description *string `pulumi:"description"`
	// Count of downloads.
	DownloadCount *int `pulumi:"downloadCount"`
	// Site extension ID.
	ExtensionId *string `pulumi:"extensionId"`
	// Site extension type.
	ExtensionType *string `pulumi:"extensionType"`
	// Extension URL.
	ExtensionUrl *string `pulumi:"extensionUrl"`
	// Feed URL.
	FeedUrl *string `pulumi:"feedUrl"`
	// Icon URL.
	IconUrl *string `pulumi:"iconUrl"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Installed timestamp.
	InstalledDateTime *string `pulumi:"installedDateTime"`
	// Installer command line parameters.
	InstallerCommandLineParams *string `pulumi:"installerCommandLineParams"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// License URL.
	LicenseUrl *string `pulumi:"licenseUrl"`
	// <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
	LocalIsLatestVersion *bool `pulumi:"localIsLatestVersion"`
	// Local path.
	LocalPath *string `pulumi:"localPath"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Project URL.
	ProjectUrl *string `pulumi:"projectUrl"`
	// Provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Published timestamp.
	PublishedDateTime *string `pulumi:"publishedDateTime"`
	// Summary description.
	Summary *string `pulumi:"summary"`
	Title   *string `pulumi:"title"`
	// Resource type.
	Type string `pulumi:"type"`
	// Version information.
	Version *string `pulumi:"version"`
}

func LookupWebAppSiteExtensionSlotOutput(ctx *pulumi.Context, args LookupWebAppSiteExtensionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSiteExtensionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSiteExtensionSlotResult, error) {
			args := v.(LookupWebAppSiteExtensionSlotArgs)
			r, err := LookupWebAppSiteExtensionSlot(ctx, &args, opts...)
			var s LookupWebAppSiteExtensionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSiteExtensionSlotResultOutput)
}

type LookupWebAppSiteExtensionSlotOutputArgs struct {
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site extension name.
	SiteExtensionId pulumi.StringInput `pulumi:"siteExtensionId"`
	// Name of the deployment slot. If a slot is not specified, the API uses the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppSiteExtensionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionSlotArgs)(nil)).Elem()
}

// Site Extension Information.
type LookupWebAppSiteExtensionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSiteExtensionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionSlotResult)(nil)).Elem()
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ToLookupWebAppSiteExtensionSlotResultOutput() LookupWebAppSiteExtensionSlotResultOutput {
	return o
}

func (o LookupWebAppSiteExtensionSlotResultOutput) ToLookupWebAppSiteExtensionSlotResultOutputWithContext(ctx context.Context) LookupWebAppSiteExtensionSlotResultOutput {
	return o
}

// List of authors.
func (o LookupWebAppSiteExtensionSlotResultOutput) Authors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) []string { return v.Authors }).(pulumi.StringArrayOutput)
}

// Site Extension comment.
func (o LookupWebAppSiteExtensionSlotResultOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

// Detailed description.
func (o LookupWebAppSiteExtensionSlotResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Count of downloads.
func (o LookupWebAppSiteExtensionSlotResultOutput) DownloadCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *int { return v.DownloadCount }).(pulumi.IntPtrOutput)
}

// Site extension ID.
func (o LookupWebAppSiteExtensionSlotResultOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

// Site extension type.
func (o LookupWebAppSiteExtensionSlotResultOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

// Extension URL.
func (o LookupWebAppSiteExtensionSlotResultOutput) ExtensionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ExtensionUrl }).(pulumi.StringPtrOutput)
}

// Feed URL.
func (o LookupWebAppSiteExtensionSlotResultOutput) FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.FeedUrl }).(pulumi.StringPtrOutput)
}

// Icon URL.
func (o LookupWebAppSiteExtensionSlotResultOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppSiteExtensionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Installed timestamp.
func (o LookupWebAppSiteExtensionSlotResultOutput) InstalledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.InstalledDateTime }).(pulumi.StringPtrOutput)
}

// Installer command line parameters.
func (o LookupWebAppSiteExtensionSlotResultOutput) InstallerCommandLineParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.InstallerCommandLineParams }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o LookupWebAppSiteExtensionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// License URL.
func (o LookupWebAppSiteExtensionSlotResultOutput) LicenseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.LicenseUrl }).(pulumi.StringPtrOutput)
}

// <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
func (o LookupWebAppSiteExtensionSlotResultOutput) LocalIsLatestVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *bool { return v.LocalIsLatestVersion }).(pulumi.BoolPtrOutput)
}

// Local path.
func (o LookupWebAppSiteExtensionSlotResultOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.LocalPath }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppSiteExtensionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project URL.
func (o LookupWebAppSiteExtensionSlotResultOutput) ProjectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ProjectUrl }).(pulumi.StringPtrOutput)
}

// Provisioning state.
func (o LookupWebAppSiteExtensionSlotResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Published timestamp.
func (o LookupWebAppSiteExtensionSlotResultOutput) PublishedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.PublishedDateTime }).(pulumi.StringPtrOutput)
}

// Summary description.
func (o LookupWebAppSiteExtensionSlotResultOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionSlotResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWebAppSiteExtensionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version information.
func (o LookupWebAppSiteExtensionSlotResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionSlotResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSiteExtensionSlotResultOutput{})
}
