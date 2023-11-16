// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the security admin configuration
type SecurityAdminConfiguration struct {
	pulumi.CustomResourceState

	// Enum list of network intent policy based services.
	ApplyOnNetworkIntentPolicyBasedServices pulumi.StringArrayOutput `pulumi:"applyOnNetworkIntentPolicyBasedServices"`
	// A description of the security configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecurityAdminConfiguration registers a new resource with the given unique name, arguments, and options.
func NewSecurityAdminConfiguration(ctx *pulumi.Context,
	name string, args *SecurityAdminConfigurationArgs, opts ...pulumi.ResourceOption) (*SecurityAdminConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:SecurityAdminConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:SecurityAdminConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecurityAdminConfiguration
	err := ctx.RegisterResource("azure-native:network/v20230601:SecurityAdminConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityAdminConfiguration gets an existing SecurityAdminConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityAdminConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityAdminConfigurationState, opts ...pulumi.ResourceOption) (*SecurityAdminConfiguration, error) {
	var resource SecurityAdminConfiguration
	err := ctx.ReadResource("azure-native:network/v20230601:SecurityAdminConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityAdminConfiguration resources.
type securityAdminConfigurationState struct {
}

type SecurityAdminConfigurationState struct {
}

func (SecurityAdminConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityAdminConfigurationState)(nil)).Elem()
}

type securityAdminConfigurationArgs struct {
	// Enum list of network intent policy based services.
	ApplyOnNetworkIntentPolicyBasedServices []string `pulumi:"applyOnNetworkIntentPolicyBasedServices"`
	// The name of the network manager Security Configuration.
	ConfigurationName *string `pulumi:"configurationName"`
	// A description of the security configuration.
	Description *string `pulumi:"description"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a SecurityAdminConfiguration resource.
type SecurityAdminConfigurationArgs struct {
	// Enum list of network intent policy based services.
	ApplyOnNetworkIntentPolicyBasedServices pulumi.StringArrayInput
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringPtrInput
	// A description of the security configuration.
	Description pulumi.StringPtrInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (SecurityAdminConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityAdminConfigurationArgs)(nil)).Elem()
}

type SecurityAdminConfigurationInput interface {
	pulumi.Input

	ToSecurityAdminConfigurationOutput() SecurityAdminConfigurationOutput
	ToSecurityAdminConfigurationOutputWithContext(ctx context.Context) SecurityAdminConfigurationOutput
}

func (*SecurityAdminConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAdminConfiguration)(nil)).Elem()
}

func (i *SecurityAdminConfiguration) ToSecurityAdminConfigurationOutput() SecurityAdminConfigurationOutput {
	return i.ToSecurityAdminConfigurationOutputWithContext(context.Background())
}

func (i *SecurityAdminConfiguration) ToSecurityAdminConfigurationOutputWithContext(ctx context.Context) SecurityAdminConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAdminConfigurationOutput)
}

type SecurityAdminConfigurationOutput struct{ *pulumi.OutputState }

func (SecurityAdminConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAdminConfiguration)(nil)).Elem()
}

func (o SecurityAdminConfigurationOutput) ToSecurityAdminConfigurationOutput() SecurityAdminConfigurationOutput {
	return o
}

func (o SecurityAdminConfigurationOutput) ToSecurityAdminConfigurationOutputWithContext(ctx context.Context) SecurityAdminConfigurationOutput {
	return o
}

// Enum list of network intent policy based services.
func (o SecurityAdminConfigurationOutput) ApplyOnNetworkIntentPolicyBasedServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAdminConfiguration) pulumi.StringArrayOutput {
		return v.ApplyOnNetworkIntentPolicyBasedServices
	}).(pulumi.StringArrayOutput)
}

// A description of the security configuration.
func (o SecurityAdminConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAdminConfiguration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o SecurityAdminConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityAdminConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource name.
func (o SecurityAdminConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityAdminConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o SecurityAdminConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityAdminConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o SecurityAdminConfigurationOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityAdminConfiguration) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o SecurityAdminConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SecurityAdminConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o SecurityAdminConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityAdminConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityAdminConfigurationOutput{})
}
