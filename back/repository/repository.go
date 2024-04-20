package repository

import (
	Repolib "photo-share/back/sharelib/domain/repository"
)

type PhotoShareRepository struct {
	trns Repolib.Transactor
}

func NewPhotoShareRepository(trns Repolib.Transactor) (*PhotoShareRepository, error) {
	repo := &PhotoShareRepository{
		trns: trns,
	}
	return repo, nil
}
