const options = {
    definition: {
        openapi: "3.0.0",
        info: {
            title: "SRIN Backend API",
            version: "0.1.0",
            description:
                "This is a simple CRUD API application made with Express and documented with Swagger",
            license: {
                name: "MIT",
                url: "https://spdx.org/licenses/MIT.html",
            },
            contact: {
                name: "LogRocket",
                url: "https://logrocket.com",
                email: "info@email.com",
            },
        },
        servers: [
            {
                url: "http://localhost:3000",
            },
        ],
    },
    apis: ["./routes/*.js"],
};

const express = require('express');
const path = require('path');
const cookieParser = require('cookie-parser');
const bodyParser = require('body-parser')
const logger = require('morgan');
const swaggerUi = require('swagger-ui-express');
const MongoClient = require('mongodb').MongoClient
const swaggerJSDoc = require('swagger-jsdoc');
const swaggerSpec = swaggerJSDoc(options);

const brandRouter = require('./routes/brand');

const app = express();

// MongoClient.connect('mongodb+srv://fahrulalwan:Pradipta21@cluster0.bb7cuyb.mongodb.net/?retryWrites=true&w=majority',
//     (err, client) => {
//
// console.log('Connected to Database')
// const db = client.db('srin-db')
// const quotesCollection = db.collection('quotes')
//     })

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({extended: false}));
app.use(cookieParser());
app.use(bodyParser.urlencoded({extended: true}))
app.use(express.static(path.join(__dirname, 'public')));

app.use('/brands', brandRouter);
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerSpec));

module.exports = app;
