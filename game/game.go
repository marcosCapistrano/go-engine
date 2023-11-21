package game

import (
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
		fmt.Println(deltaTime)
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
	g.registry = ecs.NewRegistry()

	player := ecs.NewEntity("player", []ecs.Component{
		components.NewPosition().(*components.Position).WithX(0).WithY(0),
		components.NewSize().(*components.Size).WithWidth(32).WithHeight(32),
		components.NewVelocity().(*components.Velocity).WithX(1).WithY(2),
	})

	player2 := ecs.NewEntity("player2", []ecs.Component{
		components.NewPosition().(*components.Position).WithX(0).WithY(0),
		components.NewSize().(*components.Size).WithWidth(32).WithHeight(32),
		components.NewVelocity().(*components.Velocity).WithX(30).WithY(20),
	})

	renderingSystem := systems.NewRendering(g.renderer)
	movementSystem := systems.NewMovement().(*systems.Movement).WithData(&deltaTime)

	g.registry.AddEntities(player)
	g.registry.AddEntities(player2)

	g.registry.AddSystems(renderingSystem)
	g.registry.AddSystems(movementSystem)
}
