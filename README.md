# collections
Dynamic arrays


## What are the functions?

* collections
  * struct
    * PageList （用于分页数组，包含总记录数）
    * IList（集合）
    * List（泛型集合）
      * Add（添加元素）
      * Count（集合大小）
      * ToArray（转成数组）
      * IsEmpty（集合是为空的）
      * Any（是否存在）
      * All（是否所有数据都满足fn条件）
      * Index（获取第index索引位置的元素）
      * First（查找符合条件的第一个元素）
      * Where（对数据进行筛选）
      * Contains（是否包含元素）
      * IndexOf（元素在集合的索引位置）
      * Remove（移除元素）
      * RemoveAt（移除指定索引的元素）
      * RemoveAll（移除条件=true的元素）
      * Insert（向第index索引位置插入元素）
      * Clear（清空集合）
      * MapToList（类型转换）
      * MapToArray（类型转换）
    * func
      * NewIList（创建集合）
      * NewList（创建集合）
      * NewPageList （数据分页列表及总数）