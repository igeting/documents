# explain

## example
```
explain select distinct course_id from course where course_term = 'xxx';
 
NOTICE:  QUERY PLAN:
 
Unique  (cost=12223.09..12339.76 rows=4667 width=4)
 
               -> Sort (cost=12223.09..12223.09 rows=46666 width=4)
 
               ->  Seq Scan on course  (cost=0.00..8279.99 rows=46666 width=4)
1.从下往上读，从右往左
2.explain报告查询的操作，开启的消耗，查询总的消耗，访问的行数 访问的平均宽度
3.开启时间消耗是输出开始前的时间例如排序的时间
4.消耗包括磁盘检索页，cpu时间 
5.注意，每一步的cost包括上一步的
6.重要的是，explain 不是真正的执行一次查询 只是得到查询执行的计划和估计的花费
```

## options
|执行计划运算类型|操作说明|是否有启动时间|
|---|---|---|
SeqScan|扫描表|无启动时间
IndexScan|索引扫描||无启动时间
Bitmap|IndexScan索引扫描|有启动时间
Bitmap|HeapScan索引扫描|有启动时间
Subquery|Scan子查询|无启动时间
TidScan|ctid=…条件|无启动时间
FunctionScan|函数扫描|无启动时间
NestedLoop|循环结合|无启动时间
MergeJoin|合并结合|有启动时间
HashJoin|哈希结合|有启动时间
Sort|排序，ORDERBY操作|有启动时间
Hash|哈希运算|有启动时间
Result|函数扫描，和具体的表无关|无启动时间
Unique|DISTINCT，UNION操作|有启动时间
Limit|LIMIT，OFFSET操作|有启动时间
Aggregate|count,sum,avg,stddev集约函数|有启动时间
Group|GROUPBY分组操作|有启动时间
Append|UNION操作|无启动时间
Materialize|子查询|有启动时间
SetOp|INTERCECT，EXCEPT|有启动时间