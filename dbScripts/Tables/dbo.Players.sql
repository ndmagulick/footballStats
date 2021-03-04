CREATE TABLE [dbo].[Players]
(
	PlayerID INT PRIMARY KEY IDENTITY,
	FirstName NVARCHAR(32),
	LastName NVARCHAR(32),
	DateOfBirth DATE,
	Height DECIMAL(3,2),	 -- In meters
	Nationality NVARCHAR(3), -- ISO Code
	Position NVARCHAR(2),    -- Position enum
	TeamID INT,
	Number INT,
	Foot NVARCHAR(1),	     -- L, R
)

ALTER TABLE dbo.Players ADD CONSTRAINT FK_Players_Teams_TeamID FOREIGN KEY (TeamID) REFERENCES dbo.Teams(TeamID)