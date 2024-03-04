package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Kbgjtn/notethingness-api.git/types"
)

type Task struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"Call John"`
	Priority  int       `json:"priority" example:"1"`
	Date      time.Time `json:"date" example:"2024-03-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2024-03-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-03-01T00:00:00Z"`
}

func (c *Task) toJSON() types.JSONResult {
	return types.JSONResult{
		Data:    c,
		Code:    200,
		Message: "success",
	}
}

type Tasks []Task

func (q Tasks) Len() int {
	return len(q)
}

func (c Tasks) ToJSON(pag types.Pageable) types.JSONResultWithPaginate {
	pag.Calc()
	return types.JSONResultWithPaginate{
		Data:     c,
		Code:     200,
		Message:  "success",
		Paginate: &pag,
		Length:   len(c),
	}
}

func (req *TaskURLParams) Parse(value string) error {
	p, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("param \"id\" is required and must be a number")
	}

	req.ID = p
	return nil
}

func (q Task) CreateTaskResponseDto() types.JSONResult {
	return types.JSONResult{
		Data:    q,
		Code:    200,
		Message: "success",
	}
}

type TaskURLParams struct {
	ID int `json:"id" example:"1" validate:"required"`
}

type TaskRequestPayload struct {
	// ID       string    `json:"content"     example:"I am a quote"`
	Title    string    `json:"title"   example:"Call John"`
	Priority int       `json:"priority"   example:"1"`
	Date     time.Time `json:"date" example:"2024-03-01T00:00:00Z"`
}

// Len is the number of elements in the collection.

// Less reports whether the element with
func (q Tasks) CreateTaskResponseDto(pag *types.Pageable) types.JSONResultWithPaginate {
	if pag.Total > 0 {
		return types.JSONResultWithPaginate{
			Message:  "success",
			Code:     200,
			Data:     q,
			Length:   len(q),
			Paginate: pag,
		}
	}

	return types.JSONResultWithPaginate{
		Code:    200,
		Message: "success",
		Data:    q,
		Length:  len(q),
	}
}
