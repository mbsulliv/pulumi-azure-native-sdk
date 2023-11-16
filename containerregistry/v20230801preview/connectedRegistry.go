// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents a connected registry for a container registry.
type ConnectedRegistry struct {
	pulumi.CustomResourceState

	// The activation properties of the connected registry.
	Activation ActivationPropertiesResponseOutput `pulumi:"activation"`
	// The list of the ACR token resource IDs used to authenticate clients to the connected registry.
	ClientTokenIds pulumi.StringArrayOutput `pulumi:"clientTokenIds"`
	// The current connection state of the connected registry.
	ConnectionState pulumi.StringOutput `pulumi:"connectionState"`
	// The last activity time of the connected registry.
	LastActivityTime pulumi.StringOutput `pulumi:"lastActivityTime"`
	// The logging properties of the connected registry.
	Logging LoggingPropertiesResponsePtrOutput `pulumi:"logging"`
	// The login server properties of the connected registry.
	LoginServer LoginServerPropertiesResponsePtrOutput `pulumi:"loginServer"`
	// The mode of the connected registry resource that indicates the permissions of the registry.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of notifications subscription information for the connected registry.
	NotificationsList pulumi.StringArrayOutput `pulumi:"notificationsList"`
	// The parent of the connected registry.
	Parent ParentPropertiesResponseOutput `pulumi:"parent"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The list of current statuses of the connected registry.
	StatusDetails StatusDetailPropertiesResponseArrayOutput `pulumi:"statusDetails"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The current version of ACR runtime on the connected registry.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewConnectedRegistry registers a new resource with the given unique name, arguments, and options.
func NewConnectedRegistry(ctx *pulumi.Context,
	name string, args *ConnectedRegistryArgs, opts ...pulumi.ResourceOption) (*ConnectedRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Logging != nil {
		args.Logging = args.Logging.ToLoggingPropertiesPtrOutput().ApplyT(func(v *LoggingProperties) *LoggingProperties { return v.Defaults() }).(LoggingPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230101preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:ConnectedRegistry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20231101preview:ConnectedRegistry"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectedRegistry
	err := ctx.RegisterResource("azure-native:containerregistry/v20230801preview:ConnectedRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectedRegistry gets an existing ConnectedRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectedRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedRegistryState, opts ...pulumi.ResourceOption) (*ConnectedRegistry, error) {
	var resource ConnectedRegistry
	err := ctx.ReadResource("azure-native:containerregistry/v20230801preview:ConnectedRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectedRegistry resources.
type connectedRegistryState struct {
}

type ConnectedRegistryState struct {
}

func (ConnectedRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedRegistryState)(nil)).Elem()
}

type connectedRegistryArgs struct {
	// The list of the ACR token resource IDs used to authenticate clients to the connected registry.
	ClientTokenIds []string `pulumi:"clientTokenIds"`
	// The name of the connected registry.
	ConnectedRegistryName *string `pulumi:"connectedRegistryName"`
	// The logging properties of the connected registry.
	Logging *LoggingProperties `pulumi:"logging"`
	// The mode of the connected registry resource that indicates the permissions of the registry.
	Mode string `pulumi:"mode"`
	// The list of notifications subscription information for the connected registry.
	NotificationsList []string `pulumi:"notificationsList"`
	// The parent of the connected registry.
	Parent ParentProperties `pulumi:"parent"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ConnectedRegistry resource.
type ConnectedRegistryArgs struct {
	// The list of the ACR token resource IDs used to authenticate clients to the connected registry.
	ClientTokenIds pulumi.StringArrayInput
	// The name of the connected registry.
	ConnectedRegistryName pulumi.StringPtrInput
	// The logging properties of the connected registry.
	Logging LoggingPropertiesPtrInput
	// The mode of the connected registry resource that indicates the permissions of the registry.
	Mode pulumi.StringInput
	// The list of notifications subscription information for the connected registry.
	NotificationsList pulumi.StringArrayInput
	// The parent of the connected registry.
	Parent ParentPropertiesInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ConnectedRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedRegistryArgs)(nil)).Elem()
}

type ConnectedRegistryInput interface {
	pulumi.Input

	ToConnectedRegistryOutput() ConnectedRegistryOutput
	ToConnectedRegistryOutputWithContext(ctx context.Context) ConnectedRegistryOutput
}

func (*ConnectedRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedRegistry)(nil)).Elem()
}

func (i *ConnectedRegistry) ToConnectedRegistryOutput() ConnectedRegistryOutput {
	return i.ToConnectedRegistryOutputWithContext(context.Background())
}

func (i *ConnectedRegistry) ToConnectedRegistryOutputWithContext(ctx context.Context) ConnectedRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedRegistryOutput)
}

type ConnectedRegistryOutput struct{ *pulumi.OutputState }

func (ConnectedRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedRegistry)(nil)).Elem()
}

func (o ConnectedRegistryOutput) ToConnectedRegistryOutput() ConnectedRegistryOutput {
	return o
}

func (o ConnectedRegistryOutput) ToConnectedRegistryOutputWithContext(ctx context.Context) ConnectedRegistryOutput {
	return o
}

// The activation properties of the connected registry.
func (o ConnectedRegistryOutput) Activation() ActivationPropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectedRegistry) ActivationPropertiesResponseOutput { return v.Activation }).(ActivationPropertiesResponseOutput)
}

// The list of the ACR token resource IDs used to authenticate clients to the connected registry.
func (o ConnectedRegistryOutput) ClientTokenIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringArrayOutput { return v.ClientTokenIds }).(pulumi.StringArrayOutput)
}

// The current connection state of the connected registry.
func (o ConnectedRegistryOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.ConnectionState }).(pulumi.StringOutput)
}

// The last activity time of the connected registry.
func (o ConnectedRegistryOutput) LastActivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.LastActivityTime }).(pulumi.StringOutput)
}

// The logging properties of the connected registry.
func (o ConnectedRegistryOutput) Logging() LoggingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ConnectedRegistry) LoggingPropertiesResponsePtrOutput { return v.Logging }).(LoggingPropertiesResponsePtrOutput)
}

// The login server properties of the connected registry.
func (o ConnectedRegistryOutput) LoginServer() LoginServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ConnectedRegistry) LoginServerPropertiesResponsePtrOutput { return v.LoginServer }).(LoginServerPropertiesResponsePtrOutput)
}

// The mode of the connected registry resource that indicates the permissions of the registry.
func (o ConnectedRegistryOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// The name of the resource.
func (o ConnectedRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of notifications subscription information for the connected registry.
func (o ConnectedRegistryOutput) NotificationsList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringArrayOutput { return v.NotificationsList }).(pulumi.StringArrayOutput)
}

// The parent of the connected registry.
func (o ConnectedRegistryOutput) Parent() ParentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectedRegistry) ParentPropertiesResponseOutput { return v.Parent }).(ParentPropertiesResponseOutput)
}

// Provisioning state of the resource.
func (o ConnectedRegistryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of current statuses of the connected registry.
func (o ConnectedRegistryOutput) StatusDetails() StatusDetailPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *ConnectedRegistry) StatusDetailPropertiesResponseArrayOutput { return v.StatusDetails }).(StatusDetailPropertiesResponseArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ConnectedRegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedRegistry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ConnectedRegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The current version of ACR runtime on the connected registry.
func (o ConnectedRegistryOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedRegistry) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedRegistryOutput{})
}
