package game

import (
	"engine/config"
	"engine/ecs"
	"engine/game/components"
	"engine/game/systems"
	"engine/math/vector"

	"github.com/veandco/go-sdl2/sdl"
)

const FPS uint64 = 60
const MILLISECS_PER_FRAME uint64 = 1000 / FPS
const PIXELS_PER_METER = 50

var deltaTime float64

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

		deltaTime = float64(sdl.GetTicks64()-g.millisecsPrevFrame) / 1000.0
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
		&components.Position{
			Vector: vector.Vector2{X: 0, Y: 0},
		},
		&components.Rotation{
			Value: 0,
		},
		&components.Mass{
			Value:   1,
			Inverse: 1,
		},
		&components.MomentOfInertia{
			Value:   0.5,
			Inverse: 2,
		},
		&components.LinearMotion{
			Velocity:     vector.NewVector2(0, 0),
			Acceleration: vector.NewVector2(0, 0),
			Forces:       vector.NewVector2(0, 0),
		},
		&components.AngularMotion{
			Velocity:     0,
			Acceleration: 0,
			Torque:       0,
		},
		&components.Shape{
			Type:   components.CircleType,
			Radius: 40,
		},
	})

	renderingSystem := systems.NewRendering(g.renderer)
	movementSystem := systems.NewMovement().(*systems.Movement).WithData(&deltaTime)

	g.registry.AddEntities(player)

	g.registry.AddSystems(renderingSystem)
	g.registry.AddSystems(movementSystem)
}
