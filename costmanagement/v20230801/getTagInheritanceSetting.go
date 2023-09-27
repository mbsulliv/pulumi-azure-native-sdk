// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the setting from the given scope by name.
func LookupTagInheritanceSetting(ctx *pulumi.Context, args *LookupTagInheritanceSettingArgs, opts ...pulumi.InvokeOption) (*LookupTagInheritanceSettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTagInheritanceSettingResult
	err := ctx.Invoke("azure-native:costmanagement/v20230801:getTagInheritanceSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagInheritanceSettingArgs struct {
	// The scope associated with this setting. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billing profile scope.
	Scope string `pulumi:"scope"`
	// Setting type.
	Type string `pulumi:"type"`
}

// Tag Inheritance Setting definition.
type LookupTagInheritanceSettingResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Specifies the kind of settings.
	// Expected value is 'taginheritance'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The properties of the tag inheritance setting.
	Properties TagInheritancePropertiesResponse `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTagInheritanceSettingOutput(ctx *pulumi.Context, args LookupTagInheritanceSettingOutputArgs, opts ...pulumi.InvokeOption) LookupTagInheritanceSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagInheritanceSettingResult, error) {
			args := v.(LookupTagInheritanceSettingArgs)
			r, err := LookupTagInheritanceSetting(ctx, &args, opts...)
			var s LookupTagInheritanceSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagInheritanceSettingResultOutput)
}

type LookupTagInheritanceSettingOutputArgs struct {
	// The scope associated with this setting. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billing profile scope.
	Scope pulumi.StringInput `pulumi:"scope"`
	// Setting type.
	Type pulumi.StringInput `pulumi:"type"`
}

func (LookupTagInheritanceSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagInheritanceSettingArgs)(nil)).Elem()
}

// Tag Inheritance Setting definition.
type LookupTagInheritanceSettingResultOutput struct{ *pulumi.OutputState }

func (LookupTagInheritanceSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagInheritanceSettingResult)(nil)).Elem()
}

func (o LookupTagInheritanceSettingResultOutput) ToLookupTagInheritanceSettingResultOutput() LookupTagInheritanceSettingResultOutput {
	return o
}

func (o LookupTagInheritanceSettingResultOutput) ToLookupTagInheritanceSettingResultOutputWithContext(ctx context.Context) LookupTagInheritanceSettingResultOutput {
	return o
}

func (o LookupTagInheritanceSettingResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTagInheritanceSettingResult] {
	return pulumix.Output[LookupTagInheritanceSettingResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTagInheritanceSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the kind of settings.
// Expected value is 'taginheritance'.
func (o LookupTagInheritanceSettingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTagInheritanceSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of the tag inheritance setting.
func (o LookupTagInheritanceSettingResultOutput) Properties() TagInheritancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) TagInheritancePropertiesResponse { return v.Properties }).(TagInheritancePropertiesResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTagInheritanceSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagInheritanceSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagInheritanceSettingResultOutput{})
}
