CREATE TABLE [dbo].[MatchStats]
(
	MatchStatsID INT PRIMARY KEY,
	MatchID INT,
	PossessionHome INT,
	TotalShotsHome INT,
	ShotsOnTargetHome INT,
	SavesHome INT,
	CornersHome INT,
	FoulsHome INT,
	OffsidesHome INT,
	PossessionAway INT,
	TotalShotsAway INT,
	ShotsOnTargetAway INT,
	SavesAway INT,
	CornersAway INT,
	FoulsAway INT,
	OffsidesAway INT
)

ALTER TABLE dbo.MatchStats ADD CONSTRAINT FK_MatchStats_Matches_MatchID FOREIGN KEY (MatchID) REFERENCES dbo.Matches(MatchID)