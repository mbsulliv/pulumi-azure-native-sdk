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
type DeploymentStackAtManagementGroup struct {
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

// NewDeploymentStackAtManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewDeploymentStackAtManagementGroup(ctx *pulumi.Context,
	name string, args *DeploymentStackAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*DeploymentStackAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionOnUnmanage == nil {
		return nil, errors.New("invalid value for required argument 'ActionOnUnmanage'")
	}
	if args.DenySettings == nil {
		return nil, errors.New("invalid value for required argument 'DenySettings'")
	}
	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:DeploymentStackAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DeploymentStackAtManagementGroup
	err := ctx.RegisterResource("azure-native:resources/v20220801preview:DeploymentStackAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentStackAtManagementGroup gets an existing DeploymentStackAtManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentStackAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentStackAtManagementGroupState, opts ...pulumi.ResourceOption) (*DeploymentStackAtManagementGroup, error) {
	var resource DeploymentStackAtManagementGroup
	err := ctx.ReadResource("azure-native:resources/v20220801preview:DeploymentStackAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentStackAtManagementGroup resources.
type deploymentStackAtManagementGroupState struct {
}

type DeploymentStackAtManagementGroupState struct {
}

func (DeploymentStackAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentStackAtManagementGroupState)(nil)).Elem()
}

type deploymentStackAtManagementGroupArgs struct {
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
	// Management Group.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
	Parameters interface{} `pulumi:"parameters"`
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink *DeploymentStacksParametersLink `pulumi:"parametersLink"`
	// Deployment stack resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
	Template interface{} `pulumi:"template"`
	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink *DeploymentStacksTemplateLink `pulumi:"templateLink"`
}

// The set of arguments for constructing a DeploymentStackAtManagementGroup resource.
type DeploymentStackAtManagementGroupArgs struct {
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
	// Management Group.
	ManagementGroupId pulumi.StringInput
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
	Parameters pulumi.Input
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink DeploymentStacksParametersLinkPtrInput
	// Deployment stack resource tags.
	Tags pulumi.StringMapInput
	// The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
	Template pulumi.Input
	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink DeploymentStacksTemplateLinkPtrInput
}

func (DeploymentStackAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentStackAtManagementGroupArgs)(nil)).Elem()
}

type DeploymentStackAtManagementGroupInput interface {
	pulumi.Input

	ToDeploymentStackAtManagementGroupOutput() DeploymentStackAtManagementGroupOutput
	ToDeploymentStackAtManagementGroupOutputWithContext(ctx context.Context) DeploymentStackAtManagementGroupOutput
}

func (*DeploymentStackAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentStackAtManagementGroup)(nil)).Elem()
}

func (i *DeploymentStackAtManagementGroup) ToDeploymentStackAtManagementGroupOutput() DeploymentStackAtManagementGroupOutput {
	return i.ToDeploymentStackAtManagementGroupOutputWithContext(context.Background())
}

func (i *DeploymentStackAtManagementGroup) ToDeploymentStackAtManagementGroupOutputWithContext(ctx context.Context) DeploymentStackAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentStackAtManagementGroupOutput)
}

func (i *DeploymentStackAtManagementGroup) ToOutput(ctx context.Context) pulumix.Output[*DeploymentStackAtManagementGroup] {
	return pulumix.Output[*DeploymentStackAtManagementGroup]{
		OutputState: i.ToDeploymentStackAtManagementGroupOutputWithContext(ctx).OutputState,
	}
}

type DeploymentStackAtManagementGroupOutput struct{ *pulumi.OutputState }

func (DeploymentStackAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentStackAtManagementGroup)(nil)).Elem()
}

func (o DeploymentStackAtManagementGroupOutput) ToDeploymentStackAtManagementGroupOutput() DeploymentStackAtManagementGroupOutput {
	return o
}

func (o DeploymentStackAtManagementGroupOutput) ToDeploymentStackAtManagementGroupOutputWithContext(ctx context.Context) DeploymentStackAtManagementGroupOutput {
	return o
}

func (o DeploymentStackAtManagementGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*DeploymentStackAtManagementGroup] {
	return pulumix.Output[*DeploymentStackAtManagementGroup]{
		OutputState: o.OutputState,
	}
}

// Defines the behavior of resources that are not managed immediately after the stack is updated.
func (o DeploymentStackAtManagementGroupOutput) ActionOnUnmanage() DeploymentStackPropertiesResponseActionOnUnmanageOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) DeploymentStackPropertiesResponseActionOnUnmanageOutput {
		return v.ActionOnUnmanage
	}).(DeploymentStackPropertiesResponseActionOnUnmanageOutput)
}

// The debug setting of the deployment.
func (o DeploymentStackAtManagementGroupOutput) DebugSetting() DeploymentStacksDebugSettingResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) DeploymentStacksDebugSettingResponsePtrOutput {
		return v.DebugSetting
	}).(DeploymentStacksDebugSettingResponsePtrOutput)
}

// An array of resources that were deleted during the most recent update.
func (o DeploymentStackAtManagementGroupOutput) DeletedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ResourceReferenceResponseArrayOutput {
		return v.DeletedResources
	}).(ResourceReferenceResponseArrayOutput)
}

// Defines how resources deployed by the stack are locked.
func (o DeploymentStackAtManagementGroupOutput) DenySettings() DenySettingsResponseOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) DenySettingsResponseOutput { return v.DenySettings }).(DenySettingsResponseOutput)
}

// The resourceId of the deployment resource created by the deployment stack.
func (o DeploymentStackAtManagementGroupOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
func (o DeploymentStackAtManagementGroupOutput) DeploymentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringPtrOutput { return v.DeploymentScope }).(pulumi.StringPtrOutput)
}

// Deployment stack description.
func (o DeploymentStackAtManagementGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An array of resources that were detached during the most recent update.
func (o DeploymentStackAtManagementGroupOutput) DetachedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ResourceReferenceResponseArrayOutput {
		return v.DetachedResources
	}).(ResourceReferenceResponseArrayOutput)
}

// The duration of the deployment stack update.
func (o DeploymentStackAtManagementGroupOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.Duration }).(pulumi.StringOutput)
}

// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData error response format.).
func (o DeploymentStackAtManagementGroupOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ErrorResponseResponsePtrOutput { return v.Error }).(ErrorResponseResponsePtrOutput)
}

// An array of resources that failed to reach goal state during the most recent update.
func (o DeploymentStackAtManagementGroupOutput) FailedResources() ResourceReferenceExtendedResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ResourceReferenceExtendedResponseArrayOutput {
		return v.FailedResources
	}).(ResourceReferenceExtendedResponseArrayOutput)
}

// The location of the deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
func (o DeploymentStackAtManagementGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of this resource.
func (o DeploymentStackAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The outputs of the underlying deployment.
func (o DeploymentStackAtManagementGroupOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.AnyOutput { return v.Outputs }).(pulumi.AnyOutput)
}

// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
func (o DeploymentStackAtManagementGroupOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
func (o DeploymentStackAtManagementGroupOutput) ParametersLink() DeploymentStacksParametersLinkResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) DeploymentStacksParametersLinkResponsePtrOutput {
		return v.ParametersLink
	}).(DeploymentStacksParametersLinkResponsePtrOutput)
}

// State of the deployment stack.
func (o DeploymentStackAtManagementGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// An array of resources currently managed by the deployment stack.
func (o DeploymentStackAtManagementGroupOutput) Resources() ManagedResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ManagedResourceReferenceResponseArrayOutput {
		return v.Resources
	}).(ManagedResourceReferenceResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DeploymentStackAtManagementGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Deployment stack resource tags.
func (o DeploymentStackAtManagementGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of this resource.
func (o DeploymentStackAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentStackAtManagementGroupOutput{})
}
