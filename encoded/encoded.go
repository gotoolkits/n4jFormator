package encoded

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	graph "github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/labstack/echo"
	"net/http"
)

const (
	NEO4J_URL = "bolt://yunwei:yunweiadmin@localhost:7687"
)

var (
	cypherSearchNodes         = `MATCH (n) RETURN n`
	cypherSearchRelationships = `MATCH p=()-->() RETURN p`
)

type Ne4jEncode struct {
	Drv bolt.DriverPool
}

func InitNe4j() (*Ne4jEncode, error) {

	drvPool, err := bolt.NewDriverPool(NEO4J_URL, 10)
	if err != nil {
		return nil, err
	}
	return &Ne4jEncode{Drv: drvPool}, nil

}

func (n4j *Ne4jEncode) retConn() (bolt.Conn, error) {
	return n4j.Drv.OpenPool()
}

func (n4j *Ne4jEncode) FnFormatNodes(c echo.Context) error {

	conn, err := n4j.retConn()
	if err != nil {
		return err
	}

	defer conn.Close()

	data, _, _, err := conn.QueryNeoAll(cypherSearchNodes, nil)

	if err != nil {
		return c.JSONPretty(http.StatusOK, []byte("Failed"), " ")
	} else {

		iCyto := ToCytoscape{}
		nds := Nodes{}

		for _, v := range data {
			node := v[0].(graph.Node)
			iCyto.Gn = node
			iNdata := iCyto.FormateNode()
			nds.Nodes = append(nds.Nodes, iNdata)
		}

		return c.JSONPretty(http.StatusOK, nds.Nodes, " ")

	}

}

func (n4j *Ne4jEncode) FnFormatRelationships(c echo.Context) error {

	conn, err := n4j.retConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	data, _, _, err := conn.QueryNeoAll(cypherSearchRelationships, nil)

	if err != nil {
		return c.JSONPretty(http.StatusOK, []byte("Failed"), " ")
	} else {

		iCyto := ToCytoscape{}
		edges := Edges{}

		for _, v := range data {
			path := v[0].(graph.Path)

			iCyto.Gp = path
			iEdge := iCyto.FormateEdges()

			edges.Edges = append(edges.Edges, iEdge)

		}

		// js, err := json.Marshal(edges.Edges)
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Println(string(js))

		return c.JSONPretty(http.StatusOK, edges.Edges, " ")

	}

}
