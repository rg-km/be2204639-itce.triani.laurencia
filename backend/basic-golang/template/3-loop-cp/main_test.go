package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	var leaderboardObject Leaderboard
	var expectedOutput string
	BeforeEach(func() {
		users := []*UserRank{
			{
				Name: "Roger",
				Rank: 1,
			},
			{
				Name: "Tony",
				Rank: 2,
			},
			{
				Name: "Bruce",
				Rank: 3,
			},
			{
				Name: "Natasha",
				Rank: 4,
			},
			{
				Name: "Clint",
				Rank: 5,
			},
		}

		leaderboardObject = Leaderboard{
			Users: users,
		}
		expectedOutput = `Peringkat ke-1: Roger Peringkat ke-2: Tony Peringkat ke-3: Bruce Peringkat ke-4: Natasha Peringkat ke-5: Clint `
	})
	Describe("ExecuteToByteBuffer", func() {
		It("returns slice of bytes with correct wording", func() {
			b, err := ExecuteToByteBuffer(leaderboardObject)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(b)).To(Equal(expectedOutput))
		})
	})
})
