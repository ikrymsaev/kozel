package domain

type EStage int

const (
	StagePreparing   EStage = iota
	StagePraising    EStage = iota
	StagePlayerStep  EStage = iota
	StageCalculation EStage = iota
	StageDealerStep  EStage = iota
	StageGameOver    EStage = iota
)
