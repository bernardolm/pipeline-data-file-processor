package processor

import "context"

type IWorker interface {
	Process(context.Context)
}

type ISource interface {
	GetData(ctx context.Context, offset int, limit int)
	GetName(context.Context) string
	PostProcessData(context.Context)
}

type IHook interface {
	Execute(context.Context)
	GetName(context.Context) string
}

type IOutput interface {
	Build(context.Context)
	BuildBottomContent(context.Context)
	BuildTopContent(context.Context)
}

type ISender interface {
	GetFileName(context.Context)
	GetURL(context.Context)
	GetWrittenBytesCount(context.Context)
	PostProcess(context.Context)
	PreProcess(context.Context)
	SendData(context.Context)
}
