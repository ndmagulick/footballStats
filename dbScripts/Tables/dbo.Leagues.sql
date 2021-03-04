CREATE TABLE [dbo].[Leagues]
(
	LeagueID INT PRIMARY KEY IDENTITY,
	LeagueName NVARCHAR(64),
	Country NVARCHAR(3),		-- ISO Code
	Tier INT
)