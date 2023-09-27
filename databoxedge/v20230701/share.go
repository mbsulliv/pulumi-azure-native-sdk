// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a share on the  Data Box Edge/Gateway device.
type Share struct {
	pulumi.CustomResourceState

	// Access protocol to be used by the share.
	AccessProtocol pulumi.StringOutput `pulumi:"accessProtocol"`
	// Azure container mapping for the share.
	AzureContainerInfo AzureContainerInfoResponsePtrOutput `pulumi:"azureContainerInfo"`
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights ClientAccessRightResponseArrayOutput `pulumi:"clientAccessRights"`
	// Data policy of the share.
	DataPolicy pulumi.StringPtrOutput `pulumi:"dataPolicy"`
	// Description for the share.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Current monitoring status of the share.
	MonitoringStatus pulumi.StringOutput `pulumi:"monitoringStatus"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Details of the refresh job on this share.
	RefreshDetails RefreshDetailsResponsePtrOutput `pulumi:"refreshDetails"`
	// Share mount point to the role.
	ShareMappings MountPointMapResponseArrayOutput `pulumi:"shareMappings"`
	// Current status of the share.
	ShareStatus pulumi.StringOutput `pulumi:"shareStatus"`
	// Metadata pertaining to creation and last modification of Share
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights UserAccessRightResponseArrayOutput `pulumi:"userAccessRights"`
}

// NewShare registers a new resource with the given unique name, arguments, and options.
func NewShare(ctx *pulumi.Context,
	name string, args *ShareArgs, opts ...pulumi.ResourceOption) (*Share, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessProtocol == nil {
		return nil, errors.New("invalid value for required argument 'AccessProtocol'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.MonitoringStatus == nil {
		return nil, errors.New("invalid value for required argument 'MonitoringStatus'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareStatus == nil {
		return nil, errors.New("invalid value for required argument 'ShareStatus'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:Share"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Share
	err := ctx.RegisterResource("azure-native:databoxedge/v20230701:Share", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShare gets an existing Share resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareState, opts ...pulumi.ResourceOption) (*Share, error) {
	var resource Share
	err := ctx.ReadResource("azure-native:databoxedge/v20230701:Share", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Share resources.
type shareState struct {
}

type ShareState struct {
}

func (ShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareState)(nil)).Elem()
}

type shareArgs struct {
	// Access protocol to be used by the share.
	AccessProtocol string `pulumi:"accessProtocol"`
	// Azure container mapping for the share.
	AzureContainerInfo *AzureContainerInfo `pulumi:"azureContainerInfo"`
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights []ClientAccessRight `pulumi:"clientAccessRights"`
	// Data policy of the share.
	DataPolicy *string `pulumi:"dataPolicy"`
	// Description for the share.
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Current monitoring status of the share.
	MonitoringStatus string `pulumi:"monitoringStatus"`
	// The share name.
	Name *string `pulumi:"name"`
	// Details of the refresh job on this share.
	RefreshDetails *RefreshDetails `pulumi:"refreshDetails"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Current status of the share.
	ShareStatus string `pulumi:"shareStatus"`
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights []UserAccessRight `pulumi:"userAccessRights"`
}

// The set of arguments for constructing a Share resource.
type ShareArgs struct {
	// Access protocol to be used by the share.
	AccessProtocol pulumi.StringInput
	// Azure container mapping for the share.
	AzureContainerInfo AzureContainerInfoPtrInput
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights ClientAccessRightArrayInput
	// Data policy of the share.
	DataPolicy pulumi.StringPtrInput
	// Description for the share.
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// Current monitoring status of the share.
	MonitoringStatus pulumi.StringInput
	// The share name.
	Name pulumi.StringPtrInput
	// Details of the refresh job on this share.
	RefreshDetails RefreshDetailsPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Current status of the share.
	ShareStatus pulumi.StringInput
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights UserAccessRightArrayInput
}

func (ShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareArgs)(nil)).Elem()
}

type ShareInput interface {
	pulumi.Input

	ToShareOutput() ShareOutput
	ToShareOutputWithContext(ctx context.Context) ShareOutput
}

func (*Share) ElementType() reflect.Type {
	return reflect.TypeOf((**Share)(nil)).Elem()
}

func (i *Share) ToShareOutput() ShareOutput {
	return i.ToShareOutputWithContext(context.Background())
}

func (i *Share) ToShareOutputWithContext(ctx context.Context) ShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareOutput)
}

func (i *Share) ToOutput(ctx context.Context) pulumix.Output[*Share] {
	return pulumix.Output[*Share]{
		OutputState: i.ToShareOutputWithContext(ctx).OutputState,
	}
}

type ShareOutput struct{ *pulumi.OutputState }

func (ShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Share)(nil)).Elem()
}

func (o ShareOutput) ToShareOutput() ShareOutput {
	return o
}

func (o ShareOutput) ToShareOutputWithContext(ctx context.Context) ShareOutput {
	return o
}

func (o ShareOutput) ToOutput(ctx context.Context) pulumix.Output[*Share] {
	return pulumix.Output[*Share]{
		OutputState: o.OutputState,
	}
}

// Access protocol to be used by the share.
func (o ShareOutput) AccessProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.AccessProtocol }).(pulumi.StringOutput)
}

// Azure container mapping for the share.
func (o ShareOutput) AzureContainerInfo() AzureContainerInfoResponsePtrOutput {
	return o.ApplyT(func(v *Share) AzureContainerInfoResponsePtrOutput { return v.AzureContainerInfo }).(AzureContainerInfoResponsePtrOutput)
}

// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
func (o ShareOutput) ClientAccessRights() ClientAccessRightResponseArrayOutput {
	return o.ApplyT(func(v *Share) ClientAccessRightResponseArrayOutput { return v.ClientAccessRights }).(ClientAccessRightResponseArrayOutput)
}

// Data policy of the share.
func (o ShareOutput) DataPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Share) pulumi.StringPtrOutput { return v.DataPolicy }).(pulumi.StringPtrOutput)
}

// Description for the share.
func (o ShareOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Share) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Current monitoring status of the share.
func (o ShareOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.MonitoringStatus }).(pulumi.StringOutput)
}

// The object name.
func (o ShareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Details of the refresh job on this share.
func (o ShareOutput) RefreshDetails() RefreshDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Share) RefreshDetailsResponsePtrOutput { return v.RefreshDetails }).(RefreshDetailsResponsePtrOutput)
}

// Share mount point to the role.
func (o ShareOutput) ShareMappings() MountPointMapResponseArrayOutput {
	return o.ApplyT(func(v *Share) MountPointMapResponseArrayOutput { return v.ShareMappings }).(MountPointMapResponseArrayOutput)
}

// Current status of the share.
func (o ShareOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.ShareStatus }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Share
func (o ShareOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Share) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o ShareOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Mapping of users and corresponding access rights on the share (required for SMB protocol).
func (o ShareOutput) UserAccessRights() UserAccessRightResponseArrayOutput {
	return o.ApplyT(func(v *Share) UserAccessRightResponseArrayOutput { return v.UserAccessRights }).(UserAccessRightResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ShareOutput{})
}
