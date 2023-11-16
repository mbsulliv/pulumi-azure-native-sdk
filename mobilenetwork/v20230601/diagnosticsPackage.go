// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Diagnostics package resource.
type DiagnosticsPackage struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the diagnostics package resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The reason for the current state of the diagnostics package collection.
	Reason pulumi.StringOutput `pulumi:"reason"`
	// The status of the diagnostics package collection.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDiagnosticsPackage registers a new resource with the given unique name, arguments, and options.
func NewDiagnosticsPackage(ctx *pulumi.Context,
	name string, args *DiagnosticsPackageArgs, opts ...pulumi.ResourceOption) (*DiagnosticsPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PacketCoreControlPlaneName == nil {
		return nil, errors.New("invalid value for required argument 'PacketCoreControlPlaneName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:DiagnosticsPackage"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230901:DiagnosticsPackage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DiagnosticsPackage
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20230601:DiagnosticsPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiagnosticsPackage gets an existing DiagnosticsPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiagnosticsPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticsPackageState, opts ...pulumi.ResourceOption) (*DiagnosticsPackage, error) {
	var resource DiagnosticsPackage
	err := ctx.ReadResource("azure-native:mobilenetwork/v20230601:DiagnosticsPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiagnosticsPackage resources.
type diagnosticsPackageState struct {
}

type DiagnosticsPackageState struct {
}

func (DiagnosticsPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticsPackageState)(nil)).Elem()
}

type diagnosticsPackageArgs struct {
	// The name of the diagnostics package.
	DiagnosticsPackageName *string `pulumi:"diagnosticsPackageName"`
	// The name of the packet core control plane.
	PacketCoreControlPlaneName string `pulumi:"packetCoreControlPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DiagnosticsPackage resource.
type DiagnosticsPackageArgs struct {
	// The name of the diagnostics package.
	DiagnosticsPackageName pulumi.StringPtrInput
	// The name of the packet core control plane.
	PacketCoreControlPlaneName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (DiagnosticsPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticsPackageArgs)(nil)).Elem()
}

type DiagnosticsPackageInput interface {
	pulumi.Input

	ToDiagnosticsPackageOutput() DiagnosticsPackageOutput
	ToDiagnosticsPackageOutputWithContext(ctx context.Context) DiagnosticsPackageOutput
}

func (*DiagnosticsPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsPackage)(nil)).Elem()
}

func (i *DiagnosticsPackage) ToDiagnosticsPackageOutput() DiagnosticsPackageOutput {
	return i.ToDiagnosticsPackageOutputWithContext(context.Background())
}

func (i *DiagnosticsPackage) ToDiagnosticsPackageOutputWithContext(ctx context.Context) DiagnosticsPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsPackageOutput)
}

type DiagnosticsPackageOutput struct{ *pulumi.OutputState }

func (DiagnosticsPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsPackage)(nil)).Elem()
}

func (o DiagnosticsPackageOutput) ToDiagnosticsPackageOutput() DiagnosticsPackageOutput {
	return o
}

func (o DiagnosticsPackageOutput) ToDiagnosticsPackageOutputWithContext(ctx context.Context) DiagnosticsPackageOutput {
	return o
}

// The name of the resource
func (o DiagnosticsPackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticsPackage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the diagnostics package resource.
func (o DiagnosticsPackageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticsPackage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The reason for the current state of the diagnostics package collection.
func (o DiagnosticsPackageOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticsPackage) pulumi.StringOutput { return v.Reason }).(pulumi.StringOutput)
}

// The status of the diagnostics package collection.
func (o DiagnosticsPackageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticsPackage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DiagnosticsPackageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DiagnosticsPackage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DiagnosticsPackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticsPackage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticsPackageOutput{})
}
