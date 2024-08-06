import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import App from './App.jsx';
import Page from './page/Page.jsx';
import BasketPage from './page/BasketPage.jsx';
import Login from './page/Login.jsx';
import Page2 from './page/Page2.jsx';
import Customer from './page/Customer.jsx';


ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/page/:id" element={<Page />} />
        <Route path="/page2/:type" element={<Page2 />} />
        <Route path="/basketPage" element={<BasketPage />} />
        <Route path="/login" element={<Login />} />
        <Route path="/customer" element={<Customer />} />
      </Routes>
    </Router>
  </React.StrictMode>
);
