-- Write your query below
select employee_id, 
    CASE 
        WHEN employee_id % 2 = 1 
            AND 
        employees.name NOT LIKE 'M%'
        THEN employees.salary 
        ELSE 0 
    END as bonus 
from employees 
order by employee_id
    
