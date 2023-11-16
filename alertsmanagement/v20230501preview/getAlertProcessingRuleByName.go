// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an alert processing rule by name.
func LookupAlertProcessingRuleByName(ctx *pulumi.Context, args *LookupAlertProcessingRuleByNameArgs, opts ...pulumi.InvokeOption) (*LookupAlertProcessingRuleByNameResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAlertProcessingRuleByNameResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20230501preview:getAlertProcessingRuleByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAlertProcessingRuleByNameArgs struct {
	// The name of the alert processing rule that needs to be fetched.
	AlertProcessingRuleName string `pulumi:"alertProcessingRuleName"`
	// Resource group name where the resource is created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Alert processing rule object containing target scopes, conditions and scheduling logic.
type LookupAlertProcessingRuleByNameResult struct {
	// Azure resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Alert processing rule properties.
	Properties AlertProcessingRulePropertiesResponse `pulumi:"properties"`
	// Alert processing rule system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAlertProcessingRuleByNameResult
func (val *LookupAlertProcessingRuleByNameResult) Defaults() *LookupAlertProcessingRuleByNameResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupAlertProcessingRuleByNameOutput(ctx *pulumi.Context, args LookupAlertProcessingRuleByNameOutputArgs, opts ...pulumi.InvokeOption) LookupAlertProcessingRuleByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlertProcessingRuleByNameResult, error) {
			args := v.(LookupAlertProcessingRuleByNameArgs)
			r, err := LookupAlertProcessingRuleByName(ctx, &args, opts...)
			var s LookupAlertProcessingRuleByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlertProcessingRuleByNameResultOutput)
}

type LookupAlertProcessingRuleByNameOutputArgs struct {
	// The name of the alert processing rule that needs to be fetched.
	AlertProcessingRuleName pulumi.StringInput `pulumi:"alertProcessingRuleName"`
	// Resource group name where the resource is created.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAlertProcessingRuleByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertProcessingRuleByNameArgs)(nil)).Elem()
}

// Alert processing rule object containing target scopes, conditions and scheduling logic.
type LookupAlertProcessingRuleByNameResultOutput struct{ *pulumi.OutputState }

func (LookupAlertProcessingRuleByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertProcessingRuleByNameResult)(nil)).Elem()
}

func (o LookupAlertProcessingRuleByNameResultOutput) ToLookupAlertProcessingRuleByNameResultOutput() LookupAlertProcessingRuleByNameResultOutput {
	return o
}

func (o LookupAlertProcessingRuleByNameResultOutput) ToLookupAlertProcessingRuleByNameResultOutputWithContext(ctx context.Context) LookupAlertProcessingRuleByNameResultOutput {
	return o
}

// Azure resource Id
func (o LookupAlertProcessingRuleByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupAlertProcessingRuleByNameResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o LookupAlertProcessingRuleByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

// Alert processing rule properties.
func (o LookupAlertProcessingRuleByNameResultOutput) Properties() AlertProcessingRulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) AlertProcessingRulePropertiesResponse {
		return v.Properties
	}).(AlertProcessingRulePropertiesResponseOutput)
}

// Alert processing rule system data.
func (o LookupAlertProcessingRuleByNameResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupAlertProcessingRuleByNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o LookupAlertProcessingRuleByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertProcessingRuleByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlertProcessingRuleByNameResultOutput{})
}
