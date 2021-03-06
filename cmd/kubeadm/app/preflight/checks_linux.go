// +build linux

/*
Copyright 2019 The Kubernetes Authors.

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

package preflight

import (
	"github.com/pkg/errors"
	"k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/utils/exec"
)

// Check validates if Docker is setup to use systemd as the cgroup driver.
func (idsc IsDockerSystemdCheck) Check() (warnings, errorList []error) {
	driver, err := util.GetCgroupDriverDocker(exec.New())
	if err != nil {
		return nil, []error{err}
	}
	if driver != util.CgroupDriverSystemd {
		err = errors.Errorf("detected %q as the Docker cgroup driver. "+
			"The recommended driver is %q. "+
			"Please follow the guide at https://kubernetes.io/docs/setup/cri/",
			driver,
			util.CgroupDriverSystemd)
		return []error{err}, nil
	}
	return nil, nil
}
