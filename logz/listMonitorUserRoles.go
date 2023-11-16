// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response for list of user's role for Logz.io account.
// Azure REST API version: 2022-01-01-preview.
func ListMonitorUserRoles(ctx *pulumi.Context, args *ListMonitorUserRolesArgs, opts ...pulumi.InvokeOption) (*ListMonitorUserRolesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListMonitorUserRolesResult
	err := ctx.Invoke("azure-native:logz:listMonitorUserRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorUserRolesArgs struct {
	// Email of the user used by Logz for contacting them if needed
	EmailAddress *string `pulumi:"emailAddress"`
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for list of user's role for Logz.io account.
type ListMonitorUserRolesResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// List of user roles for Logz.io account.
	Value []UserRoleResponseResponse `pulumi:"value"`
}

func ListMonitorUserRolesOutput(ctx *pulumi.Context, args ListMonitorUserRolesOutputArgs, opts ...pulumi.InvokeOption) ListMonitorUserRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorUserRolesResult, error) {
			args := v.(ListMonitorUserRolesArgs)
			r, err := ListMonitorUserRoles(ctx, &args, opts...)
			var s ListMonitorUserRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorUserRolesResultOutput)
}

type ListMonitorUserRolesOutputArgs struct {
	// Email of the user used by Logz for contacting them if needed
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitorUserRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorUserRolesArgs)(nil)).Elem()
}

// Response for list of user's role for Logz.io account.
type ListMonitorUserRolesResultOutput struct{ *pulumi.OutputState }

func (ListMonitorUserRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorUserRolesResult)(nil)).Elem()
}

func (o ListMonitorUserRolesResultOutput) ToListMonitorUserRolesResultOutput() ListMonitorUserRolesResultOutput {
	return o
}

func (o ListMonitorUserRolesResultOutput) ToListMonitorUserRolesResultOutputWithContext(ctx context.Context) ListMonitorUserRolesResultOutput {
	return o
}

// Link to the next set of results, if any.
func (o ListMonitorUserRolesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorUserRolesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of user roles for Logz.io account.
func (o ListMonitorUserRolesResultOutput) Value() UserRoleResponseResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorUserRolesResult) []UserRoleResponseResponse { return v.Value }).(UserRoleResponseResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorUserRolesResultOutput{})
}
