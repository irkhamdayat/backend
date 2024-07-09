package notification

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/entity"
	"github.com/Halalins/backend/internal/model/request"
)

func (r *Repository) GetPagination(ctx context.Context, req request.GetNotificationPaginationReq) (
	result []entity.GetNotificationHistory, count int64, err error) {
	var (
		offset = req.Pagination.CountOffset()
		lang   = util.GetAcceptLanguageFromContext(ctx)
		logger = logrus.WithFields(logrus.Fields{
			"ctx":       util.DumpIncomingContext[entity.UserClaim](ctx, constant.ListUserClaimKeys),
			"companyID": req.CompanyID.String(),
		})
	)

	query := r.db.WithContext(ctx).
		Select(`
			nh.id,
       		nh.is_read,
       		nh.image,
       		nh.action_type,
       		nh.notification_type,
       		nh.action_type,
       		nh.additional_data,
       		nh.created_at,
       		tnh.language,
       		tnh.headline,
       		tnh.message
		`).
		Table("notification_histories nh").
		Joins(`JOIN translate_notification_histories tnh ON nh.id = tnh.notification_history_id`).
		Where(`tnh.language = ?`, lang).
		Order("created_at desc")

	err = query.
		Count(&count).
		Error

	if err != nil {
		logger.Errorf("failed to get count: %v", err)
		return nil, 0, err
	}

	err = query.
		Limit(req.Pagination.Limit).
		Offset(offset).
		Scan(&result).
		Error

	if err != nil {
		logger.Errorf("failed to get pagination: %v", err)
		return nil, 0, err
	}
	return
}
