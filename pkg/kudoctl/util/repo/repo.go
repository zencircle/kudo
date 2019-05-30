package repo

import (
	"github.com/kudobuilder/kudo/pkg/apis/kudo/v1alpha1"
)

// Entry represents a collection of parameters for framework repository
type Entry struct {
	Name      string `json:"name"`
	LocalPath string `json:"localPath"`
	URL       string `json:"url"`
}

// 	Metadata for a Framework. This models the structure of a bundle.yaml file.
type Metadata struct {
	// The name of the framework
	Name string `json:"name,omitempty"`
	// A SemVer 2 conformant version string of the framework
	Version string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// The base version of the application enclosed inside of this framework.
	BaseVersion string `json:"baseVersion,omitempty"`
	// The KUDO API Version of this framework.
	ApiVersion string `json:"apiVersion,omitempty"`
	// KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	KubeVersion string `json:"kubeVersion,omitempty"`
	// A SemVer 2 conformant version string of the framework to upgrade from
	UpgradeVersion string `json:"upgradeVersion,omitempty"`
	// Version dependencies to other frameworks
	Dependencies []v1alpha1.FrameworkDependency `json:"dependencies,omitempty"`
	// The URL to a relevant project page, git repo, or contact person
	Home string `json:"home,omitempty"`
	// Source is the URL to the source code of this framework
	Sources []string `json:"sources,omitempty"`
	// A one-sentence description of the framework
	Description string `json:"description,omitempty"`
	// The tag under which the framework can be found, e.g. incubating or stable
	Tag string `json:"tag,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	Maintainers []*Maintainer `json:"maintainers,omitempty"`
	// Whether or not this framework is deprecated
	Deprecated bool `json:"deprecated,omitempty"`
}

// Maintainer describes a Framework maintainer.
type Maintainer struct {
	// Name is a user name or organization name
	Name string `json:"name,omitempty"`
	// Email is an optional email address to contact the named maintainer
	Email string `json:"email,omitempty"`
	// Url is an optional URL to an address for the named maintainer
	Url string `json:"url,omitempty"`
}