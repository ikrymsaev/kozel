package domain

type EStage int

const (
	StagePreparing  EStage = iota
	StagePraising   EStage = iota
	StagePlayerStep EStage = iota
	StageDealerStep EStage = iota
	StageEnd        EStage = iota
)
