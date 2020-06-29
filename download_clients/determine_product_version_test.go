package download_clients_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/om/download_clients"
	"github.com/pivotal-cf/om/download_clients/fakes"
)

var _ = Describe("DetermineProductVersion", func() {
	When("an exact version is provided", func() {
		When("the version exists", func() {
			It("returns the version", func() {
				versioner := &fakes.ProductVersioner{}
				versioner.GetAllProductVersionsReturns([]string{"2.2.1", "2.2.2", "2.2.0"}, nil)

				version, err := download_clients.DetermineProductVersion(
					"product",
					"2.2.2",
					"",
					versioner,
					nil,
				)
				Expect(err).NotTo(HaveOccurred())
				Expect(version).To(Equal("2.2.2"))
			})
		})

		When("the exact version does not exist", func() {
			It("returns an error", func() {
				versioner := &fakes.ProductVersioner{}
				versioner.GetAllProductVersionsReturns([]string{"2.2.1", "2.2.2", "2.2.0"}, nil)

				_, err := download_clients.DetermineProductVersion(
					"product",
					"4.5.6",
					"",
					versioner,
					nil,
				)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring("no valid versions found for product \"product\" and product version \"4.5.6\"\nexisting versions: 2.2.1, 2.2.2, 2.2.0")))
			})
		})
	})
	When("a regex is provided", func() {
		It("returns the latest version for a product", func() {
			versioner := &fakes.ProductVersioner{}
			versioner.GetAllProductVersionsReturns([]string{"2.2.1", "2.2.2", "2.2.0"}, nil)

			version, err := download_clients.DetermineProductVersion(
				"product",
				"",
				`2\.2\..*`,
				versioner,
				nil,
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(version).To(Equal("2.2.2"))
		})

		When("no versions match the regex", func() {
			It("returns the latest version for a product", func() {
				versioner := &fakes.ProductVersioner{}
				versioner.GetAllProductVersionsReturns([]string{"1.2.1"}, nil)

				_, err := download_clients.DetermineProductVersion(
					"product",
					"",
					`2\.2\..*`,
					versioner,
					nil,
				)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring("existing versions: 1.2.1")))
			})
		})

		When("there are no versions", func() {
			It("returns the latest version for a product", func() {
				versioner := &fakes.ProductVersioner{}
				versioner.GetAllProductVersionsReturns([]string{}, nil)

				_, err := download_clients.DetermineProductVersion(
					"product",
					"",
					`2\.2\..*`,
					versioner,
					nil,
				)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring("existing versions: none")))
			})
		})

		When("versions cannot be loaded", func() {
			It("returns an error", func() {
				versioner := &fakes.ProductVersioner{}
				versioner.GetAllProductVersionsReturns(nil, fmt.Errorf("some error"))

				_, err := download_clients.DetermineProductVersion(
					"product",
					"",
					`2\.2\..*`,
					versioner,
					nil,
				)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring("some error")))
			})
		})
	})
})