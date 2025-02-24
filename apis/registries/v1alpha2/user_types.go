package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type UserStatusPhaseName string

const (
	UserStatusPhaseUnknown     UserStatusPhaseName = ""
	UserStatusPhaseCreating    UserStatusPhaseName = "Creating"
	UserStatusPhaseReady       UserStatusPhaseName = "Ready"
	UserStatusPhaseTerminating UserStatusPhaseName = "Terminating"
)

type UserSpec struct {
	// +kubebuilder:validation:Optional
	SysAdmin *bool `json:"sysAdmin,omitempty"`
	// The effective length of the generated user password
	// +kubebuilder:validation:Minimum=8
	// +kubebuilder:validation:Maximum=128
	// +kubebuilder:validation:Optional
	PasswordStrength int32 `json:"passwordStrength"`
	// ParentInstance is a LocalObjectReference to the
	// name of the harbor instance the user is created for
	ParentInstance corev1.LocalObjectReference `json:"parentInstance"`
	Name           string                      `json:"name"`
	RealName       string                      `json:"realname"`
	Email          string                      `json:"email"`
	UserSecretRef  corev1.LocalObjectReference `json:"userSecretRef"`
	// +kubebuilder:validation:Optional
	Comments string `json:"comments,omitempty"`
}

// User is the Schema for the users API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=users,scope=Namespaced,shortName=users;harborusers
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="phase"
// +kubebuilder:object:root=true

type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserSpec   `json:"spec,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

// UserList contains a list of User
// +kubebuilder:object:root=true
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// UserStatus defines the state of a single user
type UserStatus struct {
	Phase UserStatusPhaseName `json:"phase"`
	// +optional
	Message string `json:"message"`
	// +optional
	PasswordHash string `json:"passwordHash"`

	// Time of last observed transition into this state
	// +kubebuilder:validation:Optional
	LastTransition *metav1.Time `json:"lastTransition,omitempty"`
}

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
