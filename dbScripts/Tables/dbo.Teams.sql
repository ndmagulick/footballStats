CREATE TABLE [dbo].[Teams]
(
	TeamID INT PRIMARY KEY IDENTITY,
	LeagueID INT,
	TeamName NVARCHAR(64),
	FoundingYear INT
)

ALTER TABLE dbo.Teams ADD CONSTRAINT FK_Teams_Leagues_LeagueID FOREIGN KEY (LeagueID) REFERENCES dbo.Leagues(LeagueID)