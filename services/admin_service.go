package services

type AdminService interface{}

type adminService struct{}

func Constructor() AdminService {
	return adminService{}
}
