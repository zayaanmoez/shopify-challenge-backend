import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import React from "react";
import { Inventory, Shipment } from "./pages";
import { Navbar, Nav, Container } from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {


  return (
    <div className="App">
      <Router>
        <Navbar bg="light" expand="lg">
          <Container>
            <Navbar.Brand href="#home">SOME COMPANY</Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
              <Nav className="me-auto">
                <Nav.Link href="/inventoryPage">Inventory</Nav.Link>
                <Nav.Link href="/shipmentPage">Shipment</Nav.Link>
              </Nav>
            </Navbar.Collapse>
          </Container>
        </Navbar>
        <Routes>
          <Route exact path="/" element={<Inventory/>} />
          <Route path="inventoryPage/" element={<Inventory/>} />
          <Route path="shipmentPage/" element={<Shipment/>} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
