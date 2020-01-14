# Write your MySQL query statement below
SELECT DISTINCT Num AS ConsecutiveNums
FROM Logs a
WHERE Num = (
		SELECT Num
		FROM Logs
		WHERE Id = a.Id + 1
	)
	AND Num = (
		SELECT Num
		FROM Logs
		WHERE Id = a.Id + 2
	)

