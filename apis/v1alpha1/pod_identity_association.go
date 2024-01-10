// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodIdentityAssociationSpec defines the desired state of PodIdentityAssociation.
//
// Amazon EKS Pod Identity associations provide the ability to manage credentials
// for your applications, similar to the way that Amazon EC2 instance profiles
// provide credentials to Amazon EC2 instances.
type PodIdentityAssociationSpec struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `json:"clientRequestToken,omitempty"`
	// The name of the cluster to create the association in.
	ClusterName *string                                  `json:"clusterName,omitempty"`
	ClusterRef  *ackv1alpha1.AWSResourceReferenceWrapper `json:"clusterRef,omitempty"`
	// The name of the Kubernetes namespace inside the cluster to create the association
	// in. The service account and the pods that use the service account must be
	// in this namespace.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with the service
	// account. The EKS Pod Identity agent manages credentials to assume this role
	// for applications in the containers in the pods that use this service account.
	RoleARN *string                                  `json:"roleARN,omitempty"`
	RoleRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"roleRef,omitempty"`
	// The name of the Kubernetes service account inside the cluster to associate
	// the IAM credentials with.
	// +kubebuilder:validation:Required
	ServiceAccount *string `json:"serviceAccount"`
	// Metadata that assists with categorization and organization. Each tag consists
	// of a key and an optional value. You define both. Tags don't propagate to
	// any other cluster or Amazon Web Services resources.
	//
	// The following basic restrictions apply to tags:
	//
	//   - Maximum number of tags per resource – 50
	//
	//   - For each resource, each tag key must be unique, and each tag key can
	//     have only one value.
	//
	//   - Maximum key length – 128 Unicode characters in UTF-8
	//
	//   - Maximum value length – 256 Unicode characters in UTF-8
	//
	//   - If your tagging schema is used across multiple services and resources,
	//     remember that other services may have restrictions on allowed characters.
	//     Generally allowed characters are: letters, numbers, and spaces representable
	//     in UTF-8, and the following characters: + - = . _ : / @.
	//
	//   - Tag keys and values are case-sensitive.
	//
	//   - Do not use aws:, AWS:, or any upper or lowercase combination of such
	//     as a prefix for either keys or values as it is reserved for Amazon Web
	//     Services use. You cannot edit or delete tag keys or values with this prefix.
	//     Tags with this prefix do not count against your tags per resource limit.
	Tags map[string]*string `json:"tags,omitempty"`
}

// PodIdentityAssociationStatus defines the observed state of PodIdentityAssociation
type PodIdentityAssociationStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The Amazon Resource Name (ARN) of the association.
	// +kubebuilder:validation:Optional
	AssociationARN *string `json:"associationARN,omitempty"`
	// The ID of the association.
	// +kubebuilder:validation:Optional
	AssociationID *string `json:"associationID,omitempty"`
	// The timestamp that the association was created at.
	// +kubebuilder:validation:Optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// The most recent timestamp that the association was modified at
	// +kubebuilder:validation:Optional
	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
}

// PodIdentityAssociation is the Schema for the PodIdentityAssociations API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type PodIdentityAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PodIdentityAssociationSpec   `json:"spec,omitempty"`
	Status            PodIdentityAssociationStatus `json:"status,omitempty"`
}

// PodIdentityAssociationList contains a list of PodIdentityAssociation
// +kubebuilder:object:root=true
type PodIdentityAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodIdentityAssociation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodIdentityAssociation{}, &PodIdentityAssociationList{})
}