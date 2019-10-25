package acceptance

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("version command", func() {
	var version string
	var pathToVersionedMain string

	BeforeEach(func() {
		version = fmt.Sprintf("v0.0.0-dev.%d", time.Now().Unix())

		var err error
		pathToVersionedMain, err = gexec.Build("github.com/pivotal-cf/om",
			"--ldflags", fmt.Sprintf("-X main.version=%s -X main.applySleepDurationString=1ms", version))
		Expect(err).ToNot(HaveOccurred())
	})

	When("given the -v short flag", func() {
		It("returns the binary version", func() {
			command := exec.Command(pathToVersionedMain, "-v")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))
			Expect(session.Out.Contents()).To(ContainSubstring(fmt.Sprintf("%s\n", version)))
		})
	})

	When("given the --version long flag", func() {
		It("returns the binary version", func() {
			command := exec.Command(pathToVersionedMain, "--version")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))
			Expect(session.Out.Contents()).To(ContainSubstring(fmt.Sprintf("%s\n", version)))
		})
	})

	When("given the version command", func() {
		It("returns the binary version", func() {
			command := exec.Command(pathToVersionedMain, "version")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))
			Expect(session.Out.Contents()).To(ContainSubstring(fmt.Sprintf("%s\n", version)))
		})
	})
})
