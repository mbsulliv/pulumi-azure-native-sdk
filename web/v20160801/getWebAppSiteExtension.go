// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get site extension information by its ID for a web site, or a deployment slot.
func LookupWebAppSiteExtension(ctx *pulumi.Context, args *LookupWebAppSiteExtensionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSiteExtensionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppSiteExtensionResult
	err := ctx.Invoke("azure-native:web/v20160801:getWebAppSiteExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSiteExtensionArgs struct {
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site extension name.
	SiteExtensionId string `pulumi:"siteExtensionId"`
}

// Site Extension Information.
type LookupWebAppSiteExtensionResult struct {
	// List of authors.
	Authors []string `pulumi:"authors"`
	// Site Extension comment.
	Comment *string `pulumi:"comment"`
	// Detailed description.
	Description *string `pulumi:"description"`
	// Count of downloads.
	DownloadCount *int `pulumi:"downloadCount"`
	// Extension URL.
	ExtensionUrl *string `pulumi:"extensionUrl"`
	// Feed URL.
	FeedUrl *string `pulumi:"feedUrl"`
	// Icon URL.
	IconUrl *string `pulumi:"iconUrl"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Installer command line parameters.
	InstallationArgs *string `pulumi:"installationArgs"`
	// Installed timestamp.
	InstalledDateTime *string `pulumi:"installedDateTime"`
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
	// Site extension title.
	Title *string `pulumi:"title"`
	// Resource type.
	Type string `pulumi:"type"`
	// Version information.
	Version *string `pulumi:"version"`
}

func LookupWebAppSiteExtensionOutput(ctx *pulumi.Context, args LookupWebAppSiteExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSiteExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSiteExtensionResult, error) {
			args := v.(LookupWebAppSiteExtensionArgs)
			r, err := LookupWebAppSiteExtension(ctx, &args, opts...)
			var s LookupWebAppSiteExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSiteExtensionResultOutput)
}

type LookupWebAppSiteExtensionOutputArgs struct {
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site extension name.
	SiteExtensionId pulumi.StringInput `pulumi:"siteExtensionId"`
}

func (LookupWebAppSiteExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionArgs)(nil)).Elem()
}

// Site Extension Information.
type LookupWebAppSiteExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSiteExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionResult)(nil)).Elem()
}

func (o LookupWebAppSiteExtensionResultOutput) ToLookupWebAppSiteExtensionResultOutput() LookupWebAppSiteExtensionResultOutput {
	return o
}

func (o LookupWebAppSiteExtensionResultOutput) ToLookupWebAppSiteExtensionResultOutputWithContext(ctx context.Context) LookupWebAppSiteExtensionResultOutput {
	return o
}

func (o LookupWebAppSiteExtensionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppSiteExtensionResult] {
	return pulumix.Output[LookupWebAppSiteExtensionResult]{
		OutputState: o.OutputState,
	}
}

// List of authors.
func (o LookupWebAppSiteExtensionResultOutput) Authors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) []string { return v.Authors }).(pulumi.StringArrayOutput)
}

// Site Extension comment.
func (o LookupWebAppSiteExtensionResultOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

// Detailed description.
func (o LookupWebAppSiteExtensionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Count of downloads.
func (o LookupWebAppSiteExtensionResultOutput) DownloadCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *int { return v.DownloadCount }).(pulumi.IntPtrOutput)
}

// Extension URL.
func (o LookupWebAppSiteExtensionResultOutput) ExtensionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ExtensionUrl }).(pulumi.StringPtrOutput)
}

// Feed URL.
func (o LookupWebAppSiteExtensionResultOutput) FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.FeedUrl }).(pulumi.StringPtrOutput)
}

// Icon URL.
func (o LookupWebAppSiteExtensionResultOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppSiteExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Installer command line parameters.
func (o LookupWebAppSiteExtensionResultOutput) InstallationArgs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.InstallationArgs }).(pulumi.StringPtrOutput)
}

// Installed timestamp.
func (o LookupWebAppSiteExtensionResultOutput) InstalledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.InstalledDateTime }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o LookupWebAppSiteExtensionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// License URL.
func (o LookupWebAppSiteExtensionResultOutput) LicenseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.LicenseUrl }).(pulumi.StringPtrOutput)
}

// <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
func (o LookupWebAppSiteExtensionResultOutput) LocalIsLatestVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *bool { return v.LocalIsLatestVersion }).(pulumi.BoolPtrOutput)
}

// Local path.
func (o LookupWebAppSiteExtensionResultOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.LocalPath }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppSiteExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project URL.
func (o LookupWebAppSiteExtensionResultOutput) ProjectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ProjectUrl }).(pulumi.StringPtrOutput)
}

// Provisioning state.
func (o LookupWebAppSiteExtensionResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Published timestamp.
func (o LookupWebAppSiteExtensionResultOutput) PublishedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.PublishedDateTime }).(pulumi.StringPtrOutput)
}

// Summary description.
func (o LookupWebAppSiteExtensionResultOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Summary }).(pulumi.StringPtrOutput)
}

// Site extension title.
func (o LookupWebAppSiteExtensionResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWebAppSiteExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version information.
func (o LookupWebAppSiteExtensionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSiteExtensionResultOutput{})
}
