import React, {Component} from 'react';
import TreeMenu from 'react-simple-tree-menu'

class Tree extends Component {
    render() {
        console.log("Tree.render() props:", this.props)

        const tree = CreateTree(this.props.path || "", this.props.dirs)
        console.log("final tree structure", tree)

        let openNodes = ["root"]
        this.props.path.split("/").map((dir) => {
            if (dir !== ""){
                openNodes.push([openNodes[openNodes.length-1], dir].join("/"))
            }
        })
        console.log("openNodes:", openNodes)
        return (
            <div>
                <TreeMenu
                    cacheSearch
                    data={[tree]}
                    debounceTime={125}
                    initialOpenNodes={openNodes}
                    // initialFocusKey={openNodes[openNodes.length-1]}
                    disableKeyboard={false}
                    hasSearch={true}
                    onClickItem={({ key, label, ...props }) => {
                        console.log("Clicked key:",key, "label:",label, "props",props)
                        window.location = window.location.href.split("=")[0] + "=" + key.replace(/^root/,"")
                    }}
                    resetOpenNodesOnDataUpdate={false}
                />
            </div>
        );
    }
}

export function CreateTree(path, dirs) {
    const pathParts = path.split("/")
    const left = pathParts[0]
    const right = pathParts.splice(1).join("/")
    let tree = {}
    if (left === "") { // we are in root
        tree = {
            key: "root",
            label: "/",
            nodes: []
        }
    } else { // we are somewhere else
        tree = {
            key: left,
            label: left,
            nodes: []
        }
    }

    if (right === "" ) { // we are at the end already
        dirs.map((dir) => {
            tree.nodes.push({
                key: dir,
                label: dir,
            })
        })
        return tree
    }

    // since we're in between the root and the end, recurse
    tree.nodes.push(CreateTree(right, dirs))
    return tree
}
// export function CreateTree(path, item, dirs) {
//     console.log("createTree called: path:", path, "item:", item,"dirs:", dirs)
//     const itemParts = item.split("/")
//     console.log("itemParts:", itemParts)
//     const key = [path,itemParts[0]].join("/").replaceAll("//","/")
//     let newTree = {
//         key: key,
//         label: itemParts[0] || "/",
//         nodes: []
//     }
//     if (itemParts.length === 1 || path === null){ // end of the tree
//         dirs.map((dir) => {
//             newTree.nodes.push({
//                 key: [path, dir].join("/").replaceAll("//","/").replace("//","/"),
//                 label: dir,
//             })
//         })
//         return newTree
//     } else { // walk further
//         newTree.nodes.push(CreateTree([path,itemParts[0]].join("/"), itemParts.splice(1).join("/"), dirs))
//         return newTree
//     }
// }

export default Tree