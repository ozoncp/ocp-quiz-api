package producer

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"

	ocp_quiz_api "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"
)

type EventType int

const (
	CreateEvent EventType = iota
	ReadEvent
	UpdateEvent
	DeleteEvent
)

type EventMsg interface {
	sarama.Encoder
}

func NewEvent(ctx context.Context, quizId uint64, eventType EventType, err error) EventMsg {
	e := &event{
		quizId:    quizId,
		eventType: eventType,
		err:       err,
	}

	spanDump := opentracing.TextMapCarrier{}
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		if err := opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, spanDump); err != nil {
			log.Warn().Msgf("failed to update event message with span info %v", err)
		} else {
			e.span = spanDump
		}
	}
	return e
}

type event struct {
	quizId      uint64
	eventType   EventType
	err         error
	traceId     string
	span        map[string]string
	encodedData []byte //caching to avoid double encoding on Length() and Encode()
	encodeErr   error
}

func (e *event) Encode() ([]byte, error) {
	if e.encodedData != nil || e.encodeErr != nil {
		return e.encodedData, e.encodeErr
	}

	message := &ocp_quiz_api.QuizAPIEvent{
		QuizId: e.quizId,
	}
	if e.err != nil {
		message.Error = e.err.Error()
	}

	switch e.eventType {
	case CreateEvent:
		message.Event = ocp_quiz_api.QuizAPIEvent_CREATE
	case ReadEvent:
		message.Event = ocp_quiz_api.QuizAPIEvent_READ
	case UpdateEvent:
		message.Event = ocp_quiz_api.QuizAPIEvent_UPDATE
	case DeleteEvent:
		message.Event = ocp_quiz_api.QuizAPIEvent_DELETE
	default:
		log.Panic().Msgf("unexpected event type: %v", e.eventType)
	}

	if len(e.span) > 0 {
		message.TraceSpan = e.span
	}
	e.encodedData, e.encodeErr = proto.Marshal(message)
	return e.encodedData, e.encodeErr
}

func (e *event) Length() int {
	data, _ := e.Encode()
	return len(data)
}
