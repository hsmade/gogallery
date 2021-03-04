import React, {Component} from 'react';

class Tree extends Component {
    render() {
        console.log("tree", this.props)

        // split of "/" results in ["",""]
        let pathParts
        if (this.props.path === "/" || this.props.path === null) {
            pathParts = [""]
        } else {
            pathParts = this.props.path.split("/")
        }
        return (
            <div>
                {treeItem([], pathParts, this.props.dirs || [])}
            </div>
        );
    }
}

// function that renders a <ul> with the first part of `right` as item
// Then calls itself with left having that part pushed into it.
function treeItem(left, right, dirs) {
    if (right.length === 0) {
        const items = dirs.map((dir) =>
            <li key={dir}><a href={"?path=" + left.join("/")  + "/" + dir}>|- {dir}</a></li>
        )
        return <ul>{items}</ul>
    }

    const current = right.shift()
    left.push(current)
    return <ul>
        <li>
            <a href={"?path=" + left.join("/")}>|- {current}/</a>
        </li>
        {treeItem(left, right, dirs)}</ul>
}

export default Tree