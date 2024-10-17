package context

import (
	c "context"
	"image-retrieval/internal/runtime"
	"log"
)

type BaseContext struct {
	Ctx            c.Context
	Canceller      c.CancelFunc
	stageList      []Stage
	deferStageList []Stage
	BaseError      error
}
type StageHandler func(ctx c.Context) (err error)

type Stage struct {
	handler   StageHandler
	stageName string
	// stageErr  string
}

func (ctx *BaseContext) Init() {
	ctx.Ctx, ctx.Canceller = c.WithCancel(c.Background())
	ctx.stageList = make([]Stage, 0)
}
func (ctx *BaseContext) AddBaseHandler(handler StageHandler, stageName string) *BaseContext {
	ctx.stageList = append(ctx.stageList, Stage{
		handler:   handler,
		stageName: stageName,
	})
	return ctx
}
func (ctx *BaseContext) AddDeferHandler(handler StageHandler, stageName string) *BaseContext {
	ctx.deferStageList = append(ctx.deferStageList, Stage{
		handler:   handler,
		stageName: stageName,
	})
	return ctx
}
func (ctx *BaseContext) Run() {
	// defer runtime.PrintPanic(&ctx.BaseError)
	defer ctx.Canceller()
	stageDeferHandlerWrapper := func(handler StageHandler, stageName string) {
		if err := handler(ctx.Ctx); err != nil {
			log.Fatal("Error in stage:", stageName, "Error:", err.Error())
			return
		}
	}
	for _, deferStage := range ctx.deferStageList {
		defer runtime.PrintPanic(&ctx.BaseError)
		defer stageDeferHandlerWrapper(deferStage.handler, deferStage.stageName)

	}
	for _, stage := range ctx.stageList {
		defer runtime.PrintPanic(&ctx.BaseError)
		if err := stage.handler(ctx.Ctx); err != nil {
			log.Println("Error in stage:", stage.stageName, "Error:", err.Error())
			ctx.BaseError = err
			return
		}
	}

}
