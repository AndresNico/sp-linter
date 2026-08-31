package main

import apiconst "github.com/openmcp-project/openmcp-operator/api/constants"

// Missing the Reconcile case — linter should report a diagnostic here.
func handle(annotations map[string]string) {
	op := annotations[apiconst.OperationAnnotation]
	switch op {
	case apiconst.OperationAnnotationValueIgnore:
		_ = "ignored"
	}
}
