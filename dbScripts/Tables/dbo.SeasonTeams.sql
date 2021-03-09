CREATE TABLE [dbo].[SeasonTeams]
(
	SeasonTeamID INT PRIMARY KEY IDENTITY,
	SeasonID INT,
	TeamID INT
)

ALTER TABLE dbo.SeasonTeams ADD CONSTRAINT FK_SeasonTeams_Seasons_SeasonID FOREIGN KEY (SeasonID) REFERENCES dbo.Seasons(SeasonID)
ALTER TABLE dbo.SeasonTeams ADD CONSTRAINT FK_SeasonTeams_Teams_TeamID FOREIGN KEY (TeamID) REFERENCES dbo.Teams(TeamID)