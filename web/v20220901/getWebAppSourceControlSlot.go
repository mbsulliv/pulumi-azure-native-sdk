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

// Description for Gets the source control configuration of an app.
func LookupWebAppSourceControlSlot(ctx *pulumi.Context, args *LookupWebAppSourceControlSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSourceControlSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppSourceControlSlotResult
	err := ctx.Invoke("azure-native:web/v20220901:getWebAppSourceControlSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSourceControlSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the source control configuration for the production slot.
	Slot string `pulumi:"slot"`
}

// Source control configuration for an app.
type LookupWebAppSourceControlSlotResult struct {
	// Name of branch to use for deployment.
	Branch *string `pulumi:"branch"`
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled *bool `pulumi:"deploymentRollbackEnabled"`
	// If GitHub Action is selected, than the associated configuration.
	GitHubActionConfiguration *GitHubActionConfigurationResponse `pulumi:"gitHubActionConfiguration"`
	// Resource Id.
	Id string `pulumi:"id"`
	// <code>true</code> if this is deployed via GitHub action.
	IsGitHubAction *bool `pulumi:"isGitHubAction"`
	// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration *bool `pulumi:"isManualIntegration"`
	// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial *bool `pulumi:"isMercurial"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Repository or source control URL.
	RepoUrl *string `pulumi:"repoUrl"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppSourceControlSlotOutput(ctx *pulumi.Context, args LookupWebAppSourceControlSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSourceControlSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSourceControlSlotResult, error) {
			args := v.(LookupWebAppSourceControlSlotArgs)
			r, err := LookupWebAppSourceControlSlot(ctx, &args, opts...)
			var s LookupWebAppSourceControlSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSourceControlSlotResultOutput)
}

type LookupWebAppSourceControlSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the source control configuration for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppSourceControlSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSourceControlSlotArgs)(nil)).Elem()
}

// Source control configuration for an app.
type LookupWebAppSourceControlSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSourceControlSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSourceControlSlotResult)(nil)).Elem()
}

func (o LookupWebAppSourceControlSlotResultOutput) ToLookupWebAppSourceControlSlotResultOutput() LookupWebAppSourceControlSlotResultOutput {
	return o
}

func (o LookupWebAppSourceControlSlotResultOutput) ToLookupWebAppSourceControlSlotResultOutputWithContext(ctx context.Context) LookupWebAppSourceControlSlotResultOutput {
	return o
}

func (o LookupWebAppSourceControlSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppSourceControlSlotResult] {
	return pulumix.Output[LookupWebAppSourceControlSlotResult]{
		OutputState: o.OutputState,
	}
}

// Name of branch to use for deployment.
func (o LookupWebAppSourceControlSlotResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
func (o LookupWebAppSourceControlSlotResultOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *bool { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

// If GitHub Action is selected, than the associated configuration.
func (o LookupWebAppSourceControlSlotResultOutput) GitHubActionConfiguration() GitHubActionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *GitHubActionConfigurationResponse {
		return v.GitHubActionConfiguration
	}).(GitHubActionConfigurationResponsePtrOutput)
}

// Resource Id.
func (o LookupWebAppSourceControlSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// <code>true</code> if this is deployed via GitHub action.
func (o LookupWebAppSourceControlSlotResultOutput) IsGitHubAction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *bool { return v.IsGitHubAction }).(pulumi.BoolPtrOutput)
}

// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
func (o LookupWebAppSourceControlSlotResultOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *bool { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
func (o LookupWebAppSourceControlSlotResultOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *bool { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o LookupWebAppSourceControlSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppSourceControlSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Repository or source control URL.
func (o LookupWebAppSourceControlSlotResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWebAppSourceControlSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSourceControlSlotResultOutput{})
}
