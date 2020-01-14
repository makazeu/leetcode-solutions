# Write your MySQL query statement below
SELECT a.Score,
       (SELECT COUNT(DISTINCT Score) FROM Scores WHERE Score > a.Score) + 1 AS Rank
FROM Scores AS a
ORDER BY a.Score DESC

