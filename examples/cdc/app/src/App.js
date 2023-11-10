import { Routes, Route } from 'react-router-dom';
import Home from "./Home";
import CreateMovie from './CreateMovie';

function App({ fetch, token }) {
  return (
    <Routes>
      <Route path='/' element={
        <Home fetch={ fetch } />
      } />

      <Route path='/create' element={
        <CreateMovie fetch={ fetch } token={ token } />
      } />
    </Routes>
  );
}

export default App;
