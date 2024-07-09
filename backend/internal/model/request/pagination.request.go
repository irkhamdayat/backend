package request

import (
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/model/response"
)

type PaginationReq struct {
	Limit int `query:"limit" form:"limit"`
	Page  int `query:"page" form:"page"`
}

func (p *PaginationReq) CountOffset() int {
	return (p.Page - 1) * p.Limit
}

func (p *PaginationReq) processLimitation() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit > constant.PaginationMaxLimit {
		p.Limit = constant.PaginationMaxLimit
	}
	if p.Limit <= constant.PaginationMinLimit {
		p.Limit = constant.PaginationMinLimit
	}
}

func (p *PaginationReq) ToMetaResp() response.MetaResp {
	p.processLimitation()
	return response.MetaResp{
		Page:  p.Page,
		Limit: p.Limit,
	}
}
