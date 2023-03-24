// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230126preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Confidential Ledger. Contains the properties of Confidential Ledger Resource.
type Ledger struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of Confidential Ledger Resource.
	Properties LedgerPropertiesResponseOutput `pulumi:"properties"`
	// Object representing RunningState for Ledger.
	RunningState pulumi.StringPtrOutput `pulumi:"runningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Additional tags for Confidential Ledger
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
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
			Type: pulumi.String("azure-native:confidentialledger:Ledger"),
		},
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
	})
	opts = append(opts, aliases)
	var resource Ledger
	err := ctx.RegisterResource("azure-native:confidentialledger/v20230126preview:Ledger", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:confidentialledger/v20230126preview:Ledger", name, id, state, &resource, opts...)
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
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Properties of Confidential Ledger Resource.
	Properties *LedgerProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Object representing RunningState for Ledger.
	RunningState *string `pulumi:"runningState"`
	// Additional tags for Confidential Ledger
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Ledger resource.
type LedgerArgs struct {
	// Name of the Confidential Ledger
	LedgerName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Properties of Confidential Ledger Resource.
	Properties LedgerPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Object representing RunningState for Ledger.
	RunningState pulumi.StringPtrInput
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

// The geo-location where the resource lives
func (o LedgerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LedgerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of Confidential Ledger Resource.
func (o LedgerOutput) Properties() LedgerPropertiesResponseOutput {
	return o.ApplyT(func(v *Ledger) LedgerPropertiesResponseOutput { return v.Properties }).(LedgerPropertiesResponseOutput)
}

// Object representing RunningState for Ledger.
func (o LedgerOutput) RunningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringPtrOutput { return v.RunningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LedgerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Ledger) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Additional tags for Confidential Ledger
func (o LedgerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LedgerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LedgerOutput{})
}