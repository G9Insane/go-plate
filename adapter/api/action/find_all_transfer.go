package action

import (
	"net/http"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/felixa1996/go-plate/usecase"
)

type FindAllTransferAction struct {
	uc  usecase.FindAllTransferUseCase
	log logger.Logger
}

func NewFindAllTransferAction(uc usecase.FindAllTransferUseCase, log logger.Logger) FindAllTransferAction {
	return FindAllTransferAction{
		uc:  uc,
		log: log,
	}
}

func (t FindAllTransferAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_transfer"

	output, err := t.uc.Execute(r.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the transfer list")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning transfer list")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
