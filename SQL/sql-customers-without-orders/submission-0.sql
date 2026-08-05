-- Write your query below
SELECT c.name
FROM customers AS c
LEFT JOIN orders AS o ON o.customer_id = c.id
GROUP BY c.id
HAVING COUNT(o.id) < 1
