// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Deployment stack object.
type DeploymentStackAtResourceGroup struct {
	pulumi.CustomResourceState

	// Defines the behavior of resources that are not managed immediately after the stack is updated.
	ActionOnUnmanage DeploymentStackPropertiesResponseActionOnUnmanageOutput `pulumi:"actionOnUnmanage"`
	// The debug setting of the deployment.
	DebugSetting DeploymentStacksDebugSettingResponsePtrOutput `pulumi:"debugSetting"`
	// An array of resources that were deleted during the most recent update.
	DeletedResources ResourceReferenceResponseArrayOutput `pulumi:"deletedResources"`
	// Defines how resources deployed by the stack are locked.
	DenySettings DenySettingsResponseOutput `pulumi:"denySettings"`
	// The resourceId of the deployment resource created by the deployment stack.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
	DeploymentScope pulumi.StringPtrOutput `pulumi:"deploymentScope"`
	// Deployment stack description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An array of resources that were detached during the most recent update.
	DetachedResources ResourceReferenceResponseArrayOutput `pulumi:"detachedResources"`
	// The duration of the deployment stack update.
	Duration pulumi.StringOutput `pulumi:"duration"`
	// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData error response format.).
	Error ErrorResponseResponsePtrOutput `pulumi:"error"`
	// An array of resources that failed to reach goal state during the most recent update.
	FailedResources ResourceReferenceExtendedResponseArrayOutput `pulumi:"failedResources"`
	// The location of the deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The outputs of the underlying deployment.
	Outputs pulumi.AnyOutput `pulumi:"outputs"`
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
	Parameters pulumi.AnyOutput `pulumi:"parameters"`
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink DeploymentStacksParametersLinkResponsePtrOutput `pulumi:"parametersLink"`
	// State of the deployment stack.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// An array of resources currently managed by the deployment stack.
	Resources ManagedResourceReferenceResponseArrayOutput `pulumi:"resources"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Deployment stack resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDeploymentStackAtResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewDeploymentStackAtResourceGroup(ctx *pulumi.Context,
	name string, args *DeploymentStackAtResourceGroupArgs, opts ...pulumi.ResourceOption) (*DeploymentStackAtResourceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionOnUnmanage == nil {
		return nil, errors.New("invalid value for required argument 'ActionOnUnmanage'")
	}
	if args.DenySettings == nil {
		return nil, errors.New("invalid value for required argument 'DenySettings'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:DeploymentStackAtResourceGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DeploymentStackAtResourceGroup
	err := ctx.RegisterResource("azure-native:resources/v20220801preview:DeploymentStackAtResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentStackAtResourceGroup gets an existing DeploymentStackAtResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentStackAtResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentStackAtResourceGroupState, opts ...pulumi.ResourceOption) (*DeploymentStackAtResourceGroup, error) {
	var resource DeploymentStackAtResourceGroup
	err := ctx.ReadResource("azure-native:resources/v20220801preview:DeploymentStackAtResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentStackAtResourceGroup resources.
type deploymentStackAtResourceGroupState struct {
}

type DeploymentStackAtResourceGroupState struct {
}

func (DeploymentStackAtResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentStackAtResourceGroupState)(nil)).Elem()
}

type deploymentStackAtResourceGroupArgs struct {
	// Defines the behavior of resources that are not managed immediately after the stack is updated.
	ActionOnUnmanage DeploymentStackPropertiesActionOnUnmanage `pulumi:"actionOnUnmanage"`
	// The debug setting of the deployment.
	DebugSetting *DeploymentStacksDebugSetting `pulumi:"debugSetting"`
	// Defines how resources deployed by the stack are locked.
	DenySettings DenySettings `pulumi:"denySettings"`
	// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
	DeploymentScope *string `pulumi:"deploymentScope"`
	// Name of the deployment stack.
	DeploymentStackName *string `pulumi:"deploymentStackName"`
	// Deployment stack description.
	Description *string `pulumi:"description"`
	// The location of the deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
	Location *string `pulumi:"location"`
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
	Parameters interface{} `pulumi:"parameters"`
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink *DeploymentStacksParametersLink `pulumi:"parametersLink"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Deployment stack resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
	Template interface{} `pulumi:"template"`
	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink *DeploymentStacksTemplateLink `pulumi:"templateLink"`
}

// The set of arguments for constructing a DeploymentStackAtResourceGroup resource.
type DeploymentStackAtResourceGroupArgs struct {
	// Defines the behavior of resources that are not managed immediately after the stack is updated.
	ActionOnUnmanage DeploymentStackPropertiesActionOnUnmanageInput
	// The debug setting of the deployment.
	DebugSetting DeploymentStacksDebugSettingPtrInput
	// Defines how resources deployed by the stack are locked.
	DenySettings DenySettingsInput
	// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
	DeploymentScope pulumi.StringPtrInput
	// Name of the deployment stack.
	DeploymentStackName pulumi.StringPtrInput
	// Deployment stack description.
	Description pulumi.StringPtrInput
	// The location of the deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
	Location pulumi.StringPtrInput
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
	Parameters pulumi.Input
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink DeploymentStacksParametersLinkPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Deployment stack resource tags.
	Tags pulumi.StringMapInput
	// The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
	Template pulumi.Input
	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink DeploymentStacksTemplateLinkPtrInput
}

func (DeploymentStackAtResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentStackAtResourceGroupArgs)(nil)).Elem()
}

