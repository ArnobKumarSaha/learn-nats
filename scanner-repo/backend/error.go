package backend

import (
	"gocloud.dev/gcerrors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ErrorToAPIStatus(err error) *metav1.Status {
	// If error is nil, set the status to Success
	// Convert the error to metav1.Status by specifying the statusCode properly (utilize below to functions)
	return nil
}

func HTTPStatusFromCode(code gcerrors.ErrorCode) int32 {
	return 0
}

func GRPCCode(err error) gcerrors.ErrorCode {
	return 0
}
