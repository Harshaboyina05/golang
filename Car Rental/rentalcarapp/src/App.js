import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { NavBar } from './components/NavBar';
import { Banner } from './components/Banner';
import DocsPage from './components/Docs';
import SwaggerUI from 'swagger-ui-react';
import 'swagger-ui-react/swagger-ui.css';
import PrivateRoute from './config/auth/privateRoute';
import NotesList from './components/notes/NotesList';
import CustomerList from './components/customer/CustomerList';
import CarSelection from './components/CarSelection';
import UserDetails from './components/UserDetails';
import Payment from './components/Payment';

function App() {
  return (
    <div className="App">
      <NavBar />
      <Router>
        <Routes>
          <Route path="/" element={<Banner />} />
          <Route path="/selectcar" element={<PrivateRoute><CarSelection /></PrivateRoute>} />
          <Route path="/booking" element={<PrivateRoute><UserDetails /></PrivateRoute>} />
          <Route path="/payment" element={<PrivateRoute><Payment /></PrivateRoute>} />
          <Route
            path="/docs"
            element={
              <PrivateRoute>
                <div className="container">
                  <DocsPage />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/swagger/carmanagement"
            element={
              <PrivateRoute>
                <div className="swagger">
                  <SwaggerUI url={process.env.REACT_APP_MICROSERVICE_CARMANAGEMENT.concat('/v3/api-docs')} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/swagger/customermanagement"
            element={
              <PrivateRoute>
                <div className="swagger">
                  <SwaggerUI url={process.env.REACT_APP_MICROSERVICE_CUSTOMERMANAGEMENT.concat('/v3/api-docs')} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/swagger/paymentmanagement"
            element={
              <PrivateRoute>
                <div className="swagger">
                  <SwaggerUI url={process.env.REACT_APP_MICROSERVICE_PAYMENTMANAGEMENT.concat('/v3/api-docs')} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/swagger/reservationmanagement"
            element={
              <PrivateRoute>
                <div className="swagger">
                  <SwaggerUI url={process.env.REACT_APP_MICROSERVICE_RESERVATIONMANAGEMENT.concat('/v3/api-docs')} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/notes/carmanagement"
            element={
              <PrivateRoute>
                <div className="component">
                  <NotesList notesApp={'carmanagement'} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/notes/customermanagement"
            element={
              <PrivateRoute>
                <div className="component">
                  <CustomerList customerApp={'customermanagement'} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/notes/paymentmanagement"
            element={
              <PrivateRoute>
                <div className="component">
                  <NotesList notesApp={'paymentmanagement'} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/notes/reservationmanagement"
            element={
              <PrivateRoute>
                <div className="component">
                  <NotesList notesApp={'reservationmanagement'} />
                </div>
              </PrivateRoute>
            }
          />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
