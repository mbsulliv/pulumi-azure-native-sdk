// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220808

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the connection type.
type ConnectionType struct {
	pulumi.CustomResourceState

	// Gets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets the field definitions of the connection type.
	FieldDefinitions FieldDefinitionResponseMapOutput `pulumi:"fieldDefinitions"`
	// Gets or sets a Boolean value to indicate if the connection type is global.
	IsGlobal pulumi.BoolPtrOutput `pulumi:"isGlobal"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// Gets the name of the connection type.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectionType registers a new resource with the given unique name, arguments, and options.
func NewConnectionType(ctx *pulumi.Context,
	name string, args *ConnectionTypeArgs, opts ...pulumi.ResourceOption) (*ConnectionType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.FieldDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'FieldDefinitions'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20231101:ConnectionType"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectionType
	err := ctx.RegisterResource("azure-native:automation/v20220808:ConnectionType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionType gets an existing ConnectionType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionTypeState, opts ...pulumi.ResourceOption) (*ConnectionType, error) {
	var resource ConnectionType
	err := ctx.ReadResource("azure-native:automation/v20220808:ConnectionType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionType resources.
type connectionTypeState struct {
}

type ConnectionTypeState struct {
}

func (ConnectionTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionTypeState)(nil)).Elem()
}

type connectionTypeArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The parameters supplied to the create or update connection type operation.
	ConnectionTypeName *string `pulumi:"connectionTypeName"`
	// Gets or sets the field definitions of the connection type.
	FieldDefinitions map[string]FieldDefinition `pulumi:"fieldDefinitions"`
	// Gets or sets a Boolean value to indicate if the connection type is global.
	IsGlobal *bool `pulumi:"isGlobal"`
	// Gets or sets the name of the connection type.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ConnectionType resource.
type ConnectionTypeArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The parameters supplied to the create or update connection type operation.
	ConnectionTypeName pulumi.StringPtrInput
	// Gets or sets the field definitions of the connection type.
	FieldDefinitions FieldDefinitionMapInput
	// Gets or sets a Boolean value to indicate if the connection type is global.
	IsGlobal pulumi.BoolPtrInput
	// Gets or sets the name of the connection type.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
}

func (ConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionTypeArgs)(nil)).Elem()
}

type ConnectionTypeInput interface {
	pulumi.Input

	ToConnectionTypeOutput() ConnectionTypeOutput
	ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput
}

func (*ConnectionType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionType)(nil)).Elem()
}

func (i *ConnectionType) ToConnectionTypeOutput() ConnectionTypeOutput {
	return i.ToConnectionTypeOutputWithContext(context.Background())
}

func (i *ConnectionType) ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeOutput)
}

type ConnectionTypeOutput struct{ *pulumi.OutputState }

func (ConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionType)(nil)).Elem()
}

func (o ConnectionTypeOutput) ToConnectionTypeOutput() ConnectionTypeOutput {
	return o
}

func (o ConnectionTypeOutput) ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput {
	return o
}

// Gets the creation time.
func (o ConnectionTypeOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionType) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets or sets the description.
func (o ConnectionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionType) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets the field definitions of the connection type.
func (o ConnectionTypeOutput) FieldDefinitions() FieldDefinitionResponseMapOutput {
	return o.ApplyT(func(v *ConnectionType) FieldDefinitionResponseMapOutput { return v.FieldDefinitions }).(FieldDefinitionResponseMapOutput)
}

// Gets or sets a Boolean value to indicate if the connection type is global.
func (o ConnectionTypeOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConnectionType) pulumi.BoolPtrOutput { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o ConnectionTypeOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionType) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// Gets the name of the connection type.
func (o ConnectionTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o ConnectionTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionTypeOutput{})
}
