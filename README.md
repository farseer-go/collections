# collections
Dynamic arrays


## What are the functions?

* collections
  * struct
    * PageList （用于分页数组，包含总记录数）
    * list
      * Index（获取第index索引位置的元素）
      * IndexOf（元素在集合的索引位置）
      * Insert（向第index索引位置插入元素）
      * RemoveAt（移除指定索引的元素）
    * collection
      * Count（集合大小）
      * Add（添加元素）
      * Clear（清空集合）
      * Remove（移除元素）
      * RemoveAll（移除条件=true的元素）
    * enumerable
      * Any（是否存在）
      * IsEmpty（集合是为空的）
      * First（查找符合条件的第一个元素）
      * Last（集合最后一个元素）
      * Contains（是否包含元素）
      * Where（对数据进行筛选）
      * All（是否所有数据都满足fn条件）
      * Take（返回前多少条数据）
      * Skip（跳过前多少条记录）
      * Sum（求总和）
      * SumItem（求总和）
      * Average（求平均数）
      * AverageItem（求平均数）
      * GroupBy（将数组进行分组后返回map）
      * OrderBy（正序排序）
      * OrderByItem（正序排序）
      * OrderByDescending（倒序排序）
      * OrderByDescendingItem（倒序排序）
      * Intersect（两个集合的交集）
      * Min（获取最小值）
      * MinItem（获取最小值）
      * Max（获取最大值）
      * MaxItem（获取最大值）
      * Select（筛选子元素字段）
      * SelectMany（筛选子元素字段）
      * ToMap（转成字典）
      * ToList（返回List集合）
      * ToArray（转成数组）
      * ToPageList（数组分页）
      * MapToList（类型转换）
      * MapToArray（类型转换）
      * Concat（合并两个集合）
      * Union（合并两个集合，并去重）
      * Distinct（集合去重）
      * Empty（返回一个新的Empty集合）
      * Except（移除参数中包含的集合元素）
    * List（泛型集合）
      * AsEnumerable（返回enumerable类型）
    * func
      * NewList（创建集合）
      * NewPageList （数据分页列表及总数）