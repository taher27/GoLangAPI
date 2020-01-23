package main

import (
	"fmt"
	"github.com/Dummy/api/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	// we can create gin server by simply calling

	g := gin.Default()

	//it allow us to very easily add endpoints

	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		// intializing a,b from url passed
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		// parameter, base 10, int64
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Parameter B"})
			return
		}

		// initialize ends here

		req := &proto.Request{A: int64(a), B: int64(b)}
		//add this value to req var using this proto.Request method

		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		// parameter, base 10, int64
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Parameter B"})
			return
		}

		// initialize ends here

		req := &proto.Request{A: int64(a), B: int64(b)}
		//add this value to req var using this proto.Request method

		if response, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server : %v", err)
	}

}
