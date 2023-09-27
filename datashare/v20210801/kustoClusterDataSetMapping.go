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

// A Kusto cluster data set mapping
type KustoClusterDataSetMapping struct {
	pulumi.CustomResourceState

	// The id of the source data set.
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus pulumi.StringOutput `pulumi:"dataSetMappingStatus"`
	// Kind of data set mapping.
	// Expected value is 'KustoCluster'.
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

// NewKustoClusterDataSetMapping registers a new resource with the given unique name, arguments, and options.
func NewKustoClusterDataSetMapping(ctx *pulumi.Context,
	name string, args *KustoClusterDataSetMappingArgs, opts ...pulumi.ResourceOption) (*KustoClusterDataSetMapping, error) {
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
	args.Kind = pulumi.String("KustoCluster")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoClusterDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KustoClusterDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20210801:KustoClusterDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKustoClusterDataSetMapping gets an existing KustoClusterDataSetMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKustoClusterDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoClusterDataSetMappingState, opts ...pulumi.ResourceOption) (*KustoClusterDataSetMapping, error) {
	var resource KustoClusterDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20210801:KustoClusterDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KustoClusterDataSetMapping resources.
type kustoClusterDataSetMappingState struct {
}

type KustoClusterDataSetMappingState struct {
}

func (KustoClusterDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoClusterDataSetMappingState)(nil)).Elem()
}

type kustoClusterDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// The name of the data set mapping to be created.
	DataSetMappingName *string `pulumi:"dataSetMappingName"`
	// Kind of data set mapping.
	// Expected value is 'KustoCluster'.
	Kind string `pulumi:"kind"`
	// Resource id of the sink kusto cluster.
	KustoClusterResourceId string `pulumi:"kustoClusterResourceId"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// The set of arguments for constructing a KustoClusterDataSetMapping resource.
type KustoClusterDataSetMappingArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The id of the source data set.
	DataSetId pulumi.StringInput
	// The name of the data set mapping to be created.
	DataSetMappingName pulumi.StringPtrInput
	// Kind of data set mapping.
	// Expected value is 'KustoCluster'.
	Kind pulumi.StringInput
	// Resource id of the sink kusto cluster.
	KustoClusterResourceId pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName pulumi.StringInput
}

func (KustoClusterDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoClusterDataSetMappingArgs)(nil)).Elem()
}

type KustoClusterDataSetMappingInput interface {
	pulumi.Input

	ToKustoClusterDataSetMappingOutput() KustoClusterDataSetMappingOutput
	ToKustoClusterDataSetMappingOutputWithContext(ctx context.Context) KustoClusterDataSetMappingOutput
}

func (*KustoClusterDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoClusterDataSetMapping)(nil)).Elem()
}

func (i *KustoClusterDataSetMapping) ToKustoClusterDataSetMappingOutput() KustoClusterDataSetMappingOutput {
	return i.ToKustoClusterDataSetMappingOutputWithContext(context.Background())
}

func (i *KustoClusterDataSetMapping) ToKustoClusterDataSetMappingOutputWithContext(ctx context.Context) KustoClusterDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoClusterDataSetMappingOutput)
}

func (i *KustoClusterDataSetMapping) ToOutput(ctx context.Context) pulumix.Output[*KustoClusterDataSetMapping] {
	return pulumix.Output[*KustoClusterDataSetMapping]{
		OutputState: i.ToKustoClusterDataSetMappingOutputWithContext(ctx).OutputState,
	}
}

type KustoClusterDataSetMappingOutput struct{ *pulumi.OutputState }

func (KustoClusterDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoClusterDataSetMapping)(nil)).Elem()
}

func (o KustoClusterDataSetMappingOutput) ToKustoClusterDataSetMappingOutput() KustoClusterDataSetMappingOutput {
	return o
}

func (o KustoClusterDataSetMappingOutput) ToKustoClusterDataSetMappingOutputWithContext(ctx context.Context) KustoClusterDataSetMappingOutput {
	return o
}

func (o KustoClusterDataSetMappingOutput) ToOutput(ctx context.Context) pulumix.Output[*KustoClusterDataSetMapping] {
	return pulumix.Output[*KustoClusterDataSetMapping]{
		OutputState: o.OutputState,
	}
}

// The id of the source data set.
func (o KustoClusterDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o KustoClusterDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'KustoCluster'.
func (o KustoClusterDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource id of the sink kusto cluster.
func (o KustoClusterDataSetMappingOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

// Location of the sink kusto cluster.
func (o KustoClusterDataSetMappingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o KustoClusterDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o KustoClusterDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o KustoClusterDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o KustoClusterDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoClusterDataSetMappingOutput{})
}
