// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a blueprint definition.
// Azure REST API version: 2018-11-01-preview.
//
// Other available API versions: 2017-11-11-preview.
func LookupBlueprint(ctx *pulumi.Context, args *LookupBlueprintArgs, opts ...pulumi.InvokeOption) (*LookupBlueprintResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBlueprintResult
	err := ctx.Invoke("azure-native:blueprint:getBlueprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlueprintArgs struct {
	// Name of the blueprint definition.
	BlueprintName string `pulumi:"blueprintName"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope string `pulumi:"resourceScope"`
}

// Represents a Blueprint definition.
type LookupBlueprintResult struct {
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Layout view of the blueprint definition for UI reference.
	Layout interface{} `pulumi:"layout"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Parameters required by this blueprint definition.
	Parameters map[string]ParameterDefinitionResponse `pulumi:"parameters"`
	// Resource group placeholders defined by this blueprint definition.
	ResourceGroups map[string]ResourceGroupDefinitionResponse `pulumi:"resourceGroups"`
	// Status of the blueprint. This field is readonly.
	Status BlueprintStatusResponse `pulumi:"status"`
	// The scope where this blueprint definition can be assigned.
	TargetScope string `pulumi:"targetScope"`
	// Type of this resource.
	Type string `pulumi:"type"`
	// Published versions of this blueprint definition.
	Versions interface{} `pulumi:"versions"`
}

func LookupBlueprintOutput(ctx *pulumi.Context, args LookupBlueprintOutputArgs, opts ...pulumi.InvokeOption) LookupBlueprintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlueprintResult, error) {
			args := v.(LookupBlueprintArgs)
			r, err := LookupBlueprint(ctx, &args, opts...)
			var s LookupBlueprintResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlueprintResultOutput)
}

type LookupBlueprintOutputArgs struct {
	// Name of the blueprint definition.
	BlueprintName pulumi.StringInput `pulumi:"blueprintName"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope pulumi.StringInput `pulumi:"resourceScope"`
}

func (LookupBlueprintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlueprintArgs)(nil)).Elem()
}

// Represents a Blueprint definition.
type LookupBlueprintResultOutput struct{ *pulumi.OutputState }

func (LookupBlueprintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlueprintResult)(nil)).Elem()
}

func (o LookupBlueprintResultOutput) ToLookupBlueprintResultOutput() LookupBlueprintResultOutput {
	return o
}

func (o LookupBlueprintResultOutput) ToLookupBlueprintResultOutputWithContext(ctx context.Context) LookupBlueprintResultOutput {
	return o
}

// Multi-line explain this resource.
func (o LookupBlueprintResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlueprintResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// One-liner string explain this resource.
func (o LookupBlueprintResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlueprintResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// String Id used to locate any resource on Azure.
func (o LookupBlueprintResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Id }).(pulumi.StringOutput)
}

// Layout view of the blueprint definition for UI reference.
func (o LookupBlueprintResultOutput) Layout() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBlueprintResult) interface{} { return v.Layout }).(pulumi.AnyOutput)
}

// Name of this resource.
func (o LookupBlueprintResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Name }).(pulumi.StringOutput)
}

// Parameters required by this blueprint definition.
func (o LookupBlueprintResultOutput) Parameters() ParameterDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupBlueprintResult) map[string]ParameterDefinitionResponse { return v.Parameters }).(ParameterDefinitionResponseMapOutput)
}

// Resource group placeholders defined by this blueprint definition.
func (o LookupBlueprintResultOutput) ResourceGroups() ResourceGroupDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupBlueprintResult) map[string]ResourceGroupDefinitionResponse { return v.ResourceGroups }).(ResourceGroupDefinitionResponseMapOutput)
}

// Status of the blueprint. This field is readonly.
func (o LookupBlueprintResultOutput) Status() BlueprintStatusResponseOutput {
	return o.ApplyT(func(v LookupBlueprintResult) BlueprintStatusResponse { return v.Status }).(BlueprintStatusResponseOutput)
}

// The scope where this blueprint definition can be assigned.
func (o LookupBlueprintResultOutput) TargetScope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.TargetScope }).(pulumi.StringOutput)
}

// Type of this resource.
func (o LookupBlueprintResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Type }).(pulumi.StringOutput)
}

// Published versions of this blueprint definition.
func (o LookupBlueprintResultOutput) Versions() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBlueprintResult) interface{} { return v.Versions }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlueprintResultOutput{})
}
