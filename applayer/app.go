package applayer

import (
	"github.com/gin-gonic/gin"
)

type App interface {
    Create(ctx *gin.Context)
    Read(ctx *gin.Context)
    Update(ctx *gin.Context)
    Delete(ctx *gin.Context)
}

/*
type app struct {
    store storelayer.Store
}

func New(store storelayer.Store) *app {
    return &app{
        store: store,
    }
}
*/