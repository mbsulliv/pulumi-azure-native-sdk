// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201120

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Application Insights workbook template definition.
type WorkbookTemplate struct {
	pulumi.CustomResourceState

	// Information about the author of the workbook template.
	Author pulumi.StringPtrOutput `pulumi:"author"`
	// Workbook galleries supported by the template.
	Galleries WorkbookTemplateGalleryResponseArrayOutput `pulumi:"galleries"`
	// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized WorkbookTemplateLocalizedGalleryResponseArrayMapOutput `pulumi:"localized"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Valid JSON object containing workbook template payload.
	TemplateData pulumi.AnyOutput `pulumi:"templateData"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkbookTemplate registers a new resource with the given unique name, arguments, and options.
func NewWorkbookTemplate(ctx *pulumi.Context,
	name string, args *WorkbookTemplateArgs, opts ...pulumi.ResourceOption) (*WorkbookTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Galleries == nil {
		return nil, errors.New("invalid value for required argument 'Galleries'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TemplateData == nil {
		return nil, errors.New("invalid value for required argument 'TemplateData'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:WorkbookTemplate"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20191017preview:WorkbookTemplate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkbookTemplate
	err := ctx.RegisterResource("azure-native:insights/v20201120:WorkbookTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkbookTemplate gets an existing WorkbookTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkbookTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookTemplateState, opts ...pulumi.ResourceOption) (*WorkbookTemplate, error) {
	var resource WorkbookTemplate
	err := ctx.ReadResource("azure-native:insights/v20201120:WorkbookTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkbookTemplate resources.
type workbookTemplateState struct {
}

type WorkbookTemplateState struct {
}

func (WorkbookTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookTemplateState)(nil)).Elem()
}

type workbookTemplateArgs struct {
	// Information about the author of the workbook template.
	Author *string `pulumi:"author"`
	// Workbook galleries supported by the template.
	Galleries []WorkbookTemplateGallery `pulumi:"galleries"`
	// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized map[string][]WorkbookTemplateLocalizedGallery `pulumi:"localized"`
	// Resource location
	Location *string `pulumi:"location"`
	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
	Priority *int `pulumi:"priority"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Valid JSON object containing workbook template payload.
	TemplateData interface{} `pulumi:"templateData"`
}

// The set of arguments for constructing a WorkbookTemplate resource.
type WorkbookTemplateArgs struct {
	// Information about the author of the workbook template.
	Author pulumi.StringPtrInput
	// Workbook galleries supported by the template.
	Galleries WorkbookTemplateGalleryArrayInput
	// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized WorkbookTemplateLocalizedGalleryArrayMapInput
	// Resource location
	Location pulumi.StringPtrInput
	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
	Priority pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Valid JSON object containing workbook template payload.
	TemplateData pulumi.Input
}

func (WorkbookTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookTemplateArgs)(nil)).Elem()
}

type WorkbookTemplateInput interface {
	pulumi.Input

	ToWorkbookTemplateOutput() WorkbookTemplateOutput
	ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput
}

func (*WorkbookTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookTemplate)(nil)).Elem()
}

func (i *WorkbookTemplate) ToWorkbookTemplateOutput() WorkbookTemplateOutput {
	return i.ToWorkbookTemplateOutputWithContext(context.Background())
}

func (i *WorkbookTemplate) ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateOutput)
}

type WorkbookTemplateOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookTemplate)(nil)).Elem()
}

func (o WorkbookTemplateOutput) ToWorkbookTemplateOutput() WorkbookTemplateOutput {
	return o
}

func (o WorkbookTemplateOutput) ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput {
	return o
}

// Information about the author of the workbook template.
func (o WorkbookTemplateOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookTemplate) pulumi.StringPtrOutput { return v.Author }).(pulumi.StringPtrOutput)
}

// Workbook galleries supported by the template.
func (o WorkbookTemplateOutput) Galleries() WorkbookTemplateGalleryResponseArrayOutput {
	return o.ApplyT(func(v *WorkbookTemplate) WorkbookTemplateGalleryResponseArrayOutput { return v.Galleries }).(WorkbookTemplateGalleryResponseArrayOutput)
}

// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
func (o WorkbookTemplateOutput) Localized() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o.ApplyT(func(v *WorkbookTemplate) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput { return v.Localized }).(WorkbookTemplateLocalizedGalleryResponseArrayMapOutput)
}

// Resource location
func (o WorkbookTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbookTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name.
func (o WorkbookTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbookTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
func (o WorkbookTemplateOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkbookTemplate) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// Resource tags
func (o WorkbookTemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkbookTemplate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Valid JSON object containing workbook template payload.
func (o WorkbookTemplateOutput) TemplateData() pulumi.AnyOutput {
	return o.ApplyT(func(v *WorkbookTemplate) pulumi.AnyOutput { return v.TemplateData }).(pulumi.AnyOutput)
}

// Azure resource type
func (o WorkbookTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkbookTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkbookTemplateOutput{})
}
