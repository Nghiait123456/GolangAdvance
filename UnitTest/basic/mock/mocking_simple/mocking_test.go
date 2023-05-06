package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Matematik interface {
	MockTopla([]int) (int, error)
}

type islem struct{}

func (isl *islem) MockTopla(sayilar []int) (int, error) {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}
	return toplam, nil
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) MockTopla(sayilar []int) (int, error) {
	args := mock.Called(sayilar)
	result := args.Get(0)

	return result.(int), args.Error(1)
}

func TestMockTopla(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("MockTopla", []int{2, 3}).Return(5, nil)

	testMatematik := Matematik(mockRepo)
	sonuc, err := testMatematik.MockTopla([]int{2, 3})
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 5, sonuc)
	assert.Nil(t, err)
}
