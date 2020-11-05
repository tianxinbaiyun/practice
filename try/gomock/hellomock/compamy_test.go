package hellomock

import (
	"github.com/golang/mock/gomock"
	mock_hellomock "github.com/tianxinbaiyun/practice/try/gomock/hellomock/mock"
	"testing"
)

func TestCompanyMeeting(t *testing.T) {
	person := NewPerson("王尼美")
	company := NewCompany(person)
	t.Log(company.Meeting("王尼玛"))
}

func TestCompany_Meeting2(t *testing.T) {
	ctl := gomock.NewController(t)
	mock_talker := mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("这是自定义的返回值，可以是任意类型。")

	company := NewCompany(mock_talker)
	t.Log(company.Meeting("王尼玛"))
	//t.Log(company.Meeting("张全蛋"))
}
