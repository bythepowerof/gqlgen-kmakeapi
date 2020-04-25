package controller

import (
	context "context"
	"fmt"

	"github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/bythepowerof/kmake-controller/gql"
	v11 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type JobType string

const (
	JobTypeJob      JobType = "JOB"
	JobTypeDummy    JobType = "DUMMY"
	JobTypeFilewait JobType = "FILEWAIT"
)

type RunType string

const (
	RunTypeStart   RunType = "START"
	RunTypeRestart RunType = "RESTART"
	RunTypeStop    RunType = "STOP"
	RunTypeDelete  RunType = "DELETE"
	RunTypeCreate  RunType = "CREATE"
	RunTypeReset   RunType = "RESET"
	RunTypeForce   RunType = "FORCE"
)

type NewReset struct {
	Namespace      string `json:"namespace"`
	Kmakescheduler string `json:"kmakescheduler"`
	Full           bool   `json:"full"`
}
type RunLevelIn struct {
	Namespace      string `json:"namespace"`
	Kmakerun       string `json:"kmakerun"`
	Kmakescheduler string `json:"kmakescheduler"`
}
type SubNamespace struct {
	Namespace string `json:"namespace"`
}
type KmakeController interface {
	Namespaces(name *string) ([]*v11.Namespace, error)
	Kmakes(namespace *string, name *string) ([]*v1.Kmake, error)
	Kmakeruns(namespace *string, kmakename *string, jobtype *JobType, name *string) ([]*v1.KmakeRun, error)
	Kmakescheduleruns(namespace string, kmake *string, kmakerun *string, kmakescheduler *string, name *string, runtype *RunType) ([]*v1.KmakeScheduleRun, error)
	Kmakenowschedulers(namespace string, name *string, monitor *string) ([]*v1.KmakeNowScheduler, error)
	CreateScheduleRun(namespace string, kmake *string, kmakerun *string, kmakescheduler *string, runtype *RunType, opts map[string]string) (*v1.KmakeScheduleRun, error)
	// AddChangeClient(ctx context.Context, namespace string) (<-chan gql.KmakeObject, error)
	AddListener()
	GetListener() *gql.KmakeListener
}

type KubernetesController struct {
	client    client.Client
	namespace string
	Listener  *gql.KmakeListener
}

func NewKubernetesController(client client.Client, namespace string) *KubernetesController {

	return &KubernetesController{
		client:    client,
		namespace: namespace,
	}
}

func (r *KubernetesController) AddListener() {
	r.Listener = gql.NewKmakeListener(r.namespace, nil)
}

func (r *KubernetesController) GetListener() *gql.KmakeListener {
	return r.Listener
}

func (r *KubernetesController) Namespaces(name *string) ([]*v11.Namespace, error) {

	ret := make([]*v11.Namespace, 0)

	nsList := &v11.NamespaceList{}
	o := &client.ListOptions{}

	if r.namespace != "all" && r.namespace != *name {
		return nil, fmt.Errorf("namespace %q not supported", *name)
	}

	if name != nil {
		fields := map[string]string{"metadata.name": *name}
		client.MatchingFields(fields).ApplyToList(o)
	}

	err := r.client.List(context.Background(), nsList, o)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(nsList.Items); i++ {
		ret = append(ret, &nsList.Items[i])
	}
	return ret, nil
}

func (r *KubernetesController) Kmakes(namespace *string, name *string) ([]*v1.Kmake, error) {
	ret := make([]*v1.Kmake, 0)

	if r.namespace != "all" && r.namespace != *namespace {
		return nil, fmt.Errorf("namespace %q not supported", *namespace)
	}

	kmakeList := &v1.KmakeList{}
	o := &client.ListOptions{}
	client.InNamespace(*namespace).ApplyToList(o)

	if name != nil {
		fields := map[string]string{"metadata.name": *name}
		client.MatchingFields(fields).ApplyToList(o)
	}

	err := r.client.List(context.Background(), kmakeList, o)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(kmakeList.Items); i++ {
		ret = append(ret, &kmakeList.Items[i])
	}
	return ret, nil
}

func (r *KubernetesController) Kmakeruns(namespace *string, kmakename *string, jobtype *JobType, name *string) ([]*v1.KmakeRun, error) {
	ret := make([]*v1.KmakeRun, 0)

	if r.namespace != "all" && r.namespace != *namespace {
		return nil, fmt.Errorf("namespace %q not supported", *namespace)
	}

	kmakerunList := &v1.KmakeRunList{}

	o := &client.ListOptions{}
	client.InNamespace(*namespace).ApplyToList(o)

	if name != nil {
		fields := map[string]string{"metadata.name": *name}
		client.MatchingFields(fields).ApplyToList(o)
	}

	if kmakename != nil {
		labels := map[string]string{"bythepowerof.github.io/kmake": *kmakename}
		client.MatchingLabels(labels).ApplyToList(o)
	}

	err := r.client.List(context.Background(), kmakerunList, o)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(kmakerunList.Items); i++ {
		if jobtype != nil {
			if *jobtype == "DUMMY" && kmakerunList.Items[i].Spec.KmakeRunOperation.Dummy == nil {
				continue
			}
			if *jobtype == "JOB" && kmakerunList.Items[i].Spec.KmakeRunOperation.Job == nil {
				continue
			}
			if *jobtype == "FILEWAIT" && kmakerunList.Items[i].Spec.KmakeRunOperation.FileWait == nil {
				continue
			}
		}
		ret = append(ret, &kmakerunList.Items[i])
	}
	return ret, nil
}

func (r *KubernetesController) Kmakescheduleruns(namespace string, kmake *string, kmakerun *string, kmakescheduler *string, name *string, runtype *RunType) ([]*v1.KmakeScheduleRun, error) {
	ret := make([]*v1.KmakeScheduleRun, 0)

	if r.namespace != "all" && r.namespace != namespace {
		return nil, fmt.Errorf("namespace %s not supported", namespace)
	}

	kmakeschedulerunList := &v1.KmakeScheduleRunList{}

	o := &client.ListOptions{}
	client.InNamespace(namespace).ApplyToList(o)

	if name != nil {
		fields := map[string]string{"metadata.name": *name}
		client.MatchingFields(fields).ApplyToList(o)
	}

	if kmake != nil {
		labels := map[string]string{"bythepowerof.github.io/kmake": *kmake}
		client.MatchingLabels(labels).ApplyToList(o)
	}

	if kmakerun != nil {
		labels := map[string]string{"bythepowerof.github.io/run": *kmakerun}
		client.MatchingLabels(labels).ApplyToList(o)
	}

	if kmakescheduler != nil {
		labels := map[string]string{"bythepowerof.github.io/schedule-instance": *kmakescheduler}
		client.MatchingLabels(labels).ApplyToList(o)
	}
	err := r.client.List(context.Background(), kmakeschedulerunList, o)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(kmakeschedulerunList.Items); i++ {
		if runtype != nil {
			if *runtype == "START" && kmakeschedulerunList.Items[i].Spec.KmakeScheduleRunOperation.Start == nil {
				continue
			}
			if *runtype == "RESTART" && kmakeschedulerunList.Items[i].Spec.KmakeScheduleRunOperation.Restart == nil {
				continue
			}
			if *runtype == "STOP" && kmakeschedulerunList.Items[i].Spec.KmakeScheduleRunOperation.Stop == nil {
				continue
			}
			if *runtype == "DELETE" && kmakeschedulerunList.Items[i].Spec.KmakeScheduleRunOperation.Delete == nil {
				continue
			}
			if *runtype == "CREATE" && kmakeschedulerunList.Items[i].Spec.KmakeScheduleRunOperation.Create == nil {
				continue
			}
			if *runtype == "RESET" && kmakeschedulerunList.Items[i].Spec.KmakeScheduleRunOperation.Reset == nil {
				continue
			}
			if *runtype == "FORCE" && kmakeschedulerunList.Items[i].Spec.KmakeScheduleRunOperation.Force == nil {
				continue
			}
		}
		ret = append(ret, &kmakeschedulerunList.Items[i])
	}
	return ret, nil
}

func (r *KubernetesController) Kmakenowschedulers(namespace string, name *string, monitor *string) ([]*v1.KmakeNowScheduler, error) {
	ret := make([]*v1.KmakeNowScheduler, 0)

	if r.namespace != "all" && r.namespace != namespace {
		return nil, fmt.Errorf("namespace %s not supported", namespace)
	}

	kmakeNowSchedulerList := &v1.KmakeNowSchedulerList{}
	o := &client.ListOptions{}
	client.InNamespace(namespace).ApplyToList(o)

	if name != nil {
		fields := map[string]string{"metadata.name": *name}
		client.MatchingFields(fields).ApplyToList(o)
	}

	err := r.client.List(context.Background(), kmakeNowSchedulerList, o)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(kmakeNowSchedulerList.Items); i++ {
		if monitor == nil {
			ret = append(ret, &kmakeNowSchedulerList.Items[i])
		} else {
			for _, m := range kmakeNowSchedulerList.Items[i].Spec.Monitor {
				if m == *monitor {
					ret = append(ret, &kmakeNowSchedulerList.Items[i])
					break
				}
			}
		}
	}
	return ret, nil
}

func (r *KubernetesController) CreateScheduleRun(namespace string, kmake *string, kmakerun *string, kmakescheduler *string, runtype *RunType, opts map[string]string) (*v1.KmakeScheduleRun, error) {
	// make sure the scheduler exists...

	// create a rset job for it

	if r.namespace != "all" && r.namespace != namespace {
		return nil, fmt.Errorf("namespace %s not supported", namespace)
	}

	op := v1.KmakeScheduleRunOperation{}

	kmakeschedulerun := &v1.KmakeScheduleRun{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Labels: map[string]string{
				"bythepowerof.github.io/workload": "no"},
		},
	}

	switch *runtype {
	case RunTypeReset:
		full := "no"
		if opts["full"] == "true" {
			full = "yes"
		}
		op.Reset = &v1.KmakeScheduleReset{Full: full}
		kmakeschedulerun.ObjectMeta.GenerateName = "kmakenowscheduler-reset-gql-"
		if kmakescheduler != nil {
			kmakeschedulerun.Labels["bythepowerof.github.io/schedule-instance"] = *kmakescheduler
		}
	case RunTypeStop:
		op.Stop = &v1.KmakeScheduleRunStop{}
		kmakeschedulerun.ObjectMeta.GenerateName = "kmakenowscheduler-stop-gql-"
		if kmakescheduler != nil {
			kmakeschedulerun.Labels["bythepowerof.github.io/schedule-instance"] = *kmakescheduler
		}
		if kmakerun != nil {
			kmakeschedulerun.Labels["bythepowerof.github.io/run"] = *kmakerun
		}
	case RunTypeRestart:
		op.Restart = &v1.KmakeScheduleRunRestart{
			Run: *kmakerun,
		}
		kmakeschedulerun.ObjectMeta.GenerateName = "kmakenowscheduler-restart-gql-"
		if kmakescheduler != nil {
			kmakeschedulerun.Labels["bythepowerof.github.io/schedule-instance"] = *kmakescheduler
		}
	}

	kmakeschedulerun.Spec.KmakeScheduleRunOperation = op

	err := r.client.Create(context.Background(), kmakeschedulerun)
	if err != nil {
		return nil, err
	}
	return kmakeschedulerun, nil
}

