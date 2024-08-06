import './App.css'
import './css/header.css';
import './css/style.css';
import Header from './components/Header.jsx';
import DataUpload from './components/DataUpload';

function App() {

  return (
    <div className="main-body">
      <Header />
      <DataUpload content={"hepsi"} />
    </div>

  )
}

export default App
