// +build tools

package tools

import (
	_ "knative.dev/test-infra/scripts"

	_ "knative.dev/serving/test/conformance/ingress"
	_ "knative.dev/serving/test/test_images/flaky"
	_ "knative.dev/serving/test/test_images/grpc-ping"
	_ "knative.dev/serving/test/test_images/httpproxy"
	_ "knative.dev/serving/test/test_images/runtime"
	_ "knative.dev/serving/test/test_images/timeout"
	_ "knative.dev/serving/test/test_images/wsserver"
)