// func (r *KubernetesController) AddChangeClient(ctx context.Context, namespace string) (<-chan gql.KmakeObject, error) {
// 	if r.namespace != "all" && r.namespace != namespace {
// 		return nil, fmt.Errorf("namespace %s not supported", namespace)
// 	}

// 	kmo := make(chan gql.KmakeObject, 1)
// 	r.mutex.Lock()
// 	r.index++

// 	if _, ok := r.changes[namespace]; !ok {
// 		r.changes[namespace] = make(map[int]chan gql.KmakeObject)
// 	}

// 	r.changes[namespace][r.index] = kmo
// 	r.mutex.Unlock()

// 	// ret := []gql.KmakeObject{}

// 	// // push the current state to the subriber
// 	// kms, _ := r.Kmakenowschedulers(namespace, nil, nil)
// 	// for _, v := range kms {
// 	// 	ret = append(ret, v)
// 	// }

// 	// km, _ := r.Kmakes(&namespace, nil)
// 	// for _, v := range km {
// 	// 	ret = append(ret, v)
// 	// }

// 	// kmr, _ := r.Kmakeruns(&namespace, nil, nil, nil)
// 	// for _, v := range kmr {
// 	// 	ret = append(ret, v)
// 	// }

// 	// kmsr, _ := r.Kmakescheduleruns(namespace, nil, nil, nil, nil, nil)

