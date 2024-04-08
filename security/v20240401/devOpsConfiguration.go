// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DevOps Configuration resource.
type DevOpsConfiguration struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// DevOps Configuration properties.
	Properties DevOpsConfigurationPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDevOpsConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDevOpsConfiguration(ctx *pulumi.Context,
	name string, args *DevOpsConfigurationArgs, opts ...pulumi.ResourceOption) (*DevOpsConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecurityConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'SecurityConnectorName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:DevOpsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:security/v20230901preview:DevOpsConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DevOpsConfiguration
	err := ctx.RegisterResource("azure-native:security/v20240401:DevOpsConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevOpsConfiguration gets an existing DevOpsConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevOpsConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevOpsConfigurationState, opts ...pulumi.ResourceOption) (*DevOpsConfiguration, error) {
	var resource DevOpsConfiguration
	err := ctx.ReadResource("azure-native:security/v20240401:DevOpsConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DevOpsConfiguration resources.
type devOpsConfigurationState struct {
}

type DevOpsConfigurationState struct {
}

func (DevOpsConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*devOpsConfigurationState)(nil)).Elem()
}

type devOpsConfigurationArgs struct {
	// DevOps Configuration properties.
	Properties *DevOpsConfigurationProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}

// The set of arguments for constructing a DevOpsConfiguration resource.
type DevOpsConfigurationArgs struct {
	// DevOps Configuration properties.
	Properties DevOpsConfigurationPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The security connector name.
	SecurityConnectorName pulumi.StringInput
}

func (DevOpsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devOpsConfigurationArgs)(nil)).Elem()
}

type DevOpsConfigurationInput interface {
	pulumi.Input

	ToDevOpsConfigurationOutput() DevOpsConfigurationOutput
	ToDevOpsConfigurationOutputWithContext(ctx context.Context) DevOpsConfigurationOutput
}

func (*DevOpsConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**DevOpsConfiguration)(nil)).Elem()
}

func (i *DevOpsConfiguration) ToDevOpsConfigurationOutput() DevOpsConfigurationOutput {
	return i.ToDevOpsConfigurationOutputWithContext(context.Background())
}

func (i *DevOpsConfiguration) ToDevOpsConfigurationOutputWithContext(ctx context.Context) DevOpsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevOpsConfigurationOutput)
}

type DevOpsConfigurationOutput struct{ *pulumi.OutputState }

func (DevOpsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevOpsConfiguration)(nil)).Elem()
}

func (o DevOpsConfigurationOutput) ToDevOpsConfigurationOutput() DevOpsConfigurationOutput {
	return o
}

func (o DevOpsConfigurationOutput) ToDevOpsConfigurationOutputWithContext(ctx context.Context) DevOpsConfigurationOutput {
	return o
}

// The name of the resource
func (o DevOpsConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevOpsConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// DevOps Configuration properties.
func (o DevOpsConfigurationOutput) Properties() DevOpsConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v *DevOpsConfiguration) DevOpsConfigurationPropertiesResponseOutput { return v.Properties }).(DevOpsConfigurationPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DevOpsConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DevOpsConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DevOpsConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DevOpsConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DevOpsConfigurationOutput{})
}
