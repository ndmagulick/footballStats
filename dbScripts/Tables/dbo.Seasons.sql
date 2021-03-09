CREATE TABLE [dbo].[Seasons]
(
	SeasonID INT PRIMARY KEY IDENTITY,
	LeagueID INT,
	StartYear INT,
	EndYear INT
)

ALTER TABLE dbo.Seasons ADD CONSTRAINT FK_Seasons_Leagues_LeagueID FOREIGN KEY (LeagueID) REFERENCES dbo.Leagues(LeagueID)