// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Application Configuration Service resource
type ConfigurationService struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Application Configuration Service properties payload
	Properties ConfigurationServicePropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationService registers a new resource with the given unique name, arguments, and options.
func NewConfigurationService(ctx *pulumi.Context,
	name string, args *ConfigurationServiceArgs, opts ...pulumi.ResourceOption) (*ConfigurationService, error) {
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
			Type: pulumi.String("azure-native:appplatform:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:ConfigurationService"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationService
	err := ctx.RegisterResource("azure-native:appplatform/v20220401:ConfigurationService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationService gets an existing ConfigurationService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationServiceState, opts ...pulumi.ResourceOption) (*ConfigurationService, error) {
	var resource ConfigurationService
	err := ctx.ReadResource("azure-native:appplatform/v20220401:ConfigurationService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationService resources.
type configurationServiceState struct {
}

type ConfigurationServiceState struct {
}

func (ConfigurationServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationServiceState)(nil)).Elem()
}

type configurationServiceArgs struct {
	// The name of Application Configuration Service.
	ConfigurationServiceName *string `pulumi:"configurationServiceName"`
	// Application Configuration Service properties payload
	Properties *ConfigurationServiceProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ConfigurationService resource.
type ConfigurationServiceArgs struct {
	// The name of Application Configuration Service.
	ConfigurationServiceName pulumi.StringPtrInput
	// Application Configuration Service properties payload
	Properties ConfigurationServicePropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (ConfigurationServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationServiceArgs)(nil)).Elem()
}

type ConfigurationServiceInput interface {
	pulumi.Input

	ToConfigurationServiceOutput() ConfigurationServiceOutput
	ToConfigurationServiceOutputWithContext(ctx context.Context) ConfigurationServiceOutput
}

func (*ConfigurationService) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationService)(nil)).Elem()
}

func (i *ConfigurationService) ToConfigurationServiceOutput() ConfigurationServiceOutput {
	return i.ToConfigurationServiceOutputWithContext(context.Background())
}

func (i *ConfigurationService) ToConfigurationServiceOutputWithContext(ctx context.Context) ConfigurationServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceOutput)
}

type ConfigurationServiceOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationService)(nil)).Elem()
}

func (o ConfigurationServiceOutput) ToConfigurationServiceOutput() ConfigurationServiceOutput {
	return o
}

func (o ConfigurationServiceOutput) ToConfigurationServiceOutputWithContext(ctx context.Context) ConfigurationServiceOutput {
	return o
}

// The name of the resource.
func (o ConfigurationServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Application Configuration Service properties payload
func (o ConfigurationServiceOutput) Properties() ConfigurationServicePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationService) ConfigurationServicePropertiesResponseOutput { return v.Properties }).(ConfigurationServicePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ConfigurationServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ConfigurationServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationServiceOutput{})
}
