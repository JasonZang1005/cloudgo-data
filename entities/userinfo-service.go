package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

var userInfoQueryByID = "SELECT *FROM userinfo where uid=?"
var userInfoInsertStmt = "InSERT userinfo SET username=?,department=?,created=?"
var userInfoQueryAll = "SELECT * FROM userinfo"

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(&u)
	CheckErr(err)
	return err
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	allUserInfo := make([]UserInfo, 0)
	err := engine.Find(&allUserInfo)
	CheckErr(err)
	return allUserInfo
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	user := &UserInfo{UID: id}
	exist, err := engine.Get(user)
	CheckErr(err)
	if exist {
		return user
	}
	return nil
}
