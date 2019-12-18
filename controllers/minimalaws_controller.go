/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"github.com/go-logr/logr"
	"github.com/muvaf/crossplane-resourcepacks/pkg/controllers"
	"github.com/muvaf/minimal-aws/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MinimalAWSReconciler reconciles a MinimalAWS object
type MinimalAWSReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

var (
	MinimalAWSKind             = reflect.TypeOf(v1alpha1.MinimalAWS{}).Name()
	MinimalAWSKindAPIVersion   = MinimalAWSKind + "." + v1alpha1.GroupVersion.String()
	MinimalAWSGroupVersionKind = v1alpha1.GroupVersion.WithKind(MinimalAWSKind)
)

func (r *MinimalAWSReconciler) SetupWithManager(mgr ctrl.Manager) error {
	csr := controllers.NewResourcePackReconciler(mgr, MinimalAWSGroupVersionKind)
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.MinimalAWS{}).
		Complete(csr)
}
