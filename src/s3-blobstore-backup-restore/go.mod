module s3-blobstore-backup-restore

go 1.14

require (
	github.com/aws/aws-sdk-go v1.43.20
	github.com/cloudfoundry-incubator/bosh-backup-and-restore v1.9.27
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.19.0
	system-tests v0.0.0
)

replace system-tests => ../system-tests

