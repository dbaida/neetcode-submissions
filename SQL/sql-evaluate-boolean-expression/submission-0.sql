-- Write your query below
SELECT
    e.left_operand, e.operator, e.right_operand,
    CASE
        WHEN e.operator = '<' THEN vl.value < vr.value
        WHEN e.operator = '>' THEN vl.value > vr.value
        WHEN e.operator = '=' THEN vl.value = vr.value
    END AS value
FROM expressions AS e
INNER JOIN variables AS vl ON vl.name = e.left_operand
INNER JOIN variables AS vr ON vr.name = e.right_operand