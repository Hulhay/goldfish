package usecase

import (
	"context"
	"errors"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/repository"
	"github.com/Hulhay/goldfish/usecase/family"
	"github.com/Hulhay/goldfish/usecase/member"
	"github.com/google/uuid"
)

type memberUC struct {
	memberRepo repository.MemberRepository
	familyUC   Family
}

type Member interface {
	InsertMember(ctx context.Context, params member.InsertMemberRequest) error
	GetMember(ctx context.Context, params member.GetMemberRequest) ([]member.MemberListResponse, error)
	GetDetailMember(ctx context.Context, params member.GetMemberDetailRequest) (*member.MemberDetailResponse, error)
}

func NewMemberUC(mr repository.MemberRepository, fu Family) Member {
	return &memberUC{
		memberRepo: mr,
		familyUC:   fu,
	}
}

func (u *memberUC) InsertMember(ctx context.Context, params member.InsertMemberRequest) error {

	var (
		err         error
		memberExist *model.Member
	)

	if err = params.Validate(); err != nil {
		return err
	}

	if !params.MemberIsHead {
		member, err := u.memberRepo.GetMemberByFamilyID(ctx, params.FamilyID)
		if err != nil || member == nil {
			return errors.New("family_id not found")
		}
	} else {
		params.FamilyID = uuid.NewString()
	}
	memberExist, _ = u.memberRepo.GetMemberByMemberNIK(ctx, params.MemberNIK)
	if memberExist != nil {
		return errors.New("nik exist. use other nik")
	}

	err = u.memberRepo.InsertMember(ctx, &model.Member{
		MemberNIK:    params.MemberNIK,
		MemberName:   params.MemberName,
		MemberIsHead: params.MemberIsHead,
		FamilyID:     params.FamilyID,
	})
	if err != nil {
		return err
	}

	if params.MemberIsHead {
		memberExist, err = u.memberRepo.GetMemberByMemberNIK(ctx, params.MemberNIK)
		if err != nil {
			return err
		}

		err = u.familyUC.InsertFamily(ctx, family.InsertFamilyRequest{
			FamilyID:           params.FamilyID,
			FamilyNIK:          params.FamilyNIK,
			FamilyMemberHeadID: memberExist.ID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *memberUC) GetMember(ctx context.Context, params member.GetMemberRequest) ([]member.MemberListResponse, error) {

	var (
		members []member.MemberListResponse
		err     error
	)

	if err = params.Validate(); err != nil {
		return nil, err
	}

	members, err = u.memberRepo.GetMember(ctx, params)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (u *memberUC) GetDetailMember(ctx context.Context, params member.GetMemberDetailRequest) (*member.MemberDetailResponse, error) {

	if err := params.Validate(); err != nil {
		return nil, err
	}

	memberData, err := u.memberRepo.GetMember(ctx, member.GetMemberRequest{
		MemberNIK: params.MemberNIK,
	})
	if err != nil {
		return nil, err
	}
	if len(memberData) == 0 || memberData == nil {
		return nil, errors.New("member not found")
	}

	res := &member.MemberDetailResponse{
		MemberNIK:    memberData[0].MemberNIK,
		MemberName:   memberData[0].MemberName,
		FamilyNIK:    memberData[0].FamilyNIK,
		FamilyID:     memberData[0].FamilyID,
		MemberIsHead: memberData[0].MemberIsHead,
	}

	return res, nil
}
