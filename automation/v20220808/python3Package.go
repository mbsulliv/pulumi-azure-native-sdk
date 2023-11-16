// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220808

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the module type.
type Python3Package struct {
	pulumi.CustomResourceState

	// Gets or sets the activity count of the module.
	ActivityCount pulumi.IntPtrOutput `pulumi:"activityCount"`
	// Gets or sets the contentLink of the module.
	ContentLink ContentLinkResponsePtrOutput `pulumi:"contentLink"`
	// Gets or sets the creation time.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets or sets the error info of the module.
	Error ModuleErrorInfoResponsePtrOutput `pulumi:"error"`
	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets or sets type of module, if its composite or not.
	IsComposite pulumi.BoolPtrOutput `pulumi:"isComposite"`
	// Gets or sets the isGlobal flag of the module.
	IsGlobal pulumi.BoolPtrOutput `pulumi:"isGlobal"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the provisioning state of the module.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets the size in bytes of the module.
	SizeInBytes pulumi.Float64PtrOutput `pulumi:"sizeInBytes"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets the version of the module.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewPython3Package registers a new resource with the given unique name, arguments, and options.
func NewPython3Package(ctx *pulumi.Context,
	name string, args *Python3PackageArgs, opts ...pulumi.ResourceOption) (*Python3Package, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ContentLink == nil {
		return nil, errors.New("invalid value for required argument 'ContentLink'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Python3Package"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:Python3Package"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20231101:Python3Package"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Python3Package
	err := ctx.RegisterResource("azure-native:automation/v20220808:Python3Package", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPython3Package gets an existing Python3Package resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPython3Package(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Python3PackageState, opts ...pulumi.ResourceOption) (*Python3Package, error) {
	var resource Python3Package
	err := ctx.ReadResource("azure-native:automation/v20220808:Python3Package", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Python3Package resources.
type python3PackageState struct {
}

type Python3PackageState struct {
}

func (Python3PackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*python3PackageState)(nil)).Elem()
}

type python3PackageArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the module content link.
	ContentLink ContentLink `pulumi:"contentLink"`
	// The name of python package.
	PackageName *string `pulumi:"packageName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Python3Package resource.
type Python3PackageArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the module content link.
	ContentLink ContentLinkInput
	// The name of python package.
	PackageName pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (Python3PackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*python3PackageArgs)(nil)).Elem()
}

type Python3PackageInput interface {
	pulumi.Input

	ToPython3PackageOutput() Python3PackageOutput
	ToPython3PackageOutputWithContext(ctx context.Context) Python3PackageOutput
}

func (*Python3Package) ElementType() reflect.Type {
	return reflect.TypeOf((**Python3Package)(nil)).Elem()
}

func (i *Python3Package) ToPython3PackageOutput() Python3PackageOutput {
	return i.ToPython3PackageOutputWithContext(context.Background())
}

func (i *Python3Package) ToPython3PackageOutputWithContext(ctx context.Context) Python3PackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Python3PackageOutput)
}

type Python3PackageOutput struct{ *pulumi.OutputState }

func (Python3PackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Python3Package)(nil)).Elem()
}

func (o Python3PackageOutput) ToPython3PackageOutput() Python3PackageOutput {
	return o
}

func (o Python3PackageOutput) ToPython3PackageOutputWithContext(ctx context.Context) Python3PackageOutput {
	return o
}

// Gets or sets the activity count of the module.
func (o Python3PackageOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.IntPtrOutput { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

// Gets or sets the contentLink of the module.
func (o Python3PackageOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Python3Package) ContentLinkResponsePtrOutput { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

// Gets or sets the creation time.
func (o Python3PackageOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o Python3PackageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the error info of the module.
func (o Python3PackageOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v *Python3Package) ModuleErrorInfoResponsePtrOutput { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

// Gets or sets the etag of the resource.
func (o Python3PackageOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Gets or sets type of module, if its composite or not.
func (o Python3PackageOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.BoolPtrOutput { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

// Gets or sets the isGlobal flag of the module.
func (o Python3PackageOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.BoolPtrOutput { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o Python3PackageOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The Azure Region where the resource lives
func (o Python3PackageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o Python3PackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state of the module.
func (o Python3PackageOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets or sets the size in bytes of the module.
func (o Python3PackageOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.Float64PtrOutput { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

// Resource tags.
func (o Python3PackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o Python3PackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the version of the module.
func (o Python3PackageOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(Python3PackageOutput{})
}
