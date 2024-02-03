package web

import (
	"CF_PROJECT/store"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	kGetRecentActions = "/recent/actions"
)
const (
	GetUserData = "/user/signup"
)

func CreateWebServer(mongoStore *store.MongoStore) *Server {
	srv := new(Server)
	srv.r = gin.Default()
	srv.store = mongoStore

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	srv.r.Use(cors.New(config))

	srv.r.GET(kGetRecentActions, srv.RecentActionsHandler)

	return srv
}

type Server struct {
	r     *gin.Engine
	store *store.MongoStore
}

func (srv *Server) RecentActionsHandler(ctx *gin.Context) {
	recentActions, err := srv.store.QueryRecentActions()
	if err != nil {
		log.Printf("Error occurred while fetching recentActions: %v", err)
		ctx.String(http.StatusBadRequest, "Error while getting recent actions")
	}

	ctx.JSON(http.StatusOK, recentActions)
}

func (srv *Server) StartListeningRequests(addr string) error {
	err := srv.r.Run(addr)
	if err != nil {
		log.Printf("Error occurred while starting server: %v", err)
		return nil
	}
	return err
}

