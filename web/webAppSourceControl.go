// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Source control configuration for an app.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2020-10-01, 2023-01-01.
type WebAppSourceControl struct {
	pulumi.CustomResourceState

	// Name of branch to use for deployment.
	Branch pulumi.StringPtrOutput `pulumi:"branch"`
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled pulumi.BoolPtrOutput `pulumi:"deploymentRollbackEnabled"`
	// If GitHub Action is selected, than the associated configuration.
	GitHubActionConfiguration GitHubActionConfigurationResponsePtrOutput `pulumi:"gitHubActionConfiguration"`
	// <code>true</code> if this is deployed via GitHub action.
	IsGitHubAction pulumi.BoolPtrOutput `pulumi:"isGitHubAction"`
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration pulumi.BoolPtrOutput `pulumi:"isManualIntegration"`
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial pulumi.BoolPtrOutput `pulumi:"isMercurial"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Repository or source control URL.
	RepoUrl pulumi.StringPtrOutput `pulumi:"repoUrl"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppSourceControl registers a new resource with the given unique name, arguments, and options.
func NewWebAppSourceControl(ctx *pulumi.Context,
	name string, args *WebAppSourceControlArgs, opts ...pulumi.ResourceOption) (*WebAppSourceControl, error) {
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
			Type: pulumi.String("azure-native:web/v20150801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppSourceControl"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppSourceControl
	err := ctx.RegisterResource("azure-native:web:WebAppSourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSourceControl gets an existing WebAppSourceControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSourceControlState, opts ...pulumi.ResourceOption) (*WebAppSourceControl, error) {
	var resource WebAppSourceControl
	err := ctx.ReadResource("azure-native:web:WebAppSourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSourceControl resources.
type webAppSourceControlState struct {
}

type WebAppSourceControlState struct {
}

func (WebAppSourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSourceControlState)(nil)).Elem()
}

type webAppSourceControlArgs struct {
	// Name of branch to use for deployment.
	Branch *string `pulumi:"branch"`
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled *bool `pulumi:"deploymentRollbackEnabled"`
	// If GitHub Action is selected, than the associated configuration.
	GitHubActionConfiguration *GitHubActionConfiguration `pulumi:"gitHubActionConfiguration"`
	// <code>true</code> if this is deployed via GitHub action.
	IsGitHubAction *bool `pulumi:"isGitHubAction"`
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration *bool `pulumi:"isManualIntegration"`
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial *bool `pulumi:"isMercurial"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Repository or source control URL.
	RepoUrl *string `pulumi:"repoUrl"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppSourceControl resource.
type WebAppSourceControlArgs struct {
	// Name of branch to use for deployment.
	Branch pulumi.StringPtrInput
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled pulumi.BoolPtrInput
	// If GitHub Action is selected, than the associated configuration.
	GitHubActionConfiguration GitHubActionConfigurationPtrInput
	// <code>true</code> if this is deployed via GitHub action.
	IsGitHubAction pulumi.BoolPtrInput
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration pulumi.BoolPtrInput
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Repository or source control URL.
	RepoUrl pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppSourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSourceControlArgs)(nil)).Elem()
}

type WebAppSourceControlInput interface {
	pulumi.Input

	ToWebAppSourceControlOutput() WebAppSourceControlOutput
	ToWebAppSourceControlOutputWithContext(ctx context.Context) WebAppSourceControlOutput
}

func (*WebAppSourceControl) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSourceControl)(nil)).Elem()
}

func (i *WebAppSourceControl) ToWebAppSourceControlOutput() WebAppSourceControlOutput {
	return i.ToWebAppSourceControlOutputWithContext(context.Background())
}

func (i *WebAppSourceControl) ToWebAppSourceControlOutputWithContext(ctx context.Context) WebAppSourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSourceControlOutput)
}

type WebAppSourceControlOutput struct{ *pulumi.OutputState }

func (WebAppSourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSourceControl)(nil)).Elem()
}

func (o WebAppSourceControlOutput) ToWebAppSourceControlOutput() WebAppSourceControlOutput {
	return o
}

func (o WebAppSourceControlOutput) ToWebAppSourceControlOutputWithContext(ctx context.Context) WebAppSourceControlOutput {
	return o
}

// Name of branch to use for deployment.
func (o WebAppSourceControlOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
func (o WebAppSourceControlOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.BoolPtrOutput { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

// If GitHub Action is selected, than the associated configuration.
func (o WebAppSourceControlOutput) GitHubActionConfiguration() GitHubActionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) GitHubActionConfigurationResponsePtrOutput {
		return v.GitHubActionConfiguration
	}).(GitHubActionConfigurationResponsePtrOutput)
}

// <code>true</code> if this is deployed via GitHub action.
func (o WebAppSourceControlOutput) IsGitHubAction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.BoolPtrOutput { return v.IsGitHubAction }).(pulumi.BoolPtrOutput)
}

// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
func (o WebAppSourceControlOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.BoolPtrOutput { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
func (o WebAppSourceControlOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.BoolPtrOutput { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o WebAppSourceControlOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppSourceControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Repository or source control URL.
func (o WebAppSourceControlOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringPtrOutput { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WebAppSourceControlOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSourceControlOutput{})
}
