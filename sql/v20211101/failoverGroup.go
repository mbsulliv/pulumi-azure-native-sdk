// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A failover group.
type FailoverGroup struct {
	pulumi.CustomResourceState

	// List of databases in the failover group.
	Databases pulumi.StringArrayOutput `pulumi:"databases"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of partner server information for the failover group.
	PartnerServers PartnerInfoResponseArrayOutput `pulumi:"partnerServers"`
	// Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint FailoverGroupReadOnlyEndpointResponsePtrOutput `pulumi:"readOnlyEndpoint"`
	// Read-write endpoint of the failover group instance.
	ReadWriteEndpoint FailoverGroupReadWriteEndpointResponseOutput `pulumi:"readWriteEndpoint"`
	// Local replication role of the failover group instance.
	ReplicationRole pulumi.StringOutput `pulumi:"replicationRole"`
	// Replication state of the failover group instance.
	ReplicationState pulumi.StringOutput `pulumi:"replicationState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFailoverGroup registers a new resource with the given unique name, arguments, and options.
func NewFailoverGroup(ctx *pulumi.Context,
	name string, args *FailoverGroupArgs, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PartnerServers == nil {
		return nil, errors.New("invalid value for required argument 'PartnerServers'")
	}
	if args.ReadWriteEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ReadWriteEndpoint'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:FailoverGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FailoverGroup
	err := ctx.RegisterResource("azure-native:sql/v20211101:FailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFailoverGroup gets an existing FailoverGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverGroupState, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	var resource FailoverGroup
	err := ctx.ReadResource("azure-native:sql/v20211101:FailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FailoverGroup resources.
type failoverGroupState struct {
}

type FailoverGroupState struct {
}

func (FailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupState)(nil)).Elem()
}

type failoverGroupArgs struct {
	// List of databases in the failover group.
	Databases []string `pulumi:"databases"`
	// The name of the failover group.
	FailoverGroupName *string `pulumi:"failoverGroupName"`
	// List of partner server information for the failover group.
	PartnerServers []PartnerInfo `pulumi:"partnerServers"`
	// Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint *FailoverGroupReadOnlyEndpoint `pulumi:"readOnlyEndpoint"`
	// Read-write endpoint of the failover group instance.
	ReadWriteEndpoint FailoverGroupReadWriteEndpoint `pulumi:"readWriteEndpoint"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server containing the failover group.
	ServerName string `pulumi:"serverName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FailoverGroup resource.
type FailoverGroupArgs struct {
	// List of databases in the failover group.
	Databases pulumi.StringArrayInput
	// The name of the failover group.
	FailoverGroupName pulumi.StringPtrInput
	// List of partner server information for the failover group.
	PartnerServers PartnerInfoArrayInput
	// Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint FailoverGroupReadOnlyEndpointPtrInput
	// Read-write endpoint of the failover group instance.
	ReadWriteEndpoint FailoverGroupReadWriteEndpointInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server containing the failover group.
	ServerName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (FailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupArgs)(nil)).Elem()
}

type FailoverGroupInput interface {
	pulumi.Input

	ToFailoverGroupOutput() FailoverGroupOutput
	ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput
}

func (*FailoverGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil)).Elem()
}

func (i *FailoverGroup) ToFailoverGroupOutput() FailoverGroupOutput {
	return i.ToFailoverGroupOutputWithContext(context.Background())
}

func (i *FailoverGroup) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupOutput)
}

func (i *FailoverGroup) ToOutput(ctx context.Context) pulumix.Output[*FailoverGroup] {
	return pulumix.Output[*FailoverGroup]{
		OutputState: i.ToFailoverGroupOutputWithContext(ctx).OutputState,
	}
}

type FailoverGroupOutput struct{ *pulumi.OutputState }

func (FailoverGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil)).Elem()
}

func (o FailoverGroupOutput) ToFailoverGroupOutput() FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*FailoverGroup] {
	return pulumix.Output[*FailoverGroup]{
		OutputState: o.OutputState,
	}
}

// List of databases in the failover group.
func (o FailoverGroupOutput) Databases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringArrayOutput { return v.Databases }).(pulumi.StringArrayOutput)
}

// Resource location.
func (o FailoverGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o FailoverGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of partner server information for the failover group.
func (o FailoverGroupOutput) PartnerServers() PartnerInfoResponseArrayOutput {
	return o.ApplyT(func(v *FailoverGroup) PartnerInfoResponseArrayOutput { return v.PartnerServers }).(PartnerInfoResponseArrayOutput)
}

// Read-only endpoint of the failover group instance.
func (o FailoverGroupOutput) ReadOnlyEndpoint() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyT(func(v *FailoverGroup) FailoverGroupReadOnlyEndpointResponsePtrOutput { return v.ReadOnlyEndpoint }).(FailoverGroupReadOnlyEndpointResponsePtrOutput)
}

// Read-write endpoint of the failover group instance.
func (o FailoverGroupOutput) ReadWriteEndpoint() FailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v *FailoverGroup) FailoverGroupReadWriteEndpointResponseOutput { return v.ReadWriteEndpoint }).(FailoverGroupReadWriteEndpointResponseOutput)
}

// Local replication role of the failover group instance.
func (o FailoverGroupOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.ReplicationRole }).(pulumi.StringOutput)
}

// Replication state of the failover group instance.
func (o FailoverGroupOutput) ReplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.ReplicationState }).(pulumi.StringOutput)
}

// Resource tags.
func (o FailoverGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o FailoverGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FailoverGroupOutput{})
}
