/*
Copyright 2020 VMware, Inc.
All Rights Reserved.

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
	"encoding/json"
	"testing"

	"github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	akov1alpha1 "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/ako-operator/api/v1alpha1"
)

var sfJson = `
{
    "apiVersion": "apps/v1",
    "kind": "StatefulSet",
    "metadata": {
        "name": "ako",
        "namespace": "avi-system"
    },
    "spec": {
        "podManagementPolicy": "OrderedReady",
        "replicas": 1,
        "revisionHistoryLimit": 10,
        "selector": {
            "matchLabels": {
                "app": "ako"
            }
        },
        "serviceName": "ako",
        "template": {
            "metadata": {
                "creationTimestamp": null,
                "labels": {
                    "app": "ako"
                }
            },
            "spec": {
                "containers": [
                    {
                        "env": [
                            {
                                "name": "SHARD_VS_SIZE",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "shardVSSize",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "NETWORK_NAME",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "networkName",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "SEG_NAME",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "serviceEngineGroupName",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "NODE_NETWORK_LIST",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "nodeNetworkList",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "CTRL_VERSION",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "controllerVersion",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "CLUSTER_NAME",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "clusterName",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "DEFAULT_DOMAIN",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "defaultDomain",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "FULL_SYNC_INTERVAL",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "fullSyncFrequency",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "AUTO_L4_FQDN",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "autoFQDN",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "CNI_PLUGIN",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "cniPlugin",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "PASSTHROUGH_SHARD_SIZE",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "passhtroughShardSize",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "CLOUD_NAME",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "cloudName",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "ENABLE_RHI",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "enableRHI",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "DISABLE_STATIC_ROUTE_SYNC",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "disableStaticRouteSync",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "DEFAULT_ING_CONTROLLER",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "defaultIngController",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "CTRL_IPADDRESS",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "controllerIP",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "AKO_API_PORT",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "apiServerPort",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "SERVICE_TYPE",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "serviceType",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "CTRL_USERNAME",
                                "valueFrom": {
                                    "secretKeyRef": {
                                        "key": "username",
                                        "name": "avi-secret"
                                    }
                                }
                            },
                            {
                                "name": "CTRL_PASSWORD",
                                "valueFrom": {
                                    "secretKeyRef": {
                                        "key": "password",
                                        "name": "avi-secret"
                                    }
                                }
                            },
                            {
                                "name": "CTRL_CA_DATA",
                                "valueFrom": {
                                    "secretKeyRef": {
                                        "key": "certificateAuthorityData",
                                        "name": "avi-secret"
                                    }
                                }
                            },
                            {
                                "name": "POD_NAME",
                                "valueFrom": {
                                    "fieldRef": {
                                        "apiVersion": "v1",
                                        "fieldPath": "metadata.name"
                                    }
                                }
                            },
                            {
                                "name": "NSXT_T1_LR",
                                "valueFrom": {
                                    "configMapKeyRef": {
                                        "key": "nsxtT1LR",
                                        "name": "avi-k8s-config"
                                    }
                                }
                            },
                            {
                                "name": "LOG_FILE_PATH",
                                "value": "/log"
                            },
                            {
                                "name": "LOG_FILE_NAME",
                                "value": "ako.log"
                            }
                        ],
                        "image": "test-repo",
                        "imagePullPolicy": "Always",
                        "lifecycle": {
                            "preStop": {
                                "exec": {
                                    "command": [
                                        "/bin/sh",
                                        "/var/pre_stop_hook.sh"
                                    ]
                                }
                            }
                        },
                        "livenessProbe": {
                            "failureThreshold": 3,
                            "httpGet": {
                                "path": "/api/status",
                                "port": 8080,
                                "scheme": "HTTP"
                            },
                            "initialDelaySeconds": 5,
                            "periodSeconds": 10,
                            "successThreshold": 1,
                            "timeoutSeconds": 1
                        },
                        "name": "ako",
                        "ports": [
                            {
                                "containerPort": 80,
                                "name": "http",
                                "protocol": "TCP"
                            }
                        ],
                        "resources": {
                            "limits": {
                                "cpu": "250m",
                                "memory": "300Mi"
                            },
                            "requests": {
                                "cpu": "100m",
                                "memory": "200Mi"
                            }
                        },
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File"
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "ako-sa",
                "serviceAccountName": "ako-sa",
                "terminationGracePeriodSeconds": 30
            }
        },
        "updateStrategy": {
            "rollingUpdate": {
                "partition": 0
            },
            "type": "RollingUpdate"
        }
    }
}
`

func getTestDefaultStatefulSet() (appsv1.StatefulSet, error) {
	defSf := appsv1.StatefulSet{}
	err := json.Unmarshal([]byte(sfJson), &defSf)
	return defSf, err
}

func buildStatefulSetAndVerify(existingSf appsv1.StatefulSet, akoConfig akov1alpha1.AKOConfig,
	update bool, errExpected bool, t *testing.T) appsv1.StatefulSet {

	g := gomega.NewGomegaWithT(t)

	secretData := map[string][]byte{
		"username": []byte("abc"),
		"password": []byte("abc"),
	}
	secretObj := v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "avi-secret",
			Namespace: AviSystemNS,
		},
		Data: secretData,
	}
	newSf, err := BuildStatefulSet(akoConfig, secretObj)
	if errExpected {
		g.Expect(err).To(gomega.Not(gomega.BeNil()))
		return appsv1.StatefulSet{}
	}

	g.Expect(err).To(gomega.BeNil())
	g.Expect(isSfUpdateRequired(existingSf, newSf)).To(gomega.Equal(update))
	return newSf
}
