package annotations

import apiconst "github.com/openmcp-project/openmcp-operator/api/constants"

func goodHandler(annotations map[string]string) {
	op := annotations[apiconst.OperationAnnotation]
	switch op {
	case apiconst.OperationAnnotationValueIgnore:
		_ = "ignored"
	case apiconst.OperationAnnotationValueReconcile:
		_ = "reconciled"
	}
}

func missingReconcile(annotations map[string]string) {
	op := annotations[apiconst.OperationAnnotation]
	switch op { // want `switch on OperationAnnotation missing case for OperationAnnotationValueReconcile`
	case apiconst.OperationAnnotationValueIgnore:
		_ = "ignored"
	}
}

func missingIgnore(annotations map[string]string) {
	op := annotations[apiconst.OperationAnnotation]
	switch op { // want `switch on OperationAnnotation missing case for OperationAnnotationValueIgnore`
	case apiconst.OperationAnnotationValueReconcile:
		_ = "reconciled"
	}
}
