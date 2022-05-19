import React, { useState, useEffect } from "react";
import { Button, Container, Form, Table, Col } from "react-bootstrap";
import axios from "axios";

function Shipment() {
    const [shipment, setShipment] = useState([]);

    useEffect(() => {
        // fetch(
        //   `/shipment`, { method: "GET" }
        // ).then((res) => res.json()).then((data) => setShipment(data))
        axios.get("/shipments").then(res => setShipment(res.data))
        .catch(err => alert(err));
      }, [])

    return (
        <div>
            <div>
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Label</th>
                        <th>Dest City</th>
                        <th>Dest Warehouse</th>
                        <th>Status</th>
                        <th>Items</th>
                    </tr>
                </thead>
                <tbody>
                    {shipment.map((ship, idx) => {
                        return(
                        <tr key={idx}>
                            <td>{ship._id}</td>
                            <td>{ship.label}</td>
                            <td>{ship.city}</td>
                            <td>{ship.warehouse}</td>
                            <td>{ship.status}</td>
                            <td>
                                <ul>
                                {ship.items.map((item, idx) => {
                                    return(
                                        <li key={idx}>{`${item.name} - ${item.stock} units`}</li>
                                    )
                                })}
                                </ul>
                            </td>
                        </tr>
                        )
                    })}
                </tbody>
                </Table>
            </div>
            {/* <div>
                <Container>
                    <Form>
                        <Form.Group controlId="shipmentId">
                            <Form.Label>Shipment ID</Form.Label>
                            <Form.Control as="select" name="id" defaultValue={"Select Shipment"}>
                                {shipment.map((ship, idx) => {
                                    return(
                                        <option key={idx} value={ship._id}>{ship.label}</option>
                                    )
                                })}
                            </Form.Control>
                        </Form.Group>
                        <Button>Deliver</Button>
                    </Form>
                </Container>
            </div>
            <div>
                <Container>
                    <Form>
                        <Form.Group controlId="shipmentId">
                            <Form.Label>Shipment ID</Form.Label>
                            <Form.Control as="select" name="id" defaultValue={"Select Shipment"}>
                                {shipment.map((ship, idx) => {
                                    return(
                                        <option key={idx} value={ship._id}>{ship.label}</option>
                                    )
                                })}
                            </Form.Control>
                        </Form.Group>
                        <Button>Recieve</Button>
                    </Form>
                </Container>
            </div> */}
            
        </div>
    )
}

export default Shipment;