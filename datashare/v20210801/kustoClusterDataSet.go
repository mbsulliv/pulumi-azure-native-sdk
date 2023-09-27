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

// A kusto cluster data set.
type KustoClusterDataSet struct {
	pulumi.CustomResourceState

	// Unique id for identifying a data set resource
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Kind of data set.
	// Expected value is 'KustoCluster'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource id of the kusto cluster.
	KustoClusterResourceId pulumi.StringOutput `pulumi:"kustoClusterResourceId"`
	// Location of the kusto cluster.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the kusto cluster data set.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKustoClusterDataSet registers a new resource with the given unique name, arguments, and options.
func NewKustoClusterDataSet(ctx *pulumi.Context,
	name string, args *KustoClusterDataSetArgs, opts ...pulumi.ResourceOption) (*KustoClusterDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
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
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	args.Kind = pulumi.String("KustoCluster")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:KustoClusterDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:KustoClusterDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoClusterDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:KustoClusterDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoClusterDataSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KustoClusterDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20210801:KustoClusterDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKustoClusterDataSet gets an existing KustoClusterDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKustoClusterDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoClusterDataSetState, opts ...pulumi.ResourceOption) (*KustoClusterDataSet, error) {
	var resource KustoClusterDataSet
	err := ctx.ReadResource("azure-native:datashare/v20210801:KustoClusterDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KustoClusterDataSet resources.
type kustoClusterDataSetState struct {
}

type KustoClusterDataSetState struct {
}

func (KustoClusterDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoClusterDataSetState)(nil)).Elem()
}

type kustoClusterDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName *string `pulumi:"dataSetName"`
	// Kind of data set.
	// Expected value is 'KustoCluster'.
	Kind string `pulumi:"kind"`
	// Resource id of the kusto cluster.
	KustoClusterResourceId string `pulumi:"kustoClusterResourceId"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share to add the data set to.
	ShareName string `pulumi:"shareName"`
}

// The set of arguments for constructing a KustoClusterDataSet resource.
type KustoClusterDataSetArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The name of the dataSet.
	DataSetName pulumi.StringPtrInput
	// Kind of data set.
	// Expected value is 'KustoCluster'.
	Kind pulumi.StringInput
	// Resource id of the kusto cluster.
	KustoClusterResourceId pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share to add the data set to.
	ShareName pulumi.StringInput
}

func (KustoClusterDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoClusterDataSetArgs)(nil)).Elem()
}

type KustoClusterDataSetInput interface {
	pulumi.Input

	ToKustoClusterDataSetOutput() KustoClusterDataSetOutput
	ToKustoClusterDataSetOutputWithContext(ctx context.Context) KustoClusterDataSetOutput
}

func (*KustoClusterDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoClusterDataSet)(nil)).Elem()
}

func (i *KustoClusterDataSet) ToKustoClusterDataSetOutput() KustoClusterDataSetOutput {
	return i.ToKustoClusterDataSetOutputWithContext(context.Background())
}

func (i *KustoClusterDataSet) ToKustoClusterDataSetOutputWithContext(ctx context.Context) KustoClusterDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoClusterDataSetOutput)
}

func (i *KustoClusterDataSet) ToOutput(ctx context.Context) pulumix.Output[*KustoClusterDataSet] {
	return pulumix.Output[*KustoClusterDataSet]{
		OutputState: i.ToKustoClusterDataSetOutputWithContext(ctx).OutputState,
	}
}

type KustoClusterDataSetOutput struct{ *pulumi.OutputState }

func (KustoClusterDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoClusterDataSet)(nil)).Elem()
}

func (o KustoClusterDataSetOutput) ToKustoClusterDataSetOutput() KustoClusterDataSetOutput {
	return o
}

func (o KustoClusterDataSetOutput) ToKustoClusterDataSetOutputWithContext(ctx context.Context) KustoClusterDataSetOutput {
	return o
}

func (o KustoClusterDataSetOutput) ToOutput(ctx context.Context) pulumix.Output[*KustoClusterDataSet] {
	return pulumix.Output[*KustoClusterDataSet]{
		OutputState: o.OutputState,
	}
}

// Unique id for identifying a data set resource
func (o KustoClusterDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'KustoCluster'.
func (o KustoClusterDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource id of the kusto cluster.
func (o KustoClusterDataSetOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

// Location of the kusto cluster.
func (o KustoClusterDataSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o KustoClusterDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the kusto cluster data set.
func (o KustoClusterDataSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o KustoClusterDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o KustoClusterDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoClusterDataSetOutput{})
}
