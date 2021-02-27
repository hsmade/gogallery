import React, {Component} from 'react';
import Tree from "./Tree";

class IndexViewer extends Component {
    constructor(props) {
        super(props);

        this.state = {
            entries: {
                Directories: [],
                Files: []
            },
        };
    }
    componentDidMount() {
        fetch("http://localhost:8080/list?path=" + this.props.path).then(res => res.json().then(data => this.setState({entries: data})))
    }

    render() {
        console.log("IndexViewer", this.props, this.state)

        let images
        if (this.state.entries.Files != null) {
            images = this.state.entries.Files.map((image) => {
                const path = this.props.path.replace("/\/$/","") + "/" + image.Name;
                return <a key={image.Name} href={"http://localhost:8080/download?path="+path}>
                    <img alt={image.Name} src={"data:image/jpeg;base64," + image.Image}/>
                </a>
            })
        }

        return (
            <div>
                <div>
                    <b>Path:</b> {this.props.path}
                </div>
                <div align={"left"}>
                    <Tree path={this.props.path} dirs={this.state.entries.Directories}/>
                </div>
                <div align={"right"}>
                    {images}
                </div>
            </div>
        )
    }
}

export default IndexViewer
