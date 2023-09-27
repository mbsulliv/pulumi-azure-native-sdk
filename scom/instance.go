// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scom

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A SCOM instance resource
// Azure REST API version: 2023-07-07-preview.
type Instance struct {
	pulumi.CustomResourceState

	// The Azure Active Directory identity of the SCOM instance
	Identity ManagedIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of a SCOM instance resource
	Properties MonitoringInstancePropertiesResponseOutput `pulumi:"properties"`
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

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToMonitoringInstancePropertiesPtrOutput().ApplyT(func(v *MonitoringInstanceProperties) *MonitoringInstanceProperties { return v.Defaults() }).(MonitoringInstancePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scom/v20230707preview:Instance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("azure-native:scom:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:scom:Instance", name, id, state, &resource, opts...)
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
	// The Azure Active Directory identity of the SCOM instance
	Identity *ManagedIdentity `pulumi:"identity"`
	// Name of the Azure Monitor Operations Manager Managed Instance (SCOM MI)
	InstanceName *string `pulumi:"instanceName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The properties of a SCOM instance resource
	Properties *MonitoringInstanceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Validation mode for the SCOM managed instance
	ValidationMode *bool `pulumi:"validationMode"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The Azure Active Directory identity of the SCOM instance
	Identity ManagedIdentityPtrInput
	// Name of the Azure Monitor Operations Manager Managed Instance (SCOM MI)
	InstanceName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The properties of a SCOM instance resource
	Properties MonitoringInstancePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Validation mode for the SCOM managed instance
	ValidationMode pulumi.BoolPtrInput
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

// The Azure Active Directory identity of the SCOM instance
func (o InstanceOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Instance) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of a SCOM instance resource
func (o InstanceOutput) Properties() MonitoringInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *Instance) MonitoringInstancePropertiesResponseOutput { return v.Properties }).(MonitoringInstancePropertiesResponseOutput)
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
