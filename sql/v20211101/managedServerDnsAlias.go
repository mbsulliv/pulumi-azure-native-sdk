// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A managed server DNS alias.
type ManagedServerDnsAlias struct {
	pulumi.CustomResourceState

	// The fully qualified DNS record for managed server alias
	AzureDnsRecord pulumi.StringOutput `pulumi:"azureDnsRecord"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The fully qualified public DNS record for managed server alias
	PublicAzureDnsRecord pulumi.StringOutput `pulumi:"publicAzureDnsRecord"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedServerDnsAlias registers a new resource with the given unique name, arguments, and options.
func NewManagedServerDnsAlias(ctx *pulumi.Context,
	name string, args *ManagedServerDnsAliasArgs, opts ...pulumi.ResourceOption) (*ManagedServerDnsAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.CreateDnsRecord == nil {
		args.CreateDnsRecord = pulumi.BoolPtr(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ManagedServerDnsAlias"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:ManagedServerDnsAlias"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedServerDnsAlias
	err := ctx.RegisterResource("azure-native:sql/v20211101:ManagedServerDnsAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedServerDnsAlias gets an existing ManagedServerDnsAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedServerDnsAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedServerDnsAliasState, opts ...pulumi.ResourceOption) (*ManagedServerDnsAlias, error) {
	var resource ManagedServerDnsAlias
	err := ctx.ReadResource("azure-native:sql/v20211101:ManagedServerDnsAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedServerDnsAlias resources.
type managedServerDnsAliasState struct {
}

type ManagedServerDnsAliasState struct {
}

func (ManagedServerDnsAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedServerDnsAliasState)(nil)).Elem()
}

type managedServerDnsAliasArgs struct {
	// Whether or not DNS record should be created for this alias.
	CreateDnsRecord *bool   `pulumi:"createDnsRecord"`
	DnsAliasName    *string `pulumi:"dnsAliasName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedServerDnsAlias resource.
type ManagedServerDnsAliasArgs struct {
	// Whether or not DNS record should be created for this alias.
	CreateDnsRecord pulumi.BoolPtrInput
	DnsAliasName    pulumi.StringPtrInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
}

func (ManagedServerDnsAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedServerDnsAliasArgs)(nil)).Elem()
}

type ManagedServerDnsAliasInput interface {
	pulumi.Input

	ToManagedServerDnsAliasOutput() ManagedServerDnsAliasOutput
	ToManagedServerDnsAliasOutputWithContext(ctx context.Context) ManagedServerDnsAliasOutput
}

func (*ManagedServerDnsAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServerDnsAlias)(nil)).Elem()
}

func (i *ManagedServerDnsAlias) ToManagedServerDnsAliasOutput() ManagedServerDnsAliasOutput {
	return i.ToManagedServerDnsAliasOutputWithContext(context.Background())
}

func (i *ManagedServerDnsAlias) ToManagedServerDnsAliasOutputWithContext(ctx context.Context) ManagedServerDnsAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServerDnsAliasOutput)
}

type ManagedServerDnsAliasOutput struct{ *pulumi.OutputState }

func (ManagedServerDnsAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServerDnsAlias)(nil)).Elem()
}

func (o ManagedServerDnsAliasOutput) ToManagedServerDnsAliasOutput() ManagedServerDnsAliasOutput {
	return o
}

func (o ManagedServerDnsAliasOutput) ToManagedServerDnsAliasOutputWithContext(ctx context.Context) ManagedServerDnsAliasOutput {
	return o
}

// The fully qualified DNS record for managed server alias
func (o ManagedServerDnsAliasOutput) AzureDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedServerDnsAlias) pulumi.StringOutput { return v.AzureDnsRecord }).(pulumi.StringOutput)
}

// Resource name.
func (o ManagedServerDnsAliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedServerDnsAlias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The fully qualified public DNS record for managed server alias
func (o ManagedServerDnsAliasOutput) PublicAzureDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedServerDnsAlias) pulumi.StringOutput { return v.PublicAzureDnsRecord }).(pulumi.StringOutput)
}

// Resource type.
func (o ManagedServerDnsAliasOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedServerDnsAlias) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedServerDnsAliasOutput{})
}
