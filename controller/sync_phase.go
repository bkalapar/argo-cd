package controller

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	. "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/util/hook"
)

func syncPhases(obj *unstructured.Unstructured) (phases []SyncPhase) {
	for _, hookType := range hook.Hooks(obj) {
		switch hookType {
		case HookTypePreSync, HookTypeSync, HookTypePostSync:
			phases = append(phases, SyncPhase(hookType))
		}
	}
	return phases
}