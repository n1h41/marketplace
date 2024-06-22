package repositories

type AdminRepository interface {
	Login()
}

type adminRepository struct{}

func Constructor() AdminRepository {
	return &adminRepository{}
}

func (r adminRepository) Login() {
	// TODO:
}
