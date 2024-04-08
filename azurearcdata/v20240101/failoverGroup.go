// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A failover group resource.
type FailoverGroup struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// null
	Properties FailoverGroupPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFailoverGroup registers a new resource with the given unique name, arguments, and options.
func NewFailoverGroup(ctx *pulumi.Context,
	name string, args *FailoverGroupArgs, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SqlManagedInstanceName'")
	}
	args.Properties = args.Properties.ToFailoverGroupPropertiesOutput().ApplyT(func(v FailoverGroupProperties) FailoverGroupProperties { return *v.Defaults() }).(FailoverGroupPropertiesOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata:FailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20230115preview:FailoverGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FailoverGroup
	err := ctx.RegisterResource("azure-native:azurearcdata/v20240101:FailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFailoverGroup gets an existing FailoverGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverGroupState, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	var resource FailoverGroup
	err := ctx.ReadResource("azure-native:azurearcdata/v20240101:FailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FailoverGroup resources.
type failoverGroupState struct {
}

type FailoverGroupState struct {
}

func (FailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupState)(nil)).Elem()
}

type failoverGroupArgs struct {
	// The name of the Failover Group
	FailoverGroupName *string `pulumi:"failoverGroupName"`
	// null
	Properties FailoverGroupProperties `pulumi:"properties"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of SQL Managed Instance
	SqlManagedInstanceName string `pulumi:"sqlManagedInstanceName"`
}

// The set of arguments for constructing a FailoverGroup resource.
type FailoverGroupArgs struct {
	// The name of the Failover Group
	FailoverGroupName pulumi.StringPtrInput
	// null
	Properties FailoverGroupPropertiesInput
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput
	// Name of SQL Managed Instance
	SqlManagedInstanceName pulumi.StringInput
}

func (FailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupArgs)(nil)).Elem()
}

type FailoverGroupInput interface {
	pulumi.Input

	ToFailoverGroupOutput() FailoverGroupOutput
	ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput
}

func (*FailoverGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil)).Elem()
}

func (i *FailoverGroup) ToFailoverGroupOutput() FailoverGroupOutput {
	return i.ToFailoverGroupOutputWithContext(context.Background())
}

func (i *FailoverGroup) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupOutput)
}

type FailoverGroupOutput struct{ *pulumi.OutputState }

func (FailoverGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil)).Elem()
}

func (o FailoverGroupOutput) ToFailoverGroupOutput() FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return o
}

// The name of the resource
func (o FailoverGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// null
func (o FailoverGroupOutput) Properties() FailoverGroupPropertiesResponseOutput {
	return o.ApplyT(func(v *FailoverGroup) FailoverGroupPropertiesResponseOutput { return v.Properties }).(FailoverGroupPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FailoverGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FailoverGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FailoverGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FailoverGroupOutput{})
}
