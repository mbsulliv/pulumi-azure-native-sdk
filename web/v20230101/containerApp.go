// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App.
type ContainerApp struct {
	pulumi.CustomResourceState

	// Non versioned Container App configuration properties.
	Configuration ConfigurationResponsePtrOutput `pulumi:"configuration"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource ID of the Container App's KubeEnvironment.
	KubeEnvironmentId pulumi.StringPtrOutput `pulumi:"kubeEnvironmentId"`
	// Fully Qualified Domain Name of the latest revision of the Container App.
	LatestRevisionFqdn pulumi.StringOutput `pulumi:"latestRevisionFqdn"`
	// Name of the latest revision of the Container App.
	LatestRevisionName pulumi.StringOutput `pulumi:"latestRevisionName"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Container App.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Container App versioned application definition.
	Template TemplateResponsePtrOutput `pulumi:"template"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContainerApp registers a new resource with the given unique name, arguments, and options.
func NewContainerApp(ctx *pulumi.Context,
	name string, args *ContainerAppArgs, opts ...pulumi.ResourceOption) (*ContainerApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Configuration != nil {
		args.Configuration = args.Configuration.ToConfigurationPtrOutput().ApplyT(func(v *Configuration) *Configuration { return v.Defaults() }).(ConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:ContainerApp"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ContainerApp
	err := ctx.RegisterResource("azure-native:web/v20230101:ContainerApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerApp gets an existing ContainerApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAppState, opts ...pulumi.ResourceOption) (*ContainerApp, error) {
	var resource ContainerApp
	err := ctx.ReadResource("azure-native:web/v20230101:ContainerApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerApp resources.
type containerAppState struct {
}

type ContainerAppState struct {
}

func (ContainerAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppState)(nil)).Elem()
}

type containerAppArgs struct {
	// Non versioned Container App configuration properties.
	Configuration *Configuration `pulumi:"configuration"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource ID of the Container App's KubeEnvironment.
	KubeEnvironmentId *string `pulumi:"kubeEnvironmentId"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Name of the Container App.
	Name *string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Container App versioned application definition.
	Template *Template `pulumi:"template"`
}

// The set of arguments for constructing a ContainerApp resource.
type ContainerAppArgs struct {
	// Non versioned Container App configuration properties.
	Configuration ConfigurationPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource ID of the Container App's KubeEnvironment.
	KubeEnvironmentId pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Name of the Container App.
	Name pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Container App versioned application definition.
	Template TemplatePtrInput
}

func (ContainerAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppArgs)(nil)).Elem()
}

type ContainerAppInput interface {
	pulumi.Input

	ToContainerAppOutput() ContainerAppOutput
	ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput
}

func (*ContainerApp) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerApp)(nil)).Elem()
}

func (i *ContainerApp) ToContainerAppOutput() ContainerAppOutput {
	return i.ToContainerAppOutputWithContext(context.Background())
}

func (i *ContainerApp) ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppOutput)
}

type ContainerAppOutput struct{ *pulumi.OutputState }

func (ContainerAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerApp)(nil)).Elem()
}

func (o ContainerAppOutput) ToContainerAppOutput() ContainerAppOutput {
	return o
}

func (o ContainerAppOutput) ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput {
	return o
}

// Non versioned Container App configuration properties.
func (o ContainerAppOutput) Configuration() ConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) ConfigurationResponsePtrOutput { return v.Configuration }).(ConfigurationResponsePtrOutput)
}

// Kind of resource.
func (o ContainerAppOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource ID of the Container App's KubeEnvironment.
func (o ContainerAppOutput) KubeEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringPtrOutput { return v.KubeEnvironmentId }).(pulumi.StringPtrOutput)
}

// Fully Qualified Domain Name of the latest revision of the Container App.
func (o ContainerAppOutput) LatestRevisionFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.LatestRevisionFqdn }).(pulumi.StringOutput)
}

// Name of the latest revision of the Container App.
func (o ContainerAppOutput) LatestRevisionName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.LatestRevisionName }).(pulumi.StringOutput)
}

// Resource Location.
func (o ContainerAppOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource Name.
func (o ContainerAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Container App.
func (o ContainerAppOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o ContainerAppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Container App versioned application definition.
func (o ContainerAppOutput) Template() TemplateResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) TemplateResponsePtrOutput { return v.Template }).(TemplateResponsePtrOutput)
}

// Resource type.
func (o ContainerAppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerAppOutput{})
}