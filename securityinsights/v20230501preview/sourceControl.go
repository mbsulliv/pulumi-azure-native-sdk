// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a SourceControl in Azure Security Insights.
type SourceControl struct {
	pulumi.CustomResourceState

	// Array of source control content types.
	ContentTypes pulumi.StringArrayOutput `pulumi:"contentTypes"`
	// A description of the source control
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the source control
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Information regarding the latest deployment for the source control.
	LastDeploymentInfo DeploymentInfoResponsePtrOutput `pulumi:"lastDeploymentInfo"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The repository type of the source control
	RepoType pulumi.StringOutput `pulumi:"repoType"`
	// Repository metadata.
	Repository RepositoryResponseOutput `pulumi:"repository"`
	// Information regarding the resources created in user's repository.
	RepositoryResourceInfo RepositoryResourceInfoResponsePtrOutput `pulumi:"repositoryResourceInfo"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The version number associated with the source control
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewSourceControl registers a new resource with the given unique name, arguments, and options.
func NewSourceControl(ctx *pulumi.Context,
	name string, args *SourceControlArgs, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentTypes == nil {
		return nil, errors.New("invalid value for required argument 'ContentTypes'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.RepoType == nil {
		return nil, errors.New("invalid value for required argument 'RepoType'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:SourceControl"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SourceControl
	err := ctx.RegisterResource("azure-native:securityinsights/v20230501preview:SourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceControl gets an existing SourceControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceControlState, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	var resource SourceControl
	err := ctx.ReadResource("azure-native:securityinsights/v20230501preview:SourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceControl resources.
type sourceControlState struct {
}

type SourceControlState struct {
}

func (SourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlState)(nil)).Elem()
}

type sourceControlArgs struct {
	// Array of source control content types.
	ContentTypes []string `pulumi:"contentTypes"`
	// A description of the source control
	Description *string `pulumi:"description"`
	// The display name of the source control
	DisplayName string `pulumi:"displayName"`
	// The id (a Guid) of the source control
	Id *string `pulumi:"id"`
	// Information regarding the latest deployment for the source control.
	LastDeploymentInfo *DeploymentInfo `pulumi:"lastDeploymentInfo"`
	// The repository type of the source control
	RepoType string `pulumi:"repoType"`
	// Repository metadata.
	Repository Repository `pulumi:"repository"`
	// Information regarding the resources created in user's repository.
	RepositoryResourceInfo *RepositoryResourceInfo `pulumi:"repositoryResourceInfo"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Source control Id
	SourceControlId *string `pulumi:"sourceControlId"`
	// The version number associated with the source control
	Version *string `pulumi:"version"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SourceControl resource.
type SourceControlArgs struct {
	// Array of source control content types.
	ContentTypes pulumi.StringArrayInput
	// A description of the source control
	Description pulumi.StringPtrInput
	// The display name of the source control
	DisplayName pulumi.StringInput
	// The id (a Guid) of the source control
	Id pulumi.StringPtrInput
	// Information regarding the latest deployment for the source control.
	LastDeploymentInfo DeploymentInfoPtrInput
	// The repository type of the source control
	RepoType pulumi.StringInput
	// Repository metadata.
	Repository RepositoryInput
	// Information regarding the resources created in user's repository.
	RepositoryResourceInfo RepositoryResourceInfoPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Source control Id
	SourceControlId pulumi.StringPtrInput
	// The version number associated with the source control
	Version pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (SourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlArgs)(nil)).Elem()
}

type SourceControlInput interface {
	pulumi.Input

	ToSourceControlOutput() SourceControlOutput
	ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput
}

func (*SourceControl) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControl)(nil)).Elem()
}

func (i *SourceControl) ToSourceControlOutput() SourceControlOutput {
	return i.ToSourceControlOutputWithContext(context.Background())
}

func (i *SourceControl) ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlOutput)
}

type SourceControlOutput struct{ *pulumi.OutputState }

func (SourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControl)(nil)).Elem()
}

func (o SourceControlOutput) ToSourceControlOutput() SourceControlOutput {
	return o
}

func (o SourceControlOutput) ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput {
	return o
}

// Array of source control content types.
func (o SourceControlOutput) ContentTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringArrayOutput { return v.ContentTypes }).(pulumi.StringArrayOutput)
}

// A description of the source control
func (o SourceControlOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the source control
func (o SourceControlOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o SourceControlOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Information regarding the latest deployment for the source control.
func (o SourceControlOutput) LastDeploymentInfo() DeploymentInfoResponsePtrOutput {
	return o.ApplyT(func(v *SourceControl) DeploymentInfoResponsePtrOutput { return v.LastDeploymentInfo }).(DeploymentInfoResponsePtrOutput)
}

// The name of the resource
func (o SourceControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The repository type of the source control
func (o SourceControlOutput) RepoType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringOutput { return v.RepoType }).(pulumi.StringOutput)
}

// Repository metadata.
func (o SourceControlOutput) Repository() RepositoryResponseOutput {
	return o.ApplyT(func(v *SourceControl) RepositoryResponseOutput { return v.Repository }).(RepositoryResponseOutput)
}

// Information regarding the resources created in user's repository.
func (o SourceControlOutput) RepositoryResourceInfo() RepositoryResourceInfoResponsePtrOutput {
	return o.ApplyT(func(v *SourceControl) RepositoryResourceInfoResponsePtrOutput { return v.RepositoryResourceInfo }).(RepositoryResourceInfoResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SourceControlOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SourceControl) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SourceControlOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The version number associated with the source control
func (o SourceControlOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SourceControlOutput{})
}
