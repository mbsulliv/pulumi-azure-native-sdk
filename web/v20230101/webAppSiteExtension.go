// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Site Extension Information.
type WebAppSiteExtension struct {
	pulumi.CustomResourceState

	// List of authors.
	Authors pulumi.StringArrayOutput `pulumi:"authors"`
	// Site Extension comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Detailed description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Count of downloads.
	DownloadCount pulumi.IntPtrOutput `pulumi:"downloadCount"`
	// Site extension ID.
	ExtensionId pulumi.StringPtrOutput `pulumi:"extensionId"`
	// Site extension type.
	ExtensionType pulumi.StringPtrOutput `pulumi:"extensionType"`
	// Extension URL.
	ExtensionUrl pulumi.StringPtrOutput `pulumi:"extensionUrl"`
	// Feed URL.
	FeedUrl pulumi.StringPtrOutput `pulumi:"feedUrl"`
	// Icon URL.
	IconUrl pulumi.StringPtrOutput `pulumi:"iconUrl"`
	// Installed timestamp.
	InstalledDateTime pulumi.StringPtrOutput `pulumi:"installedDateTime"`
	// Installer command line parameters.
	InstallerCommandLineParams pulumi.StringPtrOutput `pulumi:"installerCommandLineParams"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// License URL.
	LicenseUrl pulumi.StringPtrOutput `pulumi:"licenseUrl"`
	// <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
	LocalIsLatestVersion pulumi.BoolPtrOutput `pulumi:"localIsLatestVersion"`
	// Local path.
	LocalPath pulumi.StringPtrOutput `pulumi:"localPath"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Project URL.
	ProjectUrl pulumi.StringPtrOutput `pulumi:"projectUrl"`
	// Provisioning state.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Published timestamp.
	PublishedDateTime pulumi.StringPtrOutput `pulumi:"publishedDateTime"`
	// Summary description.
	Summary pulumi.StringPtrOutput `pulumi:"summary"`
	Title   pulumi.StringPtrOutput `pulumi:"title"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Version information.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewWebAppSiteExtension registers a new resource with the given unique name, arguments, and options.
func NewWebAppSiteExtension(ctx *pulumi.Context,
	name string, args *WebAppSiteExtensionArgs, opts ...pulumi.ResourceOption) (*WebAppSiteExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSiteExtension"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppSiteExtension"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppSiteExtension
	err := ctx.RegisterResource("azure-native:web/v20230101:WebAppSiteExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSiteExtension gets an existing WebAppSiteExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSiteExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSiteExtensionState, opts ...pulumi.ResourceOption) (*WebAppSiteExtension, error) {
	var resource WebAppSiteExtension
	err := ctx.ReadResource("azure-native:web/v20230101:WebAppSiteExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSiteExtension resources.
type webAppSiteExtensionState struct {
}

type WebAppSiteExtensionState struct {
}

func (WebAppSiteExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSiteExtensionState)(nil)).Elem()
}

type webAppSiteExtensionArgs struct {
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site extension name.
	SiteExtensionId *string `pulumi:"siteExtensionId"`
}

// The set of arguments for constructing a WebAppSiteExtension resource.
type WebAppSiteExtensionArgs struct {
	// Site name.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Site extension name.
	SiteExtensionId pulumi.StringPtrInput
}

func (WebAppSiteExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSiteExtensionArgs)(nil)).Elem()
}

type WebAppSiteExtensionInput interface {
	pulumi.Input

	ToWebAppSiteExtensionOutput() WebAppSiteExtensionOutput
	ToWebAppSiteExtensionOutputWithContext(ctx context.Context) WebAppSiteExtensionOutput
}

func (*WebAppSiteExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSiteExtension)(nil)).Elem()
}

func (i *WebAppSiteExtension) ToWebAppSiteExtensionOutput() WebAppSiteExtensionOutput {
	return i.ToWebAppSiteExtensionOutputWithContext(context.Background())
}

func (i *WebAppSiteExtension) ToWebAppSiteExtensionOutputWithContext(ctx context.Context) WebAppSiteExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSiteExtensionOutput)
}

type WebAppSiteExtensionOutput struct{ *pulumi.OutputState }

func (WebAppSiteExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSiteExtension)(nil)).Elem()
}

func (o WebAppSiteExtensionOutput) ToWebAppSiteExtensionOutput() WebAppSiteExtensionOutput {
	return o
}

func (o WebAppSiteExtensionOutput) ToWebAppSiteExtensionOutputWithContext(ctx context.Context) WebAppSiteExtensionOutput {
	return o
}

// List of authors.
func (o WebAppSiteExtensionOutput) Authors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringArrayOutput { return v.Authors }).(pulumi.StringArrayOutput)
}

// Site Extension comment.
func (o WebAppSiteExtensionOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Detailed description.
func (o WebAppSiteExtensionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Count of downloads.
func (o WebAppSiteExtensionOutput) DownloadCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.IntPtrOutput { return v.DownloadCount }).(pulumi.IntPtrOutput)
}

// Site extension ID.
func (o WebAppSiteExtensionOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

// Site extension type.
func (o WebAppSiteExtensionOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

// Extension URL.
func (o WebAppSiteExtensionOutput) ExtensionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ExtensionUrl }).(pulumi.StringPtrOutput)
}

// Feed URL.
func (o WebAppSiteExtensionOutput) FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.FeedUrl }).(pulumi.StringPtrOutput)
}

// Icon URL.
func (o WebAppSiteExtensionOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.IconUrl }).(pulumi.StringPtrOutput)
}

// Installed timestamp.
func (o WebAppSiteExtensionOutput) InstalledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.InstalledDateTime }).(pulumi.StringPtrOutput)
}

// Installer command line parameters.
func (o WebAppSiteExtensionOutput) InstallerCommandLineParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.InstallerCommandLineParams }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o WebAppSiteExtensionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// License URL.
func (o WebAppSiteExtensionOutput) LicenseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.LicenseUrl }).(pulumi.StringPtrOutput)
}

// <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
func (o WebAppSiteExtensionOutput) LocalIsLatestVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.BoolPtrOutput { return v.LocalIsLatestVersion }).(pulumi.BoolPtrOutput)
}

// Local path.
func (o WebAppSiteExtensionOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.LocalPath }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppSiteExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Project URL.
func (o WebAppSiteExtensionOutput) ProjectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ProjectUrl }).(pulumi.StringPtrOutput)
}

// Provisioning state.
func (o WebAppSiteExtensionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Published timestamp.
func (o WebAppSiteExtensionOutput) PublishedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.PublishedDateTime }).(pulumi.StringPtrOutput)
}

// Summary description.
func (o WebAppSiteExtensionOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o WebAppSiteExtensionOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WebAppSiteExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version information.
func (o WebAppSiteExtensionOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSiteExtension) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSiteExtensionOutput{})
}
