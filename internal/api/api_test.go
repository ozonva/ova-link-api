package api_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/ozonva/ova-link-api/internal/kafka"

	"github.com/rs/zerolog"

	"github.com/ozonva/ova-link-api/internal/link"

	"github.com/golang/mock/gomock"
	"github.com/ozonva/ova-link-api/internal/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	api "github.com/ozonva/ova-link-api/internal/api"
	grpc "github.com/ozonva/ova-link-api/pkg/ova-link-api"
	ova_link_api "github.com/ozonva/ova-link-api/pkg/ova-link-api"
)

type LinkMatcher struct {
	expected []link.Link
}

type UpdateMatcher struct {
	expected link.Link
}

type KafkaMatcher struct {
	expected kafka.Message
}

type KafkaMultiMatcher struct {
	expected kafka.Message
}

func (lm LinkMatcher) Matches(x interface{}) bool {
	entities := x.([]link.Link)
	if len(entities) != len(lm.expected) {
		return false
	}

	for i := 0; i < len(entities); i++ {
		if entities[i].UserID != lm.expected[i].UserID {
			return false
		}
		if entities[i].Description != lm.expected[i].Description {
			return false
		}
		if entities[i].Tags != lm.expected[i].Tags {
			return false
		}
		if entities[i].Url != lm.expected[i].Url {
			return false
		}
	}

	return true
}

func (lm LinkMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", lm.expected, lm.expected)
}

func (um UpdateMatcher) Matches(x interface{}) bool {
	entity := x.(link.Link)

	if um.expected.UserID != entity.UserID {
		return false
	}
	if um.expected.Description != entity.Description {
		return false
	}
	if um.expected.Tags != entity.Tags {
		return false
	}
	if um.expected.Url != entity.Url {
		return false
	}

	return true
}

func (um UpdateMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", um.expected, um.expected)
}

func (km KafkaMatcher) Matches(x interface{}) bool {
	message := x.(kafka.Message)

	if message.EventType != km.expected.EventType {
		return false
	}

	if val, ok := message.Value.(link.Link); ok {
		expected := km.expected.Value.(link.Link)
		if expected.UserID != val.UserID {
			return false
		}
		if expected.Description != val.Description {
			return false
		}
		if expected.Tags != val.Tags {
			return false
		}
		if expected.Url != val.Url {
			return false
		}

		return true
	}

	if !reflect.DeepEqual(message.Value, km.expected.Value) {
		return false
	}

	return true
}

func (km KafkaMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", km.expected, km.expected)
}

func (kmm KafkaMultiMatcher) Matches(x interface{}) bool {
	message := x.(kafka.Message)

	if message.EventType != kmm.expected.EventType {
		return false
	}

	if val, ok := message.Value.([]link.Link); ok {
		expected := kmm.expected.Value.([]link.Link)
		for i, expectedEntity := range expected {
			if expectedEntity.UserID != val[i].UserID {
				return false
			}
			if expectedEntity.Description != val[i].Description {
				return false
			}
			if expectedEntity.Tags != val[i].Tags {
				return false
			}
			if expectedEntity.Url != val[i].Url {
				return false
			}
		}

		return true
	}

	if !reflect.DeepEqual(message.Value, kmm.expected.Value) {
		return false
	}

	return true
}

func (kmm KafkaMultiMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", kmm.expected, kmm.expected)
}

