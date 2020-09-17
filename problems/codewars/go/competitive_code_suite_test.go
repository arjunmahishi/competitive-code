package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCompetitiveCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CompetitiveCode Suite")
}
