// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package confidentialledger

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Confidential Ledger. Contains the properties of Confidential Ledger Resource.
// Azure REST API version: 2022-05-13. Prior API version in Azure Native 1.x: 2020-12-01-preview
type Ledger struct {
	pulumi.CustomResourceState

	// The Azure location where the Confidential Ledger is running.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the Resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of Confidential Ledger Resource.
	Properties LedgerPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Additional tags for Confidential Ledger
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLedger registers a new resource with the given unique name, arguments, and options.
func NewLedger(ctx *pulumi.Context,
	name string, args *LedgerArgs, opts ...pulumi.ResourceOption) (*Ledger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:confidentialledger/v20201201preview:Ledger"),
		},
		{
			Type: pulumi.String("azure-native:confidentialledger/v20210513preview:Ledger"),
		},
		{
			Type: pulumi.String("azure-native:confidentialledger/v20220513:Ledger"),
		},
		{
			Type: pulumi.String("azure-native:confidentialledger/v20220908preview:Ledger"),
		},
		{
			Type: pulumi.String("azure-native:confidentialledger/v20230126preview:Ledger"),
		},
		{
			Type: pulumi.String("azure-native:confidentialledger/v20230628preview:Ledger"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Ledger
	err := ctx.RegisterResource("azure-native:confidentialledger:Ledger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLedger gets an existing Ledger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLedger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LedgerState, opts ...pulumi.ResourceOption) (*Ledger, error) {
	var resource Ledger
	err := ctx.ReadResource("azure-native:confidentialledger:Ledger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ledger resources.
type ledgerState struct {
}

type LedgerState struct {
}

func (LedgerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ledgerState)(nil)).Elem()
}

type ledgerArgs struct {
	// Name of the Confidential Ledger
	LedgerName *string `pulumi:"ledgerName"`
	// The Azure location where the Confidential Ledger is running.
	Location *string `pulumi:"location"`
	// Properties of Confidential Ledger Resource.
	Properties *LedgerProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Additional tags for Confidential Ledger
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Ledger resource.
type LedgerArgs struct {
	// Name of the Confidential Ledger
	LedgerName pulumi.StringPtrInput
	// The Azure location where the Confidential Ledger is running.
	Location pulumi.StringPtrInput
	// Properties of Confidential Ledger Resource.
	Properties LedgerPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Additional tags for Confidential Ledger
	Tags pulumi.StringMapInput
}

func (LedgerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ledgerArgs)(nil)).Elem()
}

type LedgerInput interface {
	pulumi.Input

	ToLedgerOutput() LedgerOutput
	ToLedgerOutputWithContext(ctx context.Context) LedgerOutput
}

func (*Ledger) ElementType() reflect.Type {
	return reflect.TypeOf((**Ledger)(nil)).Elem()
}

func (i *Ledger) ToLedgerOutput() LedgerOutput {
	return i.ToLedgerOutputWithContext(context.Background())
}

func (i *Ledger) ToLedgerOutputWithContext(ctx context.Context) LedgerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LedgerOutput)
}

func (i *Ledger) ToOutput(ctx context.Context) pulumix.Output[*Ledger] {
	return pulumix.Output[*Ledger]{
		OutputState: i.ToLedgerOutputWithContext(ctx).OutputState,
	}
}

type LedgerOutput struct{ *pulumi.OutputState }

func (LedgerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ledger)(nil)).Elem()
}

func (o LedgerOutput) ToLedgerOutput() LedgerOutput {
	return o
}

func (o LedgerOutput) ToLedgerOutputWithContext(ctx context.Context) LedgerOutput {
	return o
}

func (o LedgerOutput) ToOutput(ctx context.Context) pulumix.Output[*Ledger] {
	return pulumix.Output[*Ledger]{
		OutputState: o.OutputState,
	}
}

// The Azure location where the Confidential Ledger is running.
func (o LedgerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the Resource.
func (o LedgerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of Confidential Ledger Resource.
func (o LedgerOutput) Properties() LedgerPropertiesResponseOutput {
	return o.ApplyT(func(v *Ledger) LedgerPropertiesResponseOutput { return v.Properties }).(LedgerPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o LedgerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Ledger) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Additional tags for Confidential Ledger
func (o LedgerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LedgerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LedgerOutput{})
}
