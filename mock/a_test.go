package mock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserService_GetUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mockRepo 是自动生成的 Mock 对象
	mockRepo := NewMockUserRepository(ctrl)

	// 设置期望：预期调用 GetByID("123")，返回指定值
	mockRepo.EXPECT().
		GetByID(int64(123)).
		Return(&User{Name: "Alice"}, nil).
		Times(1) // 必须调用 exactly 1 次

	svc := &UserService{repo: mockRepo}
	name, err := svc.GetUserName(int64(123))

	if err != nil || name != "Alice" {
		t.Fatalf("unexpected result: name=%s, err=%v", name, err)
	}
}
