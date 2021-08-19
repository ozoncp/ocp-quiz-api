package saver

import (
	"sync"
	"time"

	"github.com/ozoncp/ocp-quiz-api/internal/flusher"
	"github.com/ozoncp/ocp-quiz-api/internal/models"
)

type Saver interface {
	Save(entity models.Quiz)
	Close()
}

type saver struct {
	capacity uint
	flusher  flusher.Flusher
	ticker   *time.Ticker
	entity   chan models.Quiz
	close    chan struct{}
	wait     *sync.WaitGroup
}

func (s *saver) Save(entity models.Quiz) {
	select {
	case <-s.close:
		return
	default:
		s.entity <- entity
	}
}

func (s *saver) Close() {
	close(s.close)
	s.wait.Wait()
}

func NewSaver(cap uint, flusher flusher.Flusher, timeout time.Duration) Saver {
	s := &saver{
		capacity: cap,
		flusher:  flusher,
		ticker:   time.NewTicker(timeout),
		entity:   make(chan models.Quiz),
		close:    make(chan struct{}),
		wait:     &sync.WaitGroup{},
	}
	s.start()
	return s

}

func (s *saver) start() {
	entities := make([]models.Quiz, 0, s.capacity)
	s.wait.Add(1)
	go func() {
		defer func() {
			s.ticker.Stop()
			close(s.entity)
			_, _ = s.flusher.Flush(entities)
			s.wait.Done()
		}()
		for {
			select {
			case <-s.close:
				return
			case <-s.ticker.C:
				entities, _ = s.flusher.Flush(entities)
			case entity := <-s.entity:
				entities = append(entities, entity)
			}
		}
	}()
}
