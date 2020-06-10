package ormb

import (
	v1 "k8s.io/api/core/v1"
)

const (
	ORMBUsername = "ORMB_USERNAME"
	ORMBPassword = "ORMB_PASSWORD"

	ORMBUsernameName = "ormbUsername"
	ORMBPasswordName = "ormbPassword"
)

type ORMBConfig struct {
	ORMBUsernameName string `json:"ormbUsername,omitempty"`
	ORMBPasswordName string `json:"ormbPassword,omitempty"`
}

func BuildSecretEnvs(secret *v1.Secret, ormbConfig *ORMBConfig) []v1.EnvVar {
	ormbUsernameName := ORMBUsernameName
	ormbPasswordName := ORMBPasswordName
	if ormbConfig.ORMBUsernameName != "" {
		ormbUsernameName = ormbConfig.ORMBUsernameName
	}

	if ormbConfig.ORMBPasswordName != "" {
		ormbPasswordName = ormbConfig.ORMBPasswordName
	}
	envs := []v1.EnvVar{
		{
			Name: ORMBUsername,
			ValueFrom: &v1.EnvVarSource{
				SecretKeyRef: &v1.SecretKeySelector{
					LocalObjectReference: v1.LocalObjectReference{
						Name: secret.Name,
					},
					Key: ormbUsernameName,
				},
			},
		},
		{
			Name: ORMBPassword,
			ValueFrom: &v1.EnvVarSource{
				SecretKeyRef: &v1.SecretKeySelector{
					LocalObjectReference: v1.LocalObjectReference{
						Name: secret.Name,
					},
					Key: ormbPasswordName,
				},
			},
		},
	}
	return envs
}
