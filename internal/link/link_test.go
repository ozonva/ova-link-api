package link

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Link.", func() {
	id := uint64(1)
	userId := uint64(2)
	url := "https://test.com"
	expected := &Link{id, userId, url, "", make(map[Tag]struct{}), time.Now()}

	Context("Creation.", func() {
		linkEntity := New(id, userId, url)

		It("ID should be the same", func() {
			Expect(linkEntity.GetID()).Should(BeIdenticalTo(expected.GetID()))
		})
		It("User should ID be the same", func() {
			Expect(linkEntity.GetUserID()).Should(BeIdenticalTo(expected.GetUserID()))
		})
		It("URL should be the same", func() {
			Expect(linkEntity.GetURL()).Should(BeIdenticalTo(expected.GetURL()))
		})
		It("Description should be the same", func() {
			Expect(linkEntity.GetDescription()).Should(BeIdenticalTo(expected.GetDescription()))
		})
		It("Tags should be the same", func() {
			Expect(linkEntity.GetTags()).Should(BeEquivalentTo(expected.GetTags()))
		})
		It("Time should be different", func() {
			Expect(linkEntity.GetDateCreated().After(expected.GetDateCreated())).Should(BeTrue())
		})
	})
	Context("Set values.", func() {
		linkEntity := New(id, userId, url)
		description := "some description"
		newUrl := "https://new_test.com"
		linkEntity.SetDescription(description)
		linkEntity.SetURL(newUrl)

		tags := make(map[Tag]struct{})
		tags["tag1"] = struct{}{}
		tags["tag2"] = struct{}{}
		linkEntity.SetTags(tags)

		It("Description should be updated", func() {
			Expect(linkEntity.GetDescription()).Should(BeIdenticalTo(description))
			Expect(linkEntity.GetDescription()).ShouldNot(BeIdenticalTo(expected.GetDescription()))
		})
		It("Tags should be updated", func() {
			Expect(linkEntity.GetTags()).Should(BeEquivalentTo(tags))
			Expect(linkEntity.GetTags()).ShouldNot(BeEquivalentTo(expected.GetTags()))
		})
	})
	Context("Add/remove tags.", func() {
		linkEntity := New(id, userId, url)
		tags := make(map[Tag]struct{})
		tags["tag1"] = struct{}{}
		tags["tag2"] = struct{}{}
		linkEntity.SetTags(tags)

		It("Tags should be added.", func() {
			linkEntity.AddTag("tag3")
			linkEntity.AddTag("tag4")
			linkEntity.AddTag("tag1")
			linkEntity.AddTag("tag2")

			Expect(linkEntity.GetTags()).Should(BeEquivalentTo(map[Tag]struct{}{"tag1": {}, "tag2": {}, "tag3": {}, "tag4": {}}))
		})
		It("Tags should be removed.", func() {
			linkEntity.RemoveTag("tag3")
			linkEntity.RemoveTag("tag1")
			linkEntity.RemoveTag("tag3")
			linkEntity.RemoveTag("tag1")

			Expect(linkEntity.GetTags()).Should(BeEquivalentTo(map[Tag]struct{}{"tag2": {}, "tag4": {}}))
		})
	})
	Context("Compare with another link entity.", func() {
		linkEntity1 := New(id, userId, url)
		linkCopy := *linkEntity1
		linkEntity2 := &linkCopy
		linkEntity2.SetTags(make(map[Tag]struct{}, 0))
		description := "some description"

		It("Two empty link should be equal", func() {
			Expect(linkEntity1.Equals(linkEntity2)).Should(BeTrue())
		})

		It("Add tags. Links should not be equal.", func() {
			linkEntity2.AddTag("tag1")
			linkEntity2.AddTag("tag2")
			Expect(linkEntity1.Equals(linkEntity2)).Should(BeFalse())

			linkEntity1.AddTag("tag2")
			Expect(linkEntity1.Equals(linkEntity2)).Should(BeFalse())

			linkEntity1.AddTag("tag3")
			Expect(linkEntity1.Equals(linkEntity2)).Should(BeFalse())
		})

		It("Make tags are the same. Links should be equal.", func() {
			linkEntity1.RemoveTag("tag3")
			linkEntity1.AddTag("tag1")
			Expect(linkEntity1.Equals(linkEntity2)).Should(BeTrue())
		})

		It("Change description. Links should not be equal.", func() {
			linkEntity1.SetDescription(description)
			Expect(linkEntity1.Equals(linkEntity2)).Should(BeFalse())
		})

		It("Make description the same. Links should be equal.", func() {
			linkEntity2.SetDescription(description)
			Expect(linkEntity1.Equals(linkEntity2)).Should(BeTrue())
		})
	})
	Context("Cast to string.", func() {
		defer GinkgoRecover()
		linkEntity := New(id, userId, url)
		linkEntity.SetDescription("Ozon Go School. Project.")
		tags := make(map[Tag]struct{})
		linkEntity.SetTags(tags)
		linkEntity.AddTag("tag1")
		linkEntity.AddTag("tag2")

		regexpString := `ID: 1,
UserID: 2,
URL: "https://test.com",
Description: "Ozon Go School. Project.",
Tags: map\[tag1:{} tag2:{}],
DateCreated: [0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\+[0-9]{2}:[0-9]{2}`
		Expect(linkEntity.String()).Should(MatchRegexp(regexpString))
	})
})
