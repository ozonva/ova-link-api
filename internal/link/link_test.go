package link

import (
	"regexp"
	"testing"
	"time"
)

func TestCreateLink(t *testing.T) {
	userId := uint64(1)
	id := uint64(1)
	url := "https://test.com"
	expected := &Link{
		id,
		userId,
		url,
		"",
		make(map[Tag]struct{}),
		time.Now(),
	}

	actual := New(id, userId, url)
	if expected.GetID() != actual.GetID() {
		t.Fatalf(
			`%v. Expected: %v. Actual: %v`, "create link simple link",
			expected.GetID(),
			actual.GetID(),
		)
	}

	if expected.GetUserID() != actual.GetUserID() {
		t.Fatalf(
			`%v. Expected: %v. Actual: %v`, "create link simple link",
			expected.GetUserID(),
			actual.GetUserID(),
		)
	}
	if expected.GetURL() != actual.GetURL() {
		t.Fatalf(
			`%v. Expected: %v. Actual: %v`, "create link simple link",
			expected.GetURL(),
			actual.GetURL(),
		)
	}
	if len(expected.GetTags()) != len(actual.GetTags()) {
		t.Fatalf(
			`%v. Expected len: %v. Actual len: %v`, "create link simple link",
			len(expected.GetTags()),
			len(actual.GetTags()),
		)
	}

	if !expected.GetDateCreated().Before(actual.GetDateCreated()) {
		t.Fatalf(
			`%v. Expected len: %v. Actual len: %v`, "create link simple link",
			expected.GetDateCreated(),
			actual.GetDateCreated(),
		)
	}
}

func TestLinkSetGet(t *testing.T) {
	userId := uint64(1)
	id := uint64(1)
	url := "https://test.com"

	link := New(id, userId, url)

	description := "some description"
	newUrl := "https://new_test.com"
	link.SetDescription(description)
	link.SetURL(newUrl)

	tags := make(map[Tag]struct{})
	tags["tag1"] = struct{}{}
	tags["tag2"] = struct{}{}
	link.SetTags(tags)

	if link.GetUserID() != userId {
		t.Fatalf(
			`%v. Expected: %v. Actual: %v`, "link get/set",
			userId,
			link.GetUserID(),
		)
	}
	if link.GetURL() != newUrl {
		t.Fatalf(
			`%v. Expected: %v. Actual: %v`, "link get/set",
			newUrl,
			link.GetURL(),
		)
	}
	if link.GetDescription() != description {
		t.Fatalf(
			`%v. Expected: %v. Actual: %v`, "link get/set",
			description,
			link.GetURL(),
		)
	}
	if len(link.GetTags()) != len(tags) {
		t.Fatalf(
			`%v. Expected len: %v. Actual len: %v`, "link get/set",
			len(tags),
			len(link.GetTags()),
		)
	}

	for tag := range tags {
		if _, ok := link.GetTags()[tag]; !ok {
			t.Fatalf(
				`%v. Expected key: %v doesn't exist`, "link get/set",
				tag,
			)
		}
	}
}

func TestLinkAddRemoveTags(t *testing.T) {
	userId := uint64(1)
	id := uint64(1)
	url := "https://test.com"

	link := New(id, userId, url)

	tags := make(map[Tag]struct{})
	tags["tag1"] = struct{}{}
	tags["tag2"] = struct{}{}
	link.SetTags(tags)

	link.AddTag("tag3")
	link.AddTag("tag4")
	link.AddTag("tag1")
	link.AddTag("tag2")

	if len(link.GetTags()) != 4 {
		t.Fatalf(
			`%v. Expected len: %v. Actual len: %v`, "link add/remove tag",
			4,
			len(link.GetTags()),
		)
	}

	for _, tag := range []Tag{"tag1", "tag2", "tag3", "tag4"} {
		if _, ok := link.GetTags()[tag]; !ok {
			t.Fatalf(
				`%v. Expected key: %v doesn't exist`, "link add/remove tag",
				tag,
			)
		}
	}

	link.RemoveTag("tag3")
	link.RemoveTag("tag1")
	link.RemoveTag("tag3")
	link.RemoveTag("tag1")

	if len(link.GetTags()) != 2 {
		t.Fatalf(
			`%v. Expected len: %v. Actual len: %v`, "link add/remove tag",
			2,
			len(link.GetTags()),
		)
	}

	for _, tag := range []Tag{"tag2", "tag4"} {
		if _, ok := link.GetTags()[tag]; !ok {
			t.Fatalf(
				`%v. Expected key: %v doesn't exist`, "link add/remove tag",
				tag,
			)
		}
	}
}

func TestLinkEquals(t *testing.T) {
	link1 := New(1, 1, "https://test1.com")
	linkCopy := *link1
	link2 := &linkCopy
	link2.SetTags(make(map[Tag]struct{}, 0))

	if !link1.Equals(link2) {
		t.Fatalf("%v. Structures should be equal:\n%v\n%v", "link equality", link1, link2)
	}

	link2.AddTag("tag1")
	link2.AddTag("tag2")
	if link1.Equals(link2) {
		t.Fatalf("%v. Structures should not be equal:\n%v\n%v", "link equality", link1, link2)
	}

	link1.AddTag("tag2")
	if link1.Equals(link2) {
		t.Fatalf("%v. Structures should not be equal:\n%v\n%v", "link equality", link1, link2)
	}

	link1.AddTag("tag3")
	if link1.Equals(link2) {
		t.Fatalf("%v. Structures should not be equal:\n%v\n%v", "link equality", link1, link2)
	}

	link1.RemoveTag("tag3")
	link1.AddTag("tag1")
	if !link1.Equals(link2) {
		t.Fatalf("%v. Structures should be equal:\n%v\n%v", "link equality", link1, link2)
	}

	newDescription := "New Description"
	link1.SetDescription(newDescription)
	if link1.Equals(link2) {
		t.Fatalf("%v. Structures should not be equal:\n%v\n%v", "link equality", link1, link2)
	}

	link2.SetDescription(newDescription)
	if !link1.Equals(link2) {
		t.Fatalf("%v. Structures should be equal:\n%v\n%v", "link equality", link1, link2)
	}
}

func TestLinkString(t *testing.T) {
	link1 := New(1, 2, "https://test1.com")
	link1.SetDescription("Ozon Go School. Project.")
	tags := make(map[Tag]struct{})
	link1.SetTags(tags)
	link1.AddTag("tag1")
	link1.AddTag("tag2")

	re := regexp.MustCompile(`ID: 1,
UserID: 2,
URL: "https://test1.com",
Description: "Ozon Go School. Project.",
Tags: map\[tag1:{} tag2:{}],
DateCreated: [0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\+[0-9]{2}:[0-9]{2}`)

	if !re.MatchString(link1.String()) {
		t.Fatalf("%v. Expected: %v, Actual: %v", "link string", re, link1.String())
	}
}
