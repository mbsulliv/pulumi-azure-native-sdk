// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the properties of a lab user.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2018-10-15, 2023-06-07.
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:labservices:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	// The name of the lab that uniquely identifies it within containing lab plan. Used in resource URIs.
	LabName string `pulumi:"labName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the user that uniquely identifies it within containing lab. Used in resource URIs.
	UserName string `pulumi:"userName"`
}

// User of a lab that can register for and use virtual machines within the lab.
type LookupUserResult struct {
	// The amount of usage quota time the user gets in addition to the lab usage quota.
	AdditionalUsageQuota *string `pulumi:"additionalUsageQuota"`
	// Display name of the user, for example user's full name.
	DisplayName string `pulumi:"displayName"`
	// Email address of the user.
	Email string `pulumi:"email"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Date and time when the invitation message was sent to the user.
	InvitationSent string `pulumi:"invitationSent"`
	// State of the invitation message for the user.
	InvitationState string `pulumi:"invitationState"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Current provisioning state of the user resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// State of the user's registration within the lab.
	RegistrationState string `pulumi:"registrationState"`
	// Metadata pertaining to creation and last modification of the user resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// How long the user has used their virtual machines in this lab.
	TotalUsage string `pulumi:"totalUsage"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

type LookupUserOutputArgs struct {
	// The name of the lab that uniquely identifies it within containing lab plan. Used in resource URIs.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the user that uniquely identifies it within containing lab. Used in resource URIs.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// User of a lab that can register for and use virtual machines within the lab.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupUserResult] {
	return pulumix.Output[LookupUserResult]{
		OutputState: o.OutputState,
	}
}

// The amount of usage quota time the user gets in addition to the lab usage quota.
func (o LookupUserResultOutput) AdditionalUsageQuota() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.AdditionalUsageQuota }).(pulumi.StringPtrOutput)
}

// Display name of the user, for example user's full name.
func (o LookupUserResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Email address of the user.
func (o LookupUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Date and time when the invitation message was sent to the user.
func (o LookupUserResultOutput) InvitationSent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.InvitationSent }).(pulumi.StringOutput)
}

// State of the invitation message for the user.
func (o LookupUserResultOutput) InvitationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.InvitationState }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Current provisioning state of the user resource.
func (o LookupUserResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// State of the user's registration within the lab.
func (o LookupUserResultOutput) RegistrationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.RegistrationState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the user resource.
func (o LookupUserResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUserResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// How long the user has used their virtual machines in this lab.
func (o LookupUserResultOutput) TotalUsage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.TotalUsage }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
