package link

import (
	"fmt"
	"time"
)

type Tag string

type Link struct {
	id          uint64
	userID      uint64
	url         string
	description string
	tags        map[Tag]struct{}
	dateCreated time.Time
}

func New(id uint64, userID uint64, url string) *Link {
	return &Link{id, userID, url, "", make(map[Tag]struct{}, 0), time.Now()}
}

func (l *Link) String() string {
	return fmt.Sprintf(
		"ID: %v,\nUserID: %v,\nURL: %q,\nDescription: %q,\nTags: %v,\nDateCreated: %v",
		l.id,
		l.userID,
		l.url,
		l.description,
		l.tags,
		l.dateCreated.Format(time.RFC3339),
	)
}

func (l *Link) Equals(inputLink *Link) bool {
	if l.GetID() != inputLink.GetID() {
		return false
	}
	if l.GetUserID() != inputLink.GetUserID() {
		return false
	}
	if l.GetURL() != inputLink.GetURL() {
		return false
	}
	if !l.GetDateCreated().Equal(inputLink.GetDateCreated()) {
		return false
	}
	if l.GetDescription() != inputLink.GetDescription() {
		return false
	}
	if len(l.GetTags()) != len(inputLink.GetTags()) {
		return false
	}
	for tag := range l.GetTags() {
		if _, ok := inputLink.GetTags()[tag]; !ok {
			return false
		}
	}

	return true
}

func (l *Link) GetID() uint64 {
	return l.id
}

func (l *Link) GetUserID() uint64 {
	return l.userID
}

func (l *Link) SetURL(url string) {
	l.url = url
}

func (l *Link) GetURL() string {
	return l.url
}

func (l *Link) SetDescription(description string) {
	l.description = description
}

func (l *Link) GetDescription() string {
	return l.description
}

func (l *Link) GetTags() map[Tag]struct{} {
	return l.tags
}

func (l *Link) SetTags(tags map[Tag]struct{}) {
	l.tags = tags
}

func (l *Link) GetDateCreated() time.Time {
	return l.dateCreated
}

func (l *Link) AddTag(tag Tag) {
	if _, ok := l.tags[tag]; !ok {
		l.tags[tag] = struct{}{}
	}
}

func (l *Link) RemoveTag(tag Tag) {
	if _, ok := l.tags[tag]; ok {
		delete(l.tags, tag)
	}
}
