import './App.css';
import HomeHeader from './components/header';
import Products from './components/products';
import Footer from './components/footer';
import {
  BrowserRouter,
  Routes,
  Route,
  Link
} from 'react-router-dom';


function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <Routes>
          <Route exact path='/' 
            element=
              { 
                <div className='main-page'>
                    <HomeHeader />
                    <Products />
                    <Footer />
                </div>
              }></Route>

          <Route exact path='/become_a_seller' element={< HomeHeader />}></Route>
          <Route exact path='/more' element={< HomeHeader />}></Route>
          <Route exact path='/cart' element={< HomeHeader />}></Route>
          <Route exact path='/signin' element={< HomeHeader />}></Route>
          <Route exact path='/signup' element={< HomeHeader />}></Route>
        </Routes>
         
      </div>
    </BrowserRouter>
  );
}
export default App;
