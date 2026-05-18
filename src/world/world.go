package world

import (
	"potato-bones/src/world/entities"
	"potato-bones/src/globals"
	"sync"
)

var deltaTime float64 = *globals.GameSpeed / float64(*globals.Tickrate)

type World struct{
	entities []*entities.Entity
	staticEntities []*entities.Entity
	mu sync.RWMutex
}

func (world *World) UpdateGame() {
	world.mu.RLock()
    defer world.mu.RUnlock()

	for _, entity := range(world.entities) {
		entity.Update(deltaTime)
	}
}

func (world *World) AddEntity(entity *entities.Entity) {
    world.mu.Lock()
    defer world.mu.Unlock()

    for index, entity := range world.entities {
        if entity == nil {
            world.entities[index] = entity
            return
        }
    }

    world.entities = append(world.entities, entity)
}


func CreateWorld() *World {
	world := &World{
		entities: make([]*entities.Entity, *globals.MaxEntities),
	}

	return world
}