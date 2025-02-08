package diary

import (
	"errors"
	"time"
)

type diaryModel struct {
	diaryId   int
	userId    int
	title     string
	content   string
	createdAt time.Time
	updatedAt time.Time
	isPublic  bool
}

type Diary struct {
	DiaryId   int       `json:"diaryId"`
	UserId    int       `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsPublic  bool      `json:"isPublic"`
}

type CreateDiaryRequest struct {
	UserId   int    `json:"userId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsPublic bool   `json:"isPublic"`
}

func (cdr *CreateDiaryRequest) Examine() error {
	if cdr.UserId <= 0 {
		return errors.New("unknown error")
	}
	if cdr.Title == "" || cdr.Content == "" {
		return errors.New("缺少标题或内容")
	}
	return nil
}

type GetDiaryRequest struct {
	UserId  int `json:"userId"`
	DiaryId int `json:"diaryId"`
}

func (gdr *GetDiaryRequest) Examine() error {
	if gdr.DiaryId == 0 {
		return errors.New("no diaryId")
	}
	return nil
}

type GetDiaryResponse struct {
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetDiaryNumRequest struct {
	IsPublicSearch bool `json:"isPublicSearch"`
	UserId         int  `json:"userId"`
}

func (gdnr *GetDiaryNumRequest) Examine() error {
	if !gdnr.IsPublicSearch && gdnr.UserId <= 0 {
		return errors.New("no login")
	}
	return nil
}

type GetDiaryNumResponse struct {
	Count int `json:"count"`
}

type GetDiaryListRequest struct {
	IsPublicSearch bool `json:"isPublicSearch"`
	UserId         int  `json:"userId"`
	Offset         int  `json:"offset"`
	Limit          int  `json:"limit"`
}

func (gdlr *GetDiaryListRequest) Examine() error {
	if !gdlr.IsPublicSearch && gdlr.UserId <= 0 {
		return errors.New("no login")
	}
	return nil
}

type GetDiaryListResponse struct {
	Diaries []Diary `json:"diaries"`
}

type UpdateDiaryRequest struct {
	UserId  int    `json:"userId"`
	DiaryId int    `json:"diaryId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (udr *UpdateDiaryRequest) Examine() error {
	if udr.UserId <= 0 {
		return errors.New("no login")
	}
	if udr.DiaryId <= 0 {
		return errors.New("no diaryId")
	}
	if udr.Title == "" || udr.Content == "" {
		return errors.New("no title or content")
	}
	return nil
}

type DeleteDiaryRequest struct {
	UserId  int `json:"userId"`
	DiaryId int `json:"diaryId"`
}

func (ddr *DeleteDiaryRequest) Examine() error {
	if ddr.UserId <= 0 {
		return errors.New("no login")
	}
	if ddr.DiaryId <= 0 {
		return errors.New("no diaryId")
	}
	return nil
}
