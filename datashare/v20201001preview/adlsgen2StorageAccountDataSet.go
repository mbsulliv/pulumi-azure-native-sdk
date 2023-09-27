// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An ADLSGen2 storage account data set.
type ADLSGen2StorageAccountDataSet struct {
	pulumi.CustomResourceState

	// Unique id for identifying a data set resource
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Kind of data set.
	// Expected value is 'AdlsGen2StorageAccount'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Location of the ADLSGen2 storage account.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of ADLSGen2 storage account paths.
	Paths ADLSGen2StorageAccountPathResponseArrayOutput `pulumi:"paths"`
	// Resource id of the ADLSGen2 storage account.
	StorageAccountResourceId pulumi.StringOutput `pulumi:"storageAccountResourceId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewADLSGen2StorageAccountDataSet registers a new resource with the given unique name, arguments, and options.
func NewADLSGen2StorageAccountDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen2StorageAccountDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen2StorageAccountDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Paths == nil {
		return nil, errors.New("invalid value for required argument 'Paths'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.StorageAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountResourceId'")
	}
	args.Kind = pulumi.String("AdlsGen2StorageAccount")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2StorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2StorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2StorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2StorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2StorageAccountDataSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ADLSGen2StorageAccountDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetADLSGen2StorageAccountDataSet gets an existing ADLSGen2StorageAccountDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetADLSGen2StorageAccountDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2StorageAccountDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen2StorageAccountDataSet, error) {
	var resource ADLSGen2StorageAccountDataSet
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ADLSGen2StorageAccountDataSet resources.
type adlsgen2StorageAccountDataSetState struct {
}

type ADLSGen2StorageAccountDataSetState struct {
}

func (ADLSGen2StorageAccountDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2StorageAccountDataSetState)(nil)).Elem()
}

type adlsgen2StorageAccountDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName *string `pulumi:"dataSetName"`
	// Kind of data set.
	// Expected value is 'AdlsGen2StorageAccount'.
	Kind string `pulumi:"kind"`
	// A list of ADLSGen2 storage account paths.
	Paths []ADLSGen2StorageAccountPath `pulumi:"paths"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share to add the data set to.
	ShareName string `pulumi:"shareName"`
	// Resource id of the ADLSGen2 storage account.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
}

// The set of arguments for constructing a ADLSGen2StorageAccountDataSet resource.
type ADLSGen2StorageAccountDataSetArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The name of the dataSet.
	DataSetName pulumi.StringPtrInput
	// Kind of data set.
	// Expected value is 'AdlsGen2StorageAccount'.
	Kind pulumi.StringInput
	// A list of ADLSGen2 storage account paths.
	Paths ADLSGen2StorageAccountPathArrayInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share to add the data set to.
	ShareName pulumi.StringInput
	// Resource id of the ADLSGen2 storage account.
	StorageAccountResourceId pulumi.StringInput
}

func (ADLSGen2StorageAccountDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2StorageAccountDataSetArgs)(nil)).Elem()
}

type ADLSGen2StorageAccountDataSetInput interface {
	pulumi.Input

	ToADLSGen2StorageAccountDataSetOutput() ADLSGen2StorageAccountDataSetOutput
	ToADLSGen2StorageAccountDataSetOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetOutput
}

func (*ADLSGen2StorageAccountDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2StorageAccountDataSet)(nil)).Elem()
}

func (i *ADLSGen2StorageAccountDataSet) ToADLSGen2StorageAccountDataSetOutput() ADLSGen2StorageAccountDataSetOutput {
	return i.ToADLSGen2StorageAccountDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen2StorageAccountDataSet) ToADLSGen2StorageAccountDataSetOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2StorageAccountDataSetOutput)
}

func (i *ADLSGen2StorageAccountDataSet) ToOutput(ctx context.Context) pulumix.Output[*ADLSGen2StorageAccountDataSet] {
	return pulumix.Output[*ADLSGen2StorageAccountDataSet]{
		OutputState: i.ToADLSGen2StorageAccountDataSetOutputWithContext(ctx).OutputState,
	}
}

type ADLSGen2StorageAccountDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen2StorageAccountDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2StorageAccountDataSet)(nil)).Elem()
}

func (o ADLSGen2StorageAccountDataSetOutput) ToADLSGen2StorageAccountDataSetOutput() ADLSGen2StorageAccountDataSetOutput {
	return o
}

func (o ADLSGen2StorageAccountDataSetOutput) ToADLSGen2StorageAccountDataSetOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetOutput {
	return o
}

func (o ADLSGen2StorageAccountDataSetOutput) ToOutput(ctx context.Context) pulumix.Output[*ADLSGen2StorageAccountDataSet] {
	return pulumix.Output[*ADLSGen2StorageAccountDataSet]{
		OutputState: o.OutputState,
	}
}

// Unique id for identifying a data set resource
func (o ADLSGen2StorageAccountDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'AdlsGen2StorageAccount'.
func (o ADLSGen2StorageAccountDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Location of the ADLSGen2 storage account.
func (o ADLSGen2StorageAccountDataSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o ADLSGen2StorageAccountDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of ADLSGen2 storage account paths.
func (o ADLSGen2StorageAccountDataSetOutput) Paths() ADLSGen2StorageAccountPathResponseArrayOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) ADLSGen2StorageAccountPathResponseArrayOutput { return v.Paths }).(ADLSGen2StorageAccountPathResponseArrayOutput)
}

// Resource id of the ADLSGen2 storage account.
func (o ADLSGen2StorageAccountDataSetOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o ADLSGen2StorageAccountDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o ADLSGen2StorageAccountDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2StorageAccountDataSetOutput{})
}
