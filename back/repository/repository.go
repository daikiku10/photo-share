package repository

type PhotoShareRepository struct{}

func NewPhotoShareRepository() (*PhotoShareRepository, error) {
	repo := &PhotoShareRepository{}
	return repo, nil
}
