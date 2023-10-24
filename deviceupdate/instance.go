// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceupdate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Device Update instance details.
// Azure REST API version: 2023-07-01. Prior API version in Azure Native 1.x: 2020-03-01-preview.
type Instance struct {
	pulumi.CustomResourceState

	// Parent Device Update Account name which Instance belongs to.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Customer-initiated diagnostic log collection storage properties
	DiagnosticStorageProperties DiagnosticStoragePropertiesResponsePtrOutput `pulumi:"diagnosticStorageProperties"`
	// Enables or Disables the diagnostic logs collection
	EnableDiagnostics pulumi.BoolPtrOutput `pulumi:"enableDiagnostics"`
	// List of IoT Hubs associated with the account.
	IotHubs IotHubSettingsResponseArrayOutput `pulumi:"iotHubs"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deviceupdate/v20200301preview:Instance"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20220401preview:Instance"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20221001:Instance"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20221201preview:Instance"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20230701:Instance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("azure-native:deviceupdate:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("azure-native:deviceupdate:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Account name.
	AccountName string `pulumi:"accountName"`
	// Customer-initiated diagnostic log collection storage properties
	DiagnosticStorageProperties *DiagnosticStorageProperties `pulumi:"diagnosticStorageProperties"`
	// Enables or Disables the diagnostic logs collection
	EnableDiagnostics *bool `pulumi:"enableDiagnostics"`
	// Instance name.
	InstanceName *string `pulumi:"instanceName"`
	// List of IoT Hubs associated with the account.
	IotHubs []IotHubSettings `pulumi:"iotHubs"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Account name.
	AccountName pulumi.StringInput
	// Customer-initiated diagnostic log collection storage properties
	DiagnosticStorageProperties DiagnosticStoragePropertiesPtrInput
	// Enables or Disables the diagnostic logs collection
	EnableDiagnostics pulumi.BoolPtrInput
	// Instance name.
	InstanceName pulumi.StringPtrInput
	// List of IoT Hubs associated with the account.
	IotHubs IotHubSettingsArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToOutput(ctx context.Context) pulumix.Output[*Instance] {
	return pulumix.Output[*Instance]{
		OutputState: i.ToInstanceOutputWithContext(ctx).OutputState,
	}
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToOutput(ctx context.Context) pulumix.Output[*Instance] {
	return pulumix.Output[*Instance]{
		OutputState: o.OutputState,
	}
}

// Parent Device Update Account name which Instance belongs to.
func (o InstanceOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// Customer-initiated diagnostic log collection storage properties
func (o InstanceOutput) DiagnosticStorageProperties() DiagnosticStoragePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Instance) DiagnosticStoragePropertiesResponsePtrOutput { return v.DiagnosticStorageProperties }).(DiagnosticStoragePropertiesResponsePtrOutput)
}

// Enables or Disables the diagnostic logs collection
func (o InstanceOutput) EnableDiagnostics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.EnableDiagnostics }).(pulumi.BoolPtrOutput)
}

// List of IoT Hubs associated with the account.
func (o InstanceOutput) IotHubs() IotHubSettingsResponseArrayOutput {
	return o.ApplyT(func(v *Instance) IotHubSettingsResponseArrayOutput { return v.IotHubs }).(IotHubSettingsResponseArrayOutput)
}

// The geo-location where the resource lives
func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state.
func (o InstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o InstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Instance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o InstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o InstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}
