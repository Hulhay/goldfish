package repository

import (
	"context"

	"github.com/Hulhay/goldfish/model"
	"gorm.io/gorm"
)

type memberRepository struct {
	qry *gorm.DB
}

type MemberRepository interface {
	InsertMember(ctx context.Context, params *model.Member) error
	GetMemberByFamilyID(ctx context.Context, familyID string) (*model.Member, error)
	GetMemberByMemberNIK(ctx context.Context, memberNIK string) (*model.Member, error)
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{
		qry: db,
	}
}

func (r *memberRepository) InsertMember(ctx context.Context, params *model.Member) error {
	var member *model.Member

	if err := r.qry.Model(&member).Create(map[string]interface{}{
		"member_nik":     params.MemberNIK,
		"member_name":    params.MemberName,
		"member_is_head": params.MemberIsHead,
		"family_id":      params.FamilyID,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *memberRepository) GetMemberByFamilyID(ctx context.Context, familyID string) (*model.Member, error) {
	var member *model.Member

	if err := r.qry.Model(&member).Where("family_id = ?", familyID).First(&member).Error; err != nil {
		return nil, err
	}

	return member, nil
}

func (r *memberRepository) GetMemberByMemberNIK(ctx context.Context, memberNIK string) (*model.Member, error) {
	var member *model.Member

	if err := r.qry.Model(&member).Where("member_nik = ?", memberNIK).First(&member).Error; err != nil {
		return nil, err
	}

	return member, nil
}