var _ = Describe("Api", func() {
	Context("Database", func() {
		var API grpc.LinkAPIServer
		var ctrl *gomock.Controller
		var mockRepo *mocks.MockRepo
		var mockProducer *mocks.MockProducer
		var mockMetrics *mocks.MockMetrics

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockRepo = mocks.NewMockRepo(ctrl)
			mockProducer = mocks.NewMockProducer(ctrl)
			mockMetrics = mocks.NewMockMetrics(ctrl)
			API = api.NewLinkAPI(mockRepo, zerolog.Nop(), mockProducer, mockMetrics)
		})

		AfterEach(func() {
			ctrl.Finish()
		})

		It("Describe success", func() {
			selectTime := time.Now()
			expected := &link.Link{
				ID:          1,
				UserID:      1,
				Url:         "https://test.com",
				Description: "test description",
				Tags:        "tag1#tag2",
				CreatedAt:   selectTime,
			}
			mockRepo.EXPECT().DescribeEntity(gomock.Eq(uint64(1))).Times(1).Return(expected, nil)
			mockMetrics.EXPECT().DescribeSuccessResponseCounter().Times(1)

			_, err := API.DescribeLink(
				context.Background(),
				&ova_link_api.DescribeLinkRequest{
					Id: 1,
				},
			)

			Expect(err).Should(Succeed())
		})

		It("Describe error", func() {
			mockRepo.EXPECT().DescribeEntity(gomock.Eq(uint64(1))).Times(1).
				Return(nil, errors.New("something goes wrong"))

			_, err := API.DescribeLink(
				context.Background(),
				&ova_link_api.DescribeLinkRequest{
					Id: 1,
				},
			)

			Expect(err).Should(HaveOccurred())
		})

		It("List success", func() {
			selectTime1 := time.Now()
			selectTime2 := time.Now()
			selectTime3 := time.Now()

			expected := []link.Link{
				{
					ID:          3,
					UserID:      1,
					Url:         "https://test.com3",
					Description: "test description3",
					Tags:        "tag3#tag6",
					CreatedAt:   selectTime1,
				},
				{
					ID:          4,
					UserID:      1,
					Url:         "https://test.com4",
					Description: "test description4",
					Tags:        "tag4#tag7",
					CreatedAt:   selectTime2,
				},
				{
					ID:          5,
					UserID:      3,
					Url:         "https://test.com5",
					Description: "test description5",
					Tags:        "tag5#tag8",
					CreatedAt:   selectTime3,
				},
			}

			mockRepo.EXPECT().ListEntities(gomock.Eq(uint64(2)), gomock.Eq(uint64(2))).
				Times(1).
				Return(expected, nil)
			mockMetrics.EXPECT().ListSuccessResponseCounter().Times(1)

			limit := uint64(2)
			offset := uint64(2)
			_, err := API.ListLink(
				context.Background(),
				&ova_link_api.ListLinkRequest{
					Limit:  &limit,
					Offset: &offset,
				},
			)

			Expect(err).Should(Succeed())
		})

		It("List error", func() {
			mockRepo.EXPECT().ListEntities(gomock.Eq(uint64(2)), gomock.Eq(uint64(2))).
				Times(1).
				Return(nil, errors.New("something goes wrong"))

			limit := uint64(2)
			offset := uint64(2)
			_, err := API.ListLink(
				context.Background(),
				&ova_link_api.ListLinkRequest{
					Limit:  &limit,
					Offset: &offset,
				},
			)

			Expect(err).Should(HaveOccurred())
		})

		It("Delete success", func() {
			mockRepo.EXPECT().DeleteEntity(gomock.Eq(uint64(1))).Times(1).Return(nil)

			kafkaMatcher := KafkaMatcher{
				expected: kafka.Message{EventType: kafka.Remove, Value: uint64(1)},
			}
			mockProducer.EXPECT().Send(kafkaMatcher).Times(1).Return(nil)
			mockMetrics.EXPECT().RemoveSuccessResponseCounter().Times(1)

			_, err := API.DeleteLink(
				context.Background(),
				&ova_link_api.DeleteLinkRequest{
					Id: 1,
				},
			)

			Expect(err).Should(Succeed())
		})

		It("Delete error", func() {
			mockRepo.EXPECT().DeleteEntity(gomock.Eq(uint64(1))).
				Times(1).Return(errors.New("something goes wrong"))

			_, err := API.DeleteLink(
				context.Background(),
				&ova_link_api.DeleteLinkRequest{
					Id: 1,
				},
			)

			Expect(err).Should(HaveOccurred())
		})

		It("Create success", func() {
			selectTime1 := time.Now()
			insert := []link.Link{
				{
					ID:          0,
					UserID:      1,
					Url:         "https://test.com3",
					Description: "test description3",
					Tags:        "tag3#tag6",
					CreatedAt:   selectTime1,
				},
			}

			linkMatcher := LinkMatcher{
				expected: insert,
			}

			kafkaMatcher := KafkaMatcher{
				expected: kafka.Message{EventType: kafka.Create, Value: insert[0]},
			}

			mockRepo.EXPECT().AddEntities(linkMatcher).Times(1).Return(nil)
			mockProducer.EXPECT().Send(kafkaMatcher).Times(1).Return(nil)
			mockMetrics.EXPECT().CreateSuccessResponseCounter().Times(1)

			_, err := API.CreateLink(
				context.Background(),
				&ova_link_api.CreateLinkRequest{
					UserId:      1,
					Url:         "https://test.com3",
					Description: "test description3",
					Tags:        []string{"tag3", "tag6"},
				},
			)
			time.Sleep(2 * time.Second)
			Expect(err).Should(Succeed())
		})

		It("Multi create success", func() {
			selectTime1 := time.Now()
			entity := link.Link{
				ID:          0,
				UserID:      1,
				Url:         "https://test.com3",
				Description: "test description3",
				Tags:        "tag3#tag6",
				CreatedAt:   selectTime1,
			}

			insert := []link.Link{
				entity,
				entity,
				entity,
				entity,
				entity,
				entity,
			}

			linkMatcher1 := LinkMatcher{
				expected: insert[0:3],
			}
			linkMatcher2 := LinkMatcher{
				expected: insert[3:],
			}

			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(linkMatcher1).Return(nil),
				mockRepo.EXPECT().AddEntities(linkMatcher2).Return(nil),
			)

			kafkaMatcher1 := KafkaMultiMatcher{
				expected: kafka.Message{EventType: kafka.MultiCreate, Value: insert[0:3]},
			}
			kafkaMatcher2 := KafkaMultiMatcher{
				expected: kafka.Message{EventType: kafka.MultiCreate, Value: insert[3:]},
			}

			gomock.InOrder(
				mockProducer.EXPECT().Send(kafkaMatcher1).Return(nil),
				mockProducer.EXPECT().Send(kafkaMatcher2).Return(nil),
			)

			mockMetrics.EXPECT().MultiCreateSuccessResponseCounter().Times(1)

			createRequest := &ova_link_api.CreateLinkRequest{
				UserId:      1,
				Url:         "https://test.com3",
				Description: "test description3",
				Tags:        []string{"tag3", "tag6"},
			}

			_, err := API.MultiCreateLink(
				context.Background(),
				&ova_link_api.MultiCreateLinkRequest{
					Items: []*ova_link_api.CreateLinkRequest{
						createRequest,
						createRequest,
						createRequest,
						createRequest,
						createRequest,
						createRequest,
					},
				},
			)
			time.Sleep(1 * time.Second)
			Expect(err).Should(Succeed())
		})

		It("Update success", func() {
			selectTime1 := time.Now()
			insert := []link.Link{
				{
					ID:          0,
					UserID:      1,
					Url:         "https://test.com3",
					Description: "test description3",
					Tags:        "tag3#tag6",
					CreatedAt:   selectTime1,
				},
			}

			updateMatcher := UpdateMatcher{
				expected: insert[0],
			}

			kafkaMatcher := KafkaMatcher{
				expected: kafka.Message{EventType: kafka.Update, Value: insert[0]},
			}

			mockRepo.EXPECT().UpdateEntity(updateMatcher, gomock.Eq(uint64(10))).Times(1).Return(nil)
			mockProducer.EXPECT().Send(kafkaMatcher).Times(1).Return(nil)
			mockMetrics.EXPECT().UpdateSuccessResponseCounter().Times(1)

			_, err := API.UpdateLink(
				context.Background(),
				&ova_link_api.UpdateLinkRequest{
					Id:          10,
					UserId:      1,
					Url:         "https://test.com3",
					Description: "test description3",
					Tags:        []string{"tag3", "tag6"},
				},
			)
			Expect(err).Should(Succeed())
		})
	})
})
