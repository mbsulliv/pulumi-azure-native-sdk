// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Config Server resource
type ConfigServer struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Config Server resource
	Properties ConfigServerPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigServer registers a new resource with the given unique name, arguments, and options.
func NewConfigServer(ctx *pulumi.Context,
	name string, args *ConfigServerArgs, opts ...pulumi.ResourceOption) (*ConfigServer, error) {
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
			Type: pulumi.String("azure-native:appplatform:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230501preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:ConfigServer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigServer
	err := ctx.RegisterResource("azure-native:appplatform/v20231101preview:ConfigServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigServer gets an existing ConfigServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigServerState, opts ...pulumi.ResourceOption) (*ConfigServer, error) {
	var resource ConfigServer
	err := ctx.ReadResource("azure-native:appplatform/v20231101preview:ConfigServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigServer resources.
type configServerState struct {
}

type ConfigServerState struct {
}

func (ConfigServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*configServerState)(nil)).Elem()
}

type configServerArgs struct {
	// Properties of the Config Server resource
	Properties *ConfigServerProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ConfigServer resource.
type ConfigServerArgs struct {
	// Properties of the Config Server resource
	Properties ConfigServerPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (ConfigServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configServerArgs)(nil)).Elem()
}

type ConfigServerInput interface {
	pulumi.Input

	ToConfigServerOutput() ConfigServerOutput
	ToConfigServerOutputWithContext(ctx context.Context) ConfigServerOutput
}

func (*ConfigServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServer)(nil)).Elem()
}

func (i *ConfigServer) ToConfigServerOutput() ConfigServerOutput {
	return i.ToConfigServerOutputWithContext(context.Background())
}

func (i *ConfigServer) ToConfigServerOutputWithContext(ctx context.Context) ConfigServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerOutput)
}

type ConfigServerOutput struct{ *pulumi.OutputState }

func (ConfigServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServer)(nil)).Elem()
}

func (o ConfigServerOutput) ToConfigServerOutput() ConfigServerOutput {
	return o
}

func (o ConfigServerOutput) ToConfigServerOutputWithContext(ctx context.Context) ConfigServerOutput {
	return o
}

// The name of the resource.
func (o ConfigServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the Config Server resource
func (o ConfigServerOutput) Properties() ConfigServerPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigServer) ConfigServerPropertiesResponseOutput { return v.Properties }).(ConfigServerPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ConfigServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigServer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ConfigServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigServerOutput{})
}
