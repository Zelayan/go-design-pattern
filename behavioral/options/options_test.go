package options

import "testing"

func Test(t *testing.T) {
	

	u := NewUser("test", WithAddress("hangzhou"), WithAge(13), WithEmail("zeze@qq.com"), WithQq("111111"))
	t.Log(u)
}