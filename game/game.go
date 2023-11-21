package game

import (
	"engine/config"
	"engine/ecs"
	"engine/game/components"
	"engine/game/systems"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const FPS uint64 = 60
const MILLISECS_PER_FRAME uint64 = 1000 / FPS

var deltaTime float32

type Game struct {
	isRunning          bool
	millisecsPrevFrame uint64
	window             *sdl.Window
	renderer           *sdl.Renderer

	config   config.Configuration
	registry ecs.Registry
}

func New() *Game {
	return &Game{
		isRunning: false,
		registry:  ecs.NewRegistry(),
	}
}

func (g *Game) Initialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	windowWidth := 800
	windowHeight := 600

	window, err := sdl.CreateWindow("2D Game Engine", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(windowWidth), int32(windowHeight), sdl.WINDOW_BORDERLESS)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}

	window.SetFullscreen(sdl.WINDOW_FULLSCREEN)

	g.window = window
	g.renderer = renderer
	g.isRunning = true
	g.millisecsPrevFrame = sdl.GetTicks64()
}

func (g *Game) Run() {
	g.setup()

	for g.isRunning {
		timeToWait := MILLISECS_PER_FRAME - (sdl.GetTicks64() - g.millisecsPrevFrame)
		if timeToWait > 0 && timeToWait <= MILLISECS_PER_FRAME {
			sdl.Delay(uint32(timeToWait))
		}

		deltaTime = float32(sdl.GetTicks64()-g.millisecsPrevFrame) / 1000.0
		g.millisecsPrevFrame = sdl.GetTicks64()
		g.processInput()

		for _, system := range g.registry.Systems() {
			system.Process(g.registry)
		}
	}
}

func (g *Game) Destroy() {
	g.renderer.Destroy()
	g.window.Destroy()
	sdl.Quit()
}

func (g *Game) setup() {
	config, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}

	g.config = config
	g.registry = ecs.NewRegistry()

	player := ecs.NewEntity("player", []ecs.Component{
		components.NewPosition().(*components.Position).WithX(0).WithY(0),
		components.NewSize().(*components.Size).WithWidth(32).WithHeight(32),
		components.NewVelocity().(*components.Velocity).WithX(1).WithY(2),
		components.NewAcceleration().(*components.Acceleration).WithX(0).WithY(0),
	})

	player2 := ecs.NewEntity("player2", []ecs.Component{
		components.NewPosition().(*components.Position).WithX(0).WithY(0),
		components.NewSize().(*components.Size).WithWidth(32).WithHeight(32),
		components.NewVelocity().(*components.Velocity).WithX(30).WithY(100),
		components.NewAcceleration().(*components.Acceleration).WithX(0).WithY(0),
		components.NewBoxCollider().(*components.BoxCollider).WithWidth(32).WithHeight(32),
	})

	platformCount, _ := config.GetInt("platforms:count")
	xs, _ := config.GetSliceOfInt("platforms:x")
	ys, _ := config.GetSliceOfInt("platforms:y")
	ws, _ := config.GetSliceOfInt("platforms:w")
	hs, _ := config.GetSliceOfInt("platforms:h")

	fmt.Println(platformCount, xs, ys)

	for i := 0; i < platformCount; i++ {
		platform := ecs.NewEntity("platform", []ecs.Component{
			components.NewPosition().(*components.Position).WithX(float32(xs[i])).WithY(float32(ys[i])),
			components.NewSize().(*components.Size).WithWidth(float32(ws[i])).WithHeight(float32(hs[i])),
			components.NewBoxCollider().(*components.BoxCollider).WithWidth(int(ws[i])).WithHeight(int(hs[i])),
		})

		g.registry.AddEntities(platform)

	}

	renderingSystem := systems.NewRendering(g.renderer)
	movementSystem := systems.NewMovement().(*systems.Movement).WithData(&deltaTime)
	collisionSystem := systems.NewCollision().(*systems.Collision)

	g.registry.AddEntities(player)
	g.registry.AddEntities(player2)

	g.registry.AddSystems(renderingSystem)
	g.registry.AddSystems(movementSystem)
	g.registry.AddSystems(collisionSystem)
}
