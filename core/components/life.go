package components

type Life interface {
	LifeCycle() *LifeCycle
}

////////////////////////////////////////////////////////////////////////////////

type LifeCyclePhaseFunc func(lc *LifeCycle) error

type OnInitFunc LifeCyclePhaseFunc
type OnCreateFunc LifeCyclePhaseFunc
type OnStartFunc LifeCyclePhaseFunc
type OnResumeFunc LifeCyclePhaseFunc

type OnLoopFunc LifeCyclePhaseFunc

type OnPauseFunc LifeCyclePhaseFunc
type OnStopFunc LifeCyclePhaseFunc
type OnDestroyFunc LifeCyclePhaseFunc
type OnReleaseFunc LifeCyclePhaseFunc

////////////////////////////////////////////////////////////////////////////////

type LifeCycle struct {
	Order   int
	Enabled bool

	OnInit   OnInitFunc
	OnCreate OnCreateFunc
	OnStart  OnStartFunc
	OnResume OnResumeFunc

	OnLoop OnLoopFunc

	OnPause   OnPauseFunc
	OnStop    OnStopFunc
	OnDestroy OnDestroyFunc
	OnRelease OnReleaseFunc
}