// 	// for _, v := range kmsr {
// 	// 	ret = append(ret, v)
// 	// }

// 	// for _, v := range ret {
// 	// 	kmo <- v
// 	// }

// 	// Delete channel when done
// 	go func() {
// 		<-ctx.Done()
// 		r.mutex.Lock()
// 		delete(r.changes[r.namespace], r.index)
// 		r.mutex.Unlock()
// 	}()
// 	return kmo, nil
// }

// func (r *KubernetesController) KmakeChanges(namespace string) error {
// 	// Create a new Controller that will call the provided Reconciler function in response
// 	// to events.

// 	if r.namespace != "all" && r.namespace != namespace {
// 		return fmt.Errorf("namespace %q not supported", namespace)
// 	}

// 	err := r.prepareKmakeWatch()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = r.prepareKmakeRunWatch()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = r.prepareKmakeScheduleRunWatch()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// err := r.prepareKmakeNowSchedulerWatch()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// Start the Controllers through the manager.
// 	go func() {
// 		if err := r.manager.Start(signals.SetupSignalHandler()); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	return nil
// }

// func (r *KubernetesController) prepareKmakeWatch() error {
// 	c, err := controller.New("kmake-watch", r.manager, controller.Options{
// 		Reconciler: reconcile.Func(r.watchKmake),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	// Watch for kmake objects create / update / delete events and call Reconcile
// 	return c.Watch(&source.Kind{Type: &v1.Kmake{}}, &handler.EnqueueRequestForObject{})
// }

