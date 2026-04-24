/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package v3

import (
	"path/filepath"
)

const (
	CommitterVersion        = "v3"
	ScalableCommitterImage  = "hyperledger/fabric-x-committer-test-node:0.1.9"
	SidecarDefaultPort      = "4001/tcp"
	QueryServiceDefaultPort = "7001/tcp"
)

var ContainerCmd = []string{"run", "db", "orderer", "committer"}

func ContainerEnvVars(scMSPDir, scTLSDir, scMSPID, channelName, ordererEndpoint string, tlsEnabled bool, ordererTLSCACert string) []string {
	env := []string{
		"SC_SIDECAR_LOGGING_LOGSPEC=debug",
		"SC_SIDECAR_ORDERER_CHANNEL_ID=" + channelName,
		"SC_SIDECAR_ORDERER_SIGNED_ENVELOPES=true",
		"SC_SIDECAR_ORDERER_IDENTITY_MSP_ID=" + scMSPID,
		"SC_SIDECAR_ORDERER_IDENTITY_MSP_DIR=" + scMSPDir,
		"SC_QUERY_SERVICE_SERVER_ENDPOINT=:7001",
		"SC_QUERY_SERVICE_LOGGING_LOGSPEC=DEBUG",
		"SC_COORDINATOR_LOGGING_LOGSPEC=DEBUG",
		"SC_ORDERER_LOGGING_LOGSPEC=debug",
		"SC_ORDERER_BLOCK_SIZE=1",
		"SC_VC_LOGGING_LOGSPEC=DEBUG",
		"SC_VERIFIER_LOGGING_LOGSPEC=INFO",
		"SC_SIDECAR_SERVER_MAX_CONCURRENT_STREAMS=0",
	}
	if tlsEnabled {
		env = append(env,
			"SC_SIDECAR_ORDERER_TLS_MODE=mtls",
			"SC_SIDECAR_ORDERER_TLS_CERT_FILE="+filepath.Join(scTLSDir, "server.crt"),
			"SC_SIDECAR_ORDERER_TLS_KEY_FILE="+filepath.Join(scTLSDir, "server.key"),
			"SC_SIDECAR_ORDERER_TLS_ROOT_CERT_FILE="+ordererTLSCACert,

			"SC_ORDERER_GENERAL_TLS_ENABLED=true",
			"SC_ORDERER_GENERAL_TLS_CERTIFICATE="+filepath.Join(scTLSDir, "server.crt"),
			"SC_ORDERER_GENERAL_TLS_PRIVATE_KEY="+filepath.Join(scTLSDir, "server.key"),
			"SC_ORDERER_GENERAL_TLS_ROOTCAS="+filepath.Join(scTLSDir, "ca.crt"),

			"SC_QUERY_SERVICE_SERVER_TLS_MODE=tls",
			"SC_QUERY_SERVICE_SERVER_TLS_CERT_FILE="+filepath.Join(scTLSDir, "server.crt"),
			"SC_QUERY_SERVICE_SERVER_TLS_KEY_FILE="+filepath.Join(scTLSDir, "server.key"),
			"SC_QUERY_SERVICE_SERVER_TLS_CLIENT_CA_FILES="+filepath.Join(scTLSDir, "ca.crt"),

			"SC_SIDECAR_SERVER_TLS_MODE=tls",
			"SC_SIDECAR_SERVER_TLS_CERT_FILE="+filepath.Join(scTLSDir, "server.crt"),
			"SC_SIDECAR_SERVER_TLS_KEY_FILE="+filepath.Join(scTLSDir, "server.key"),
		)
	} else {
		env = append(env, "SC_SIDECAR_ORDERER_TLS_MODE=none")
	}
	return env
}
