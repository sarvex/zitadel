package grpc

import (
	"github.com/caos/logging"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	"github.com/caos/zitadel/internal/model"
	"github.com/golang/protobuf/ptypes"
)

func addIamMemberToModel(member *AddIamMemberRequest) *iam_model.IamMember {
	memberModel := &iam_model.IamMember{
		UserID: member.UserId,
	}
	memberModel.Roles = member.Roles

	return memberModel
}

func changeIamMemberToModel(member *ChangeIamMemberRequest) *iam_model.IamMember {
	memberModel := &iam_model.IamMember{
		UserID: member.UserId,
	}
	memberModel.Roles = member.Roles

	return memberModel
}

func iamMemberFromModel(member *iam_model.IamMember) *IamMember {
	creationDate, err := ptypes.TimestampProto(member.CreationDate)
	logging.Log("GRPC-Lsp76").OnError(err).Debug("date parse failed")

	changeDate, err := ptypes.TimestampProto(member.ChangeDate)
	logging.Log("GRPC-3fG5s").OnError(err).Debug("date parse failed")

	return &IamMember{
		UserId:       member.UserID,
		CreationDate: creationDate,
		ChangeDate:   changeDate,
		Roles:        member.Roles,
		Sequence:     member.Sequence,
	}
}

func iamMemberSearchRequestToModel(request *IamMemberSearchRequest) *iam_model.IamMemberSearchRequest {
	return &iam_model.IamMemberSearchRequest{
		Limit:   request.Limit,
		Offset:  request.Offset,
		Queries: iamMemberSearchQueriesToModel(request.Queries),
	}
}

func iamMemberSearchQueriesToModel(queries []*IamMemberSearchQuery) []*iam_model.IamMemberSearchQuery {
	modelQueries := make([]*iam_model.IamMemberSearchQuery, len(queries))
	for i, query := range queries {
		modelQueries[i] = iamMemberSearchQueryToModel(query)
	}

	return modelQueries
}

func iamMemberSearchQueryToModel(query *IamMemberSearchQuery) *iam_model.IamMemberSearchQuery {
	return &iam_model.IamMemberSearchQuery{
		Key:    iamMemberSearchKeyToModel(query.Key),
		Method: iamMemberSearchMethodToModel(query.Method),
		Value:  query.Value,
	}
}

func iamMemberSearchKeyToModel(key IamMemberSearchKey) iam_model.IamMemberSearchKey {
	switch key {
	case IamMemberSearchKey_IAMMEMBERSEARCHKEY_EMAIL:
		return iam_model.IamMemberSearchKeyEmail
	case IamMemberSearchKey_IAMMEMBERSEARCHKEY_FIRST_NAME:
		return iam_model.IamMemberSearchKeyFirstName
	case IamMemberSearchKey_IAMMEMBERSEARCHKEY_LAST_NAME:
		return iam_model.IamMemberSearchKeyLastName
	case IamMemberSearchKey_IAMMEMBERSEARCHKEY_USER_ID:
		return iam_model.IamMemberSearchKeyUserID
	default:
		return iam_model.IamMemberSearchKeyUnspecified
	}
}

func iamMemberSearchMethodToModel(key SearchMethod) model.SearchMethod {
	switch key {
	case SearchMethod_SEARCHMETHOD_CONTAINS:
		return model.SearchMethodContains
	case SearchMethod_SEARCHMETHOD_CONTAINS_IGNORE_CASE:
		return model.SearchMethodContainsIgnoreCase
	case SearchMethod_SEARCHMETHOD_EQUALS:
		return model.SearchMethodEquals
	case SearchMethod_SEARCHMETHOD_EQUALS_IGNORE_CASE:
		return model.SearchMethodEqualsIgnoreCase
	case SearchMethod_SEARCHMETHOD_STARTS_WITH:
		return model.SearchMethodStartsWith
	case SearchMethod_SEARCHMETHOD_STARTS_WITH_IGNORE_CASE:
		return model.SearchMethodStartsWithIgnoreCase
	default:
		return -1
	}
}

func iamMemberSearchResponseFromModel(resp *iam_model.IamMemberSearchResponse) *IamMemberSearchResponse {
	return &IamMemberSearchResponse{
		Limit:       resp.Limit,
		Offset:      resp.Offset,
		TotalResult: resp.TotalResult,
		Result:      iamMembersFromView(resp.Result),
	}
}
func iamMembersFromView(viewMembers []*iam_model.IamMemberView) []*IamMemberView {
	members := make([]*IamMemberView, len(viewMembers))

	for i, member := range viewMembers {
		members[i] = iamMemberFromView(member)
	}

	return members
}

func iamMemberFromView(member *iam_model.IamMemberView) *IamMemberView {
	changeDate, err := ptypes.TimestampProto(member.ChangeDate)
	logging.Log("GRPC-Lso9c").OnError(err).Debug("unable to parse changedate")
	creationDate, err := ptypes.TimestampProto(member.CreationDate)
	logging.Log("GRPC-6szE").OnError(err).Debug("unable to parse creation date")

	return &IamMemberView{
		ChangeDate:   changeDate,
		CreationDate: creationDate,
		Roles:        member.Roles,
		Sequence:     member.Sequence,
		UserId:       member.UserID,
		UserName:     member.UserName,
		Email:        member.Email,
		FirstName:    member.FirstName,
		LastName:     member.LastName,
	}
}
