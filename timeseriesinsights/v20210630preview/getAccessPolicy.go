// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the access policy with the specified name in the specified environment.
func LookupAccessPolicy(ctx *pulumi.Context, args *LookupAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessPolicyResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20210630preview:getAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPolicyArgs struct {
	// The name of the Time Series Insights access policy associated with the specified environment.
	AccessPolicyName string `pulumi:"accessPolicyName"`
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An access policy is used to grant users and applications access to the environment. Roles are assigned to service principals in Azure Active Directory. These roles define the actions the principal can perform through the Time Series Insights data plane APIs.
type LookupAccessPolicyResult struct {
	// An description of the access policy.
	Description *string `pulumi:"description"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// The objectId of the principal in Azure Active Directory.
	PrincipalObjectId *string `pulumi:"principalObjectId"`
	// The list of roles the principal is assigned on the environment.
	Roles []string `pulumi:"roles"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupAccessPolicyOutput(ctx *pulumi.Context, args LookupAccessPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessPolicyResult, error) {
			args := v.(LookupAccessPolicyArgs)
			r, err := LookupAccessPolicy(ctx, &args, opts...)
			var s LookupAccessPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessPolicyResultOutput)
}

type LookupAccessPolicyOutputArgs struct {
	// The name of the Time Series Insights access policy associated with the specified environment.
	AccessPolicyName pulumi.StringInput `pulumi:"accessPolicyName"`
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccessPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyArgs)(nil)).Elem()
}

// An access policy is used to grant users and applications access to the environment. Roles are assigned to service principals in Azure Active Directory. These roles define the actions the principal can perform through the Time Series Insights data plane APIs.
type LookupAccessPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyResult)(nil)).Elem()
}

func (o LookupAccessPolicyResultOutput) ToLookupAccessPolicyResultOutput() LookupAccessPolicyResultOutput {
	return o
}

func (o LookupAccessPolicyResultOutput) ToLookupAccessPolicyResultOutputWithContext(ctx context.Context) LookupAccessPolicyResultOutput {
	return o
}

// An description of the access policy.
func (o LookupAccessPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupAccessPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupAccessPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The objectId of the principal in Azure Active Directory.
func (o LookupAccessPolicyResultOutput) PrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) *string { return v.PrincipalObjectId }).(pulumi.StringPtrOutput)
}

// The list of roles the principal is assigned on the environment.
func (o LookupAccessPolicyResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// Resource type
func (o LookupAccessPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPolicyResultOutput{})
}
