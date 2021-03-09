CREATE TABLE [dbo].[Matches]
(
	MatchID INT PRIMARY KEY IDENTITY,
	SeasonID INT,
	HomeTeamID INT,
	AwayTeamID INT,
	MatchDay INT,
	MatchDate DATETIME
)

ALTER TABLE dbo.Matches ADD CONSTRAINT FK_Matches_Seasons_SeasonID FOREIGN KEY (SeasonID) REFERENCES dbo.Seasons(SeasonID)
ALTER TABLE dbo.Matches ADD CONSTRAINT FK_SeasonTeams_Teams_HomeTeamID FOREIGN KEY (HomeTeamID) REFERENCES dbo.Teams(TeamID)
ALTER TABLE dbo.Matches ADD CONSTRAINT FK_SeasonTeams_Teams_AwayTeamID FOREIGN KEY (AwayTeamID) REFERENCES dbo.Teams(TeamID)