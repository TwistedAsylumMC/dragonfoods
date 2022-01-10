package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenBeans struct{}

// AlwaysConsumable ...
func (GreenBeans) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenBeans) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenBeans) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenBeans) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenBeans) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenBeans) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_beans", 0
}

// Name ...
func (GreenBeans) Name() string {
	return "Green Beans"
}

// Texture ...
func (GreenBeans) Texture() image.Image {
	return textureFromName("green_bean")
}
