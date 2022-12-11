package usecase

import (
	"context"
	"errors"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/repository"
	"github.com/Hulhay/goldfish/usecase/family"
)

type familyUC struct {
	familyRepo repository.FamilyRepository
}

type Family interface {
	InsertFamily(ctx context.Context, params family.InsertFamilyRequest) error
}

func NewFamilyUC(fr repository.FamilyRepository) Family {
	return &familyUC{
		familyRepo: fr,
	}
}

func (u *familyUC) InsertFamily(ctx context.Context, params family.InsertFamilyRequest) error {

	var (
		err         error
		familyExist *model.Family
	)
	if err = params.Validate(); err != nil {
		return err
	}

	familyExist, _ = u.familyRepo.GetFamilyByFamilyNIK(ctx, params.FamilyNIK)
	if familyExist != nil {
		return errors.New("nik (keluarga) exist. use other nik")
	}

	err = u.familyRepo.InsertFamily(ctx, &model.Family{
		FamilyID:           params.FamilyID,
		FamilyNIK:          params.FamilyNIK,
		FamilyMemberHeadID: params.FamilyMemberHeadID,
	})
	if err != nil {
		return err
	}

	return nil
}
