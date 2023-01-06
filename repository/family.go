package repository

import (
	"context"

	"github.com/Hulhay/goldfish/model"
	"gorm.io/gorm"
)

type familyRepository struct {
	qry *gorm.DB
}

type FamilyRepository interface {
	InsertFamily(ctx context.Context, params *model.Family) error
	GetFamilyByFamilyNIK(ctx context.Context, familyNIK string) (*model.Family, error)
	GetFamilyByFamilyID(ctx context.Context, familyID string) (*model.Family, error)
}

func NewFamilyRepository(db *gorm.DB) FamilyRepository {
	return &familyRepository{
		qry: db,
	}
}

func (r *familyRepository) InsertFamily(ctx context.Context, params *model.Family) error {
	var family *model.Family

	if err := r.qry.Model(&family).Create(map[string]interface{}{
		"family_id":             params.FamilyID,
		"family_nik":            params.FamilyNIK,
		"family_member_head_id": params.FamilyMemberHeadID,
	}).Error; err != nil {
		return nil
	}

	return nil
}

func (r *familyRepository) GetFamilyByFamilyNIK(ctx context.Context, familyNIK string) (*model.Family, error) {
	var family *model.Family

	if err := r.qry.Model(&family).Where("family_nik = ?", familyNIK).First(&family).Error; err != nil {
		return nil, err
	}

	return family, nil
}

func (r *familyRepository) GetFamilyByFamilyID(ctx context.Context, familyID string) (*model.Family, error) {
	var family *model.Family

	if err := r.qry.Model(&family).Where("family_id = ?", familyID).First(&family).Error; err != nil {
		return nil, err
	}

	return family, nil
}
