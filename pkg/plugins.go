package pkg

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

const Name = "sample-plugins"

type Args struct {
	FavoriteColor  string `json:"favorite_color,omitempty"`
	FavoriteNumber int    `json:"favorite_number,omitempty"`
	ThanksTo       string `json:"thanks_to,omitempty"`
}

type Sample struct {
	args   *Args
	handle framework.FrameworkHandle
}

func (s *Sample) Name() string {
	return Name
}

func (s *Sample) PreFilter(ctx context.Context, state *framework.CycleState, p *corev1.Pod) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v", p.Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

func (s *Sample) Filter(ctx context.Context, state *framework.CycleState, pod *corev1.Pod, nodeInfo *schedulernodeinfo.NodeInfo) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v\tnode: %v", pod.Name, nodeInfo.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) PreBind(ctx context.Context, state *framework.CycleState, p *corev1.Pod, nodeName string) *framework.Status {
	node, err := s.handle.ClientSet().CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("get node error: %v", nodeName))
	} else {
		klog.V(3).Infof("prebind node info: %+v", node)
		return framework.NewStatus(framework.Success, "")
	}
}

func New(config *runtime.Unknown, f framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(config, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("get plugin config args: %+v", args)
	return &Sample{
		args:   args,
		handle: f,
	}, nil
}
