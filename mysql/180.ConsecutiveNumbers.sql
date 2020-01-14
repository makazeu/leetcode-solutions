# Write your MySQL query statement below
SELECT DISTINCT a.Num AS ConsecutiveNums
FROM Logs a
	JOIN Logs b
	ON a.Num = b.Num
		AND a.Id + 1 = b.Id
	JOIN Logs c
	ON a.Num = c.Num
		AND a.Id + 2 = c.Id

