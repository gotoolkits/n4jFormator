package encoded

import (
	"strconv"

	graph "github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

type ToCytoscape struct {
	Gn graph.Node
	Gp graph.Path
}

type Nodes struct {
	Nodes []NodeData `json:"nodes"`
}
type NodeData struct {
	Data NodeProperties `json:"data"`
}
type NodeProperties struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Ipaddr string `json:"ipaddr"`
}

type Edges struct {
	Edges []EdgeData `json:"edges"`
}
type EdgeData struct {
	Data EdgeProperties `json:"data"`
}
type EdgeProperties struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label"`
}

func (t *ToCytoscape) FormateNode() NodeData {

	ndata := NodeData{}
	ndata.Data.Id = strconv.FormatInt(t.Gn.NodeIdentity, 10)
	ndata.Data.Label = t.Gn.Properties["hostname"].(string)
	ndata.Data.Ipaddr = t.Gn.Properties["ipaddr"].(string)

	return ndata
}

func (t *ToCytoscape) FormateEdges() EdgeData {
	eData := EdgeData{}
	srcId := strconv.FormatInt(t.Gp.Nodes[0].NodeIdentity, 10)
	dstId := strconv.FormatInt(t.Gp.Nodes[1].NodeIdentity, 10)
	eData.Data.Id = srcId + "-" + dstId
	eData.Data.Source = srcId
	eData.Data.Target = dstId
	eData.Data.Label = t.Gp.Relationships[0].Type

	return eData

}
