CREATE TABLE [dbo].[MatchEvents]
(
	MatchEventID INT PRIMARY KEY IDENTITY,
	MatchID INT,
	TeamID INT,
	PlayerID INT,
	EventType INT, --This will be referencing an enum
	EventMinute INT,
	StoppageTime INT
)

ALTER TABLE dbo.MatchEvents ADD CONSTRAINT FK_MatchEvents_Matches_MatchID FOREIGN KEY (MatchID) REFERENCES dbo.Matches(MatchID)
ALTER TABLE dbo.MatchEvents ADD CONSTRAINT FK_MatchEvents_Teams_TeamID FOREIGN KEY  (TeamID) REFERENCES dbo.Teams(TeamID)
ALTER TABLE dbo.MatchEvents ADD CONSTRAINT FK_MatchEvents_Players_PlayerID FOREIGN KEY (PlayerID) REFERENCES dbo.Players(PlayerID)