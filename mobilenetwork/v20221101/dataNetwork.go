// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Data network resource. Must be created in the same location as its parent mobile network.
type DataNetwork struct {
	pulumi.CustomResourceState

	// An optional description for this data network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the data network resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataNetwork registers a new resource with the given unique name, arguments, and options.
func NewDataNetwork(ctx *pulumi.Context,
	name string, args *DataNetworkArgs, opts ...pulumi.ResourceOption) (*DataNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:DataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:DataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:DataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230601:DataNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataNetwork
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20221101:DataNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataNetwork gets an existing DataNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataNetworkState, opts ...pulumi.ResourceOption) (*DataNetwork, error) {
	var resource DataNetwork
	err := ctx.ReadResource("azure-native:mobilenetwork/v20221101:DataNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataNetwork resources.
type dataNetworkState struct {
}

type DataNetworkState struct {
}

func (DataNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataNetworkState)(nil)).Elem()
}

type dataNetworkArgs struct {
	// The name of the data network.
	DataNetworkName *string `pulumi:"dataNetworkName"`
	// An optional description for this data network.
	Description *string `pulumi:"description"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataNetwork resource.
type DataNetworkArgs struct {
	// The name of the data network.
	DataNetworkName pulumi.StringPtrInput
	// An optional description for this data network.
	Description pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DataNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataNetworkArgs)(nil)).Elem()
}

type DataNetworkInput interface {
	pulumi.Input

	ToDataNetworkOutput() DataNetworkOutput
	ToDataNetworkOutputWithContext(ctx context.Context) DataNetworkOutput
}

func (*DataNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**DataNetwork)(nil)).Elem()
}

func (i *DataNetwork) ToDataNetworkOutput() DataNetworkOutput {
	return i.ToDataNetworkOutputWithContext(context.Background())
}

func (i *DataNetwork) ToDataNetworkOutputWithContext(ctx context.Context) DataNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataNetworkOutput)
}

func (i *DataNetwork) ToOutput(ctx context.Context) pulumix.Output[*DataNetwork] {
	return pulumix.Output[*DataNetwork]{
		OutputState: i.ToDataNetworkOutputWithContext(ctx).OutputState,
	}
}

type DataNetworkOutput struct{ *pulumi.OutputState }

func (DataNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataNetwork)(nil)).Elem()
}

func (o DataNetworkOutput) ToDataNetworkOutput() DataNetworkOutput {
	return o
}

func (o DataNetworkOutput) ToDataNetworkOutputWithContext(ctx context.Context) DataNetworkOutput {
	return o
}

func (o DataNetworkOutput) ToOutput(ctx context.Context) pulumix.Output[*DataNetwork] {
	return pulumix.Output[*DataNetwork]{
		OutputState: o.OutputState,
	}
}

// An optional description for this data network.
func (o DataNetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o DataNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DataNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the data network resource.
func (o DataNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DataNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DataNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataNetworkOutput{})
}
