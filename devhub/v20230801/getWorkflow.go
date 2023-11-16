// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representation of a workflow
func LookupWorkflow(ctx *pulumi.Context, args *LookupWorkflowArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkflowResult
	err := ctx.Invoke("azure-native:devhub/v20230801:getWorkflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workflow resource.
	WorkflowName string `pulumi:"workflowName"`
}

// Resource representation of a workflow
type LookupWorkflowResult struct {
	// Information on the azure container registry
	Acr *ACRResponse `pulumi:"acr"`
	// The Azure Kubernetes Cluster Resource the application will be deployed to.
	AksResourceId *string `pulumi:"aksResourceId"`
	// The name of the app.
	AppName *string `pulumi:"appName"`
	// Determines the authorization status of requests.
	AuthStatus string `pulumi:"authStatus"`
	// Repository Branch Name
	BranchName *string `pulumi:"branchName"`
	// The version of the language image used for building the code in the generated dockerfile.
	BuilderVersion       *string                       `pulumi:"builderVersion"`
	DeploymentProperties *DeploymentPropertiesResponse `pulumi:"deploymentProperties"`
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
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the image to be generated.
	ImageName *string `pulumi:"imageName"`
	// The tag to apply to the generated image.
	ImageTag *string `pulumi:"imageTag"`
	// The version of the language image used for execution in the generated dockerfile.
	LanguageVersion *string              `pulumi:"languageVersion"`
	LastWorkflowRun *WorkflowRunResponse `pulumi:"lastWorkflowRun"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The mode of generation to be used for generating Manifest.
	ManifestGenerationMode *string `pulumi:"manifestGenerationMode"`
	// The directory to output the generated manifests to.
	ManifestOutputDirectory *string `pulumi:"manifestOutputDirectory"`
	// Determines the type of manifests to be generated.
	ManifestType *string `pulumi:"manifestType"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Kubernetes namespace the application is deployed to.
	Namespace *string `pulumi:"namespace"`
	// The fields needed for OIDC with GitHub.
	OidcCredentials *GitHubWorkflowProfileResponseOidcCredentials `pulumi:"oidcCredentials"`
	// The port the application is exposed on.
	Port *string `pulumi:"port"`
	// The status of the Pull Request submitted against the users repository.
	PrStatus string `pulumi:"prStatus"`
	// The URL to the Pull Request submitted against the users repository.
	PrURL string `pulumi:"prURL"`
	// The number associated with the submitted pull request.
	PullNumber int `pulumi:"pullNumber"`
	// Repository Name
	RepositoryName *string `pulumi:"repositoryName"`
	// Repository Owner
	RepositoryOwner *string `pulumi:"repositoryOwner"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkflowOutput(ctx *pulumi.Context, args LookupWorkflowOutputArgs, opts ...pulumi.InvokeOption) LookupWorkflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkflowResult, error) {
			args := v.(LookupWorkflowArgs)
			r, err := LookupWorkflow(ctx, &args, opts...)
			var s LookupWorkflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkflowResultOutput)
}

type LookupWorkflowOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workflow resource.
	WorkflowName pulumi.StringInput `pulumi:"workflowName"`
}

func (LookupWorkflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowArgs)(nil)).Elem()
}

// Resource representation of a workflow
type LookupWorkflowResultOutput struct{ *pulumi.OutputState }

func (LookupWorkflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowResult)(nil)).Elem()
}

func (o LookupWorkflowResultOutput) ToLookupWorkflowResultOutput() LookupWorkflowResultOutput {
	return o
}

func (o LookupWorkflowResultOutput) ToLookupWorkflowResultOutputWithContext(ctx context.Context) LookupWorkflowResultOutput {
	return o
}

// Information on the azure container registry
func (o LookupWorkflowResultOutput) Acr() ACRResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *ACRResponse { return v.Acr }).(ACRResponsePtrOutput)
}

// The Azure Kubernetes Cluster Resource the application will be deployed to.
func (o LookupWorkflowResultOutput) AksResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.AksResourceId }).(pulumi.StringPtrOutput)
}

// The name of the app.
func (o LookupWorkflowResultOutput) AppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.AppName }).(pulumi.StringPtrOutput)
}

// Determines the authorization status of requests.
func (o LookupWorkflowResultOutput) AuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.AuthStatus }).(pulumi.StringOutput)
}

// Repository Branch Name
func (o LookupWorkflowResultOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.BranchName }).(pulumi.StringPtrOutput)
}

// The version of the language image used for building the code in the generated dockerfile.
func (o LookupWorkflowResultOutput) BuilderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.BuilderVersion }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) DeploymentProperties() DeploymentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *DeploymentPropertiesResponse { return v.DeploymentProperties }).(DeploymentPropertiesResponsePtrOutput)
}

// Path to Dockerfile Build Context within the repository.
func (o LookupWorkflowResultOutput) DockerBuildContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.DockerBuildContext }).(pulumi.StringPtrOutput)
}

// Path to the Dockerfile within the repository.
func (o LookupWorkflowResultOutput) Dockerfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Dockerfile }).(pulumi.StringPtrOutput)
}

// The mode of generation to be used for generating Dockerfiles.
func (o LookupWorkflowResultOutput) DockerfileGenerationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.DockerfileGenerationMode }).(pulumi.StringPtrOutput)
}

// The directory to output the generated Dockerfile to.
func (o LookupWorkflowResultOutput) DockerfileOutputDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.DockerfileOutputDirectory }).(pulumi.StringPtrOutput)
}

// The programming language used.
func (o LookupWorkflowResultOutput) GenerationLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.GenerationLanguage }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkflowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the image to be generated.
func (o LookupWorkflowResultOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

// The tag to apply to the generated image.
func (o LookupWorkflowResultOutput) ImageTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.ImageTag }).(pulumi.StringPtrOutput)
}

// The version of the language image used for execution in the generated dockerfile.
func (o LookupWorkflowResultOutput) LanguageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.LanguageVersion }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) LastWorkflowRun() WorkflowRunResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *WorkflowRunResponse { return v.LastWorkflowRun }).(WorkflowRunResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupWorkflowResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Location }).(pulumi.StringOutput)
}

// The mode of generation to be used for generating Manifest.
func (o LookupWorkflowResultOutput) ManifestGenerationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.ManifestGenerationMode }).(pulumi.StringPtrOutput)
}

// The directory to output the generated manifests to.
func (o LookupWorkflowResultOutput) ManifestOutputDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.ManifestOutputDirectory }).(pulumi.StringPtrOutput)
}

// Determines the type of manifests to be generated.
func (o LookupWorkflowResultOutput) ManifestType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.ManifestType }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupWorkflowResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Name }).(pulumi.StringOutput)
}

// Kubernetes namespace the application is deployed to.
func (o LookupWorkflowResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The fields needed for OIDC with GitHub.
func (o LookupWorkflowResultOutput) OidcCredentials() GitHubWorkflowProfileResponseOidcCredentialsPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *GitHubWorkflowProfileResponseOidcCredentials { return v.OidcCredentials }).(GitHubWorkflowProfileResponseOidcCredentialsPtrOutput)
}

// The port the application is exposed on.
func (o LookupWorkflowResultOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Port }).(pulumi.StringPtrOutput)
}

// The status of the Pull Request submitted against the users repository.
func (o LookupWorkflowResultOutput) PrStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.PrStatus }).(pulumi.StringOutput)
}

// The URL to the Pull Request submitted against the users repository.
func (o LookupWorkflowResultOutput) PrURL() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.PrURL }).(pulumi.StringOutput)
}

// The number associated with the submitted pull request.
func (o LookupWorkflowResultOutput) PullNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkflowResult) int { return v.PullNumber }).(pulumi.IntOutput)
}

// Repository Name
func (o LookupWorkflowResultOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

// Repository Owner
func (o LookupWorkflowResultOutput) RepositoryOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.RepositoryOwner }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWorkflowResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkflowResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupWorkflowResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkflowResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkflowResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkflowResultOutput{})
}
