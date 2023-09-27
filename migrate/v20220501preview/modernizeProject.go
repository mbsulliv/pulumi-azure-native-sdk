// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ModernizeProject model.
type ModernizeProject struct {
	pulumi.CustomResourceState

	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// Gets or sets the location of the modernizeProject.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets or sets the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ModernizeProject properties.
	Properties ModernizeProjectModelPropertiesResponseOutput `pulumi:"properties"`
	SystemData ModernizeProjectModelResponseSystemDataOutput `pulumi:"systemData"`
	// Gets or sets the resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewModernizeProject registers a new resource with the given unique name, arguments, and options.
func NewModernizeProject(ctx *pulumi.Context,
	name string, args *ModernizeProjectArgs, opts ...pulumi.ResourceOption) (*ModernizeProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate:ModernizeProject"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ModernizeProject
	err := ctx.RegisterResource("azure-native:migrate/v20220501preview:ModernizeProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModernizeProject gets an existing ModernizeProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModernizeProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModernizeProjectState, opts ...pulumi.ResourceOption) (*ModernizeProject, error) {
	var resource ModernizeProject
	err := ctx.ReadResource("azure-native:migrate/v20220501preview:ModernizeProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModernizeProject resources.
type modernizeProjectState struct {
}

type ModernizeProjectState struct {
}

func (ModernizeProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*modernizeProjectState)(nil)).Elem()
}

type modernizeProjectArgs struct {
	Identity *ResourceIdentity `pulumi:"identity"`
	// Gets or sets the location of the modernizeProject.
	Location *string `pulumi:"location"`
	// ModernizeProject Name.
	ModernizeProjectName *string `pulumi:"modernizeProjectName"`
	// ModernizeProject properties.
	Properties *ModernizeProjectModelProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Subscription Id in which project was created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Gets or sets the resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ModernizeProject resource.
type ModernizeProjectArgs struct {
	Identity ResourceIdentityPtrInput
	// Gets or sets the location of the modernizeProject.
	Location pulumi.StringPtrInput
	// ModernizeProject Name.
	ModernizeProjectName pulumi.StringPtrInput
	// ModernizeProject properties.
	Properties ModernizeProjectModelPropertiesPtrInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
	// Azure Subscription Id in which project was created.
	SubscriptionId pulumi.StringPtrInput
	// Gets or sets the resource tags.
	Tags pulumi.StringMapInput
}

func (ModernizeProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modernizeProjectArgs)(nil)).Elem()
}

type ModernizeProjectInput interface {
	pulumi.Input

	ToModernizeProjectOutput() ModernizeProjectOutput
	ToModernizeProjectOutputWithContext(ctx context.Context) ModernizeProjectOutput
}

func (*ModernizeProject) ElementType() reflect.Type {
	return reflect.TypeOf((**ModernizeProject)(nil)).Elem()
}

func (i *ModernizeProject) ToModernizeProjectOutput() ModernizeProjectOutput {
	return i.ToModernizeProjectOutputWithContext(context.Background())
}

func (i *ModernizeProject) ToModernizeProjectOutputWithContext(ctx context.Context) ModernizeProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModernizeProjectOutput)
}

func (i *ModernizeProject) ToOutput(ctx context.Context) pulumix.Output[*ModernizeProject] {
	return pulumix.Output[*ModernizeProject]{
		OutputState: i.ToModernizeProjectOutputWithContext(ctx).OutputState,
	}
}

type ModernizeProjectOutput struct{ *pulumi.OutputState }

func (ModernizeProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModernizeProject)(nil)).Elem()
}

func (o ModernizeProjectOutput) ToModernizeProjectOutput() ModernizeProjectOutput {
	return o
}

func (o ModernizeProjectOutput) ToModernizeProjectOutputWithContext(ctx context.Context) ModernizeProjectOutput {
	return o
}

func (o ModernizeProjectOutput) ToOutput(ctx context.Context) pulumix.Output[*ModernizeProject] {
	return pulumix.Output[*ModernizeProject]{
		OutputState: o.OutputState,
	}
}

func (o ModernizeProjectOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ModernizeProject) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Gets or sets the location of the modernizeProject.
func (o ModernizeProjectOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModernizeProject) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets the name of the resource.
func (o ModernizeProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ModernizeProject) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ModernizeProject properties.
func (o ModernizeProjectOutput) Properties() ModernizeProjectModelPropertiesResponseOutput {
	return o.ApplyT(func(v *ModernizeProject) ModernizeProjectModelPropertiesResponseOutput { return v.Properties }).(ModernizeProjectModelPropertiesResponseOutput)
}

func (o ModernizeProjectOutput) SystemData() ModernizeProjectModelResponseSystemDataOutput {
	return o.ApplyT(func(v *ModernizeProject) ModernizeProjectModelResponseSystemDataOutput { return v.SystemData }).(ModernizeProjectModelResponseSystemDataOutput)
}

// Gets or sets the resource tags.
func (o ModernizeProjectOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ModernizeProject) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o ModernizeProjectOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ModernizeProject) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ModernizeProjectOutput{})
}
