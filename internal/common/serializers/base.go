package serializers

import "context"

type BaseSerializer[TInput any, TOutput any] interface {
	Serialize(ctx context.Context, value TInput) TOutput
}
