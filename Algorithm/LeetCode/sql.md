- [175.组合两个表](https://leetcode.cn/problems/combine-two-tables/)
```mysql
SELECT Person.FirstName As firstName, Person.LastName AS lastName, Address.City AS city, Address.State As state
FROM Person LEFT JOIN Address
ON Person.PersonId = Address.PersonId;
```
- [176.第二高的薪水](https://leetcode.cn/problems/second-highest-salary/)
```mysql
# 使用临时表解决null的问题, 使用distinct去重
SELECT 
    (SELECT DISTINCT Salary FROM Employee ORDER BY Salary DESC LIMIT 1 OFFSET 1 )
    AS SecondHighestSalary;
# 使用ifnull函数解决null的问题, 使用distinct去重
SELECT
    IFNULL(
            (SELECT DISTINCT Salary
             FROM Employee
             ORDER BY Salary DESC
             LIMIT 1 OFFSET 1),
            NULL) AS SecondHighestSalary;
```