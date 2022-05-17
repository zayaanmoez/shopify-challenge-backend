inventory_init = [
    {
        "name": "Product1",
        "quantity": 100,
        "value": 1000,
        "costPerUnit": 10,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 50}, 
            {"label": "variant2", "stock": 50}
        ]
    },
    {
        "name": "Product2",
        "quantity": 10,
        "value": 3000,
        "costPerUnit": 300,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 10}, 
            {"label": "variant2", "stock": 0}
        ]
    },
    {
        "name": "Product3",
        "quantity": 500,
        "value": 7500,
        "costPerUnit": 75,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 500}
        ]
    },
    {
        "name": "Product94",
        "quantity": 5,
        "value": 75,
        "costPerUnit": 15,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 2}, 
            {"label": "variant2", "stock": 3}
        ]
    },
    {
        "name": "Product5",
        "quantity": 100,
        "value": 1700,
        "costPerUnit": 17,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 70}, 
            {"label": "variant1", "stock": 30}
        ]
    },
    {
        "name": "Product6",
        "quantity": 500,
        "value": 50000,
        "costPerUnit": 100,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 500}
        ]
    },
    {
        "name": "Product7",
        "quantity": 0,
        "value": 0,
        "costPerUnit": 55,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 0}
        ]
    },
    {
        "name": "Product8",
        "quantity": 100,
        "value": 1000,
        "costPerUnit": 10,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 100}
        ]
    },
    {
        "name": "Product9",
        "quantity": 160,
        "value": 4800,
        "costPerUnit": 30,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 100}, 
            {"label": "variant2", "stock": 15},
            {"label": "variant3", "stock": 45}
        ]
    },
    {
        "name": "Product10",
        "quantity": 0,
        "value": 0,
        "costPerUnit": 1000,
        "warehouse": "warehouse1",
        "variants": [
            {"label": "variant1", "stock": 0}, 
            {"label": "variant2", "stock": 0},
            {"label": "variant3", "stock": 0}
        ]
    }
]

db.auth('root', 'root-password')

db = db.getSiblingDB('logistics');

db.createCollection('inventory');
db.createCollection('shipments');

db.inventory.insertMany(inventory_init)