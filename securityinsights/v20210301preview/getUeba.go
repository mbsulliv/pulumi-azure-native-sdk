// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a setting.
func LookupUeba(ctx *pulumi.Context, args *LookupUebaArgs, opts ...pulumi.InvokeOption) (*LookupUebaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUebaResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getUeba", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUebaArgs struct {
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName string `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Settings with single toggle.
type LookupUebaResult struct {
	// The relevant data sources that enriched by ueba
	DataSources []string `pulumi:"dataSources"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// The kind of the setting
	// Expected value is 'Ueba'.
	Kind string `pulumi:"kind"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource type
	Type string `pulumi:"type"`
}

func LookupUebaOutput(ctx *pulumi.Context, args LookupUebaOutputArgs, opts ...pulumi.InvokeOption) LookupUebaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUebaResult, error) {
			args := v.(LookupUebaArgs)
			r, err := LookupUeba(ctx, &args, opts...)
			var s LookupUebaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUebaResultOutput)
}

type LookupUebaOutputArgs struct {
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName pulumi.StringInput `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupUebaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUebaArgs)(nil)).Elem()
}

// Settings with single toggle.
type LookupUebaResultOutput struct{ *pulumi.OutputState }

func (LookupUebaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUebaResult)(nil)).Elem()
}

func (o LookupUebaResultOutput) ToLookupUebaResultOutput() LookupUebaResultOutput {
	return o
}

func (o LookupUebaResultOutput) ToLookupUebaResultOutputWithContext(ctx context.Context) LookupUebaResultOutput {
	return o
}

// The relevant data sources that enriched by ueba
func (o LookupUebaResultOutput) DataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUebaResult) []string { return v.DataSources }).(pulumi.StringArrayOutput)
}

// Etag of the azure resource
func (o LookupUebaResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUebaResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Azure resource Id
func (o LookupUebaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUebaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the setting
// Expected value is 'Ueba'.
func (o LookupUebaResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUebaResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Azure resource name
func (o LookupUebaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUebaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupUebaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUebaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource type
func (o LookupUebaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUebaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUebaResultOutput{})
}
