package usecase

import (
	"context"
	"errors"
	"fmt"

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
		fmt.Printf("memberExist.ID: %v\n", memberExist.ID)

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
