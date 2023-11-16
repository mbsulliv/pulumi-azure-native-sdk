// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings about where we should store your security data and logs. If the result is empty, it means that no custom-workspace configuration was set
func LookupWorkspaceSetting(ctx *pulumi.Context, args *LookupWorkspaceSettingArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceSettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceSettingResult
	err := ctx.Invoke("azure-native:security/v20170801preview:getWorkspaceSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceSettingArgs struct {
	// Name of the security setting
	WorkspaceSettingName string `pulumi:"workspaceSettingName"`
}

// Configures where to store the OMS agent data for workspaces under a scope
type LookupWorkspaceSettingResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// All the VMs in this scope will send their security data to the mentioned workspace unless overridden by a setting with more specific scope
	Scope string `pulumi:"scope"`
	// Resource type
	Type string `pulumi:"type"`
	// The full Azure ID of the workspace to save the data in
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupWorkspaceSettingOutput(ctx *pulumi.Context, args LookupWorkspaceSettingOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceSettingResult, error) {
			args := v.(LookupWorkspaceSettingArgs)
			r, err := LookupWorkspaceSetting(ctx, &args, opts...)
			var s LookupWorkspaceSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceSettingResultOutput)
}

type LookupWorkspaceSettingOutputArgs struct {
	// Name of the security setting
	WorkspaceSettingName pulumi.StringInput `pulumi:"workspaceSettingName"`
}

func (LookupWorkspaceSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSettingArgs)(nil)).Elem()
}

// Configures where to store the OMS agent data for workspaces under a scope
type LookupWorkspaceSettingResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSettingResult)(nil)).Elem()
}

func (o LookupWorkspaceSettingResultOutput) ToLookupWorkspaceSettingResultOutput() LookupWorkspaceSettingResultOutput {
	return o
}

func (o LookupWorkspaceSettingResultOutput) ToLookupWorkspaceSettingResultOutputWithContext(ctx context.Context) LookupWorkspaceSettingResultOutput {
	return o
}

// Resource Id
func (o LookupWorkspaceSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupWorkspaceSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// All the VMs in this scope will send their security data to the mentioned workspace unless overridden by a setting with more specific scope
func (o LookupWorkspaceSettingResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Resource type
func (o LookupWorkspaceSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

// The full Azure ID of the workspace to save the data in
func (o LookupWorkspaceSettingResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSettingResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceSettingResultOutput{})
}
