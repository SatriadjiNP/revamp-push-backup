package hrController

import (
	"net/http"

	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
)

type TalentsDetailMockupController struct {
	talentDetailService *hrService.TalentsDetailMockupService
}

// declare constructor
func NewTalentDetailMockupController(talentDetailService *hrService.TalentsDetailMockupService) *TalentsDetailMockupController {
	return &TalentsDetailMockupController{
		// struct 				parameter
		talentDetailService: talentDetailService,
	}
}

func (talentDetailController TalentsDetailMockupController) GetListTalentDetailMockupDetail(ctx *gin.Context) {
	responses, responseErr := talentDetailController.talentDetailService.GetListTalentDetailMockup(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, responses)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}
