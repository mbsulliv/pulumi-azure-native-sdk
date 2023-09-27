// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the source control configuration of an app.
func LookupWebAppSourceControl(ctx *pulumi.Context, args *LookupWebAppSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSourceControlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppSourceControlResult
	err := ctx.Invoke("azure-native:web/v20201001:getWebAppSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSourceControlArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Source control configuration for an app.
type LookupWebAppSourceControlResult struct {
	// Name of branch to use for deployment.
	Branch *string `pulumi:"branch"`
	// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled *bool `pulumi:"deploymentRollbackEnabled"`
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
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppSourceControlOutput(ctx *pulumi.Context, args LookupWebAppSourceControlOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSourceControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSourceControlResult, error) {
			args := v.(LookupWebAppSourceControlArgs)
			r, err := LookupWebAppSourceControl(ctx, &args, opts...)
			var s LookupWebAppSourceControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSourceControlResultOutput)
}

type LookupWebAppSourceControlOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppSourceControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSourceControlArgs)(nil)).Elem()
}

// Source control configuration for an app.
type LookupWebAppSourceControlResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSourceControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSourceControlResult)(nil)).Elem()
}

func (o LookupWebAppSourceControlResultOutput) ToLookupWebAppSourceControlResultOutput() LookupWebAppSourceControlResultOutput {
	return o
}

func (o LookupWebAppSourceControlResultOutput) ToLookupWebAppSourceControlResultOutputWithContext(ctx context.Context) LookupWebAppSourceControlResultOutput {
	return o
}

func (o LookupWebAppSourceControlResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppSourceControlResult] {
	return pulumix.Output[LookupWebAppSourceControlResult]{
		OutputState: o.OutputState,
	}
}

// Name of branch to use for deployment.
func (o LookupWebAppSourceControlResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

// <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
func (o LookupWebAppSourceControlResultOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *bool { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

// Resource Id.
func (o LookupWebAppSourceControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) string { return v.Id }).(pulumi.StringOutput)
}

// <code>true</code> if this is deployed via GitHub action.
func (o LookupWebAppSourceControlResultOutput) IsGitHubAction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *bool { return v.IsGitHubAction }).(pulumi.BoolPtrOutput)
}

// <code>true</code> to limit to manual integration; <code>false</code> to enable continuous integration (which configures webhooks into online repos like GitHub).
func (o LookupWebAppSourceControlResultOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *bool { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

// <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
func (o LookupWebAppSourceControlResultOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *bool { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o LookupWebAppSourceControlResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppSourceControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) string { return v.Name }).(pulumi.StringOutput)
}

// Repository or source control URL.
func (o LookupWebAppSourceControlResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupWebAppSourceControlResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupWebAppSourceControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSourceControlResultOutput{})
}
