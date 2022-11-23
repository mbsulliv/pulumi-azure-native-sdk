// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Linker of source and target resource
type Connector struct {
	pulumi.CustomResourceState

	// The authentication type.
	AuthInfo pulumi.AnyOutput `pulumi:"authInfo"`
	// The application client type
	ClientType pulumi.StringPtrOutput `pulumi:"clientType"`
	// The connection information consumed by applications, including secrets, connection strings.
	ConfigurationInfo ConfigurationInfoResponsePtrOutput `pulumi:"configurationInfo"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The network solution.
	PublicNetworkSolution PublicNetworkSolutionResponsePtrOutput `pulumi:"publicNetworkSolution"`
	// connection scope in source service.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// An option to store secret value in secure place
	SecretStore SecretStoreResponsePtrOutput `pulumi:"secretStore"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The target service properties
	TargetService pulumi.AnyOutput `pulumi:"targetService"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The VNet solution.
	VNetSolution VNetSolutionResponsePtrOutput `pulumi:"vNetSolution"`
}

// NewConnector registers a new resource with the given unique name, arguments, and options.
func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Connector
	err := ctx.RegisterResource("azure-native:servicelinker/v20221101preview:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnector gets an existing Connector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("azure-native:servicelinker/v20221101preview:Connector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connector resources.
type connectorState struct {
}

type ConnectorState struct {
}

func (ConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorState)(nil)).Elem()
}

type connectorArgs struct {
	// The authentication type.
	AuthInfo interface{} `pulumi:"authInfo"`
	// The application client type
	ClientType *string `pulumi:"clientType"`
	// The connection information consumed by applications, including secrets, connection strings.
	ConfigurationInfo *ConfigurationInfo `pulumi:"configurationInfo"`
	// The name of resource.
	ConnectorName *string `pulumi:"connectorName"`
	// The name of Azure region.
	Location string `pulumi:"location"`
	// The network solution.
	PublicNetworkSolution *PublicNetworkSolution `pulumi:"publicNetworkSolution"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// connection scope in source service.
	Scope *string `pulumi:"scope"`
	// An option to store secret value in secure place
	SecretStore *SecretStore `pulumi:"secretStore"`
	// The target service properties
	TargetService interface{} `pulumi:"targetService"`
	// The VNet solution.
	VNetSolution *VNetSolution `pulumi:"vNetSolution"`
}

// The set of arguments for constructing a Connector resource.
type ConnectorArgs struct {
	// The authentication type.
	AuthInfo pulumi.Input
	// The application client type
	ClientType pulumi.StringPtrInput
	// The connection information consumed by applications, including secrets, connection strings.
	ConfigurationInfo ConfigurationInfoPtrInput
	// The name of resource.
	ConnectorName pulumi.StringPtrInput
	// The name of Azure region.
	Location pulumi.StringInput
	// The network solution.
	PublicNetworkSolution PublicNetworkSolutionPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// connection scope in source service.
	Scope pulumi.StringPtrInput
	// An option to store secret value in secure place
	SecretStore SecretStorePtrInput
	// The target service properties
	TargetService pulumi.Input
	// The VNet solution.
	VNetSolution VNetSolutionPtrInput
}

func (ConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorArgs)(nil)).Elem()
}

type ConnectorInput interface {
	pulumi.Input

	ToConnectorOutput() ConnectorOutput
	ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput
}

func (*Connector) ElementType() reflect.Type {
	return reflect.TypeOf((**Connector)(nil)).Elem()
}

func (i *Connector) ToConnectorOutput() ConnectorOutput {
	return i.ToConnectorOutputWithContext(context.Background())
}

func (i *Connector) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorOutput)
}

type ConnectorOutput struct{ *pulumi.OutputState }

func (ConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connector)(nil)).Elem()
}

func (o ConnectorOutput) ToConnectorOutput() ConnectorOutput {
	return o
}

func (o ConnectorOutput) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return o
}

// The authentication type.
func (o ConnectorOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v *Connector) pulumi.AnyOutput { return v.AuthInfo }).(pulumi.AnyOutput)
}

// The application client type
func (o ConnectorOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringPtrOutput { return v.ClientType }).(pulumi.StringPtrOutput)
}

// The connection information consumed by applications, including secrets, connection strings.
func (o ConnectorOutput) ConfigurationInfo() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v *Connector) ConfigurationInfoResponsePtrOutput { return v.ConfigurationInfo }).(ConfigurationInfoResponsePtrOutput)
}

// The name of the resource
func (o ConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o ConnectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The network solution.
func (o ConnectorOutput) PublicNetworkSolution() PublicNetworkSolutionResponsePtrOutput {
	return o.ApplyT(func(v *Connector) PublicNetworkSolutionResponsePtrOutput { return v.PublicNetworkSolution }).(PublicNetworkSolutionResponsePtrOutput)
}

// connection scope in source service.
func (o ConnectorOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// An option to store secret value in secure place
func (o ConnectorOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v *Connector) SecretStoreResponsePtrOutput { return v.SecretStore }).(SecretStoreResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Connector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The target service properties
func (o ConnectorOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v *Connector) pulumi.AnyOutput { return v.TargetService }).(pulumi.AnyOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The VNet solution.
func (o ConnectorOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v *Connector) VNetSolutionResponsePtrOutput { return v.VNetSolution }).(VNetSolutionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectorOutput{})
}
