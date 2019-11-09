package models

type UserInfo struct {
	Name string
	Age  int
}

//생성자 함수 정의
func NewUserInfo() *UserInfo {
	d := UserInfo{}

	//d.data = map[int]string{}
	d.age = 0
	d.name = ""

	return &d //포인터 전달
}
