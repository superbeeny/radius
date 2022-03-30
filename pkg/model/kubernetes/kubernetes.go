// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package kubernetes

import (
	"github.com/project-radius/radius/pkg/handlers"
	"github.com/project-radius/radius/pkg/model"
	"github.com/project-radius/radius/pkg/providers"
	"github.com/project-radius/radius/pkg/renderers/containerv1alpha3"
	"github.com/project-radius/radius/pkg/renderers/dapr"
	"github.com/project-radius/radius/pkg/renderers/daprhttproutev1alpha3"
	"github.com/project-radius/radius/pkg/renderers/daprpubsubv1alpha3"
	"github.com/project-radius/radius/pkg/renderers/daprsecretstorev1alpha3"
	"github.com/project-radius/radius/pkg/renderers/daprstatestorev1alpha3"
	"github.com/project-radius/radius/pkg/renderers/extenderv1alpha3"
	"github.com/project-radius/radius/pkg/renderers/gateway"
	"github.com/project-radius/radius/pkg/renderers/httproutev1alpha3"
	"github.com/project-radius/radius/pkg/renderers/microsoftsqlv1alpha3"
	"github.com/project-radius/radius/pkg/renderers/mongodbv1alpha3"
	"github.com/project-radius/radius/pkg/renderers/rabbitmqv1alpha3"
	"github.com/project-radius/radius/pkg/renderers/redisv1alpha3"
	"github.com/project-radius/radius/pkg/renderers/volumev1alpha3"
	"github.com/project-radius/radius/pkg/resourcekinds"
	"github.com/project-radius/radius/pkg/resourcemodel"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewKubernetesModel(k8s client.Client) model.ApplicationModel {

	radiusResources := []model.RadiusResourceModel{
		{
			ResourceType: containerv1alpha3.ResourceType,
			Renderer:     &dapr.Renderer{Inner: &containerv1alpha3.Renderer{}},
		},
		{
			ResourceType: daprhttproutev1alpha3.ResourceType,
			Renderer:     &daprhttproutev1alpha3.Renderer{},
		},
		{
			ResourceType: daprstatestorev1alpha3.ResourceType,
			Renderer:     &daprstatestorev1alpha3.Renderer{StateStores: daprstatestorev1alpha3.SupportedKubernetesStateStoreKindValues},
		},
		{
			ResourceType: daprsecretstorev1alpha3.ResourceType,
			Renderer:     &daprsecretstorev1alpha3.Renderer{SecretStores: daprsecretstorev1alpha3.SupportedKubernetesSecretStoreKindValues},
		},
		{
			ResourceType: mongodbv1alpha3.ResourceType,
			Renderer:     &mongodbv1alpha3.KubernetesRenderer{},
		},
		{
			ResourceType: rabbitmqv1alpha3.ResourceType,
			Renderer:     &rabbitmqv1alpha3.KubernetesRenderer{},
		},
		{
			ResourceType: redisv1alpha3.ResourceType,
			Renderer:     &redisv1alpha3.KubernetesRenderer{},
		},
		{
			ResourceType: httproutev1alpha3.ResourceType,
			Renderer:     &httproutev1alpha3.Renderer{},
		},
		{
			ResourceType: gateway.ResourceType,
			Renderer:     &gateway.Renderer{},
		},
		{
			ResourceType: volumev1alpha3.ResourceType,
			Renderer:     &volumev1alpha3.KubernetesRenderer{VolumeRenderers: nil},
		},
		{
			ResourceType: microsoftsqlv1alpha3.ResourceType,
			Renderer:     &microsoftsqlv1alpha3.Renderer{Kubernetes: true},
		},
		{
			ResourceType: daprpubsubv1alpha3.ResourceType,
			Renderer: &daprpubsubv1alpha3.Renderer{
				PubSubs: daprpubsubv1alpha3.SupportedKubernetesPubSubKindValues,
			},
		},
		{
			ResourceType: extenderv1alpha3.ResourceType,
			Renderer:     &extenderv1alpha3.KubernetesRenderer{},
		},
	}
	outputResources := []model.OutputResourceModel{
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.Deployment,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.Service,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.Secret,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.Gateway,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.KubernetesHTTPRoute,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.DaprPubSubGeneric,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.DaprStateStoreGeneric,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.DaprSecretStoreGeneric,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
		{
			ResourceType: resourcemodel.ResourceType{
				Type:     resourcekinds.SecretProviderClass,
				Provider: providers.ProviderKubernetes,
			},
			ResourceHandler: handlers.NewKubernetesHandler(k8s),
		},
	}

	// Configure the providers supported by the appmodel
	supportedProviders := map[string]bool{
		providers.ProviderKubernetes: true,
	}

	return model.NewModel(radiusResources, outputResources, supportedProviders)
}
