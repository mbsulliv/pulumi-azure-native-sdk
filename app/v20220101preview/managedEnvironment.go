// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An environment for hosting container apps
type ManagedEnvironment struct {
	pulumi.CustomResourceState

	// Cluster configuration which enables the log daemon to export
	// app logs to a destination. Currently only "log-analytics" is
	// supported
	AppLogsConfiguration AppLogsConfigurationResponsePtrOutput `pulumi:"appLogsConfiguration"`
	// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
	DaprAIInstrumentationKey pulumi.StringPtrOutput `pulumi:"daprAIInstrumentationKey"`
	// Default Domain Name for the cluster
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// Any errors that occurred during deployment or deployment validation
	DeploymentErrors pulumi.StringOutput `pulumi:"deploymentErrors"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Environment.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Static IP of the Environment
	StaticIp pulumi.StringOutput `pulumi:"staticIp"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Vnet configuration for the environment
	VnetConfiguration VnetConfigurationResponsePtrOutput `pulumi:"vnetConfiguration"`
}

// NewManagedEnvironment registers a new resource with the given unique name, arguments, and options.
func NewManagedEnvironment(ctx *pulumi.Context,
	name string, args *ManagedEnvironmentArgs, opts ...pulumi.ResourceOption) (*ManagedEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ManagedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:ManagedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ManagedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ManagedEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedEnvironment
	err := ctx.RegisterResource("azure-native:app/v20220101preview:ManagedEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedEnvironment gets an existing ManagedEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedEnvironmentState, opts ...pulumi.ResourceOption) (*ManagedEnvironment, error) {
	var resource ManagedEnvironment
	err := ctx.ReadResource("azure-native:app/v20220101preview:ManagedEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedEnvironment resources.
type managedEnvironmentState struct {
}

type ManagedEnvironmentState struct {
}

func (ManagedEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedEnvironmentState)(nil)).Elem()
}

type managedEnvironmentArgs struct {
	// Cluster configuration which enables the log daemon to export
	// app logs to a destination. Currently only "log-analytics" is
	// supported
	AppLogsConfiguration *AppLogsConfiguration `pulumi:"appLogsConfiguration"`
	// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
	DaprAIInstrumentationKey *string `pulumi:"daprAIInstrumentationKey"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of the Environment.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Vnet configuration for the environment
	VnetConfiguration *VnetConfiguration `pulumi:"vnetConfiguration"`
}

// The set of arguments for constructing a ManagedEnvironment resource.
type ManagedEnvironmentArgs struct {
	// Cluster configuration which enables the log daemon to export
	// app logs to a destination. Currently only "log-analytics" is
	// supported
	AppLogsConfiguration AppLogsConfigurationPtrInput
	// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
	DaprAIInstrumentationKey pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of the Environment.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Vnet configuration for the environment
	VnetConfiguration VnetConfigurationPtrInput
}

func (ManagedEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedEnvironmentArgs)(nil)).Elem()
}

type ManagedEnvironmentInput interface {
	pulumi.Input

	ToManagedEnvironmentOutput() ManagedEnvironmentOutput
	ToManagedEnvironmentOutputWithContext(ctx context.Context) ManagedEnvironmentOutput
}

func (*ManagedEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironment)(nil)).Elem()
}

func (i *ManagedEnvironment) ToManagedEnvironmentOutput() ManagedEnvironmentOutput {
	return i.ToManagedEnvironmentOutputWithContext(context.Background())
}

func (i *ManagedEnvironment) ToManagedEnvironmentOutputWithContext(ctx context.Context) ManagedEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedEnvironmentOutput)
}

type ManagedEnvironmentOutput struct{ *pulumi.OutputState }

func (ManagedEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironment)(nil)).Elem()
}

func (o ManagedEnvironmentOutput) ToManagedEnvironmentOutput() ManagedEnvironmentOutput {
	return o
}

func (o ManagedEnvironmentOutput) ToManagedEnvironmentOutputWithContext(ctx context.Context) ManagedEnvironmentOutput {
	return o
}

// Cluster configuration which enables the log daemon to export
// app logs to a destination. Currently only "log-analytics" is
// supported
func (o ManagedEnvironmentOutput) AppLogsConfiguration() AppLogsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ManagedEnvironment) AppLogsConfigurationResponsePtrOutput { return v.AppLogsConfiguration }).(AppLogsConfigurationResponsePtrOutput)
}

// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
func (o ManagedEnvironmentOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringPtrOutput { return v.DaprAIInstrumentationKey }).(pulumi.StringPtrOutput)
}

// Default Domain Name for the cluster
func (o ManagedEnvironmentOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

// Any errors that occurred during deployment or deployment validation
func (o ManagedEnvironmentOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.DeploymentErrors }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ManagedEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ManagedEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Environment.
func (o ManagedEnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Static IP of the Environment
func (o ManagedEnvironmentOutput) StaticIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.StaticIp }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ManagedEnvironmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedEnvironment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ManagedEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ManagedEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Vnet configuration for the environment
func (o ManagedEnvironmentOutput) VnetConfiguration() VnetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ManagedEnvironment) VnetConfigurationResponsePtrOutput { return v.VnetConfiguration }).(VnetConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedEnvironmentOutput{})
}
