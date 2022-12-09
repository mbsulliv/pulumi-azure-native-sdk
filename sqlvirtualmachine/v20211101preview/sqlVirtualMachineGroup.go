// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A SQL virtual machine group.
type SqlVirtualMachineGroup struct {
	pulumi.CustomResourceState

	// Cluster type.
	ClusterConfiguration pulumi.StringOutput `pulumi:"clusterConfiguration"`
	// Type of cluster manager: Windows Server Failover Cluster (WSFC), implied by the scale type of the group and the OS type.
	ClusterManagerType pulumi.StringOutput `pulumi:"clusterManagerType"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state to track the async operation status.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Scale type.
	ScaleType pulumi.StringOutput `pulumi:"scaleType"`
	// SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
	SqlImageOffer pulumi.StringPtrOutput `pulumi:"sqlImageOffer"`
	// SQL image sku.
	SqlImageSku pulumi.StringPtrOutput `pulumi:"sqlImageSku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Cluster Active Directory domain profile.
	WsfcDomainProfile WsfcDomainProfileResponsePtrOutput `pulumi:"wsfcDomainProfile"`
}

// NewSqlVirtualMachineGroup registers a new resource with the given unique name, arguments, and options.
func NewSqlVirtualMachineGroup(ctx *pulumi.Context,
	name string, args *SqlVirtualMachineGroupArgs, opts ...pulumi.ResourceOption) (*SqlVirtualMachineGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20170301preview:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201preview:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220701preview:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220801preview:SqlVirtualMachineGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlVirtualMachineGroup
	err := ctx.RegisterResource("azure-native:sqlvirtualmachine/v20211101preview:SqlVirtualMachineGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlVirtualMachineGroup gets an existing SqlVirtualMachineGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlVirtualMachineGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlVirtualMachineGroupState, opts ...pulumi.ResourceOption) (*SqlVirtualMachineGroup, error) {
	var resource SqlVirtualMachineGroup
	err := ctx.ReadResource("azure-native:sqlvirtualmachine/v20211101preview:SqlVirtualMachineGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlVirtualMachineGroup resources.
type sqlVirtualMachineGroupState struct {
}

type SqlVirtualMachineGroupState struct {
}

func (SqlVirtualMachineGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlVirtualMachineGroupState)(nil)).Elem()
}

type sqlVirtualMachineGroupArgs struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
	SqlImageOffer *string `pulumi:"sqlImageOffer"`
	// SQL image sku.
	SqlImageSku *string `pulumi:"sqlImageSku"`
	// Name of the SQL virtual machine group.
	SqlVirtualMachineGroupName *string `pulumi:"sqlVirtualMachineGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Cluster Active Directory domain profile.
	WsfcDomainProfile *WsfcDomainProfile `pulumi:"wsfcDomainProfile"`
}

// The set of arguments for constructing a SqlVirtualMachineGroup resource.
type SqlVirtualMachineGroupArgs struct {
	// Resource location.
	Location pulumi.StringPtrInput
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
	SqlImageOffer pulumi.StringPtrInput
	// SQL image sku.
	SqlImageSku pulumi.StringPtrInput
	// Name of the SQL virtual machine group.
	SqlVirtualMachineGroupName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Cluster Active Directory domain profile.
	WsfcDomainProfile WsfcDomainProfilePtrInput
}

func (SqlVirtualMachineGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlVirtualMachineGroupArgs)(nil)).Elem()
}

type SqlVirtualMachineGroupInput interface {
	pulumi.Input

	ToSqlVirtualMachineGroupOutput() SqlVirtualMachineGroupOutput
	ToSqlVirtualMachineGroupOutputWithContext(ctx context.Context) SqlVirtualMachineGroupOutput
}

func (*SqlVirtualMachineGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVirtualMachineGroup)(nil)).Elem()
}

func (i *SqlVirtualMachineGroup) ToSqlVirtualMachineGroupOutput() SqlVirtualMachineGroupOutput {
	return i.ToSqlVirtualMachineGroupOutputWithContext(context.Background())
}

func (i *SqlVirtualMachineGroup) ToSqlVirtualMachineGroupOutputWithContext(ctx context.Context) SqlVirtualMachineGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlVirtualMachineGroupOutput)
}

type SqlVirtualMachineGroupOutput struct{ *pulumi.OutputState }

func (SqlVirtualMachineGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVirtualMachineGroup)(nil)).Elem()
}

func (o SqlVirtualMachineGroupOutput) ToSqlVirtualMachineGroupOutput() SqlVirtualMachineGroupOutput {
	return o
}

func (o SqlVirtualMachineGroupOutput) ToSqlVirtualMachineGroupOutputWithContext(ctx context.Context) SqlVirtualMachineGroupOutput {
	return o
}

// Cluster type.
func (o SqlVirtualMachineGroupOutput) ClusterConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.ClusterConfiguration }).(pulumi.StringOutput)
}

// Type of cluster manager: Windows Server Failover Cluster (WSFC), implied by the scale type of the group and the OS type.
func (o SqlVirtualMachineGroupOutput) ClusterManagerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.ClusterManagerType }).(pulumi.StringOutput)
}

// Resource location.
func (o SqlVirtualMachineGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o SqlVirtualMachineGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state to track the async operation status.
func (o SqlVirtualMachineGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Scale type.
func (o SqlVirtualMachineGroupOutput) ScaleType() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.ScaleType }).(pulumi.StringOutput)
}

// SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
func (o SqlVirtualMachineGroupOutput) SqlImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringPtrOutput { return v.SqlImageOffer }).(pulumi.StringPtrOutput)
}

// SQL image sku.
func (o SqlVirtualMachineGroupOutput) SqlImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringPtrOutput { return v.SqlImageSku }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o SqlVirtualMachineGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SqlVirtualMachineGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o SqlVirtualMachineGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Cluster Active Directory domain profile.
func (o SqlVirtualMachineGroupOutput) WsfcDomainProfile() WsfcDomainProfileResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) WsfcDomainProfileResponsePtrOutput { return v.WsfcDomainProfile }).(WsfcDomainProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlVirtualMachineGroupOutput{})
}
