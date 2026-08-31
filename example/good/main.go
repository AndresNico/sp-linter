package main

import apiconst "github.com/openmcp-project/openmcp-operator/api/constants"

// Both cases handled — linter should pass with no diagnostics.
func handle(annotations map[string]string) {
	op := annotations[apiconst.OperationAnnotation]
	switch op {
	case apiconst.OperationAnnotationValueIgnore:
		_ = "ignored"
	case apiconst.OperationAnnotationValueReconcile:
		_ = "reconciled"
	}
}
