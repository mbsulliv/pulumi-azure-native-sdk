// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of the connection.
// Azure REST API version: 2022-08-08. Prior API version in Azure Native 1.x: 2019-06-01.
type Connection struct {
	pulumi.CustomResourceState

	// Gets or sets the connectionType of the connection.
	ConnectionType ConnectionTypeAssociationPropertyResponsePtrOutput `pulumi:"connectionType"`
	// Gets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets the field definition values of the connection.
	FieldDefinitionValues pulumi.StringMapOutput `pulumi:"fieldDefinitionValues"`
	// Gets the last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ConnectionType == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionType'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation/v20151031:Connection"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Connection"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Connection"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:Connection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("azure-native:automation:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("azure-native:automation:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The parameters supplied to the create or update connection operation.
	ConnectionName *string `pulumi:"connectionName"`
	// Gets or sets the connectionType of the connection.
	ConnectionType ConnectionTypeAssociationProperty `pulumi:"connectionType"`
	// Gets or sets the description of the connection.
	Description *string `pulumi:"description"`
	// Gets or sets the field definition properties of the connection.
	FieldDefinitionValues map[string]string `pulumi:"fieldDefinitionValues"`
	// Gets or sets the name of the connection.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The parameters supplied to the create or update connection operation.
	ConnectionName pulumi.StringPtrInput
	// Gets or sets the connectionType of the connection.
	ConnectionType ConnectionTypeAssociationPropertyInput
	// Gets or sets the description of the connection.
	Description pulumi.StringPtrInput
	// Gets or sets the field definition properties of the connection.
	FieldDefinitionValues pulumi.StringMapInput
	// Gets or sets the name of the connection.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

func (i *Connection) ToOutput(ctx context.Context) pulumix.Output[*Connection] {
	return pulumix.Output[*Connection]{
		OutputState: i.ToConnectionOutputWithContext(ctx).OutputState,
	}
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*Connection] {
	return pulumix.Output[*Connection]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the connectionType of the connection.
func (o ConnectionOutput) ConnectionType() ConnectionTypeAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v *Connection) ConnectionTypeAssociationPropertyResponsePtrOutput { return v.ConnectionType }).(ConnectionTypeAssociationPropertyResponsePtrOutput)
}

// Gets the creation time.
func (o ConnectionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets or sets the description.
func (o ConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets the field definition values of the connection.
func (o ConnectionOutput) FieldDefinitionValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringMapOutput { return v.FieldDefinitionValues }).(pulumi.StringMapOutput)
}

// Gets the last modified time.
func (o ConnectionOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource.
func (o ConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
}
