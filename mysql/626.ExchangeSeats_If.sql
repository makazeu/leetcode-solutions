# Write your MySQL query statement below
SELECT id
	, IF(id % 2 = 1, (
		SELECT student
		FROM seat
		WHERE id = IF(a.id = (
			SELECT MAX(id)
			FROM seat
		), a.id, a.id + 1)
	), (
		SELECT student
		FROM seat
		WHERE id = a.id - 1
	)) AS student
FROM seat a

