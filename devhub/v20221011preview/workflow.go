// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221011preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representation of a workflow
type Workflow struct {
	pulumi.CustomResourceState

	// Information on the azure container registry
	Acr ACRResponsePtrOutput `pulumi:"acr"`
	// The Azure Kubernetes Cluster Resource the application will be deployed to.
	AksResourceId pulumi.StringPtrOutput `pulumi:"aksResourceId"`
	// The name of the app.
	AppName pulumi.StringPtrOutput `pulumi:"appName"`
	// Determines the authorization status of requests.
	AuthStatus pulumi.StringOutput `pulumi:"authStatus"`
	// Repository Branch Name
	BranchName pulumi.StringPtrOutput `pulumi:"branchName"`
	// The version of the language image used for building the code in the generated dockerfile.
	BuilderVersion       pulumi.StringPtrOutput                `pulumi:"builderVersion"`
	DeploymentProperties DeploymentPropertiesResponsePtrOutput `pulumi:"deploymentProperties"`
	// Path to Dockerfile Build Context within the repository.
	DockerBuildContext pulumi.StringPtrOutput `pulumi:"dockerBuildContext"`
	// Path to the Dockerfile within the repository.
	Dockerfile pulumi.StringPtrOutput `pulumi:"dockerfile"`
	// The mode of generation to be used for generating Dockerfiles.
	DockerfileGenerationMode pulumi.StringPtrOutput `pulumi:"dockerfileGenerationMode"`
	// The directory to output the generated Dockerfile to.
	DockerfileOutputDirectory pulumi.StringPtrOutput `pulumi:"dockerfileOutputDirectory"`
	// The programming language used.
	GenerationLanguage pulumi.StringPtrOutput `pulumi:"generationLanguage"`
	// The name of the image to be generated.
	ImageName pulumi.StringPtrOutput `pulumi:"imageName"`
	// The tag to apply to the generated image.
	ImageTag pulumi.StringPtrOutput `pulumi:"imageTag"`
	// The version of the language image used for execution in the generated dockerfile.
	LanguageVersion pulumi.StringPtrOutput       `pulumi:"languageVersion"`
	LastWorkflowRun WorkflowRunResponsePtrOutput `pulumi:"lastWorkflowRun"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The mode of generation to be used for generating Manifest.
	ManifestGenerationMode pulumi.StringPtrOutput `pulumi:"manifestGenerationMode"`
	// The directory to output the generated manifests to.
	ManifestOutputDirectory pulumi.StringPtrOutput `pulumi:"manifestOutputDirectory"`
	// Determines the type of manifests to be generated.
	ManifestType pulumi.StringPtrOutput `pulumi:"manifestType"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Kubernetes namespace the application is deployed to.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The fields needed for OIDC with GitHub.
	OidcCredentials GitHubWorkflowProfileResponseOidcCredentialsPtrOutput `pulumi:"oidcCredentials"`
	// The port the application is exposed on.
	Port pulumi.StringPtrOutput `pulumi:"port"`
	// The status of the Pull Request submitted against the users repository.
	PrStatus pulumi.StringOutput `pulumi:"prStatus"`
	// The URL to the Pull Request submitted against the users repository.
	PrURL pulumi.StringOutput `pulumi:"prURL"`
	// The number associated with the submitted pull request.
	PullNumber pulumi.IntOutput `pulumi:"pullNumber"`
	// Repository Name
	RepositoryName pulumi.StringPtrOutput `pulumi:"repositoryName"`
	// Repository Owner
	RepositoryOwner pulumi.StringPtrOutput `pulumi:"repositoryOwner"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devhub:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:devhub/v20220401preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:devhub/v20230801:Workflow"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Workflow
	err := ctx.RegisterResource("azure-native:devhub/v20221011preview:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azure-native:devhub/v20221011preview:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workflow resources.
type workflowState struct {
}

type WorkflowState struct {
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	// Information on the azure container registry
	Acr *ACR `pulumi:"acr"`
	// The Azure Kubernetes Cluster Resource the application will be deployed to.
	AksResourceId *string `pulumi:"aksResourceId"`
	// The name of the app.
	AppName *string `pulumi:"appName"`
	// Repository Branch Name
	BranchName *string `pulumi:"branchName"`
	// The version of the language image used for building the code in the generated dockerfile.
	BuilderVersion       *string               `pulumi:"builderVersion"`
	DeploymentProperties *DeploymentProperties `pulumi:"deploymentProperties"`
	// Path to Dockerfile Build Context within the repository.
	DockerBuildContext *string `pulumi:"dockerBuildContext"`
	// Path to the Dockerfile within the repository.
	Dockerfile *string `pulumi:"dockerfile"`
	// The mode of generation to be used for generating Dockerfiles.
	DockerfileGenerationMode *string `pulumi:"dockerfileGenerationMode"`
	// The directory to output the generated Dockerfile to.
	DockerfileOutputDirectory *string `pulumi:"dockerfileOutputDirectory"`
	// The programming language used.
	GenerationLanguage *string `pulumi:"generationLanguage"`
	// The name of the image to be generated.
	ImageName *string `pulumi:"imageName"`
	// The tag to apply to the generated image.
	ImageTag *string `pulumi:"imageTag"`
	// The version of the language image used for execution in the generated dockerfile.
	LanguageVersion *string      `pulumi:"languageVersion"`
	LastWorkflowRun *WorkflowRun `pulumi:"lastWorkflowRun"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The mode of generation to be used for generating Manifest.
	ManifestGenerationMode *string `pulumi:"manifestGenerationMode"`
	// The directory to output the generated manifests to.
	ManifestOutputDirectory *string `pulumi:"manifestOutputDirectory"`
	// Determines the type of manifests to be generated.
	ManifestType *string `pulumi:"manifestType"`
	// Kubernetes namespace the application is deployed to.
	Namespace *string `pulumi:"namespace"`
	// The fields needed for OIDC with GitHub.
	OidcCredentials *GitHubWorkflowProfileOidcCredentials `pulumi:"oidcCredentials"`
	// The port the application is exposed on.
	Port *string `pulumi:"port"`
	// Repository Name
	RepositoryName *string `pulumi:"repositoryName"`
	// Repository Owner
	RepositoryOwner *string `pulumi:"repositoryOwner"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workflow resource.
	WorkflowName *string `pulumi:"workflowName"`
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// Information on the azure container registry
	Acr ACRPtrInput
	// The Azure Kubernetes Cluster Resource the application will be deployed to.
	AksResourceId pulumi.StringPtrInput
	// The name of the app.
	AppName pulumi.StringPtrInput
	// Repository Branch Name
	BranchName pulumi.StringPtrInput
	// The version of the language image used for building the code in the generated dockerfile.
	BuilderVersion       pulumi.StringPtrInput
	DeploymentProperties DeploymentPropertiesPtrInput
	// Path to Dockerfile Build Context within the repository.
	DockerBuildContext pulumi.StringPtrInput
	// Path to the Dockerfile within the repository.
	Dockerfile pulumi.StringPtrInput
	// The mode of generation to be used for generating Dockerfiles.
	DockerfileGenerationMode pulumi.StringPtrInput
	// The directory to output the generated Dockerfile to.
	DockerfileOutputDirectory pulumi.StringPtrInput
	// The programming language used.
	GenerationLanguage pulumi.StringPtrInput
	// The name of the image to be generated.
	ImageName pulumi.StringPtrInput
	// The tag to apply to the generated image.
	ImageTag pulumi.StringPtrInput
	// The version of the language image used for execution in the generated dockerfile.
	LanguageVersion pulumi.StringPtrInput
	LastWorkflowRun WorkflowRunPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The mode of generation to be used for generating Manifest.
	ManifestGenerationMode pulumi.StringPtrInput
	// The directory to output the generated manifests to.
	ManifestOutputDirectory pulumi.StringPtrInput
	// Determines the type of manifests to be generated.
	ManifestType pulumi.StringPtrInput
	// Kubernetes namespace the application is deployed to.
	Namespace pulumi.StringPtrInput
	// The fields needed for OIDC with GitHub.
	OidcCredentials GitHubWorkflowProfileOidcCredentialsPtrInput
	// The port the application is exposed on.
	Port pulumi.StringPtrInput
	// Repository Name
	RepositoryName pulumi.StringPtrInput
	// Repository Owner
	RepositoryOwner pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workflow resource.
	WorkflowName pulumi.StringPtrInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}

type WorkflowInput interface {
	pulumi.Input

	ToWorkflowOutput() WorkflowOutput
	ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput
}

func (*Workflow) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

type WorkflowOutput struct{ *pulumi.OutputState }

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

// Information on the azure container registry
func (o WorkflowOutput) Acr() ACRResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ACRResponsePtrOutput { return v.Acr }).(ACRResponsePtrOutput)
}

// The Azure Kubernetes Cluster Resource the application will be deployed to.
func (o WorkflowOutput) AksResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.AksResourceId }).(pulumi.StringPtrOutput)
}

// The name of the app.
func (o WorkflowOutput) AppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.AppName }).(pulumi.StringPtrOutput)
}

// Determines the authorization status of requests.
func (o WorkflowOutput) AuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.AuthStatus }).(pulumi.StringOutput)
}

// Repository Branch Name
func (o WorkflowOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.BranchName }).(pulumi.StringPtrOutput)
}

// The version of the language image used for building the code in the generated dockerfile.
func (o WorkflowOutput) BuilderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.BuilderVersion }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) DeploymentProperties() DeploymentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) DeploymentPropertiesResponsePtrOutput { return v.DeploymentProperties }).(DeploymentPropertiesResponsePtrOutput)
}

// Path to Dockerfile Build Context within the repository.
func (o WorkflowOutput) DockerBuildContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.DockerBuildContext }).(pulumi.StringPtrOutput)
}

// Path to the Dockerfile within the repository.
func (o WorkflowOutput) Dockerfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Dockerfile }).(pulumi.StringPtrOutput)
}

// The mode of generation to be used for generating Dockerfiles.
func (o WorkflowOutput) DockerfileGenerationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.DockerfileGenerationMode }).(pulumi.StringPtrOutput)
}

// The directory to output the generated Dockerfile to.
func (o WorkflowOutput) DockerfileOutputDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.DockerfileOutputDirectory }).(pulumi.StringPtrOutput)
}

// The programming language used.
func (o WorkflowOutput) GenerationLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.GenerationLanguage }).(pulumi.StringPtrOutput)
}

// The name of the image to be generated.
func (o WorkflowOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.ImageName }).(pulumi.StringPtrOutput)
}

// The tag to apply to the generated image.
func (o WorkflowOutput) ImageTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.ImageTag }).(pulumi.StringPtrOutput)
}

// The version of the language image used for execution in the generated dockerfile.
func (o WorkflowOutput) LanguageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.LanguageVersion }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) LastWorkflowRun() WorkflowRunResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) WorkflowRunResponsePtrOutput { return v.LastWorkflowRun }).(WorkflowRunResponsePtrOutput)
}

// The geo-location where the resource lives
func (o WorkflowOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The mode of generation to be used for generating Manifest.
func (o WorkflowOutput) ManifestGenerationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.ManifestGenerationMode }).(pulumi.StringPtrOutput)
}

// The directory to output the generated manifests to.
func (o WorkflowOutput) ManifestOutputDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.ManifestOutputDirectory }).(pulumi.StringPtrOutput)
}

// Determines the type of manifests to be generated.
func (o WorkflowOutput) ManifestType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.ManifestType }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkflowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Kubernetes namespace the application is deployed to.
func (o WorkflowOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The fields needed for OIDC with GitHub.
func (o WorkflowOutput) OidcCredentials() GitHubWorkflowProfileResponseOidcCredentialsPtrOutput {
	return o.ApplyT(func(v *Workflow) GitHubWorkflowProfileResponseOidcCredentialsPtrOutput { return v.OidcCredentials }).(GitHubWorkflowProfileResponseOidcCredentialsPtrOutput)
}

// The port the application is exposed on.
func (o WorkflowOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Port }).(pulumi.StringPtrOutput)
}

// The status of the Pull Request submitted against the users repository.
func (o WorkflowOutput) PrStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.PrStatus }).(pulumi.StringOutput)
}

// The URL to the Pull Request submitted against the users repository.
func (o WorkflowOutput) PrURL() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.PrURL }).(pulumi.StringOutput)
}

// The number associated with the submitted pull request.
func (o WorkflowOutput) PullNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Workflow) pulumi.IntOutput { return v.PullNumber }).(pulumi.IntOutput)
}

// Repository Name
func (o WorkflowOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

// Repository Owner
func (o WorkflowOutput) RepositoryOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.RepositoryOwner }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkflowOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workflow) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o WorkflowOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkflowOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkflowOutput{})
}
