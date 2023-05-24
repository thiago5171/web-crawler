package crawler

import (
	"backend_template/src/core/domain"
	"backend_template/src/core/domain/errors"
	"time"

	"github.com/google/uuid"
)

type VisitedLinks interface {
	domain.Model

	ID() *uuid.UUID
	Url() string
	Website() string
	CheckedDate() time.Time

	SetID(uuid.UUID)
	SetUrl(string)
	SetWebSite(string)
	SetCheckedDate(time2 time.Time)
}

type visitedLinks struct {
	id      *uuid.UUID
	website string
	url     string

	checkedDate time.Time
}

func New(id *uuid.UUID, url, webSite string, checkedDate time.Time) VisitedLinks {
	data := &visitedLinks{id, webSite, url, checkedDate}
	return data
}

func (acc *visitedLinks) ID() *uuid.UUID {
	return acc.id
}

func (acc *visitedLinks) Url() string {
	return acc.url
}

func (acc *visitedLinks) CheckedDate() time.Time {
	return acc.checkedDate
}

func (acc *visitedLinks) Website() string {
	return acc.website
}
func (acc *visitedLinks) SetID(id uuid.UUID) {
	acc.id = &id
}

func (acc *visitedLinks) SetUrl(url string) {
	acc.url = url
}

func (acc *visitedLinks) SetWebSite(webSite string) {
	acc.website = webSite
}

func (acc *visitedLinks) SetCheckedDate(date time.Time) {
	acc.checkedDate = date
}

func (acc *visitedLinks) IsValid() errors.Error {
	//TODO implement me
	panic("implement me")
}
