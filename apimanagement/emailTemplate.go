// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Email Template details.
// Azure REST API version: 2022-08-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
type EmailTemplate struct {
	pulumi.CustomResourceState

	// Email Template Body. This should be a valid XDocument
	Body pulumi.StringOutput `pulumi:"body"`
	// Description of the Email Template.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the template is the default template provided by API Management or has been edited.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Email Template Parameter values.
	Parameters EmailTemplateParametersContractPropertiesResponseArrayOutput `pulumi:"parameters"`
	// Subject of the Template.
	Subject pulumi.StringOutput `pulumi:"subject"`
	// Title of the Template.
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEmailTemplate registers a new resource with the given unique name, arguments, and options.
func NewEmailTemplate(ctx *pulumi.Context,
	name string, args *EmailTemplateArgs, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:EmailTemplate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EmailTemplate
	err := ctx.RegisterResource("azure-native:apimanagement:EmailTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailTemplate gets an existing EmailTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailTemplateState, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	var resource EmailTemplate
	err := ctx.ReadResource("azure-native:apimanagement:EmailTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailTemplate resources.
type emailTemplateState struct {
}

type EmailTemplateState struct {
}

func (EmailTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateState)(nil)).Elem()
}

type emailTemplateArgs struct {
	// Email Template Body. This should be a valid XDocument
	Body *string `pulumi:"body"`
	// Description of the Email Template.
	Description *string `pulumi:"description"`
	// Email Template Parameter values.
	Parameters []EmailTemplateParametersContractProperties `pulumi:"parameters"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subject of the Template.
	Subject *string `pulumi:"subject"`
	// Email Template Name Identifier.
	TemplateName *string `pulumi:"templateName"`
	// Title of the Template.
	Title *string `pulumi:"title"`
}

// The set of arguments for constructing a EmailTemplate resource.
type EmailTemplateArgs struct {
	// Email Template Body. This should be a valid XDocument
	Body pulumi.StringPtrInput
	// Description of the Email Template.
	Description pulumi.StringPtrInput
	// Email Template Parameter values.
	Parameters EmailTemplateParametersContractPropertiesArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Subject of the Template.
	Subject pulumi.StringPtrInput
	// Email Template Name Identifier.
	TemplateName pulumi.StringPtrInput
	// Title of the Template.
	Title pulumi.StringPtrInput
}

func (EmailTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateArgs)(nil)).Elem()
}

type EmailTemplateInput interface {
	pulumi.Input

	ToEmailTemplateOutput() EmailTemplateOutput
	ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput
}

func (*EmailTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailTemplate)(nil)).Elem()
}

func (i *EmailTemplate) ToEmailTemplateOutput() EmailTemplateOutput {
	return i.ToEmailTemplateOutputWithContext(context.Background())
}

func (i *EmailTemplate) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateOutput)
}

func (i *EmailTemplate) ToOutput(ctx context.Context) pulumix.Output[*EmailTemplate] {
	return pulumix.Output[*EmailTemplate]{
		OutputState: i.ToEmailTemplateOutputWithContext(ctx).OutputState,
	}
}

type EmailTemplateOutput struct{ *pulumi.OutputState }

func (EmailTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailTemplate)(nil)).Elem()
}

func (o EmailTemplateOutput) ToEmailTemplateOutput() EmailTemplateOutput {
	return o
}

func (o EmailTemplateOutput) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return o
}

func (o EmailTemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*EmailTemplate] {
	return pulumix.Output[*EmailTemplate]{
		OutputState: o.OutputState,
	}
}

// Email Template Body. This should be a valid XDocument
func (o EmailTemplateOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Body }).(pulumi.StringOutput)
}

// Description of the Email Template.
func (o EmailTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the template is the default template provided by API Management or has been edited.
func (o EmailTemplateOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// The name of the resource
func (o EmailTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Email Template Parameter values.
func (o EmailTemplateOutput) Parameters() EmailTemplateParametersContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *EmailTemplate) EmailTemplateParametersContractPropertiesResponseArrayOutput {
		return v.Parameters
	}).(EmailTemplateParametersContractPropertiesResponseArrayOutput)
}

// Subject of the Template.
func (o EmailTemplateOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

// Title of the Template.
func (o EmailTemplateOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EmailTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EmailTemplateOutput{})
}
