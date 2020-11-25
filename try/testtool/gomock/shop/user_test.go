package shop

import (
	"errors"
	"github.com/golang/mock/gomock"
	mock_user "github.com/tianxinbaiyun/practice/try/gomock/shop/mock"
	"log"
	"testing"
)

func TestShopMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_user.NewMockGuest(ctrl)
	// 默认期望调用一次
	repo.EXPECT().Shopping("123").Return(&User{Name: "123"}, nil)
	// 期望调用2次
	repo.EXPECT().Shopping("1244").Return(&User{Name: "1244"}, nil).Times(2)
	// 调用多少次可以,包括0次
	repo.EXPECT().Shopping("333").Return(nil, errors.New("user not found")).AnyTimes()

	// 验证一下结果
	log.Println(repo.Shopping("1")) // 这是张三
	log.Println(repo.Shopping("2")) // 这是李四
	log.Println(repo.Shopping("3")) // FindOne(2) 需调用两次,注释本行代码将导致测试不通过
	log.Println(repo.Shopping("4")) // user not found, 不限调用次数，注释掉本行也能通过测试
}
