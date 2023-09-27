// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230606

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A host resource belonging to a site resource.
type HypervHostController struct {
	pulumi.CustomResourceState

	// Gets the timestamp marking Hyper-V host creation.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// Gets the errors.
	Errors HealthErrorDetailsResponseArrayOutput `pulumi:"errors"`
	// Gets or sets the FQDN/IPAddress of the Hyper-V host.
	Fqdn pulumi.StringPtrOutput `pulumi:"fqdn"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets the run as account ID of the Hyper-V host.
	RunAsAccountId pulumi.StringPtrOutput `pulumi:"runAsAccountId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets the timestamp marking last updated on the Hyper-V host.
	UpdatedTimestamp pulumi.StringOutput `pulumi:"updatedTimestamp"`
	// Gets the version of the Hyper-V host.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewHypervHostController registers a new resource with the given unique name, arguments, and options.
func NewHypervHostController(ctx *pulumi.Context,
	name string, args *HypervHostControllerArgs, opts ...pulumi.ResourceOption) (*HypervHostController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SiteName == nil {
		return nil, errors.New("invalid value for required argument 'SiteName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazure:HypervHostController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HypervHostController
	err := ctx.RegisterResource("azure-native:offazure/v20230606:HypervHostController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHypervHostController gets an existing HypervHostController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHypervHostController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HypervHostControllerState, opts ...pulumi.ResourceOption) (*HypervHostController, error) {
	var resource HypervHostController
	err := ctx.ReadResource("azure-native:offazure/v20230606:HypervHostController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HypervHostController resources.
type hypervHostControllerState struct {
}

type HypervHostControllerState struct {
}

func (HypervHostControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*hypervHostControllerState)(nil)).Elem()
}

type hypervHostControllerArgs struct {
	// Gets or sets the FQDN/IPAddress of the Hyper-V host.
	Fqdn *string `pulumi:"fqdn"`
	//  Host name
	HostName *string `pulumi:"hostName"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the run as account ID of the Hyper-V host.
	RunAsAccountId *string `pulumi:"runAsAccountId"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// The set of arguments for constructing a HypervHostController resource.
type HypervHostControllerArgs struct {
	// Gets or sets the FQDN/IPAddress of the Hyper-V host.
	Fqdn pulumi.StringPtrInput
	//  Host name
	HostName pulumi.StringPtrInput
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the run as account ID of the Hyper-V host.
	RunAsAccountId pulumi.StringPtrInput
	// Site name
	SiteName pulumi.StringInput
}

func (HypervHostControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hypervHostControllerArgs)(nil)).Elem()
}

type HypervHostControllerInput interface {
	pulumi.Input

	ToHypervHostControllerOutput() HypervHostControllerOutput
	ToHypervHostControllerOutputWithContext(ctx context.Context) HypervHostControllerOutput
}

func (*HypervHostController) ElementType() reflect.Type {
	return reflect.TypeOf((**HypervHostController)(nil)).Elem()
}

func (i *HypervHostController) ToHypervHostControllerOutput() HypervHostControllerOutput {
	return i.ToHypervHostControllerOutputWithContext(context.Background())
}

func (i *HypervHostController) ToHypervHostControllerOutputWithContext(ctx context.Context) HypervHostControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HypervHostControllerOutput)
}

func (i *HypervHostController) ToOutput(ctx context.Context) pulumix.Output[*HypervHostController] {
	return pulumix.Output[*HypervHostController]{
		OutputState: i.ToHypervHostControllerOutputWithContext(ctx).OutputState,
	}
}

type HypervHostControllerOutput struct{ *pulumi.OutputState }

func (HypervHostControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HypervHostController)(nil)).Elem()
}

func (o HypervHostControllerOutput) ToHypervHostControllerOutput() HypervHostControllerOutput {
	return o
}

func (o HypervHostControllerOutput) ToHypervHostControllerOutputWithContext(ctx context.Context) HypervHostControllerOutput {
	return o
}

func (o HypervHostControllerOutput) ToOutput(ctx context.Context) pulumix.Output[*HypervHostController] {
	return pulumix.Output[*HypervHostController]{
		OutputState: o.OutputState,
	}
}

// Gets the timestamp marking Hyper-V host creation.
func (o HypervHostControllerOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervHostController) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Gets the errors.
func (o HypervHostControllerOutput) Errors() HealthErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v *HypervHostController) HealthErrorDetailsResponseArrayOutput { return v.Errors }).(HealthErrorDetailsResponseArrayOutput)
}

// Gets or sets the FQDN/IPAddress of the Hyper-V host.
func (o HypervHostControllerOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HypervHostController) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o HypervHostControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervHostController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o HypervHostControllerOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HypervHostController) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets or sets the run as account ID of the Hyper-V host.
func (o HypervHostControllerOutput) RunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HypervHostController) pulumi.StringPtrOutput { return v.RunAsAccountId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o HypervHostControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HypervHostController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HypervHostControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervHostController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets the timestamp marking last updated on the Hyper-V host.
func (o HypervHostControllerOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervHostController) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

// Gets the version of the Hyper-V host.
func (o HypervHostControllerOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervHostController) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HypervHostControllerOutput{})
}
