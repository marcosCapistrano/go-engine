package game

import (
	"engine/config"
	"engine/ecs"
	"engine/game/entities"
	"engine/game/systems"
	"math/rand"
	"time"

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

	windowWidth := 1920
	windowHeight := 600

	window, err := sdl.CreateWindow("2D Game Engine", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(windowWidth), int32(windowHeight), sdl.WINDOW_BORDERLESS)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}

	//window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)

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

	go func() {
		for i := 0; i < 1000; i++ {
			go func() {
				e := ecs.NewEntity("player", entities.NewCircle(rand.Float64()*1280, rand.Float64()*300, 0, 1+rand.Float64()*8, 8))
				g.registry.AddEntities(e)
			}()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	player := ecs.NewEntity("player", entities.NewCircle(0, 0, 0, 4, 16))
	player2 := ecs.NewEntity("player2", entities.NewCircle(100, 0, 0, 4, 14))
	player3 := ecs.NewEntity("player3", entities.NewCircle(200, 0, 0, 4, 40))
	player4 := ecs.NewEntity("player4", entities.NewCircle(300, 0, 0, 4, 24))
	player5 := ecs.NewEntity("player5", entities.NewBox(400, 0, 0, 4, 100, 100))

	renderingSystem := systems.NewRendering(g.renderer)
	movementSystem := systems.NewMovement().(*systems.Movement).WithData(&deltaTime)

	g.registry.AddEntities(player)
	g.registry.AddEntities(player2)
	g.registry.AddEntities(player3)
	g.registry.AddEntities(player4)
	g.registry.AddEntities(player5)

	g.registry.AddSystems(renderingSystem)
	g.registry.AddSystems(movementSystem)
}
