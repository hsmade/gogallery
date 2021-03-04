import './App.css';
import React, {Component} from 'react';
import IndexViewer from "./IndexViewer";

class App extends Component {

    render() {
        const path = new URLSearchParams(window.location.search).get('path')
        if (path === null ) {
            window.location = window.location.href + "/"
        }
        return (
            <div className="App">
                <header className="App-header">
                    <IndexViewer path={path}/>
                </header>
            </div>
        );
    }
}

export default App;
