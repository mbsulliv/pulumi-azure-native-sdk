// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the user specified by its identifier.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2016-07-07, 2016-10-10, 2017-03-01, 2018-01-01, 2022-09-01-preview, 2023-03-01-preview.
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:apimanagement:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupUserArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId string `pulumi:"userId"`
}

// User details.
type LookupUserResult struct {
	// Email address.
	Email *string `pulumi:"email"`
	// First name.
	FirstName *string `pulumi:"firstName"`
	// Collection of groups user is part of.
	Groups []GroupContractPropertiesResponse `pulumi:"groups"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Collection of user identities.
	Identities []UserIdentityContractResponse `pulumi:"identities"`
	// Last name.
	LastName *string `pulumi:"lastName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Optional note about a user set by the administrator.
	Note *string `pulumi:"note"`
	// Date of user registration. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	RegistrationDate *string `pulumi:"registrationDate"`
	// Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
	State *string `pulumi:"state"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupUserResult
func (val *LookupUserResult) Defaults() *LookupUserResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.State == nil {
		state_ := "active"
		tmp.State = &state_
	}
	return &tmp
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
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// User details.
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

// Email address.
func (o LookupUserResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// First name.
func (o LookupUserResultOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

// Collection of groups user is part of.
func (o LookupUserResultOutput) Groups() GroupContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []GroupContractPropertiesResponse { return v.Groups }).(GroupContractPropertiesResponseArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Collection of user identities.
func (o LookupUserResultOutput) Identities() UserIdentityContractResponseArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []UserIdentityContractResponse { return v.Identities }).(UserIdentityContractResponseArrayOutput)
}

// Last name.
func (o LookupUserResultOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional note about a user set by the administrator.
func (o LookupUserResultOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Note }).(pulumi.StringPtrOutput)
}

// Date of user registration. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupUserResultOutput) RegistrationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.RegistrationDate }).(pulumi.StringPtrOutput)
}

// Account state. Specifies whether the user is active or not. Blocked users are unable to sign into the developer portal or call any APIs of subscribed products. Default state is Active.
func (o LookupUserResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
