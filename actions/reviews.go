package actions

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	reviews "k8s-simple-golang/pb"
	"net/http"
	"strconv"
)

func GetReviews(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	conn, err := grpc.Dial("liwenqiang.site:19999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer conn.Close()
	client := reviews.NewReviewsClient(conn)
	limit := wrapperspb.Int64Value{
		Value: int64(i),
	}
	reviewList, err := client.List(c, &limit)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movies": reviewList.Reviews,
	})
}
