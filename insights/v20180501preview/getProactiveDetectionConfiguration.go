// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the ProactiveDetection configuration for this configuration id.
func LookupProactiveDetectionConfiguration(ctx *pulumi.Context, args *LookupProactiveDetectionConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupProactiveDetectionConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProactiveDetectionConfigurationResult
	err := ctx.Invoke("azure-native:insights/v20180501preview:getProactiveDetectionConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProactiveDetectionConfigurationArgs struct {
	// The ProactiveDetection configuration ID. This is unique within a Application Insights component.
	ConfigurationId string `pulumi:"configurationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
}

// A ProactiveDetection configuration definition.
type LookupProactiveDetectionConfigurationResult struct {
	// Custom email addresses for this rule notifications
	CustomEmails []string `pulumi:"customEmails"`
	// A flag that indicates whether this rule is enabled by the user
	Enabled *bool `pulumi:"enabled"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// The last time this rule was updated
	LastUpdatedTime string `pulumi:"lastUpdatedTime"`
	// Resource location
	Location *string `pulumi:"location"`
	// The rule name
	Name string `pulumi:"name"`
	// Static definitions of the ProactiveDetection configuration rule (same values for all components).
	RuleDefinitions *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions `pulumi:"ruleDefinitions"`
	// A flag that indicated whether notifications on this rule should be sent to subscription owners
	SendEmailsToSubscriptionOwners *bool `pulumi:"sendEmailsToSubscriptionOwners"`
	// Azure resource type
	Type string `pulumi:"type"`
}

func LookupProactiveDetectionConfigurationOutput(ctx *pulumi.Context, args LookupProactiveDetectionConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupProactiveDetectionConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProactiveDetectionConfigurationResult, error) {
			args := v.(LookupProactiveDetectionConfigurationArgs)
			r, err := LookupProactiveDetectionConfiguration(ctx, &args, opts...)
			var s LookupProactiveDetectionConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProactiveDetectionConfigurationResultOutput)
}

type LookupProactiveDetectionConfigurationOutputArgs struct {
	// The ProactiveDetection configuration ID. This is unique within a Application Insights component.
	ConfigurationId pulumi.StringInput `pulumi:"configurationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupProactiveDetectionConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProactiveDetectionConfigurationArgs)(nil)).Elem()
}

// A ProactiveDetection configuration definition.
type LookupProactiveDetectionConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupProactiveDetectionConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProactiveDetectionConfigurationResult)(nil)).Elem()
}

func (o LookupProactiveDetectionConfigurationResultOutput) ToLookupProactiveDetectionConfigurationResultOutput() LookupProactiveDetectionConfigurationResultOutput {
	return o
}

func (o LookupProactiveDetectionConfigurationResultOutput) ToLookupProactiveDetectionConfigurationResultOutputWithContext(ctx context.Context) LookupProactiveDetectionConfigurationResultOutput {
	return o
}

// Custom email addresses for this rule notifications
func (o LookupProactiveDetectionConfigurationResultOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) []string { return v.CustomEmails }).(pulumi.StringArrayOutput)
}

// A flag that indicates whether this rule is enabled by the user
func (o LookupProactiveDetectionConfigurationResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Azure resource Id
func (o LookupProactiveDetectionConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last time this rule was updated
func (o LookupProactiveDetectionConfigurationResultOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// Resource location
func (o LookupProactiveDetectionConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The rule name
func (o LookupProactiveDetectionConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Static definitions of the ProactiveDetection configuration rule (same values for all components).
func (o LookupProactiveDetectionConfigurationResultOutput) RuleDefinitions() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions {
		return v.RuleDefinitions
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput)
}

// A flag that indicated whether notifications on this rule should be sent to subscription owners
func (o LookupProactiveDetectionConfigurationResultOutput) SendEmailsToSubscriptionOwners() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *bool { return v.SendEmailsToSubscriptionOwners }).(pulumi.BoolPtrOutput)
}

// Azure resource type
func (o LookupProactiveDetectionConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProactiveDetectionConfigurationResultOutput{})
}
