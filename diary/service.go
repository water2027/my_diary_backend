package diary

import (
	"errors"
	"my_diary/user"
)

func CreateDiaryService(createDiaryReq CreateDiaryRequest) error {
	var store diaryHandler = &diaryModel{
		userId: createDiaryReq.UserId,
		title: createDiaryReq.Title,
		content: createDiaryReq.Content,
		isPublic: createDiaryReq.IsPublic,
	}
	err := store.createDiary()
	if err != nil {
		return err
	}
	return nil
}

func GetDiaryByDiaryIdService(getDiaryReq GetDiaryRequest) (GetDiaryResponse, error) {
	getDiaryResp := GetDiaryResponse{}

	diary := diaryModel{
		diaryId: getDiaryReq.DiaryId,
	}
	var store diaryHandler = &diary
	err := store.getDiaryById()
	if err != nil {
		return getDiaryResp, err
	}

	if diary.isPublic {
		user, err := user.GetUserInfoService(user.GetUserInfoRequest{
			UserId: diary.userId,
		})
		if err != nil {
			return getDiaryResp, err
		}
		getDiaryResp.Author = user.Username
		getDiaryResp.Title = diary.title
		getDiaryResp.Content = diary.content
		getDiaryResp.CreatedAt = diary.createdAt
		getDiaryResp.UpdatedAt = diary.updatedAt
		return getDiaryResp, nil
	}
	if diary.userId != getDiaryReq.UserId {
		return getDiaryResp, errors.New("无权限")
	}
	getDiaryResp.Title = diary.title
	getDiaryResp.Content = diary.content
	getDiaryResp.CreatedAt = diary.createdAt
	getDiaryResp.UpdatedAt = diary.updatedAt
	return getDiaryResp, nil
}

func GetDiaryNumService(getDiaryReq GetDiaryNumRequest) (GetDiaryNumResponse, error) {
	resp := GetDiaryNumResponse{}
	if getDiaryReq.IsPublicSearch {
		// 获取所有isPublic为true日记的数量
		count, err := getPublicDiaryNum()
		if err != nil {
			return resp, nil
		}
		resp.Count = count
		return resp, nil
	}
	count, err := getUserDiaryNum(getDiaryReq.UserId)
	if err != nil {
		return resp, nil
	}
	resp.Count = count
	return resp, nil
}

func GetDiaryListService(getDiaryListReq GetDiaryListRequest) (GetDiaryListResponse, error) {
	resp := GetDiaryListResponse{}
	if getDiaryListReq.IsPublicSearch {
		diaries, err := getPublicDiaryList(getDiaryListReq.Offset, getDiaryListReq.Limit)
		if err != nil {
			return resp, err
		}
		resp.Diaries = diaries
		return resp, nil
	}
	diaries, err := getUserDiaryList(getDiaryListReq.UserId, getDiaryListReq.Offset, getDiaryListReq.Limit)
	if err != nil {
		return resp, err
	}
	resp.Diaries = diaries
	return resp, nil
}

func UpdateDiaryService(updateDiaryReq UpdateDiaryRequest) error {
	diary := diaryModel{
		diaryId: updateDiaryReq.DiaryId,
	}
	var store diaryHandler = &diary
	err := store.getDiaryById()
	if err != nil {
		return err
	}
	if diary.userId != updateDiaryReq.UserId {
		return errors.New("无权限")
	}
	diary.title = updateDiaryReq.Title
	diary.content = updateDiaryReq.Content
	err = store.updateDiary()
	if err != nil {
		return err
	}
	return nil
}

func DeleteDiaryService(deleteDiaryReq DeleteDiaryRequest) error {
	diary := diaryModel{
		diaryId: deleteDiaryReq.DiaryId,
	}
	var store diaryHandler = &diary
	err := store.getDiaryById()
	if err != nil {
		return err
	}
	if diary.userId != deleteDiaryReq.UserId {
		return errors.New("无权限")
	}
	err = store.deleteDiary()
	if err != nil {
		return err
	}
	return nil
}