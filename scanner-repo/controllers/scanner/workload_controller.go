package scanner

import (
	"context"
	"github.com/nats-io/nats.go"
	"kmodules.xyz/client-go/client/duck"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type WorkloadReconciler struct {
	client.Client
	nc *nats.Conn
}

var _ duck.Reconciler = &WorkloadReconciler{}

func NewWorkloadReconciler(nc *nats.Conn) *WorkloadReconciler {
	return &WorkloadReconciler{
		nc: nc,
	}
}

func (r *WorkloadReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	// Get the Workload object
	// If its Kind is Job or CronJob :
	//     Get the Image from all the containers of workloadObject.spec.Template.spec.Containers + InitContainers + EphemeralContainers
	// Otherwise, list the pods with workloadObject's spec.Selectors
	// For all the containers from the PodStatus :
	//     Parse the reference of c.Image & c.ImageID
	//     If parsed Repository is same for this two: use c.Image, otherwise use c.ImageID
	panic("implement me")
}

func (r *WorkloadReconciler) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WorkloadReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// Initiate the Duck Client
	// For(Workload{})
	// WithUnderlyingTypes(Deployment, Statefulset, Job etc)
	// Also set the nats connection
	return nil
}
