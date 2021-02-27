import './App.css';
import IndexViewer from "./IndexViewer";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <IndexViewer path={new URLSearchParams(window.location.search).get('path')}/>
      </header>
    </div>
  );
}

export default App;
