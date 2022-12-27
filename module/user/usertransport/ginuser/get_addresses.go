package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/user/userbusiness"
	"food_delivery/module/user/userstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAddresses(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewGetAddressBiz(store, requester)

		data, err := biz.GetAddresses(c.Request.Context(), requester.GetUserId())

		if err != nil {
			panic(err)
		}

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ListAddress))
	}

}
