// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about the model of a security automation.
// Azure REST API version: 2019-01-01-preview.
func LookupAutomation(ctx *pulumi.Context, args *LookupAutomationArgs, opts ...pulumi.InvokeOption) (*LookupAutomationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAutomationResult
	err := ctx.Invoke("azure-native:security:getAutomation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationArgs struct {
	// The security automation name.
	AutomationName string `pulumi:"automationName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The security automation resource.
type LookupAutomationResult struct {
	// A collection of the actions which are triggered if all the configured rules evaluations, within at least one rule set, are true.
	Actions []interface{} `pulumi:"actions"`
	// The security automation description.
	Description *string `pulumi:"description"`
	// Entity tag is used for comparing two or more entities from the same requested resource.
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id string `pulumi:"id"`
	// Indicates whether the security automation is enabled.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// A collection of scopes on which the security automations logic is applied. Supported scopes are the subscription itself or a resource group under that subscription. The automation will only apply on defined scopes.
	Scopes []AutomationScopeResponse `pulumi:"scopes"`
	// A collection of the source event types which evaluate the security automation set of rules.
	Sources []AutomationSourceResponse `pulumi:"sources"`
	// A list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupAutomationOutput(ctx *pulumi.Context, args LookupAutomationOutputArgs, opts ...pulumi.InvokeOption) LookupAutomationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutomationResult, error) {
			args := v.(LookupAutomationArgs)
			r, err := LookupAutomation(ctx, &args, opts...)
			var s LookupAutomationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutomationResultOutput)
}

type LookupAutomationOutputArgs struct {
	// The security automation name.
	AutomationName pulumi.StringInput `pulumi:"automationName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAutomationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationArgs)(nil)).Elem()
}

// The security automation resource.
type LookupAutomationResultOutput struct{ *pulumi.OutputState }

func (LookupAutomationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationResult)(nil)).Elem()
}

func (o LookupAutomationResultOutput) ToLookupAutomationResultOutput() LookupAutomationResultOutput {
	return o
}

func (o LookupAutomationResultOutput) ToLookupAutomationResultOutputWithContext(ctx context.Context) LookupAutomationResultOutput {
	return o
}

// A collection of the actions which are triggered if all the configured rules evaluations, within at least one rule set, are true.
func (o LookupAutomationResultOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupAutomationResult) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

// The security automation description.
func (o LookupAutomationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Entity tag is used for comparing two or more entities from the same requested resource.
func (o LookupAutomationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupAutomationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whether the security automation is enabled.
func (o LookupAutomationResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// Kind of the resource
func (o LookupAutomationResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o LookupAutomationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupAutomationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationResult) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of scopes on which the security automations logic is applied. Supported scopes are the subscription itself or a resource group under that subscription. The automation will only apply on defined scopes.
func (o LookupAutomationResultOutput) Scopes() AutomationScopeResponseArrayOutput {
	return o.ApplyT(func(v LookupAutomationResult) []AutomationScopeResponse { return v.Scopes }).(AutomationScopeResponseArrayOutput)
}

// A collection of the source event types which evaluate the security automation set of rules.
func (o LookupAutomationResultOutput) Sources() AutomationSourceResponseArrayOutput {
	return o.ApplyT(func(v LookupAutomationResult) []AutomationSourceResponse { return v.Sources }).(AutomationSourceResponseArrayOutput)
}

// A list of key value pairs that describe the resource.
func (o LookupAutomationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAutomationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupAutomationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutomationResultOutput{})
}
