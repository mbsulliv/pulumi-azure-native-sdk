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

// An ADLS Gen 1 file data set.
type ADLSGen1FileDataSet struct {
	pulumi.CustomResourceState

	// The ADLS account name.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Unique id for identifying a data set resource
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// The file name in the ADLS account.
	FileName pulumi.StringOutput `pulumi:"fileName"`
	// The folder path within the ADLS account.
	FolderPath pulumi.StringOutput `pulumi:"folderPath"`
	// Kind of data set.
	// Expected value is 'AdlsGen1File'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource group of ADLS account.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// Subscription id of ADLS account.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewADLSGen1FileDataSet registers a new resource with the given unique name, arguments, and options.
func NewADLSGen1FileDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen1FileDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen1FileDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.FileName == nil {
		return nil, errors.New("invalid value for required argument 'FileName'")
	}
	if args.FolderPath == nil {
		return nil, errors.New("invalid value for required argument 'FolderPath'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("AdlsGen1File")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen1FileDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen1FileDataSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ADLSGen1FileDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20210801:ADLSGen1FileDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetADLSGen1FileDataSet gets an existing ADLSGen1FileDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetADLSGen1FileDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen1FileDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen1FileDataSet, error) {
	var resource ADLSGen1FileDataSet
	err := ctx.ReadResource("azure-native:datashare/v20210801:ADLSGen1FileDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ADLSGen1FileDataSet resources.
type adlsgen1FileDataSetState struct {
}

type ADLSGen1FileDataSetState struct {
}

func (ADLSGen1FileDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen1FileDataSetState)(nil)).Elem()
}

type adlsgen1FileDataSetArgs struct {
	// The ADLS account name.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName *string `pulumi:"dataSetName"`
	// The file name in the ADLS account.
	FileName string `pulumi:"fileName"`
	// The folder path within the ADLS account.
	FolderPath string `pulumi:"folderPath"`
	// Kind of data set.
	// Expected value is 'AdlsGen1File'.
	Kind string `pulumi:"kind"`
	// Resource group of ADLS account.
	ResourceGroup string `pulumi:"resourceGroup"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share to add the data set to.
	ShareName string `pulumi:"shareName"`
	// Subscription id of ADLS account.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a ADLSGen1FileDataSet resource.
type ADLSGen1FileDataSetArgs struct {
	// The ADLS account name.
	AccountName pulumi.StringInput
	// The name of the dataSet.
	DataSetName pulumi.StringPtrInput
	// The file name in the ADLS account.
	FileName pulumi.StringInput
	// The folder path within the ADLS account.
	FolderPath pulumi.StringInput
	// Kind of data set.
	// Expected value is 'AdlsGen1File'.
	Kind pulumi.StringInput
	// Resource group of ADLS account.
	ResourceGroup pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share to add the data set to.
	ShareName pulumi.StringInput
	// Subscription id of ADLS account.
	SubscriptionId pulumi.StringInput
}

func (ADLSGen1FileDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen1FileDataSetArgs)(nil)).Elem()
}

type ADLSGen1FileDataSetInput interface {
	pulumi.Input

	ToADLSGen1FileDataSetOutput() ADLSGen1FileDataSetOutput
	ToADLSGen1FileDataSetOutputWithContext(ctx context.Context) ADLSGen1FileDataSetOutput
}

func (*ADLSGen1FileDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen1FileDataSet)(nil)).Elem()
}

func (i *ADLSGen1FileDataSet) ToADLSGen1FileDataSetOutput() ADLSGen1FileDataSetOutput {
	return i.ToADLSGen1FileDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen1FileDataSet) ToADLSGen1FileDataSetOutputWithContext(ctx context.Context) ADLSGen1FileDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen1FileDataSetOutput)
}

func (i *ADLSGen1FileDataSet) ToOutput(ctx context.Context) pulumix.Output[*ADLSGen1FileDataSet] {
	return pulumix.Output[*ADLSGen1FileDataSet]{
		OutputState: i.ToADLSGen1FileDataSetOutputWithContext(ctx).OutputState,
	}
}

type ADLSGen1FileDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen1FileDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen1FileDataSet)(nil)).Elem()
}

func (o ADLSGen1FileDataSetOutput) ToADLSGen1FileDataSetOutput() ADLSGen1FileDataSetOutput {
	return o
}

func (o ADLSGen1FileDataSetOutput) ToADLSGen1FileDataSetOutputWithContext(ctx context.Context) ADLSGen1FileDataSetOutput {
	return o
}

func (o ADLSGen1FileDataSetOutput) ToOutput(ctx context.Context) pulumix.Output[*ADLSGen1FileDataSet] {
	return pulumix.Output[*ADLSGen1FileDataSet]{
		OutputState: o.OutputState,
	}
}

// The ADLS account name.
func (o ADLSGen1FileDataSetOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// Unique id for identifying a data set resource
func (o ADLSGen1FileDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// The file name in the ADLS account.
func (o ADLSGen1FileDataSetOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.FileName }).(pulumi.StringOutput)
}

// The folder path within the ADLS account.
func (o ADLSGen1FileDataSetOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.FolderPath }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'AdlsGen1File'.
func (o ADLSGen1FileDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o ADLSGen1FileDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource group of ADLS account.
func (o ADLSGen1FileDataSetOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Subscription id of ADLS account.
func (o ADLSGen1FileDataSetOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o ADLSGen1FileDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o ADLSGen1FileDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen1FileDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen1FileDataSetOutput{})
}
