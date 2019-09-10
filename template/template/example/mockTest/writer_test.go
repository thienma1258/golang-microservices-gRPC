package mockTest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MyMockedWriter struct {
	mock.Mock
}

func (m *MyMockedWriter) ReadUserInfo(i int) int {

	args := m.Called(i)
	return args.Int(0)
}

func TestRealRW_ReadUserInfoWithMockMyMockedWriter(t *testing.T) {
	rw := &realRW{}
	repo := UserRepository{
		userReader: rw,
		userWriter: rw,
	}
	assert.Equal(t,repo.ReadUserInfo(100),100)
	// when you want to write test.
	fmt.Println("Begin test....................")
	testObj := new(MyMockedWriter)
	testObj.On("ReadUserInfo", 123).Return(250)

	testRepo := UserRepository{
		userReader: testObj,
		userWriter: rw,
	}
	assert.Equal(t,testRepo.ReadUserInfo(123),250)
}