const express = require('express');
const cors = require('cors');
const port = process.env.PORT || 4000;
const router = require('./server/router');
const app = express();

app.use(express.json());
router(app);

app.listen(port);
