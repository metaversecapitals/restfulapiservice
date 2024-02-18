const express = require('express');
const app = express();
const bodyParser = require('body-parser');
app.use(bodyParser.json());
let resources = [];
app.get('/resources', function(req, res) {
    res.json(resources);
});
app.post('/resources', function(req, res) {
    resources.push(req.body);
    res.json(req.body);
});
app.get('/resources/:id', function(req, res) {
    let resourceIndex = resources.findIndex(resource => resource.id === parseInt(req.params.id));
    if (resourceIndex !== -1) {
        res.json(resources[resourceIndex]);
    } else {
        res.sendStatus(404);
    }
});
app.put('/resources/:id', function(req, res) {
    let resourceIndex = resources.findIndex(resource => resource.id === parseInt(req.params.id));
    if (resourceIndex !== -1) {
        resources[resourceIndex] = req.body;
        res.json(resources[resourceIndex]);
    } else {
        res.sendStatus(404);
    }
});
app.delete('/resources/:id', function(req, res) {
    let resourceIndex = resources.findIndex(resource => resource.id === parseInt(req.params.id));
    if (resourceIndex !== -1) {
        resources.splice(resourceIndex, 1);
        res.send('Deleted Successfully');
    } else {
        res.sendStatus(404);
    }
});
app.listen(3000, function() {
    console.log('App is listening on port 3000!');
});