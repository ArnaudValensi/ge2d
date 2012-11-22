package ge2d

import "fmt"

type SceneManager struct {
	rootNode	INode
	nodeMap		map[uint]INode
}

func NewSceneManager() *SceneManager {
	rootNode := NewBaseNode("root", nil, Vector2d {0, 0})
	nodeMap := make(map[uint]INode)
	nodeMap[rootNode.GetId()] = rootNode
	return &SceneManager {rootNode, nodeMap}
}

func (this *SceneManager) AddNode(node INode) {
	this.nodeMap[node.GetId()] = node
}

func (this *SceneManager) HandleMessage(message IMessage) {
	fmt.Print("[SceneManager] handleMessage: ", message)
	fmt.Print("[SceneManager] destination object: ", message.GetDestinationObjectId())

	this.nodeMap[message.GetDestinationObjectId()].HandleMessage(message)
}
