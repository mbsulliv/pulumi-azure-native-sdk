// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230115preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Postgres Instance.
type PostgresInstance struct {
	pulumi.CustomResourceState

	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// null
	Properties PostgresInstancePropertiesResponseOutput `pulumi:"properties"`
	// Resource sku.
	Sku PostgresInstanceSkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPostgresInstance registers a new resource with the given unique name, arguments, and options.
func NewPostgresInstance(ctx *pulumi.Context,
	name string, args *PostgresInstanceArgs, opts ...pulumi.ResourceOption) (*PostgresInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToPostgresInstanceSkuPtrOutput().ApplyT(func(v *PostgresInstanceSku) *PostgresInstanceSku { return v.Defaults() }).(PostgresInstanceSkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210601preview:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210701preview:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220301preview:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220615preview:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20240101:PostgresInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PostgresInstance
	err := ctx.RegisterResource("azure-native:azurearcdata/v20230115preview:PostgresInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPostgresInstance gets an existing PostgresInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPostgresInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PostgresInstanceState, opts ...pulumi.ResourceOption) (*PostgresInstance, error) {
	var resource PostgresInstance
	err := ctx.ReadResource("azure-native:azurearcdata/v20230115preview:PostgresInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PostgresInstance resources.
type postgresInstanceState struct {
}

type PostgresInstanceState struct {
}

func (PostgresInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*postgresInstanceState)(nil)).Elem()
}

type postgresInstanceArgs struct {
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of Postgres Instance
	PostgresInstanceName *string `pulumi:"postgresInstanceName"`
	// null
	Properties PostgresInstanceProperties `pulumi:"properties"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource sku.
	Sku *PostgresInstanceSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PostgresInstance resource.
type PostgresInstanceArgs struct {
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of Postgres Instance
	PostgresInstanceName pulumi.StringPtrInput
	// null
	Properties PostgresInstancePropertiesInput
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput
	// Resource sku.
	Sku PostgresInstanceSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PostgresInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*postgresInstanceArgs)(nil)).Elem()
}

type PostgresInstanceInput interface {
	pulumi.Input

	ToPostgresInstanceOutput() PostgresInstanceOutput
	ToPostgresInstanceOutputWithContext(ctx context.Context) PostgresInstanceOutput
}

func (*PostgresInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstance)(nil)).Elem()
}

func (i *PostgresInstance) ToPostgresInstanceOutput() PostgresInstanceOutput {
	return i.ToPostgresInstanceOutputWithContext(context.Background())
}

func (i *PostgresInstance) ToPostgresInstanceOutputWithContext(ctx context.Context) PostgresInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceOutput)
}

type PostgresInstanceOutput struct{ *pulumi.OutputState }

func (PostgresInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstance)(nil)).Elem()
}

func (o PostgresInstanceOutput) ToPostgresInstanceOutput() PostgresInstanceOutput {
	return o
}

func (o PostgresInstanceOutput) ToPostgresInstanceOutputWithContext(ctx context.Context) PostgresInstanceOutput {
	return o
}

// The extendedLocation of the resource.
func (o PostgresInstanceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *PostgresInstance) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The geo-location where the resource lives
func (o PostgresInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PostgresInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PostgresInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PostgresInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// null
func (o PostgresInstanceOutput) Properties() PostgresInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *PostgresInstance) PostgresInstancePropertiesResponseOutput { return v.Properties }).(PostgresInstancePropertiesResponseOutput)
}

// Resource sku.
func (o PostgresInstanceOutput) Sku() PostgresInstanceSkuResponsePtrOutput {
	return o.ApplyT(func(v *PostgresInstance) PostgresInstanceSkuResponsePtrOutput { return v.Sku }).(PostgresInstanceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PostgresInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PostgresInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PostgresInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PostgresInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PostgresInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PostgresInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PostgresInstanceOutput{})
}
