// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apicenter

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// API definition entity.
// Azure REST API version: 2024-03-01.
type ApiDefinition struct {
	pulumi.CustomResourceState

	// API definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// API specification details.
	Specification ApiDefinitionPropertiesResponseSpecificationOutput `pulumi:"specification"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// API definition title.
	Title pulumi.StringOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApiDefinition registers a new resource with the given unique name, arguments, and options.
func NewApiDefinition(ctx *pulumi.Context,
	name string, args *ApiDefinitionArgs, opts ...pulumi.ResourceOption) (*ApiDefinition, error) {
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
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.VersionName == nil {
		return nil, errors.New("invalid value for required argument 'VersionName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apicenter/v20240301:ApiDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiDefinition
	err := ctx.RegisterResource("azure-native:apicenter:ApiDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiDefinition gets an existing ApiDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDefinitionState, opts ...pulumi.ResourceOption) (*ApiDefinition, error) {
	var resource ApiDefinition
	err := ctx.ReadResource("azure-native:apicenter:ApiDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiDefinition resources.
type apiDefinitionState struct {
}

type ApiDefinitionState struct {
}

func (ApiDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDefinitionState)(nil)).Elem()
}

type apiDefinitionArgs struct {
	// The name of the API.
	ApiName string `pulumi:"apiName"`
	// The name of the API definition.
	DefinitionName *string `pulumi:"definitionName"`
	// API definition description.
	Description *string `pulumi:"description"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
	// API definition title.
	Title string `pulumi:"title"`
	// The name of the API version.
	VersionName string `pulumi:"versionName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ApiDefinition resource.
type ApiDefinitionArgs struct {
	// The name of the API.
	ApiName pulumi.StringInput
	// The name of the API definition.
	DefinitionName pulumi.StringPtrInput
	// API definition description.
	Description pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput
	// API definition title.
	Title pulumi.StringInput
	// The name of the API version.
	VersionName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ApiDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDefinitionArgs)(nil)).Elem()
}

type ApiDefinitionInput interface {
	pulumi.Input

	ToApiDefinitionOutput() ApiDefinitionOutput
	ToApiDefinitionOutputWithContext(ctx context.Context) ApiDefinitionOutput
}

func (*ApiDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinition)(nil)).Elem()
}

func (i *ApiDefinition) ToApiDefinitionOutput() ApiDefinitionOutput {
	return i.ToApiDefinitionOutputWithContext(context.Background())
}

func (i *ApiDefinition) ToApiDefinitionOutputWithContext(ctx context.Context) ApiDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionOutput)
}

func (i *ApiDefinition) ToOutput(ctx context.Context) pulumix.Output[*ApiDefinition] {
	return pulumix.Output[*ApiDefinition]{
		OutputState: i.ToApiDefinitionOutputWithContext(ctx).OutputState,
	}
}

type ApiDefinitionOutput struct{ *pulumi.OutputState }

func (ApiDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinition)(nil)).Elem()
}

func (o ApiDefinitionOutput) ToApiDefinitionOutput() ApiDefinitionOutput {
	return o
}

func (o ApiDefinitionOutput) ToApiDefinitionOutputWithContext(ctx context.Context) ApiDefinitionOutput {
	return o
}

func (o ApiDefinitionOutput) ToOutput(ctx context.Context) pulumix.Output[*ApiDefinition] {
	return pulumix.Output[*ApiDefinition]{
		OutputState: o.OutputState,
	}
}

// API definition description.
func (o ApiDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ApiDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// API specification details.
func (o ApiDefinitionOutput) Specification() ApiDefinitionPropertiesResponseSpecificationOutput {
	return o.ApplyT(func(v *ApiDefinition) ApiDefinitionPropertiesResponseSpecificationOutput { return v.Specification }).(ApiDefinitionPropertiesResponseSpecificationOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ApiDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApiDefinition) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// API definition title.
func (o ApiDefinitionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDefinition) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiDefinitionOutput{})
}
