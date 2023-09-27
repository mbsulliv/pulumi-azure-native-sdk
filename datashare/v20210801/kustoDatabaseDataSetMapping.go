// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Kusto database data set mapping
type KustoDatabaseDataSetMapping struct {
	pulumi.CustomResourceState

	// The id of the source data set.
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus pulumi.StringOutput `pulumi:"dataSetMappingStatus"`
	// Kind of data set mapping.
	// Expected value is 'KustoDatabase'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource id of the sink kusto cluster.
	KustoClusterResourceId pulumi.StringOutput `pulumi:"kustoClusterResourceId"`
	// Location of the sink kusto cluster.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the data set mapping.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKustoDatabaseDataSetMapping registers a new resource with the given unique name, arguments, and options.
func NewKustoDatabaseDataSetMapping(ctx *pulumi.Context,
	name string, args *KustoDatabaseDataSetMappingArgs, opts ...pulumi.ResourceOption) (*KustoDatabaseDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoClusterResourceId == nil {
		return nil, errors.New("invalid value for required argument 'KustoClusterResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	args.Kind = pulumi.String("KustoDatabase")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoDatabaseDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KustoDatabaseDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20210801:KustoDatabaseDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKustoDatabaseDataSetMapping gets an existing KustoDatabaseDataSetMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKustoDatabaseDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoDatabaseDataSetMappingState, opts ...pulumi.ResourceOption) (*KustoDatabaseDataSetMapping, error) {
	var resource KustoDatabaseDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20210801:KustoDatabaseDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KustoDatabaseDataSetMapping resources.
type kustoDatabaseDataSetMappingState struct {
}

type KustoDatabaseDataSetMappingState struct {
}

func (KustoDatabaseDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoDatabaseDataSetMappingState)(nil)).Elem()
}

type kustoDatabaseDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// The name of the data set mapping to be created.
	DataSetMappingName *string `pulumi:"dataSetMappingName"`
	// Kind of data set mapping.
	// Expected value is 'KustoDatabase'.
	Kind string `pulumi:"kind"`
	// Resource id of the sink kusto cluster.
	KustoClusterResourceId string `pulumi:"kustoClusterResourceId"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// The set of arguments for constructing a KustoDatabaseDataSetMapping resource.
type KustoDatabaseDataSetMappingArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The id of the source data set.
	DataSetId pulumi.StringInput
	// The name of the data set mapping to be created.
	DataSetMappingName pulumi.StringPtrInput
	// Kind of data set mapping.
	// Expected value is 'KustoDatabase'.
	Kind pulumi.StringInput
	// Resource id of the sink kusto cluster.
	KustoClusterResourceId pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName pulumi.StringInput
}

func (KustoDatabaseDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoDatabaseDataSetMappingArgs)(nil)).Elem()
}

type KustoDatabaseDataSetMappingInput interface {
	pulumi.Input

	ToKustoDatabaseDataSetMappingOutput() KustoDatabaseDataSetMappingOutput
	ToKustoDatabaseDataSetMappingOutputWithContext(ctx context.Context) KustoDatabaseDataSetMappingOutput
}

func (*KustoDatabaseDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoDatabaseDataSetMapping)(nil)).Elem()
}

func (i *KustoDatabaseDataSetMapping) ToKustoDatabaseDataSetMappingOutput() KustoDatabaseDataSetMappingOutput {
	return i.ToKustoDatabaseDataSetMappingOutputWithContext(context.Background())
}

func (i *KustoDatabaseDataSetMapping) ToKustoDatabaseDataSetMappingOutputWithContext(ctx context.Context) KustoDatabaseDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoDatabaseDataSetMappingOutput)
}

func (i *KustoDatabaseDataSetMapping) ToOutput(ctx context.Context) pulumix.Output[*KustoDatabaseDataSetMapping] {
	return pulumix.Output[*KustoDatabaseDataSetMapping]{
		OutputState: i.ToKustoDatabaseDataSetMappingOutputWithContext(ctx).OutputState,
	}
}

type KustoDatabaseDataSetMappingOutput struct{ *pulumi.OutputState }

func (KustoDatabaseDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoDatabaseDataSetMapping)(nil)).Elem()
}

func (o KustoDatabaseDataSetMappingOutput) ToKustoDatabaseDataSetMappingOutput() KustoDatabaseDataSetMappingOutput {
	return o
}

func (o KustoDatabaseDataSetMappingOutput) ToKustoDatabaseDataSetMappingOutputWithContext(ctx context.Context) KustoDatabaseDataSetMappingOutput {
	return o
}

func (o KustoDatabaseDataSetMappingOutput) ToOutput(ctx context.Context) pulumix.Output[*KustoDatabaseDataSetMapping] {
	return pulumix.Output[*KustoDatabaseDataSetMapping]{
		OutputState: o.OutputState,
	}
}

// The id of the source data set.
func (o KustoDatabaseDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o KustoDatabaseDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'KustoDatabase'.
func (o KustoDatabaseDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource id of the sink kusto cluster.
func (o KustoDatabaseDataSetMappingOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) pulumi.StringOutput { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

// Location of the sink kusto cluster.
func (o KustoDatabaseDataSetMappingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o KustoDatabaseDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o KustoDatabaseDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o KustoDatabaseDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o KustoDatabaseDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoDatabaseDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoDatabaseDataSetMappingOutput{})
}
