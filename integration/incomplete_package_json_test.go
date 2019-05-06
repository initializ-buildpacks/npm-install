package integration_test

import (
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/dagger"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func init() {
	suite("IncompletePackageJSON", testIncompletePackageJSON)
}

func testIncompletePackageJSON(t *testing.T, when spec.G, it spec.S) {
	var (
		Expect func(interface{}, ...interface{}) Assertion
	)

	it.Before(func() {
		Expect = NewWithT(t).Expect
	})

	it("should build a working OCI image for a simple app when there is an incomplete package json", func() {
		app, err := dagger.PackBuild(filepath.Join("testdata", "incomplete_package_json"), nodeURI, npmURI)
		Expect(err).ToNot(HaveOccurred())
		defer app.Destroy()

		Expect(app.Start()).To(Succeed())
		response, _, err := app.HTTPGet("/")
		Expect(err).NotTo(HaveOccurred())
		Expect(response).To(ContainSubstring("Hello, World!"))
	})
}