// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API deployment entity.
type Deployment struct {
	pulumi.CustomResourceState

	// The custom metadata defined for API catalog entities.
	CustomProperties pulumi.AnyOutput `pulumi:"customProperties"`
	// API center-scoped definition resource ID.
	DefinitionId pulumi.StringPtrOutput `pulumi:"definitionId"`
	// Description of the deployment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// API center-scoped environment resource ID.
	EnvironmentId pulumi.StringPtrOutput `pulumi:"environmentId"`
	// The name of the resource
	Name   pulumi.StringOutput               `pulumi:"name"`
	Server DeploymentServerResponsePtrOutput `pulumi:"server"`
	// State of API deployment.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// API deployment title
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiName == nil {
		return nil, errors.New("invalid value for required argument 'ApiName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apicenter:Deployment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Deployment
	err := ctx.RegisterResource("azure-native:apicenter/v20240301:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("azure-native:apicenter/v20240301:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
}

type DeploymentState struct {
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// The name of the API.
	ApiName string `pulumi:"apiName"`
	// The custom metadata defined for API catalog entities.
	CustomProperties interface{} `pulumi:"customProperties"`
	// API center-scoped definition resource ID.
	DefinitionId *string `pulumi:"definitionId"`
	// The name of the API deployment.
	DeploymentName *string `pulumi:"deploymentName"`
	// Description of the deployment.
	Description *string `pulumi:"description"`
	// API center-scoped environment resource ID.
	EnvironmentId *string `pulumi:"environmentId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Server            *DeploymentServer `pulumi:"server"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
	// State of API deployment.
	State *string `pulumi:"state"`
	// API deployment title
	Title *string `pulumi:"title"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// The name of the API.
	ApiName pulumi.StringInput
	// The custom metadata defined for API catalog entities.
	CustomProperties pulumi.Input
	// API center-scoped definition resource ID.
	DefinitionId pulumi.StringPtrInput
	// The name of the API deployment.
	DeploymentName pulumi.StringPtrInput
	// Description of the deployment.
	Description pulumi.StringPtrInput
	// API center-scoped environment resource ID.
	EnvironmentId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	Server            DeploymentServerPtrInput
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput
	// State of API deployment.
	State pulumi.StringPtrInput
	// API deployment title
	Title pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

// The custom metadata defined for API catalog entities.
func (o DeploymentOutput) CustomProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Deployment) pulumi.AnyOutput { return v.CustomProperties }).(pulumi.AnyOutput)
}

// API center-scoped definition resource ID.
func (o DeploymentOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Description of the deployment.
func (o DeploymentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// API center-scoped environment resource ID.
func (o DeploymentOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o DeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentOutput) Server() DeploymentServerResponsePtrOutput {
	return o.ApplyT(func(v *Deployment) DeploymentServerResponsePtrOutput { return v.Server }).(DeploymentServerResponsePtrOutput)
}

// State of API deployment.
func (o DeploymentOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DeploymentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Deployment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// API deployment title
func (o DeploymentOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentOutput{})
}
