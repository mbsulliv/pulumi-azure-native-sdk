// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Geo backup policy.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2014-04-01.
//
// Other available API versions: 2022-11-01-preview, 2023-02-01-preview.
type GeoBackupPolicy struct {
	pulumi.CustomResourceState

	// Kind of geo backup policy.  This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Backup policy location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The state of the geo backup policy.
	State pulumi.StringOutput `pulumi:"state"`
	// The storage type of the geo backup policy.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGeoBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewGeoBackupPolicy(ctx *pulumi.Context,
	name string, args *GeoBackupPolicyArgs, opts ...pulumi.ResourceOption) (*GeoBackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20140401:GeoBackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:GeoBackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:GeoBackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:GeoBackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:GeoBackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:GeoBackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:GeoBackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GeoBackupPolicy
	err := ctx.RegisterResource("azure-native:sql:GeoBackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeoBackupPolicy gets an existing GeoBackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeoBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeoBackupPolicyState, opts ...pulumi.ResourceOption) (*GeoBackupPolicy, error) {
	var resource GeoBackupPolicy
	err := ctx.ReadResource("azure-native:sql:GeoBackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeoBackupPolicy resources.
type geoBackupPolicyState struct {
}

type GeoBackupPolicyState struct {
}

func (GeoBackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*geoBackupPolicyState)(nil)).Elem()
}

type geoBackupPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Geo backup policy. This should always be 'Default'.
	GeoBackupPolicyName *string `pulumi:"geoBackupPolicyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The state of the geo backup policy.
	State GeoBackupPolicyStateEnum `pulumi:"state"`
}

// The set of arguments for constructing a GeoBackupPolicy resource.
type GeoBackupPolicyArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The name of the Geo backup policy. This should always be 'Default'.
	GeoBackupPolicyName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The state of the geo backup policy.
	State GeoBackupPolicyStateEnumInput
}

func (GeoBackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geoBackupPolicyArgs)(nil)).Elem()
}

type GeoBackupPolicyInput interface {
	pulumi.Input

	ToGeoBackupPolicyOutput() GeoBackupPolicyOutput
	ToGeoBackupPolicyOutputWithContext(ctx context.Context) GeoBackupPolicyOutput
}

func (*GeoBackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoBackupPolicy)(nil)).Elem()
}

func (i *GeoBackupPolicy) ToGeoBackupPolicyOutput() GeoBackupPolicyOutput {
	return i.ToGeoBackupPolicyOutputWithContext(context.Background())
}

func (i *GeoBackupPolicy) ToGeoBackupPolicyOutputWithContext(ctx context.Context) GeoBackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoBackupPolicyOutput)
}

func (i *GeoBackupPolicy) ToOutput(ctx context.Context) pulumix.Output[*GeoBackupPolicy] {
	return pulumix.Output[*GeoBackupPolicy]{
		OutputState: i.ToGeoBackupPolicyOutputWithContext(ctx).OutputState,
	}
}

type GeoBackupPolicyOutput struct{ *pulumi.OutputState }

func (GeoBackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoBackupPolicy)(nil)).Elem()
}

func (o GeoBackupPolicyOutput) ToGeoBackupPolicyOutput() GeoBackupPolicyOutput {
	return o
}

func (o GeoBackupPolicyOutput) ToGeoBackupPolicyOutputWithContext(ctx context.Context) GeoBackupPolicyOutput {
	return o
}

func (o GeoBackupPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*GeoBackupPolicy] {
	return pulumix.Output[*GeoBackupPolicy]{
		OutputState: o.OutputState,
	}
}

// Kind of geo backup policy.  This is metadata used for the Azure portal experience.
func (o GeoBackupPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Backup policy location.
func (o GeoBackupPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o GeoBackupPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The state of the geo backup policy.
func (o GeoBackupPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The storage type of the geo backup policy.
func (o GeoBackupPolicyOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// Resource type.
func (o GeoBackupPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GeoBackupPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GeoBackupPolicyOutput{})
}