// func (r *KubernetesController) watchKmake(o reconcile.Request) (reconcile.Result, error) {
// 	// Your business logic to implement the API by creating, updating, deleting objects goes here.
// 	ret := &v1.Kmake{}

// 	err := r.client.Get(context.Background(), o.NamespacedName, ret)
// 	if err != nil {
// 		// if errors.IsNotFound(err) {
// 		// 	return reconcile.Result{}, nil
// 		// }
// 		return reconcile.Result{}, err
// 	}
// 	if ret.IsBeingDeleted() {
// 		ret.Status.Status = "Deleting"
// 	}

// 	// Notify new message
// 	r.mutex.Lock()
// 	for _, ch := range r.changes[o.Namespace] {
// 		ch <- ret
// 	}
// 	r.mutex.Unlock()
// 	return reconcile.Result{}, nil
// }

// func (r *KubernetesController) prepareKmakeRunWatch() error {
// 	c, err := controller.New("kmakerun-watch", r.manager, controller.Options{
// 		Reconciler: reconcile.Func(r.watchKmakeRun),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	// Watch for kmake objects create / update / delete events and call Reconcile
// 	return c.Watch(&source.Kind{Type: &v1.KmakeRun{}}, &handler.EnqueueRequestForObject{})
// }

// func (r *KubernetesController) watchKmakeRun(o reconcile.Request) (reconcile.Result, error) {
// 	// Your business logic to implement the API by creating, updating, deleting objects goes here.
// 	ret := &v1.KmakeRun{}

// 	err := r.client.Get(context.Background(), o.NamespacedName, ret)
// 	if err != nil {
// 		// if errors.IsNotFound(err) {
// 		// 	return reconcile.Result{}, nil
// 		// }
// 		return reconcile.Result{}, err
// 	}
// 	if ret.IsBeingDeleted() {
// 		ret.Status.Status = "Deleting"
// 	}

// 	// Notify new message
// 	r.mutex.Lock()
// 	for _, ch := range r.changes[o.Namespace] {
// 		ch <- ret
// 	}
// 	r.mutex.Unlock()
// 	return reconcile.Result{}, nil
// }

// func (r *KubernetesController) prepareKmakeScheduleRunWatch() error {
// 	c, err := controller.New("kmakeschedulerun-watch", r.manager, controller.Options{
// 		Reconciler: reconcile.Func(r.watchKmakeScheduleRun),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	// Watch for kmake objects create / update / delete events and call Reconcile
// 	return c.Watch(&source.Kind{Type: &v1.KmakeScheduleRun{}}, &handler.EnqueueRequestForObject{})
// }

// func (r *KubernetesController) watchKmakeScheduleRun(o reconcile.Request) (reconcile.Result, error) {
// 	// Your business logic to implement the API by creating, updating, deleting objects goes here.
// 	ret := &v1.KmakeScheduleRun{}

// 	err := r.client.Get(context.Background(), o.NamespacedName, ret)
// 	if err != nil {
// 		// if errors.IsNotFound(err) {
// 		// 	return reconcile.Result{}, nil
// 		// }
// 		return reconcile.Result{}, err
// 	}
// 	if ret.IsBeingDeleted() {
// 		ret.Status.Status = "Deleting"
// 	}

// 	// Notify new message
// 	r.mutex.Lock()
// 	for _, ch := range r.changes[o.Namespace] {
// 		ch <- ret
// 	}
// 	r.mutex.Unlock()
// 	return reconcile.Result{}, nil
// }
