// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package education

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the details for a specific student in the specified lab by student alias
// Azure REST API version: 2021-12-01-preview.
func LookupStudent(ctx *pulumi.Context, args *LookupStudentArgs, opts ...pulumi.InvokeOption) (*LookupStudentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStudentResult
	err := ctx.Invoke("azure-native:education:getStudent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudentArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a billing profile.
	BillingProfileName string `pulumi:"billingProfileName"`
	// The ID that uniquely identifies an invoice section.
	InvoiceSectionName string `pulumi:"invoiceSectionName"`
	// Student alias.
	StudentAlias string `pulumi:"studentAlias"`
}

// Student details.
type LookupStudentResult struct {
	// Student Budget
	Budget AmountResponse `pulumi:"budget"`
	// Date student was added to the lab
	EffectiveDate string `pulumi:"effectiveDate"`
	// Student Email
	Email string `pulumi:"email"`
	// Date this student is set to expire from the lab.
	ExpirationDate string `pulumi:"expirationDate"`
	// First Name
	FirstName string `pulumi:"firstName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Last Name
	LastName string `pulumi:"lastName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Student Role
	Role string `pulumi:"role"`
	// Student Lab Status
	Status string `pulumi:"status"`
	// Subscription alias
	SubscriptionAlias *string `pulumi:"subscriptionAlias"`
	// Subscription Id
	SubscriptionId string `pulumi:"subscriptionId"`
	// subscription invite last sent date
	SubscriptionInviteLastSentDate *string `pulumi:"subscriptionInviteLastSentDate"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupStudentOutput(ctx *pulumi.Context, args LookupStudentOutputArgs, opts ...pulumi.InvokeOption) LookupStudentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStudentResult, error) {
			args := v.(LookupStudentArgs)
			r, err := LookupStudent(ctx, &args, opts...)
			var s LookupStudentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStudentResultOutput)
}

type LookupStudentOutputArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a billing profile.
	BillingProfileName pulumi.StringInput `pulumi:"billingProfileName"`
	// The ID that uniquely identifies an invoice section.
	InvoiceSectionName pulumi.StringInput `pulumi:"invoiceSectionName"`
	// Student alias.
	StudentAlias pulumi.StringInput `pulumi:"studentAlias"`
}

func (LookupStudentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudentArgs)(nil)).Elem()
}

// Student details.
type LookupStudentResultOutput struct{ *pulumi.OutputState }

func (LookupStudentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudentResult)(nil)).Elem()
}

func (o LookupStudentResultOutput) ToLookupStudentResultOutput() LookupStudentResultOutput {
	return o
}

func (o LookupStudentResultOutput) ToLookupStudentResultOutputWithContext(ctx context.Context) LookupStudentResultOutput {
	return o
}

func (o LookupStudentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStudentResult] {
	return pulumix.Output[LookupStudentResult]{
		OutputState: o.OutputState,
	}
}

// Student Budget
func (o LookupStudentResultOutput) Budget() AmountResponseOutput {
	return o.ApplyT(func(v LookupStudentResult) AmountResponse { return v.Budget }).(AmountResponseOutput)
}

// Date student was added to the lab
func (o LookupStudentResultOutput) EffectiveDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.EffectiveDate }).(pulumi.StringOutput)
}

// Student Email
func (o LookupStudentResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Email }).(pulumi.StringOutput)
}

// Date this student is set to expire from the lab.
func (o LookupStudentResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

// First Name
func (o LookupStudentResultOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.FirstName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupStudentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Last Name
func (o LookupStudentResultOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.LastName }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupStudentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Student Role
func (o LookupStudentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Role }).(pulumi.StringOutput)
}

// Student Lab Status
func (o LookupStudentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Status }).(pulumi.StringOutput)
}

// Subscription alias
func (o LookupStudentResultOutput) SubscriptionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudentResult) *string { return v.SubscriptionAlias }).(pulumi.StringPtrOutput)
}

// Subscription Id
func (o LookupStudentResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// subscription invite last sent date
func (o LookupStudentResultOutput) SubscriptionInviteLastSentDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudentResult) *string { return v.SubscriptionInviteLastSentDate }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupStudentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStudentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStudentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStudentResultOutput{})
}
