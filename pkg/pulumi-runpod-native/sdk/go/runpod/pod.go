// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package runpod

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"internal"
)

type Pod struct {
	pulumi.CustomResourceState

	CloudType pulumi.StringOutput `pulumi:"cloudType"`
	GpuCount  pulumi.IntOutput    `pulumi:"gpuCount"`
	GpuTypeId pulumi.StringOutput `pulumi:"gpuTypeId"`
}

// NewPod registers a new resource with the given unique name, arguments, and options.
func NewPod(ctx *pulumi.Context,
	name string, args *PodArgs, opts ...pulumi.ResourceOption) (*Pod, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudType == nil {
		return nil, errors.New("invalid value for required argument 'CloudType'")
	}
	if args.GpuCount == nil {
		return nil, errors.New("invalid value for required argument 'GpuCount'")
	}
	if args.GpuTypeId == nil {
		return nil, errors.New("invalid value for required argument 'GpuTypeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Pod
	err := ctx.RegisterResource("runpod:index:Pod", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPod gets an existing Pod resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPod(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodState, opts ...pulumi.ResourceOption) (*Pod, error) {
	var resource Pod
	err := ctx.ReadResource("runpod:index:Pod", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pod resources.
type podState struct {
}

type PodState struct {
}

func (PodState) ElementType() reflect.Type {
	return reflect.TypeOf((*podState)(nil)).Elem()
}

type podArgs struct {
	CloudType string `pulumi:"cloudType"`
	GpuCount  int    `pulumi:"gpuCount"`
	GpuTypeId string `pulumi:"gpuTypeId"`
}

// The set of arguments for constructing a Pod resource.
type PodArgs struct {
	CloudType pulumi.StringInput
	GpuCount  pulumi.IntInput
	GpuTypeId pulumi.StringInput
}

func (PodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podArgs)(nil)).Elem()
}

type PodInput interface {
	pulumi.Input

	ToPodOutput() PodOutput
	ToPodOutputWithContext(ctx context.Context) PodOutput
}

func (*Pod) ElementType() reflect.Type {
	return reflect.TypeOf((**Pod)(nil)).Elem()
}

func (i *Pod) ToPodOutput() PodOutput {
	return i.ToPodOutputWithContext(context.Background())
}

func (i *Pod) ToPodOutputWithContext(ctx context.Context) PodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodOutput)
}

func (i *Pod) ToOutput(ctx context.Context) pulumix.Output[*Pod] {
	return pulumix.Output[*Pod]{
		OutputState: i.ToPodOutputWithContext(ctx).OutputState,
	}
}

type PodOutput struct{ *pulumi.OutputState }

func (PodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pod)(nil)).Elem()
}

func (o PodOutput) ToPodOutput() PodOutput {
	return o
}

func (o PodOutput) ToPodOutputWithContext(ctx context.Context) PodOutput {
	return o
}

func (o PodOutput) ToOutput(ctx context.Context) pulumix.Output[*Pod] {
	return pulumix.Output[*Pod]{
		OutputState: o.OutputState,
	}
}

func (o PodOutput) CloudType() pulumi.StringOutput {
	return o.ApplyT(func(v *Pod) pulumi.StringOutput { return v.CloudType }).(pulumi.StringOutput)
}

func (o PodOutput) GpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Pod) pulumi.IntOutput { return v.GpuCount }).(pulumi.IntOutput)
}

func (o PodOutput) GpuTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pod) pulumi.StringOutput { return v.GpuTypeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodInput)(nil)).Elem(), &Pod{})
	pulumi.RegisterOutputType(PodOutput{})
}
