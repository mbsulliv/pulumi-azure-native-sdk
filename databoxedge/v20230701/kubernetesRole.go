// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The limited preview of Kubernetes Cluster Management from the Azure supports:
//  1. Using a simple turn-key option in Azure Portal, deploy a Kubernetes cluster on your Azure Stack Edge device.
//  2. Configure Kubernetes cluster running on your device with Arc enabled Kubernetes with a click of a button in the Azure Portal.
//     Azure Arc enables organizations to view, manage, and govern their on-premises Kubernetes clusters using the Azure Portal, command line tools, and APIs.
//  3. Easily configure Persistent Volumes using SMB and NFS shares for storing container data.
//     For more information, refer to the document here: https://databoxupdatepackages.blob.core.windows.net/documentation/Microsoft-Azure-Stack-Edge-K8-Cloud-Management-20210323.pdf
//     Or Demo: https://databoxupdatepackages.blob.core.windows.net/documentation/Microsoft-Azure-Stack-Edge-K8S-Cloud-Management-20210323.mp4
//     By using this feature, you agree to the preview legal terms. See the https://azure.microsoft.com/en-us/support/legal/preview-supplemental-terms/
type KubernetesRole struct {
	pulumi.CustomResourceState

	// Host OS supported by the Kubernetes role.
	HostPlatform pulumi.StringOutput `pulumi:"hostPlatform"`
	// Platform where the runtime is hosted.
	HostPlatformType pulumi.StringOutput `pulumi:"hostPlatformType"`
	// Role type.
	// Expected value is 'Kubernetes'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Kubernetes cluster configuration
	KubernetesClusterInfo KubernetesClusterInfoResponseOutput `pulumi:"kubernetesClusterInfo"`
	// Kubernetes role resources
	KubernetesRoleResources KubernetesRoleResourcesResponseOutput `pulumi:"kubernetesRoleResources"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of Kubernetes deployment
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Role status.
	RoleStatus pulumi.StringOutput `pulumi:"roleStatus"`
	// Metadata pertaining to creation and last modification of Role
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKubernetesRole registers a new resource with the given unique name, arguments, and options.
func NewKubernetesRole(ctx *pulumi.Context,
	name string, args *KubernetesRoleArgs, opts ...pulumi.ResourceOption) (*KubernetesRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.HostPlatform == nil {
		return nil, errors.New("invalid value for required argument 'HostPlatform'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KubernetesClusterInfo == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesClusterInfo'")
	}
	if args.KubernetesRoleResources == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesRoleResources'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RoleStatus == nil {
		return nil, errors.New("invalid value for required argument 'RoleStatus'")
	}
	args.Kind = pulumi.String("Kubernetes")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:KubernetesRole"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KubernetesRole
	err := ctx.RegisterResource("azure-native:databoxedge/v20230701:KubernetesRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesRole gets an existing KubernetesRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesRoleState, opts ...pulumi.ResourceOption) (*KubernetesRole, error) {
	var resource KubernetesRole
	err := ctx.ReadResource("azure-native:databoxedge/v20230701:KubernetesRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesRole resources.
type kubernetesRoleState struct {
}

type KubernetesRoleState struct {
}

func (KubernetesRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesRoleState)(nil)).Elem()
}

type kubernetesRoleArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Host OS supported by the Kubernetes role.
	HostPlatform string `pulumi:"hostPlatform"`
	// Role type.
	// Expected value is 'Kubernetes'.
	Kind string `pulumi:"kind"`
	// Kubernetes cluster configuration
	KubernetesClusterInfo KubernetesClusterInfo `pulumi:"kubernetesClusterInfo"`
	// Kubernetes role resources
	KubernetesRoleResources KubernetesRoleResources `pulumi:"kubernetesRoleResources"`
	// The role name.
	Name *string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Role status.
	RoleStatus string `pulumi:"roleStatus"`
}

// The set of arguments for constructing a KubernetesRole resource.
type KubernetesRoleArgs struct {
	// The device name.
	DeviceName pulumi.StringInput
	// Host OS supported by the Kubernetes role.
	HostPlatform pulumi.StringInput
	// Role type.
	// Expected value is 'Kubernetes'.
	Kind pulumi.StringInput
	// Kubernetes cluster configuration
	KubernetesClusterInfo KubernetesClusterInfoInput
	// Kubernetes role resources
	KubernetesRoleResources KubernetesRoleResourcesInput
	// The role name.
	Name pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Role status.
	RoleStatus pulumi.StringInput
}

func (KubernetesRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesRoleArgs)(nil)).Elem()
}

type KubernetesRoleInput interface {
	pulumi.Input

	ToKubernetesRoleOutput() KubernetesRoleOutput
	ToKubernetesRoleOutputWithContext(ctx context.Context) KubernetesRoleOutput
}

func (*KubernetesRole) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRole)(nil)).Elem()
}

func (i *KubernetesRole) ToKubernetesRoleOutput() KubernetesRoleOutput {
	return i.ToKubernetesRoleOutputWithContext(context.Background())
}

func (i *KubernetesRole) ToKubernetesRoleOutputWithContext(ctx context.Context) KubernetesRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleOutput)
}

type KubernetesRoleOutput struct{ *pulumi.OutputState }

func (KubernetesRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRole)(nil)).Elem()
}

func (o KubernetesRoleOutput) ToKubernetesRoleOutput() KubernetesRoleOutput {
	return o
}

func (o KubernetesRoleOutput) ToKubernetesRoleOutputWithContext(ctx context.Context) KubernetesRoleOutput {
	return o
}

// Host OS supported by the Kubernetes role.
func (o KubernetesRoleOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesRole) pulumi.StringOutput { return v.HostPlatform }).(pulumi.StringOutput)
}

// Platform where the runtime is hosted.
func (o KubernetesRoleOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesRole) pulumi.StringOutput { return v.HostPlatformType }).(pulumi.StringOutput)
}

// Role type.
// Expected value is 'Kubernetes'.
func (o KubernetesRoleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesRole) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Kubernetes cluster configuration
func (o KubernetesRoleOutput) KubernetesClusterInfo() KubernetesClusterInfoResponseOutput {
	return o.ApplyT(func(v *KubernetesRole) KubernetesClusterInfoResponseOutput { return v.KubernetesClusterInfo }).(KubernetesClusterInfoResponseOutput)
}

// Kubernetes role resources
func (o KubernetesRoleOutput) KubernetesRoleResources() KubernetesRoleResourcesResponseOutput {
	return o.ApplyT(func(v *KubernetesRole) KubernetesRoleResourcesResponseOutput { return v.KubernetesRoleResources }).(KubernetesRoleResourcesResponseOutput)
}

// The object name.
func (o KubernetesRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of Kubernetes deployment
func (o KubernetesRoleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesRole) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Role status.
func (o KubernetesRoleOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesRole) pulumi.StringOutput { return v.RoleStatus }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Role
func (o KubernetesRoleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KubernetesRole) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o KubernetesRoleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesRole) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KubernetesRoleOutput{})
}