type DeploymentStackAtResourceGroupInput interface {
	pulumi.Input

	ToDeploymentStackAtResourceGroupOutput() DeploymentStackAtResourceGroupOutput
	ToDeploymentStackAtResourceGroupOutputWithContext(ctx context.Context) DeploymentStackAtResourceGroupOutput
}

func (*DeploymentStackAtResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentStackAtResourceGroup)(nil)).Elem()
}

func (i *DeploymentStackAtResourceGroup) ToDeploymentStackAtResourceGroupOutput() DeploymentStackAtResourceGroupOutput {
	return i.ToDeploymentStackAtResourceGroupOutputWithContext(context.Background())
}

func (i *DeploymentStackAtResourceGroup) ToDeploymentStackAtResourceGroupOutputWithContext(ctx context.Context) DeploymentStackAtResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentStackAtResourceGroupOutput)
}

func (i *DeploymentStackAtResourceGroup) ToOutput(ctx context.Context) pulumix.Output[*DeploymentStackAtResourceGroup] {
	return pulumix.Output[*DeploymentStackAtResourceGroup]{
		OutputState: i.ToDeploymentStackAtResourceGroupOutputWithContext(ctx).OutputState,
	}
}

type DeploymentStackAtResourceGroupOutput struct{ *pulumi.OutputState }

func (DeploymentStackAtResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentStackAtResourceGroup)(nil)).Elem()
}

func (o DeploymentStackAtResourceGroupOutput) ToDeploymentStackAtResourceGroupOutput() DeploymentStackAtResourceGroupOutput {
	return o
}

func (o DeploymentStackAtResourceGroupOutput) ToDeploymentStackAtResourceGroupOutputWithContext(ctx context.Context) DeploymentStackAtResourceGroupOutput {
	return o
}

func (o DeploymentStackAtResourceGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*DeploymentStackAtResourceGroup] {
	return pulumix.Output[*DeploymentStackAtResourceGroup]{
		OutputState: o.OutputState,
	}
}

// Defines the behavior of resources that are not managed immediately after the stack is updated.
func (o DeploymentStackAtResourceGroupOutput) ActionOnUnmanage() DeploymentStackPropertiesResponseActionOnUnmanageOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) DeploymentStackPropertiesResponseActionOnUnmanageOutput {
		return v.ActionOnUnmanage
	}).(DeploymentStackPropertiesResponseActionOnUnmanageOutput)
}

// The debug setting of the deployment.
func (o DeploymentStackAtResourceGroupOutput) DebugSetting() DeploymentStacksDebugSettingResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) DeploymentStacksDebugSettingResponsePtrOutput {
		return v.DebugSetting
	}).(DeploymentStacksDebugSettingResponsePtrOutput)
}

// An array of resources that were deleted during the most recent update.
func (o DeploymentStackAtResourceGroupOutput) DeletedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) ResourceReferenceResponseArrayOutput {
		return v.DeletedResources
	}).(ResourceReferenceResponseArrayOutput)
}

// Defines how resources deployed by the stack are locked.
func (o DeploymentStackAtResourceGroupOutput) DenySettings() DenySettingsResponseOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) DenySettingsResponseOutput { return v.DenySettings }).(DenySettingsResponseOutput)
}

// The resourceId of the deployment resource created by the deployment stack.
func (o DeploymentStackAtResourceGroupOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
func (o DeploymentStackAtResourceGroupOutput) DeploymentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringPtrOutput { return v.DeploymentScope }).(pulumi.StringPtrOutput)
}

// Deployment stack description.
func (o DeploymentStackAtResourceGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An array of resources that were detached during the most recent update.
func (o DeploymentStackAtResourceGroupOutput) DetachedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) ResourceReferenceResponseArrayOutput {
		return v.DetachedResources
	}).(ResourceReferenceResponseArrayOutput)
}

// The duration of the deployment stack update.
func (o DeploymentStackAtResourceGroupOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringOutput { return v.Duration }).(pulumi.StringOutput)
}

// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData error response format.).
func (o DeploymentStackAtResourceGroupOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) ErrorResponseResponsePtrOutput { return v.Error }).(ErrorResponseResponsePtrOutput)
}

// An array of resources that failed to reach goal state during the most recent update.
func (o DeploymentStackAtResourceGroupOutput) FailedResources() ResourceReferenceExtendedResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) ResourceReferenceExtendedResponseArrayOutput {
		return v.FailedResources
	}).(ResourceReferenceExtendedResponseArrayOutput)
}

// The location of the deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
func (o DeploymentStackAtResourceGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of this resource.
func (o DeploymentStackAtResourceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The outputs of the underlying deployment.
func (o DeploymentStackAtResourceGroupOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.AnyOutput { return v.Outputs }).(pulumi.AnyOutput)
}

// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
func (o DeploymentStackAtResourceGroupOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
func (o DeploymentStackAtResourceGroupOutput) ParametersLink() DeploymentStacksParametersLinkResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) DeploymentStacksParametersLinkResponsePtrOutput {
		return v.ParametersLink
	}).(DeploymentStacksParametersLinkResponsePtrOutput)
}

// State of the deployment stack.
func (o DeploymentStackAtResourceGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// An array of resources currently managed by the deployment stack.
func (o DeploymentStackAtResourceGroupOutput) Resources() ManagedResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) ManagedResourceReferenceResponseArrayOutput {
		return v.Resources
	}).(ManagedResourceReferenceResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DeploymentStackAtResourceGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Deployment stack resource tags.
func (o DeploymentStackAtResourceGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of this resource.
func (o DeploymentStackAtResourceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtResourceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentStackAtResourceGroupOutput{})
}
