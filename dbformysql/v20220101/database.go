// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a Database.
type Database struct {
	pulumi.CustomResourceState

	// The charset of the database.
	Charset pulumi.StringPtrOutput `pulumi:"charset"`
	// The collation of the database.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbformysql:Database"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20200701preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20200701privatepreview:Database"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20210501:Database"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20210501preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20211201preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20230601preview:Database"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("azure-native:dbformysql/v20220101:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:dbformysql/v20220101:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
}

type DatabaseState struct {
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The charset of the database.
	Charset *string `pulumi:"charset"`
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// The name of the database.
	DatabaseName *string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The charset of the database.
	Charset pulumi.StringPtrInput
	// The collation of the database.
	Collation pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

func (i *Database) ToOutput(ctx context.Context) pulumix.Output[*Database] {
	return pulumix.Output[*Database]{
		OutputState: i.ToDatabaseOutputWithContext(ctx).OutputState,
	}
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToOutput(ctx context.Context) pulumix.Output[*Database] {
	return pulumix.Output[*Database]{
		OutputState: o.OutputState,
	}
}

// The charset of the database.
func (o DatabaseOutput) Charset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Charset }).(pulumi.StringPtrOutput)
}

// The collation of the database.
func (o DatabaseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o DatabaseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Database) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
