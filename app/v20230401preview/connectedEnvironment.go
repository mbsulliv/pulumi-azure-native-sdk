// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An environment for Kubernetes cluster specialized for web workloads by Azure App Service
type ConnectedEnvironment struct {
	pulumi.CustomResourceState

	// Custom domain configuration for the environment
	CustomDomainConfiguration CustomDomainConfigurationResponsePtrOutput `pulumi:"customDomainConfiguration"`
	// Application Insights connection string used by Dapr to export Service to Service communication telemetry
	DaprAIConnectionString pulumi.StringPtrOutput `pulumi:"daprAIConnectionString"`
	// Default Domain Name for the cluster
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// Any errors that occurred during deployment or deployment validation
	DeploymentErrors pulumi.StringOutput `pulumi:"deploymentErrors"`
	// The complex type of the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Kubernetes Environment.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Static IP of the connectedEnvironment
	StaticIp pulumi.StringPtrOutput `pulumi:"staticIp"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectedEnvironment registers a new resource with the given unique name, arguments, and options.
func NewConnectedEnvironment(ctx *pulumi.Context,
	name string, args *ConnectedEnvironmentArgs, opts ...pulumi.ResourceOption) (*ConnectedEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ConnectedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ConnectedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ConnectedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:ConnectedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230501:ConnectedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230502preview:ConnectedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230801preview:ConnectedEnvironment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectedEnvironment
	err := ctx.RegisterResource("azure-native:app/v20230401preview:ConnectedEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectedEnvironment gets an existing ConnectedEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectedEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedEnvironmentState, opts ...pulumi.ResourceOption) (*ConnectedEnvironment, error) {
	var resource ConnectedEnvironment
	err := ctx.ReadResource("azure-native:app/v20230401preview:ConnectedEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectedEnvironment resources.
type connectedEnvironmentState struct {
}

type ConnectedEnvironmentState struct {
}

func (ConnectedEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentState)(nil)).Elem()
}

type connectedEnvironmentArgs struct {
	// Name of the connectedEnvironment.
	ConnectedEnvironmentName *string `pulumi:"connectedEnvironmentName"`
	// Custom domain configuration for the environment
	CustomDomainConfiguration *CustomDomainConfiguration `pulumi:"customDomainConfiguration"`
	// Application Insights connection string used by Dapr to export Service to Service communication telemetry
	DaprAIConnectionString *string `pulumi:"daprAIConnectionString"`
	// The complex type of the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Static IP of the connectedEnvironment
	StaticIp *string `pulumi:"staticIp"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConnectedEnvironment resource.
type ConnectedEnvironmentArgs struct {
	// Name of the connectedEnvironment.
	ConnectedEnvironmentName pulumi.StringPtrInput
	// Custom domain configuration for the environment
	CustomDomainConfiguration CustomDomainConfigurationPtrInput
	// Application Insights connection string used by Dapr to export Service to Service communication telemetry
	DaprAIConnectionString pulumi.StringPtrInput
	// The complex type of the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Static IP of the connectedEnvironment
	StaticIp pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConnectedEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentArgs)(nil)).Elem()
}

type ConnectedEnvironmentInput interface {
	pulumi.Input

	ToConnectedEnvironmentOutput() ConnectedEnvironmentOutput
	ToConnectedEnvironmentOutputWithContext(ctx context.Context) ConnectedEnvironmentOutput
}

func (*ConnectedEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironment)(nil)).Elem()
}

func (i *ConnectedEnvironment) ToConnectedEnvironmentOutput() ConnectedEnvironmentOutput {
	return i.ToConnectedEnvironmentOutputWithContext(context.Background())
}

func (i *ConnectedEnvironment) ToConnectedEnvironmentOutputWithContext(ctx context.Context) ConnectedEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedEnvironmentOutput)
}

type ConnectedEnvironmentOutput struct{ *pulumi.OutputState }

func (ConnectedEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironment)(nil)).Elem()
}

func (o ConnectedEnvironmentOutput) ToConnectedEnvironmentOutput() ConnectedEnvironmentOutput {
	return o
}

func (o ConnectedEnvironmentOutput) ToConnectedEnvironmentOutputWithContext(ctx context.Context) ConnectedEnvironmentOutput {
	return o
}

// Custom domain configuration for the environment
func (o ConnectedEnvironmentOutput) CustomDomainConfiguration() CustomDomainConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) CustomDomainConfigurationResponsePtrOutput {
		return v.CustomDomainConfiguration
	}).(CustomDomainConfigurationResponsePtrOutput)
}

// Application Insights connection string used by Dapr to export Service to Service communication telemetry
func (o ConnectedEnvironmentOutput) DaprAIConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringPtrOutput { return v.DaprAIConnectionString }).(pulumi.StringPtrOutput)
}

// Default Domain Name for the cluster
func (o ConnectedEnvironmentOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

// Any errors that occurred during deployment or deployment validation
func (o ConnectedEnvironmentOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.DeploymentErrors }).(pulumi.StringOutput)
}

// The complex type of the extended location.
func (o ConnectedEnvironmentOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ConnectedEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConnectedEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Kubernetes Environment.
func (o ConnectedEnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Static IP of the connectedEnvironment
func (o ConnectedEnvironmentOutput) StaticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringPtrOutput { return v.StaticIp }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConnectedEnvironmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ConnectedEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConnectedEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedEnvironmentOutput{})
}
