package inMemoryPersistence

import (
	"sync"
	"testing"
	"warm-up/internal/domain"

	"github.com/google/uuid"
)

func TestInMemoryOrderRepository_ConcurrentAccess(t *testing.T) {
	repo := NewInMemoryOrderRepository()

	seedIDs := make([]string, 0, 50)
	for i := 0; i < 50; i++ {
		id := uuid.New()
		seedIDs = append(seedIDs, id.String())
		if err := repo.Save(domain.Order{ID: id, Amount: 10}); err != nil {
			t.Fatalf("seed save failed: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				id := uuid.New()
				_ = repo.Save(domain.Order{ID: id, Amount: float64(i)})
				return
			}

			id := seedIDs[i%len(seedIDs)]
			_, err := repo.FindById(id)
			if err != nil {
				t.Errorf("expected order for id %s, got error: %v", id, err)
			}
		}(i)
	}

	wg.Wait()
}
