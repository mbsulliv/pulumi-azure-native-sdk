// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230606

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A cluster resource belonging to a site resource.
type HypervClusterControllerCluster struct {
	pulumi.CustomResourceState

	// Gets the timestamp marking Hyper-V cluster creation.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// Gets the errors.
	Errors HealthErrorDetailsResponseArrayOutput `pulumi:"errors"`
	// Gets or sets the FQDN/IPAddress of the Hyper-V cluster.
	Fqdn pulumi.StringPtrOutput `pulumi:"fqdn"`
	// Gets the functional level of the Hyper-V cluster.
	FunctionalLevel pulumi.IntOutput `pulumi:"functionalLevel"`
	// Gets or sets list of hosts (FQDN) currently being tracked by the cluster.
	HostFqdnList pulumi.StringArrayOutput `pulumi:"hostFqdnList"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets Run as account ID of the Hyper-V cluster.
	RunAsAccountId pulumi.StringPtrOutput `pulumi:"runAsAccountId"`
	// Gets the status of the Hyper-V cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets the timestamp marking last updated on the Hyper-V cluster.
	UpdatedTimestamp pulumi.StringOutput `pulumi:"updatedTimestamp"`
}

// NewHypervClusterControllerCluster registers a new resource with the given unique name, arguments, and options.
func NewHypervClusterControllerCluster(ctx *pulumi.Context,
	name string, args *HypervClusterControllerClusterArgs, opts ...pulumi.ResourceOption) (*HypervClusterControllerCluster, error) {
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
			Type: pulumi.String("azure-native:offazure:HypervClusterControllerCluster"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20231001preview:HypervClusterControllerCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HypervClusterControllerCluster
	err := ctx.RegisterResource("azure-native:offazure/v20230606:HypervClusterControllerCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHypervClusterControllerCluster gets an existing HypervClusterControllerCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHypervClusterControllerCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HypervClusterControllerClusterState, opts ...pulumi.ResourceOption) (*HypervClusterControllerCluster, error) {
	var resource HypervClusterControllerCluster
	err := ctx.ReadResource("azure-native:offazure/v20230606:HypervClusterControllerCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HypervClusterControllerCluster resources.
type hypervClusterControllerClusterState struct {
}

type HypervClusterControllerClusterState struct {
}

func (HypervClusterControllerClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hypervClusterControllerClusterState)(nil)).Elem()
}

type hypervClusterControllerClusterArgs struct {
	//  Cluster ARM name
	ClusterName *string `pulumi:"clusterName"`
	// Gets or sets the FQDN/IPAddress of the Hyper-V cluster.
	Fqdn *string `pulumi:"fqdn"`
	// Gets or sets list of hosts (FQDN) currently being tracked by the cluster.
	HostFqdnList []string `pulumi:"hostFqdnList"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets Run as account ID of the Hyper-V cluster.
	RunAsAccountId *string `pulumi:"runAsAccountId"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// The set of arguments for constructing a HypervClusterControllerCluster resource.
type HypervClusterControllerClusterArgs struct {
	//  Cluster ARM name
	ClusterName pulumi.StringPtrInput
	// Gets or sets the FQDN/IPAddress of the Hyper-V cluster.
	Fqdn pulumi.StringPtrInput
	// Gets or sets list of hosts (FQDN) currently being tracked by the cluster.
	HostFqdnList pulumi.StringArrayInput
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Gets or sets Run as account ID of the Hyper-V cluster.
	RunAsAccountId pulumi.StringPtrInput
	// Site name
	SiteName pulumi.StringInput
}

func (HypervClusterControllerClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hypervClusterControllerClusterArgs)(nil)).Elem()
}

type HypervClusterControllerClusterInput interface {
	pulumi.Input

	ToHypervClusterControllerClusterOutput() HypervClusterControllerClusterOutput
	ToHypervClusterControllerClusterOutputWithContext(ctx context.Context) HypervClusterControllerClusterOutput
}

func (*HypervClusterControllerCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**HypervClusterControllerCluster)(nil)).Elem()
}

func (i *HypervClusterControllerCluster) ToHypervClusterControllerClusterOutput() HypervClusterControllerClusterOutput {
	return i.ToHypervClusterControllerClusterOutputWithContext(context.Background())
}

func (i *HypervClusterControllerCluster) ToHypervClusterControllerClusterOutputWithContext(ctx context.Context) HypervClusterControllerClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HypervClusterControllerClusterOutput)
}

type HypervClusterControllerClusterOutput struct{ *pulumi.OutputState }

func (HypervClusterControllerClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HypervClusterControllerCluster)(nil)).Elem()
}

func (o HypervClusterControllerClusterOutput) ToHypervClusterControllerClusterOutput() HypervClusterControllerClusterOutput {
	return o
}

func (o HypervClusterControllerClusterOutput) ToHypervClusterControllerClusterOutputWithContext(ctx context.Context) HypervClusterControllerClusterOutput {
	return o
}

// Gets the timestamp marking Hyper-V cluster creation.
func (o HypervClusterControllerClusterOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Gets the errors.
func (o HypervClusterControllerClusterOutput) Errors() HealthErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) HealthErrorDetailsResponseArrayOutput { return v.Errors }).(HealthErrorDetailsResponseArrayOutput)
}

// Gets or sets the FQDN/IPAddress of the Hyper-V cluster.
func (o HypervClusterControllerClusterOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

// Gets the functional level of the Hyper-V cluster.
func (o HypervClusterControllerClusterOutput) FunctionalLevel() pulumi.IntOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.IntOutput { return v.FunctionalLevel }).(pulumi.IntOutput)
}

// Gets or sets list of hosts (FQDN) currently being tracked by the cluster.
func (o HypervClusterControllerClusterOutput) HostFqdnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringArrayOutput { return v.HostFqdnList }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o HypervClusterControllerClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o HypervClusterControllerClusterOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets or sets Run as account ID of the Hyper-V cluster.
func (o HypervClusterControllerClusterOutput) RunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringPtrOutput { return v.RunAsAccountId }).(pulumi.StringPtrOutput)
}

// Gets the status of the Hyper-V cluster.
func (o HypervClusterControllerClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o HypervClusterControllerClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HypervClusterControllerClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets the timestamp marking last updated on the Hyper-V cluster.
func (o HypervClusterControllerClusterOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *HypervClusterControllerCluster) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HypervClusterControllerClusterOutput{})
}
