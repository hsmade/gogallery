import React, {Component} from 'react';

class Tree extends Component {
    render() {

        let dir_up
        if (this.props.path !== "/") {
            let dir_components = this.props.path.split("/")
            dir_components.pop()
            dir_up = <li key="up"><a href={"?path=" + dir_components.join("/")}>Terug</a></li>
        }
        const dirs = this.props.dirs.map((dir) =>
            <li key={dir}><a href={"?path=" + this.props.path  + "/" + dir}>{dir}</a></li>
        )

        // show link per path's component, with indents
        // show dirs at the bottom, as list
        return (
            <div>
                {dir_up}
                {dirs}
            </div>
        );
    }
}
export default Tree