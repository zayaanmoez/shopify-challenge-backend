import React, { useState, useEffect} from "react";
import { Button, Container, Form, Table, Col } from "react-bootstrap";
import axios from "axios";


function Inventory() {
    const [inventory, setInventory] = useState([]);

    useEffect(() => {
        // fetch(
        //   `/inventory`, { method: "GET" }
        // ).then((res) => res.json()).then((data) => setInventory(data))
        axios.get("/inventory").then(res => setInventory(res.data))
        .catch(err => alert(err));
      }, [])

    return (
        <div>
            <div>
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Stock</th>
                        <th>CostPerUnit</th>
                        <th>City</th>
                        <th>Warehouse</th>
                    </tr>
                </thead>
                <tbody>
                    {inventory.map((item, idx) => {
                        return(
                        <tr key={idx}>
                            <td>{item._id}</td>
                            <td>{item.name}</td>
                            <td>{item.stock}</td>
                            <td>{item.costPerUnit}</td>
                            <td>{item.city}</td>
                            <td>{item.warehouse}</td>
                        </tr>
                        )
                    })}
                </tbody>
                </Table>
            </div>
            {/* <div>
                <Button>Create</Button>
            </div>
            <div>
                <Button>Update</Button>
            </div>
            <div>
            <div>
                <Container>
                    <Form>
                        <Form.Group controlId="inventoryId">
                            <Form.Label>Inventory ID</Form.Label>
                            <Form.Control as="select" name="id" defaultValue={"Select Inventory"}>
                                {inventory.map((item, idx) => {
                                    return(
                                        <option key={idx} value={item._id}>{`${item.name} - ${item.warehouse}`}</option>
                                    )
                                })}
                            </Form.Control>
                        </Form.Group>
                        <Button>Delete</Button>
                    </Form>
                </Container>
            </div>
            </div> */}
            
        </div>
    )
}

export default Inventory;