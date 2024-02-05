// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mobilepacketcore

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure for Operators 5G Core Session Management Function (SMF) Deployment Resource
// Azure REST API version: 2023-10-15-preview.
type SmfDeployment struct {
	pulumi.CustomResourceState

	// Reference to cluster where the Network Function is deployed
	ClusterService pulumi.StringOutput `pulumi:"clusterService"`
	// Azure for Operators 5G Core SMF component parameters
	ComponentParameters pulumi.StringPtrOutput `pulumi:"componentParameters"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Operational status
	OperationalStatus OperationalStatusResponseOutput `pulumi:"operationalStatus"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Release version. This is inherited from the cluster
	ReleaseVersion pulumi.StringOutput `pulumi:"releaseVersion"`
	// Azure for Operators 5G Core SMF secrets parameters
	SecretsParameters pulumi.StringPtrOutput `pulumi:"secretsParameters"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSmfDeployment registers a new resource with the given unique name, arguments, and options.
func NewSmfDeployment(ctx *pulumi.Context,
	name string, args *SmfDeploymentArgs, opts ...pulumi.ResourceOption) (*SmfDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterService == nil {
		return nil, errors.New("invalid value for required argument 'ClusterService'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilepacketcore/v20231015preview:SmfDeployment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SmfDeployment
	err := ctx.RegisterResource("azure-native:mobilepacketcore:SmfDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmfDeployment gets an existing SmfDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmfDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmfDeploymentState, opts ...pulumi.ResourceOption) (*SmfDeployment, error) {
	var resource SmfDeployment
	err := ctx.ReadResource("azure-native:mobilepacketcore:SmfDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmfDeployment resources.
type smfDeploymentState struct {
}

type SmfDeploymentState struct {
}

func (SmfDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*smfDeploymentState)(nil)).Elem()
}

type smfDeploymentArgs struct {
	// Reference to cluster where the Network Function is deployed
	ClusterService string `pulumi:"clusterService"`
	// Azure for Operators 5G Core SMF component parameters
	ComponentParameters *string `pulumi:"componentParameters"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure for Operators 5G Core SMF secrets parameters
	SecretsParameters *string `pulumi:"secretsParameters"`
	// The name of the SmfDeployment
	SmfDeploymentName *string `pulumi:"smfDeploymentName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SmfDeployment resource.
type SmfDeploymentArgs struct {
	// Reference to cluster where the Network Function is deployed
	ClusterService pulumi.StringInput
	// Azure for Operators 5G Core SMF component parameters
	ComponentParameters pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Azure for Operators 5G Core SMF secrets parameters
	SecretsParameters pulumi.StringPtrInput
	// The name of the SmfDeployment
	SmfDeploymentName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SmfDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smfDeploymentArgs)(nil)).Elem()
}

type SmfDeploymentInput interface {
	pulumi.Input

	ToSmfDeploymentOutput() SmfDeploymentOutput
	ToSmfDeploymentOutputWithContext(ctx context.Context) SmfDeploymentOutput
}

func (*SmfDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**SmfDeployment)(nil)).Elem()
}

func (i *SmfDeployment) ToSmfDeploymentOutput() SmfDeploymentOutput {
	return i.ToSmfDeploymentOutputWithContext(context.Background())
}

func (i *SmfDeployment) ToSmfDeploymentOutputWithContext(ctx context.Context) SmfDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmfDeploymentOutput)
}

type SmfDeploymentOutput struct{ *pulumi.OutputState }

func (SmfDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmfDeployment)(nil)).Elem()
}

func (o SmfDeploymentOutput) ToSmfDeploymentOutput() SmfDeploymentOutput {
	return o
}

func (o SmfDeploymentOutput) ToSmfDeploymentOutputWithContext(ctx context.Context) SmfDeploymentOutput {
	return o
}

// Reference to cluster where the Network Function is deployed
func (o SmfDeploymentOutput) ClusterService() pulumi.StringOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringOutput { return v.ClusterService }).(pulumi.StringOutput)
}

// Azure for Operators 5G Core SMF component parameters
func (o SmfDeploymentOutput) ComponentParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringPtrOutput { return v.ComponentParameters }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o SmfDeploymentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SmfDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Operational status
func (o SmfDeploymentOutput) OperationalStatus() OperationalStatusResponseOutput {
	return o.ApplyT(func(v *SmfDeployment) OperationalStatusResponseOutput { return v.OperationalStatus }).(OperationalStatusResponseOutput)
}

// The status of the last operation.
func (o SmfDeploymentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Release version. This is inherited from the cluster
func (o SmfDeploymentOutput) ReleaseVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringOutput { return v.ReleaseVersion }).(pulumi.StringOutput)
}

// Azure for Operators 5G Core SMF secrets parameters
func (o SmfDeploymentOutput) SecretsParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringPtrOutput { return v.SecretsParameters }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SmfDeploymentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SmfDeployment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SmfDeploymentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SmfDeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SmfDeployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SmfDeploymentOutput{})
}
