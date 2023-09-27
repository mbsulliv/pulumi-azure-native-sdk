// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an integration account map.
// Azure REST API version: 2019-05-01.
func LookupIntegrationAccountMap(ctx *pulumi.Context, args *LookupIntegrationAccountMapArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountMapResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationAccountMapResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccountMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountMapArgs struct {
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The integration account map name.
	MapName string `pulumi:"mapName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The integration account map.
type LookupIntegrationAccountMapResult struct {
	// The changed time.
	ChangedTime string `pulumi:"changedTime"`
	// The content.
	Content *string `pulumi:"content"`
	// The content link.
	ContentLink ContentLinkResponse `pulumi:"contentLink"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The created time.
	CreatedTime string `pulumi:"createdTime"`
	// The resource id.
	Id string `pulumi:"id"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The map type.
	MapType string `pulumi:"mapType"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// The parameters schema of integration account map.
	ParametersSchema *IntegrationAccountMapPropertiesResponseParametersSchema `pulumi:"parametersSchema"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}

func LookupIntegrationAccountMapOutput(ctx *pulumi.Context, args LookupIntegrationAccountMapOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountMapResult, error) {
			args := v.(LookupIntegrationAccountMapArgs)
			r, err := LookupIntegrationAccountMap(ctx, &args, opts...)
			var s LookupIntegrationAccountMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountMapResultOutput)
}

type LookupIntegrationAccountMapOutputArgs struct {
	// The integration account name.
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	// The integration account map name.
	MapName pulumi.StringInput `pulumi:"mapName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountMapArgs)(nil)).Elem()
}

// The integration account map.
type LookupIntegrationAccountMapResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountMapResult)(nil)).Elem()
}

func (o LookupIntegrationAccountMapResultOutput) ToLookupIntegrationAccountMapResultOutput() LookupIntegrationAccountMapResultOutput {
	return o
}

func (o LookupIntegrationAccountMapResultOutput) ToLookupIntegrationAccountMapResultOutputWithContext(ctx context.Context) LookupIntegrationAccountMapResultOutput {
	return o
}

func (o LookupIntegrationAccountMapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIntegrationAccountMapResult] {
	return pulumix.Output[LookupIntegrationAccountMapResult]{
		OutputState: o.OutputState,
	}
}

// The changed time.
func (o LookupIntegrationAccountMapResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

// The content.
func (o LookupIntegrationAccountMapResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

// The content link.
func (o LookupIntegrationAccountMapResultOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) ContentLinkResponse { return v.ContentLink }).(ContentLinkResponseOutput)
}

// The content type.
func (o LookupIntegrationAccountMapResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The created time.
func (o LookupIntegrationAccountMapResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// The resource id.
func (o LookupIntegrationAccountMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource location.
func (o LookupIntegrationAccountMapResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The map type.
func (o LookupIntegrationAccountMapResultOutput) MapType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.MapType }).(pulumi.StringOutput)
}

// The metadata.
func (o LookupIntegrationAccountMapResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// Gets the resource name.
func (o LookupIntegrationAccountMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.Name }).(pulumi.StringOutput)
}

// The parameters schema of integration account map.
func (o LookupIntegrationAccountMapResultOutput) ParametersSchema() IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) *IntegrationAccountMapPropertiesResponseParametersSchema {
		return v.ParametersSchema
	}).(IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput)
}

// The resource tags.
func (o LookupIntegrationAccountMapResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o LookupIntegrationAccountMapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountMapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountMapResultOutput{})
}
