// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the email template specified by its identifier.
func LookupEmailTemplate(ctx *pulumi.Context, args *LookupEmailTemplateArgs, opts ...pulumi.InvokeOption) (*LookupEmailTemplateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEmailTemplateResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getEmailTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEmailTemplateArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Email Template Name Identifier.
	TemplateName string `pulumi:"templateName"`
}

// Email Template details.
type LookupEmailTemplateResult struct {
	// Email Template Body. This should be a valid XDocument
	Body string `pulumi:"body"`
	// Description of the Email Template.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Whether the template is the default template provided by API Management or has been edited.
	IsDefault bool `pulumi:"isDefault"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Email Template Parameter values.
	Parameters []EmailTemplateParametersContractPropertiesResponse `pulumi:"parameters"`
	// Subject of the Template.
	Subject string `pulumi:"subject"`
	// Title of the Template.
	Title *string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEmailTemplateOutput(ctx *pulumi.Context, args LookupEmailTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupEmailTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEmailTemplateResult, error) {
			args := v.(LookupEmailTemplateArgs)
			r, err := LookupEmailTemplate(ctx, &args, opts...)
			var s LookupEmailTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEmailTemplateResultOutput)
}

type LookupEmailTemplateOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Email Template Name Identifier.
	TemplateName pulumi.StringInput `pulumi:"templateName"`
}

func (LookupEmailTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailTemplateArgs)(nil)).Elem()
}

// Email Template details.
type LookupEmailTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupEmailTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailTemplateResult)(nil)).Elem()
}

func (o LookupEmailTemplateResultOutput) ToLookupEmailTemplateResultOutput() LookupEmailTemplateResultOutput {
	return o
}

func (o LookupEmailTemplateResultOutput) ToLookupEmailTemplateResultOutputWithContext(ctx context.Context) LookupEmailTemplateResultOutput {
	return o
}

// Email Template Body. This should be a valid XDocument
func (o LookupEmailTemplateResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Body }).(pulumi.StringOutput)
}

// Description of the Email Template.
func (o LookupEmailTemplateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEmailTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether the template is the default template provided by API Management or has been edited.
func (o LookupEmailTemplateResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// The name of the resource
func (o LookupEmailTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Email Template Parameter values.
func (o LookupEmailTemplateResultOutput) Parameters() EmailTemplateParametersContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) []EmailTemplateParametersContractPropertiesResponse {
		return v.Parameters
	}).(EmailTemplateParametersContractPropertiesResponseArrayOutput)
}

// Subject of the Template.
func (o LookupEmailTemplateResultOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Subject }).(pulumi.StringOutput)
}

// Title of the Template.
func (o LookupEmailTemplateResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEmailTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEmailTemplateResultOutput{})
}
