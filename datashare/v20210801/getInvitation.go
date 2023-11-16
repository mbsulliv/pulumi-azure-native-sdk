// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an invitation in a share
func LookupInvitation(ctx *pulumi.Context, args *LookupInvitationArgs, opts ...pulumi.InvokeOption) (*LookupInvitationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInvitationResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getInvitation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInvitationArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the invitation.
	InvitationName string `pulumi:"invitationName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// A Invitation data transfer object.
type LookupInvitationResult struct {
	// The expiration date for the invitation and share subscription.
	ExpirationDate *string `pulumi:"expirationDate"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// unique invitation id
	InvitationId string `pulumi:"invitationId"`
	// The status of the invitation.
	InvitationStatus string `pulumi:"invitationStatus"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// The time the recipient responded to the invitation.
	RespondedAt string `pulumi:"respondedAt"`
	// Gets the time at which the invitation was sent.
	SentAt string `pulumi:"sentAt"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The target Azure AD Id. Can't be combined with email.
	TargetActiveDirectoryId *string `pulumi:"targetActiveDirectoryId"`
	// The email the invitation is directed to.
	TargetEmail *string `pulumi:"targetEmail"`
	// The target user or application Id that invitation is being sent to.
	// Must be specified along TargetActiveDirectoryId. This enables sending
	// invitations to specific users or applications in an AD tenant.
	TargetObjectId *string `pulumi:"targetObjectId"`
	// Type of the azure resource
	Type string `pulumi:"type"`
	// Email of the user who created the resource
	UserEmail string `pulumi:"userEmail"`
	// Name of the user who created the resource
	UserName string `pulumi:"userName"`
}

func LookupInvitationOutput(ctx *pulumi.Context, args LookupInvitationOutputArgs, opts ...pulumi.InvokeOption) LookupInvitationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInvitationResult, error) {
			args := v.(LookupInvitationArgs)
			r, err := LookupInvitation(ctx, &args, opts...)
			var s LookupInvitationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInvitationResultOutput)
}

type LookupInvitationOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the invitation.
	InvitationName pulumi.StringInput `pulumi:"invitationName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName pulumi.StringInput `pulumi:"shareName"`
}

func (LookupInvitationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInvitationArgs)(nil)).Elem()
}

// A Invitation data transfer object.
type LookupInvitationResultOutput struct{ *pulumi.OutputState }

func (LookupInvitationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInvitationResult)(nil)).Elem()
}

func (o LookupInvitationResultOutput) ToLookupInvitationResultOutput() LookupInvitationResultOutput {
	return o
}

func (o LookupInvitationResultOutput) ToLookupInvitationResultOutputWithContext(ctx context.Context) LookupInvitationResultOutput {
	return o
}

// The expiration date for the invitation and share subscription.
func (o LookupInvitationResultOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvitationResult) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

// The resource id of the azure resource
func (o LookupInvitationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.Id }).(pulumi.StringOutput)
}

// unique invitation id
func (o LookupInvitationResultOutput) InvitationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.InvitationId }).(pulumi.StringOutput)
}

// The status of the invitation.
func (o LookupInvitationResultOutput) InvitationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.InvitationStatus }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupInvitationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The time the recipient responded to the invitation.
func (o LookupInvitationResultOutput) RespondedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.RespondedAt }).(pulumi.StringOutput)
}

// Gets the time at which the invitation was sent.
func (o LookupInvitationResultOutput) SentAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.SentAt }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupInvitationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInvitationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The target Azure AD Id. Can't be combined with email.
func (o LookupInvitationResultOutput) TargetActiveDirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvitationResult) *string { return v.TargetActiveDirectoryId }).(pulumi.StringPtrOutput)
}

// The email the invitation is directed to.
func (o LookupInvitationResultOutput) TargetEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvitationResult) *string { return v.TargetEmail }).(pulumi.StringPtrOutput)
}

// The target user or application Id that invitation is being sent to.
// Must be specified along TargetActiveDirectoryId. This enables sending
// invitations to specific users or applications in an AD tenant.
func (o LookupInvitationResultOutput) TargetObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvitationResult) *string { return v.TargetObjectId }).(pulumi.StringPtrOutput)
}

// Type of the azure resource
func (o LookupInvitationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Email of the user who created the resource
func (o LookupInvitationResultOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.UserEmail }).(pulumi.StringOutput)
}

// Name of the user who created the resource
func (o LookupInvitationResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInvitationResultOutput{})
}
