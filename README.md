# n4jFormator
Transfer Neo4j format data to   Cytoscape.js Elements format


#### Description
 
 获取Neo4j数据转换为Cytoscape.js Elements格式数据
 Transfer Neo4j format data to   Cytoscape.js Elements format

#### API And Cypher
获取所有节点信息 
To get all formatted nodes data 

> API:    `/api/getNodes`
 Cypher:  `MATCH (n) RETURN n`

获取所有所有关系信息
To get all formatted nodes relationships(paths) 
> API:  `/api/getPaths`
 Cypher:  `MATCH p=()-->() RETURN p`
 
