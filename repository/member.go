package repository

import (
	"context"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/usecase/member"
	"gorm.io/gorm"
)

type memberRepository struct {
	qry *gorm.DB
}

type MemberRepository interface {
	InsertMember(ctx context.Context, params *model.Member) error
	GetMember(ctx context.Context, params member.GetMemberRequest) ([]member.MemberListResponse, error)
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

func (r *memberRepository) GetMember(ctx context.Context, params member.GetMemberRequest) ([]member.MemberListResponse, error) {
	var res []member.MemberListResponse

	db := r.qry.Model(&model.Member{}).Select("*").Joins("JOIN families ON families.family_id = members.family_id")

	if params.MemberNIK != `` {
		db = db.Where("members.member_nik = ?", params.MemberNIK)
	}

	if params.FamilyNIK != `` {
		db = db.Where("families.family_nik = ?", params.FamilyNIK)
	}

	if params.MemberName != `` {
		db = db.Where("members.member_name LIKE ?", `%`+params.MemberName+`%`)
	}

	if params.IsHead {
		db = db.Where("members.member_is_head = ?", params.IsHead)
	}

	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
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
