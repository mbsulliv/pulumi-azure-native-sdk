// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Connector definition for kind 'Customizable'.
type CustomizableConnectorDefinition struct {
	pulumi.CustomResourceState

	// The UiConfig for 'Customizable' connector definition kind.
	ConnectionsConfig CustomizableConnectionsConfigResponsePtrOutput `pulumi:"connectionsConfig"`
	// The UiConfig for 'Customizable' connector definition kind.
	ConnectorUiConfig CustomizableConnectorUiConfigResponseOutput `pulumi:"connectorUiConfig"`
	// Gets or sets the connector definition created date in UTC format.
	CreatedTimeUtc pulumi.StringPtrOutput `pulumi:"createdTimeUtc"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector definitions
	// Expected value is 'Customizable'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Gets or sets the connector definition last modified date in UTC format.
	LastModifiedUtc pulumi.StringPtrOutput `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomizableConnectorDefinition registers a new resource with the given unique name, arguments, and options.
func NewCustomizableConnectorDefinition(ctx *pulumi.Context,
	name string, args *CustomizableConnectorDefinitionArgs, opts ...pulumi.ResourceOption) (*CustomizableConnectorDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectorUiConfig == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorUiConfig'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Customizable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:CustomizableConnectorDefinition"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:CustomizableConnectorDefinition"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:CustomizableConnectorDefinition"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:CustomizableConnectorDefinition"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:CustomizableConnectorDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomizableConnectorDefinition
	err := ctx.RegisterResource("azure-native:securityinsights/v20230901preview:CustomizableConnectorDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomizableConnectorDefinition gets an existing CustomizableConnectorDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomizableConnectorDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomizableConnectorDefinitionState, opts ...pulumi.ResourceOption) (*CustomizableConnectorDefinition, error) {
	var resource CustomizableConnectorDefinition
	err := ctx.ReadResource("azure-native:securityinsights/v20230901preview:CustomizableConnectorDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomizableConnectorDefinition resources.
type customizableConnectorDefinitionState struct {
}

type CustomizableConnectorDefinitionState struct {
}

func (CustomizableConnectorDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*customizableConnectorDefinitionState)(nil)).Elem()
}

type customizableConnectorDefinitionArgs struct {
	// The UiConfig for 'Customizable' connector definition kind.
	ConnectionsConfig *CustomizableConnectionsConfig `pulumi:"connectionsConfig"`
	// The UiConfig for 'Customizable' connector definition kind.
	ConnectorUiConfig CustomizableConnectorUiConfig `pulumi:"connectorUiConfig"`
	// Gets or sets the connector definition created date in UTC format.
	CreatedTimeUtc *string `pulumi:"createdTimeUtc"`
	// The data connector definition name.
	DataConnectorDefinitionName *string `pulumi:"dataConnectorDefinitionName"`
	// The kind of the data connector definitions
	// Expected value is 'Customizable'.
	Kind string `pulumi:"kind"`
	// Gets or sets the connector definition last modified date in UTC format.
	LastModifiedUtc *string `pulumi:"lastModifiedUtc"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a CustomizableConnectorDefinition resource.
type CustomizableConnectorDefinitionArgs struct {
	// The UiConfig for 'Customizable' connector definition kind.
	ConnectionsConfig CustomizableConnectionsConfigPtrInput
	// The UiConfig for 'Customizable' connector definition kind.
	ConnectorUiConfig CustomizableConnectorUiConfigInput
	// Gets or sets the connector definition created date in UTC format.
	CreatedTimeUtc pulumi.StringPtrInput
	// The data connector definition name.
	DataConnectorDefinitionName pulumi.StringPtrInput
	// The kind of the data connector definitions
	// Expected value is 'Customizable'.
	Kind pulumi.StringInput
	// Gets or sets the connector definition last modified date in UTC format.
	LastModifiedUtc pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (CustomizableConnectorDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customizableConnectorDefinitionArgs)(nil)).Elem()
}

type CustomizableConnectorDefinitionInput interface {
	pulumi.Input

	ToCustomizableConnectorDefinitionOutput() CustomizableConnectorDefinitionOutput
	ToCustomizableConnectorDefinitionOutputWithContext(ctx context.Context) CustomizableConnectorDefinitionOutput
}

func (*CustomizableConnectorDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizableConnectorDefinition)(nil)).Elem()
}

func (i *CustomizableConnectorDefinition) ToCustomizableConnectorDefinitionOutput() CustomizableConnectorDefinitionOutput {
	return i.ToCustomizableConnectorDefinitionOutputWithContext(context.Background())
}

func (i *CustomizableConnectorDefinition) ToCustomizableConnectorDefinitionOutputWithContext(ctx context.Context) CustomizableConnectorDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizableConnectorDefinitionOutput)
}

type CustomizableConnectorDefinitionOutput struct{ *pulumi.OutputState }

func (CustomizableConnectorDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizableConnectorDefinition)(nil)).Elem()
}

func (o CustomizableConnectorDefinitionOutput) ToCustomizableConnectorDefinitionOutput() CustomizableConnectorDefinitionOutput {
	return o
}

func (o CustomizableConnectorDefinitionOutput) ToCustomizableConnectorDefinitionOutputWithContext(ctx context.Context) CustomizableConnectorDefinitionOutput {
	return o
}

// The UiConfig for 'Customizable' connector definition kind.
func (o CustomizableConnectorDefinitionOutput) ConnectionsConfig() CustomizableConnectionsConfigResponsePtrOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) CustomizableConnectionsConfigResponsePtrOutput {
		return v.ConnectionsConfig
	}).(CustomizableConnectionsConfigResponsePtrOutput)
}

// The UiConfig for 'Customizable' connector definition kind.
func (o CustomizableConnectorDefinitionOutput) ConnectorUiConfig() CustomizableConnectorUiConfigResponseOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) CustomizableConnectorUiConfigResponseOutput {
		return v.ConnectorUiConfig
	}).(CustomizableConnectorUiConfigResponseOutput)
}

// Gets or sets the connector definition created date in UTC format.
func (o CustomizableConnectorDefinitionOutput) CreatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) pulumi.StringPtrOutput { return v.CreatedTimeUtc }).(pulumi.StringPtrOutput)
}

// Etag of the azure resource
func (o CustomizableConnectorDefinitionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector definitions
// Expected value is 'Customizable'.
func (o CustomizableConnectorDefinitionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Gets or sets the connector definition last modified date in UTC format.
func (o CustomizableConnectorDefinitionOutput) LastModifiedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) pulumi.StringPtrOutput { return v.LastModifiedUtc }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o CustomizableConnectorDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CustomizableConnectorDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CustomizableConnectorDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizableConnectorDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomizableConnectorDefinitionOutput{})
}
